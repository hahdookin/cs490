<template>
    <h3>Assign Exams</h3>
    <form @submit.prevent="onSubmit">
        <input v-model="username" type="text" required>
        <div>
            <select v-model="examID" name="exam">
                <option value="">--Select an exam--</option>
                <option v-for="exam in exams" :key="exam.id" :value="exam.id">
                {{ exam.name }}
                </option>
            </select>
        </div>
        <input type="submit" value="Assign Exam">
    </form>
    <p :style="{ color: success ? 'green' : 'red' }">{{ errorMessage }}</p>
    {{ username }} {{ examID }}
</template>

<script>
export default {
    name: 'AssignExam',
    inject: ['fetchExams', 'fetchTeacher', 'fetchStudentExams'],
    data() {
        return {
            errorMessage: '',
            success: false,
            exams: [],
            // Assign exam with examID to username
            username: '',
            examID: '',
        };
    },
    methods: {
        setError(msg) {
            this.success = false;
            this.errorMessage = msg;
        },
        setSuccess(msg) {
            this.success = true;
            this.errorMessage = msg;
        },
        async onSubmit() {
            if (this.examID === '') {
                this.setError('Please select an exam');
                return;
            }

            // Check if username is a student
            // ... TODO

            const teacherUser = this.$route.params.username;

            // Post to teacherexams
            const payload = {
                assigner: teacherUser,
                assignee: this.username,
                examid: this.examID
            }
            console.log(payload);
            const res = await fetch('http://localhost:5000/teacherexams', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(payload)
            });

            // Add to students incompleted exams
            const studentExams = await this.fetchStudentExams(this.username);
            studentExams.incompleted.push(this.examID);
            console.log(studentExams);
            const res2 = await fetch('http://localhost:5000/teacherexams', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(payload)
            });

            // Get exam name
            const name = this.exams.find(e => e.id === this.examID).name;
            // Assign examID to username
            this.setSuccess(`Successfully assigned '${name}' to ${this.username}`);
        }
    },
    async created() {
        this.exams = await this.fetchExams();
    }
}
</script>
