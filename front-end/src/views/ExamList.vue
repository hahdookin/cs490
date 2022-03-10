<template>
    <h3>Exams:</h3>
    <div v-if="noexams">
        <p>No exams to show!</p>
    </div>
    <!-- Two column display of exams -->
    <div v-else class="columns-container">

        <!-- Completed exams -->
        <div class="exams-container">
            <div>
                <h3 class="column-title">COMPLETED</h3>
            </div>
            <div v-for="exam in completedExams" class="exam-item">
                <div v-if="exam.reviewed">
                    <p><router-link :to="{ path: `/student/${$route.params.userid}/viewgraded/${exam.studentexamresultid}` }">{{ exam.name }}</router-link></p>
                    <p>View Exam</p>
                </div>
                <div v-else>
                    <p>{{ exam.name }}</p>
                    <p>Not yet graded</p>
                </div>
            </div>
        </div>

        <!-- Incompleted exams -->
        <div class="exams-container">
            <div>
                <h3 class="column-title">INCOMPLETED</h3>
            </div>
            <div v-for="exam in incompletedExams" class="exam-item">
                <p><router-link @click="emitExamStarted" :to="{ path: `/student/${$route.params.userid}/exam/${exam.id}`}">{{ exam.name }}</router-link></p>
            </div>
        </div>
        
    </div>

</template>

<script>
export default {
    name: 'ExamList',
    inject: [
        'fetchExam', 
        'fetchStudentExams',
        'fetchStudentExamResult'
    ],
    emits: ['student-exam-active'],
    data() {
        return {
            student: {
                username: '',
            },
            studentExams: {
                incompleted: [],
                completed: []
            },
            noexams: false,
        };
    },
    methods: {
        // TODO: Emit up that the exam started, hide footer and header.
        emitExamStarted() {
            this.$emit('student-exam-active');
        },
    },
    computed: {
        completedExams() {
            return this.studentExams.completed;
        },
        incompletedExams() {
            return this.studentExams.incompleted;
        }
    },
    async created() {
        const studentUserID = this.$route.params.userid;

        // Fetch student information
        const studentExams = await this.fetchStudentExams(studentUserID);
        if (!studentExams || Object.keys(studentExams).length === 0) {
            this.noexams = true;
            return;
        }
        this.studentExams = studentExams;

        // Remember the exam IDs
        const completed = studentExams.completed;
        const incompleted = studentExams.incompleted;

        // Fetch extra exam info like name
        const completedInfo = [];
        const incompletedInfo = [];
        for (const id of completed) {
            const examInfo = await this.fetchExam(id);
            // Fetch if exam has been reviewed and autograded
            const studentExamResult = await this.fetchStudentExamResult(studentUserID, id);
            examInfo.studentexamresultid = studentExamResult.id;
            examInfo.reviewed = studentExamResult.reviewed;
            completedInfo.push(examInfo);
        }
        for (const id of incompleted) {
            const examInfo = await this.fetchExam(id);
            incompletedInfo.push(examInfo);
        }
        this.studentExams.completed = completedInfo;
        this.studentExams.incompleted = incompletedInfo;
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

.exams-container {
    background-color: #344;
    padding: 1rem;
    margin-top: 1rem;
    flex: 1;
}

.exam-item {
    padding: 1rem;
    background-color: white;
    border: 1px solid blue;
}

.columns-container {
    display: flex;
    justify-content: safe center;
    width: 75%;
    margin: 0 auto;
    max-width: 800px;
}
</style>
