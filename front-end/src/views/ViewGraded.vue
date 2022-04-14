<template>
    <h1>View Graded</h1>
    <div v-if="loaded">
        <h1>{{ exam.name }} (ID: {{ exam.id }})</h1>
        <h1>Score: {{ earnedPoints }}/{{ totalPoints }}</h1>

        <!-- Minimap for questions -->
        <div class="minimap">
            <h3>Questions: </h3>
            <ol class="minimap-inner">
                <li :key="question.id" 
                    v-for="(question, i) in questions">
                    <a href="" 
                       v-if="currentQuestionIndex !== i"
                       @click.prevent="currentQuestionIndex = i">
                       {{ question.title }}
                    </a>
                    <span v-else>{{ question.title }}</span>
                    ({{ question.points }})
                </li>
            </ol>
        </div>

        <div :key="question.id" v-for="(question, i) in questions">
            <div v-if="i == currentQuestionIndex">
                <!-- Question info -->
                <QuestionDescription :question="question" :number="i + 1"/>

                <!-- Student answer and point selector -->
                <div class="two-column-container">
                    <!-- Students answer -->
                    <div class="single-column-container">
                        <AnswerBox class="single-column-item" 
                                   disabled 
                                   :question="studentsAnswer(question.id)"/>
                    </div>

                    <!-- Point distribution table -->
                    <div class="single-column-container">
                        <PointTable disabled class="single-column-item" :question="question"/>
                        <CommentBox disabled class="single-column-item" :question="question"/>
                    </div>
                </div>
            </div>
        </div>

        <!-- Navigation buttons -->
        <div v-if="questions.length > 1">
            <button @click="onPrev" :disabled="navLeftDisabled">Previous</button>
            <button @click="onNext" :disabled="navRightDisabled">Next</button>
        </div>

    </div>
</template>

<script>
import PointTable from '../components/PointTable';
import AnswerBox from '../components/AnswerBox';
import QuestionDescription from '../components/QuestionDescription';
import CommentBox from '../components/CommentBox';

export default {
    name: 'ReviewExam',
    components: {
        PointTable,
        AnswerBox,
        QuestionDescription,
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
    ],
    data() {
        return {
            studentUserID: this.$route.params.userid,
            studentExamResultID: this.$route.params.studentexamresultid,
            studentExamResult: {},
            studentExamAnswers: [],
            exam: {},
            questions: [],
            loaded: false,
            totalPoints: 0,
            earnedPoints: 0,

            currentQuestionIndex: 0,
        };
    },
    computed: {
        navLeftDisabled() {
            return this.currentQuestionIndex === 0;
        },
        navRightDisabled() {
            return this.currentQuestionIndex === this.questions.length - 1
        },
    },
    methods: {
        onNext() {
            this.currentQuestionIndex++;
        },
        onPrev() {
            this.currentQuestionIndex--;
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
                qTest.studentoutput = sTest.studentoutput;
            }

            // Attach max points for questions
            const testsCount = question.tests.length;
            const pointDist = this.split(question.points - (question.points === 0 ? 0 : 1), testsCount);
            for (const [qTest, maxPoints] of this.zip(question.tests, pointDist))
                qTest.maxpoints = maxPoints;

            question.runs = studentsAnswer.runs;
            question.namecorrect = studentsAnswer.namecorrect;
            question.namecorrectpoints = studentsAnswer.namecorrectpoints;
            question.constraintmet = studentsAnswer.constraintmet;
            question.comment = studentsAnswer.comment;
        }

        // Get total points
        // Get earned points
        this.questions.forEach(q => {
            this.totalPoints += q.points;
            q.tests.forEach(t => {
                // Add points earned from each test case
                this.earnedPoints += t.points;
            });
            this.earnedPoints += q.namecorrectpoints;
        });

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


