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
            rows="20"
            cols="60"
            autocapitalize="off" 
            spellcheck="false" 
            contenteditable="true">{{ examFunctionSignature(question) }}</textarea>
    </div>

    <button @click="onSubmit">Submit Exam</button>

</template>

<script>
export default {
    name: 'Exam',
    inject: ['fetchExam', 'fetchQuestionsFromExam', 'fetchQuestion'],
    components: {},
    /* emits: ['student-exam-active'], */
    data() {
        return {
            examID: this.$route.params.id,
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
            // Post answers, and examID to grader
            const payload = {
                examid: this.examID,
                userid: this.$route.params.username,
                autograded: false,
                points: this.totalExamPoints(this.exam),
            };
            console.log(payload);
            // Post question to DATABASE
            /* const res = await fetch('http://localhost:5000/questions', { */
            /*     method: 'post', */
            /*     headers: { */
            /*         'Content-type': 'application/json', */
            /*     }, */
            /*     body: JSON.stringify(payload) */
            /* }); */

            // Move examID from incompleted to completed
            // Redirect to student dashboard

        },
        totalExamPoints(exam) {
            if (!exam.questions) return 0;
            let total = 0;
            for (const question of exam.questions) 
                total += question.points;
            return total;
        }
    },
    async created() {
        this.questions = await this.fetchQuestionsFromExam(this.examID);
        this.exam = await this.fetchExam(this.examID);
    }
}
</script>
