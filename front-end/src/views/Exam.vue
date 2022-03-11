<template>
    <h1>{{ exam.name }} (ID: {{ examID }}) Total Points: {{ totalExamPoints(exam) }}</h1>

    <div :key="question.id" v-for="(question, i) in questions">

        <!-- Question info --> 
        <h3>{{ i + 1 }}) {{ question.title }}</h3> 
        <h5>Points: {{ question.points }}</h5>
        <p>{{ question.desc }}</p>
        <p>Signature: {{ question.functionname }}({{ question.parameters.join(', ') }})</p>
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
        'zip',
        'split',
        'fetchExam', 
        'fetchQuestionsFromExam', 
        'fetchQuestion',
        'fetchStudentExams',
        'postStudentExamResult',
        'postStudentExamAnswer',
        'updateStudentExams',
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
                reviewed: false,
                points: this.totalExamPoints(this.exam),
            };
            const res = await this.postStudentExamResult(payload);
            const studentexamresult = await res.json();

            // Post studentexamanswers
            for (const question of this.questions) {
                const testsPayload = [];
                const testsCount = question.tests.length;
                const points = question.points;
                // Subtract 1 point for the correct fname
                const pointDist = this.split(points - (points === 0 ? 0 : 1), testsCount);
                for (const [test, p] of this.zip(question.tests, pointDist)) {
                    // TODO: pass will be handled by querying autograder,
                    // for now just put random pass or fail
                    testsPayload.push({
                        pass: false,
                        points: p || 0,
                        studentoutput: ''
                    });
                }
                const payload = {
                    questionid: question.id,
                    studentexamresultid: studentexamresult.id,
                    runs: false,
                    namecorrect: false,
                    namecorrectpoints: 0,
                    code: question.code,
                    tests: testsPayload,
                    comment: '',
                };
                const res = await this.postStudentExamAnswer(payload);
            }

            // Move examID from incompleted to completed
            const studentExams = await this.fetchStudentExams(studentUserID);
            const i = studentExams.incompleted.indexOf(this.examID);
            studentExams.incompleted.splice(i, 1);
            studentExams.completed.push(this.examID);
            const res2 = await this.updateStudentExams(studentExams);

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
            //q.code = this.examFunctionSignature(q);
            q.code = "";
        });
        this.exam = await this.fetchExam(this.examID);
    }
}
</script>
