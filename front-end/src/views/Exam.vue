<template>
    <h1>{{ exam.name }}</h1>
    <h3>Total Points: {{ totalExamPoints(exam) }}</h3>

    <!-- Minimap for questions -->
    <div class="minimap">
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

    <!-- Show current question -->
    <div :key="question.id" v-for="(question, i) in questions">
        <div v-if="i == currentQuestionIndex">
            <!-- Question info --> 
            <QuestionDescription :question="question" :number="i + 1"/>

            <!-- Student answer section -->
            <AnswerBox :question="question" @student-input="handleStudentInput"/>
        </div>
    </div>

    <!-- Navigation buttons -->
    <div v-if="questions.length > 1">
        <button @click="onPrev" :disabled="navLeftDisabled">Previous</button>
        <button @click="onNext" :disabled="navRightDisabled">Next</button>
    </div>

    <!-- Submit exam button -->
    <button @click="onSubmit">Submit Exam</button>

</template>

<script>
import QuestionDescription from '../components/QuestionDescription';
import AnswerBox from '../components/AnswerBox';

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
    components: {
        QuestionDescription,
        AnswerBox
    },
    /* emits: ['student-exam-active'], */
    data() {
        return {
            examID: +this.$route.params.examid,
            questions: [],
            exam: {},
            currentQuestionIndex: 0,
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
        onNext() {
            this.currentQuestionIndex++;
        },
        onPrev() {
            this.currentQuestionIndex--;
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
        handleStudentInput(input, question) {
            question.code = input;
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
    },
    computed: {
        navLeftDisabled() {
            return this.currentQuestionIndex === 0;
        },
        navRightDisabled() {
            return this.currentQuestionIndex === this.questions.length - 1
        },
    }
}
</script>

<style scoped>
.minimap {
    border: 1px solid black;
    width: 300px;
    /* position: absolute; */
    justify-content: right;
}
.minimap-inner {
    text-align: left;
}
</style>
