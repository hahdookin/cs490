<template>
    <h1>Review Exam</h1>
    <div v-if="loaded">
        <h1>{{ exam.name }} (ID: {{ exam.id }})</h1>

        <div :key="question.id" v-for="(question, i) in questions">

            <!-- Question info -->
            <QuestionDescription :question="question" :number="i + 1"/>

            <!-- Student answer and point selector -->
            <div class="two-column-container">

                <!-- Students answer -->
                <div class="single-column-container">
                    <AnswerBox class="single-column-item" 
                               :disabled="true" 
                               :question="studentsAnswer(question.id)"/>
                </div>

                <!-- Point distribution table -->
                <div class="single-column-container">
                    <PointTable class="single-column-item" :question="question"/>
                    <CommentBox class="single-column-item" :question="question"/>
                </div>
            </div>
        </div>
    </div>
    <button @click="onSubmit">Finalize Grades</button>
</template>

<script>
import PointTable from '../components/PointTable';
import QuestionDescription from '../components/QuestionDescription';
import AnswerBox from '../components/AnswerBox';
import CommentBox from '../components/CommentBox';

export default {
    name: 'ReviewExam',
    components: {
        PointTable,
        QuestionDescription,
        AnswerBox,
        CommentBox
    },
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

</style>


