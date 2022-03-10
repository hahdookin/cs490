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
                <!-- Students answer -->
                <div class="grades-container">
                    <textarea class="grade-item" disabled="true" rows="20" cols="60">{{ studentsAnswer(question.id).code }}</textarea>
                </div>
                <!-- Instructors point column table -->
                <div class="grades-container">
                    <table>
                        <tr>
                            <th colspan="2">Runs?</th>
                            <td colspan="1">{{ question.runs ? "Yes" : "No" }}</td>
                        </tr>
                        <tr>
                            <th colspan="2">Correct Name?</th>
                            <td colspan="1">{{ question.namecorrect ? "Yes" : "No" }}</td>
                            <td><input type="number" size="4" :max="1" v-model="question.namecorrectpoints" min="0"></td>
                        </tr>
                    </table>
                    <!-- Test cases and point assignment -->
                    <table>
                        <tr>
                            <th>Expected</th>
                            <th>Run</th>
                            <th>Pass?</th>
                            <th>Points</th>
                        </tr>
                        <tr v-for="test in question.tests">
                            <td>{{ testStr(test, question.functionname) }}</td>
                            <td>{{ test.studentoutput }}</td>
                            <td>{{ test.pass ? "Yes" : "No" }}</td>
                            <td><input type="number" size="4" :max="test.maxpoints" v-model="test.points" min="0"></td>
                        </tr>
                    </table>
                    <!-- Instructor comment box -->
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
        'putStudentExamAnswer',
        'putStudentExamResult'
    ],
    data() {
        return {
            teacherUserID: +this.$route.params.userid,
            studentExamResultID: +this.$route.params.studentexamresultid,
            studentExamResult: {},
            studentExamAnswers: [],
            exam: {},
            questions: [],
            loaded: false,
        };
    },
    methods: {
        async onSubmit() {
            // Get total points awarded
            let totalPoints = 0;
            this.questions.forEach(q => {
                q.tests.forEach(t => {
                    totalPoints += t.points;
                });
                totalPoints += q.namecorrectpoints;
            });

            // Update studentexamanswer points amount and the comment
            for (const q of this.questions) {
                const studentsAnswer = this.studentsAnswer(q.id);
                for (const [qTest, sTest] of this.zip(q.tests, studentsAnswer.tests)) {
                    sTest.points = qTest.points;
                }
                studentsAnswer.namecorrectpoints = q.namecorrectpoints;
                studentsAnswer.comment = q.comment;
            }

            // Set the total points for studentexamresult and mark it reviewed
            this.studentExamResult.points = totalPoints;
            this.studentExamResult.autograded = true; // REMOVEME:
            this.studentExamResult.reviewed = true;

            // Put the studentexamanswers and studentexamresult

            // Put the studentexamanswers
            for (const sea of this.studentExamAnswers) {
                const res = await this.putStudentExamAnswer(sea);
            }

            // Put the studentexamresult
            const res = await this.putStudentExamResult(this.studentExamResult);

            // Redirect back to home
            this.$router.push(`/teacher/${this.userid}/`);
        },
        studentsAnswer(qid) {
            const answer = this.studentExamAnswers.find(
                a => a.questionid === qid && a.studentexamresultid === this.studentExamResultID
            );
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
                qTest.studentoutput = sTest.studentoutput;
            }
            const testsCount = question.tests.length;
            const pointDist = this.split(question.points - (question.points === 0 ? 0 : 1), testsCount);
            for (const [qTest, maxPoints] of this.zip(question.tests, pointDist)) {
                qTest.maxpoints = maxPoints;
            }
            question.runs = studentsAnswer.runs;
            question.namecorrect = studentsAnswer.namecorrect;
            question.namecorrectpoints = studentsAnswer.namecorrectpoints;
            question.comment = studentsAnswer.comment;
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


