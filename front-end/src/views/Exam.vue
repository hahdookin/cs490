<template>
    <h1>{{ exam.name }} (ID: {{ examID }}) Total Points: {{ totalExamPoints(exam) }}</h1>

    <div :key="question.id" v-for="(question, i) in questions">

        <!-- Question info --> 
        <h3>{{ i + 1 }}) {{ question.title }}</h3> 
        <h5>Points: {{ question.points }}</h5>
        <p>{{ question.desc }}</p>
        <p>Examples:</p>
        <div>
            <p v-for="test in question.tests">
            <code>
            {{question.functionname}}({{test.arguments.join(',')}}) -&gt; {{test.output}}
            </code>
            </p>
        </div>

        <!-- Student answer section -->
        <textarea 
            @keydown="onKeydown($event)"
            v-model="question.code"
            rows="20"
            cols="60"
            autocapitalize="off" 
            spellcheck="false" 
            contenteditable="true">{{ question.code }}</textarea>
    </div>

    <button @click="onSubmit">Submit Exam</button>

</template>

<script>
export default {
    name: 'Exam',
    inject: [
        'fetchExam', 
        'fetchQuestionsFromExam', 
        'fetchQuestion',
        'fetchStudentExams',
    ],
    components: {},
    /* emits: ['student-exam-active'], */
    data() {
        return {
            examID: +this.$route.params.examid,
            questions: [],
            exam: {}
        };
    },
    methods: {
        examFunctionSignature(question) {
            let res = 'def ';
            res += question.functionname;
            res += `(${question.parameters.join(',')}):`;
            res += '\n\t';
            return res;
        },
        async onSubmit() {
            const studentUserID = +this.$route.params.userid;

            // Post studentexamresult and grab the ID
            const payload = {
                examid: this.examID,
                userid: studentUserID,
                autograded: false,
                points: this.totalExamPoints(this.exam),
            };
            const res = await fetch('http://localhost:5000/studentexamresult', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(payload)
            });
            const studentexamresult = await res.json();

            // Post studentexamanswers
            for (const question of this.questions) {
                const payload = {
                    questionid: question.id,
                    studentexamresultid: studentexamresult.id,
                    code: question.code,
                };
                const res = await fetch('http://localhost:5000/studentexamanswers', {
                    method: 'post',
                    headers: {
                        'Content-type': 'application/json',
                    },
                    body: JSON.stringify(payload)
                });
            }
            // Move examID from incompleted to completed
            const studentExams = await this.fetchStudentExams(studentUserID);
            const i = studentExams.incompleted.indexOf(this.examID);
            studentExams.incompleted.splice(i, 1);
            studentExams.completed.push(this.examID);
            const res2 = await fetch(`http://localhost:5000/studentexams/${studentExams.id}`, {
                method: 'put',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(studentExams)
            });
            // Redirect to student dashboard
            this.$router.push(`/student/${studentUserID}`);
        },
        totalExamPoints(exam) {
            if (!exam.questions) return 0;
            let total = 0;
            for (const question of exam.questions) 
                total += question.points;
            return total;
        },
        onKeydown(e) {
            if (e.key == 'Tab') {
                const tar = e.target;
                e.preventDefault();
                var start = tar.selectionStart;
                var end = tar.selectionEnd;

                // set textarea value to: text before caret + tab + text after caret
                tar.value = tar.value.substring(0, start) +
                "\t" + tar.value.substring(end);

                // put caret at right position again
                tar.selectionStart = tar.selectionEnd = start + 1;
            }
        }
    },
    async created() {
        this.questions = await this.fetchQuestionsFromExam(this.examID);
        this.questions.forEach(q => {
            // Have each questions answer box start with the function signature
            q.code = this.examFunctionSignature(q);
        });
        this.exam = await this.fetchExam(this.examID);
    }
}
</script>
