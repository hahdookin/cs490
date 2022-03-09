<template>
    <h3>Exams:</h3>
    <!-- Two column display of exams -->
    <div class="columns-container">

        <!-- Incompleted exams -->
        <div class="exams-container">
            <div>
                <h3 class="column-title">COMPLETED</h3>
            </div>
            <div v-for="exam in completedExams" class="exam-item">
                <p><router-link to="">{{ exam.name }}</router-link></p>
            </div>
        </div>

        <!-- Completed exams -->
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
    inject: ['fetchExam', 'fetchStudentExams'],
    emits: ['student-exam-active'],
    data() {
        return {
            student: {
                username: '',
            },
            studentExams: {
                incompleted: [],
                completed: []
            }
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
        this.studentExams = studentExams;

        // Remember the exam IDs
        const completed = studentExams.completed;
        const incompleted = studentExams.incompleted;
        const completedInfo = [];
        const incompletedInfo = [];
        for (const id of completed) {
            const examInfo = await this.fetchExam(id);
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