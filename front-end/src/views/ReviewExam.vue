<template>
    <h1>Review Exam</h1>
    <div v-if="loaded">
        <h1>{{ exam.name }} (ID: {{ exam.id }})</h1>

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
                    <PointTable class="single-column-item" :question="question"/>
                    <CommentBox class="single-column-item" :question="question"/>
                </div>
            </div>
        </div>
        </div>

    </div>

    <!-- Navigation buttons -->
    <div v-if="questions.length > 1">
        <button @click="onPrev" :disabled="navLeftDisabled">Previous</button>
        <button @click="onNext" :disabled="navRightDisabled">Next</button>
    </div>

    <button @click="onSubmit">Finalize Grades</button>
    <p :style="{ color: success ? 'green' : 'red' }">{{ errorMessage }}</p>
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

            currentQuestionIndex: 0,

            errorMessage: '',
            success: false,
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
        setError(msg) {
            this.success = false;
            this.errorMessage = msg;
            return false;
        },
        setSuccess(msg) {
            this.success = true;
            this.errorMessage = msg;
        },
        validateOverrides() {
            for (const question of this.questions) {
                // Validate namecorrect input
                let nc_override = question.override.trim();
                if (nc_override !== '') {
                    if (!nc_override.match(/[0-9]+/g))
                        return this.setError('Illegal format in override input');
                    // TODO: Check if input is in range
                }
                let constraint_override = question.constraint_override.trim();
                if (constraint_override !== '') {
                    if (!constraint_override.match(/[0-9]+/g))
                        return this.setError('Illegal format in override input');
                    // TODO: Check if input is in range
                }

                // Validate all tests input
                for (const test of question.tests) {
                    let test_override = test.override.trim();
                    if (test_override !== '') {
                        if (!test_override.match(/[0-9]+/g))
                            return this.setError('Illegal format in override input');
                        // TODO: Check if input is in range
                    }
                }
            }
            return true;
        },
        async onSubmit() {
            // Get total points awarded
            let totalPoints = 0;
            this.questions.forEach(q => {
                q.tests.forEach(t => {
                    totalPoints += t.points;
                });
                totalPoints += q.namecorrectpoints;
            });

            // Validate that override inputs are OK
            if (!this.validateOverrides())
                return;

            // If override is not empty, set as the current points
            for (const question of this.questions) {
                let nc_override = question.override.trim();
                if (nc_override !== '')
                    question.namecorrectpoints = Number(nc_override);
                for (const test of question.tests) {
                    let test_override = test.override.trim();
                    if (test_override !== '')
                        test.points = Number(test_override)
                }
            }

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
            for (const sea of this.studentExamAnswers)
                await this.putStudentExamAnswer(sea);

            // Put the studentexamresult
            await this.putStudentExamResult(this.studentExamResult);

            // Redirect back to home
            this.$router.push(`/teacher/${this.userid}/`);
        },
        studentsAnswer(qid) {
            return this.studentExamAnswers.find(
                a => a.questionid === qid && 
                     a.studentexamresultid === this.studentExamResultID
            );
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
            question.override = '';
            question.constraint_override = '';
            for (const [qTest, sTest] of this.zip(question.tests, studentsAnswer.tests)) {
                qTest.points = sTest.points;
                qTest.pass = sTest.pass;
                qTest.studentoutput = sTest.studentoutput;
                
                qTest.override = '';
            }
            const testsCount = question.tests.length;
            const pointDist = this.split(question.points - (question.points === 0 ? 0 : 1), testsCount);
            for (const [qTest, maxPoints] of this.zip(question.tests, pointDist))
                qTest.maxpoints = maxPoints;

            // Attach extra info to the question object (used in point table)
            question.runs = studentsAnswer.runs;
            question.namecorrect = studentsAnswer.namecorrect;
            question.namecorrectpoints = studentsAnswer.namecorrectpoints;
            question.constraintdeduction;
            question.comment = studentsAnswer.comment;
            question.constraintmet = studentsAnswer.constraintmet;
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


