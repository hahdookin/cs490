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
                <p><router-link @click="emitExamStarted" :to="{ path: `/student/${$route.params.username}/exam/${exam.id}`}">{{ exam.name }}</router-link></p>
            </div>
        </div>
        
    </div>

</template>

<script>
export default {
    name: 'ExamList',
    inject: ['fetchStudent', 'fetchExam'],
    emits: ['student-exam-active'],
    data() {
        return {
            student: {
                username: '',
                exams: {
                    completed: [],
                    incompleted: [],
                }
            },
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
            //return this.exams.completed;
            return this.student.exams.completed;
        },
        incompletedExams() {
            return this.student.exams.incompleted;
        }
    },
    async created() {
        // Fetch student information
        const data = await this.fetchStudent('john');
        this.student = data;

        // Remember the exam IDs
        const completed = data.exams.completed;
        const incompleted = data.exams.incompleted;
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
        this.student.exams.completed = completedInfo;
        this.student.exams.incompleted = incompletedInfo;
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
