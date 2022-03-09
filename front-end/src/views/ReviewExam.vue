<template>
    <h1>Review Exam</h1>
    <div v-if="loaded">
        <h1>{{ exam.name }} (ID: {{ exam.id }})</h1>

        <div :key="question.id" v-for="(question, i) in questions">

            <!-- Question info -->
            <h3>{{ i + 1 }}) {{ question.title }}</h3>
            <h5>Points: {{ question.points }}</h5>
            <p>{{ question.desc }}</p>
            <p>Examples:</p>
            <div>
                <p v-for="test in question.tests">
                <code>
                {{ testStr(test, question.functionname) }}
                </code>
                </p>
            </div>

            <!-- Student answer and point selector -->
            <div class="columns-container">
                <div class="grades-container">
                    <textarea class="grade-item" disabled="true" rows="20" cols="60">{{ studentsAnswer(question.id).code }}</textarea>
                </div>
                <div class="grades-container">
                    <table>
                        <tr>
                            <th>Expected</th>
                            <th>Run</th>
                            <th>Pass?</th>
                            <th>Points</th>
                        </tr>
                        <tr v-for="test in question.tests">
                            <td>{{ testStr(test, question.functionname) }}</td>
                            <td>4</td>
                            <td>{{ test.pass ? "Yes" : "No" }}</td>
                            <td><input type="number" size="4" :max="test.points" v-model="test.points" min="0"></td>
                        </tr>
                    </table>
                    <div class="comment-box">
                        <p>Comment:</p>
                        <textarea v-model="question.comment"></textarea>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <button @click="onSubmit">Finalize Grades</button>
</template>

<script>
export default {
    name: 'ReviewExam',
    inject: [
        'zip',
        'split',
        'fetchExam',
        'fetchQuestionsFromExam',
        'fetchStudentExamResult',
        'fetchStudentExamResultByID',
        'fetchStudentExamAnswers',
    ],
    data() {
        return {
            teacherUserID: this.$route.params.userid,
            studentExamResultID: this.$route.params.studentexamresultid,
            studentExamResult: {},
            studentExamAnswers: [],
            exam: {},
            questions: [],
            loaded: false,
        };
    },
    methods: {
        async onSubmit() {
            let totalPoints = 0;
            this.questions.forEach(q => {
                q.tests.forEach(t => {
                    totalPoints += t.points;
                });
            });
            this.questions.forEach(q => {
                console.log(q);
            });
            console.log(totalPoints);
            // Update student exam results
            // TODO: When posting, get the test case points 
            // from question.tests[i].points and put to 
            // studentexamanswers
        },
        studentsAnswer(qid) {
            const answer = this.studentExamAnswers.find(a => a.questionid === qid);
            return answer;
        },
        testStr(test, fname) {
            return `${fname}(${test.arguments.join(',')}) -> ${test.output}`;
        }
    },
    async created() {
        this.studentExamResult = await this.fetchStudentExamResultByID(this.studentExamResultID);
        this.exam = await this.fetchExam(this.studentExamResult.examid);
        this.studentExamAnswers = await this.fetchStudentExamAnswers(this.studentExamResultID);
        this.questions = await this.fetchQuestionsFromExam(this.studentExamResult.examid);

        // Add the points from the students answer and whether
        // or not they passed to the questions test (for rendering)
        for (const question of this.questions) {
            const studentsAnswer = this.studentsAnswer(question.id);
            for (const [qTest, sTest] of this.zip(question.tests, studentsAnswer.tests)) {
                qTest.points = sTest.points;
                qTest.pass = sTest.pass;
            }
        }

        this.loaded = true;
    }
}
</script>


<style scoped>
    
body {
    margin: 0;
}
.column-title {
    color: white;
}

.grades-container {
    background-color: #344;
    padding: 1rem;
    margin-top: 1rem;
    flex: 1;
}

.grade-item {
    padding: 1rem;
    background-color: white;
    border: 1px solid blue;
}

.columns-container {
    display: flex;
    justify-content: safe center;
    /* justify-content: right; /1* REMOVE ME *1/ */
    width: 75%;
    margin: 0 auto;
    max-width: 1000px;
}

.comment-box {
    background-color: white;
    padding: 10px;
}

table {
    background-color: white;
    width: 100%;
    border: 1px solid black;
}
th, tr, td {
    border: 1px solid black;
}
</style>


