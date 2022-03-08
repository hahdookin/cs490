<template>
    <h3>View Exams</h3>
    <div v-for="exam in teacherExams" class="exam-item">
        <p>Student: {{ exam.assignee }}</p>
        <p>Exam: {{ exam.exam?.name }}</p>
        <div>
            <p>Status: {{ exam.completed ? 'Completed' : 'Not yet completed' }}</p>
            <button @click="postAutoGrade" 
                    v-if="exam.completed"
                    :disabled="exam.autograded">AutoGrade</button>
        </div>
    </div>
</template>

<script>
export default {
    name: 'ViewExams',
    inject: [
        'fetchUser',
        'fetchTeacherExams', 
        'fetchExam',
        'fetchStudentExams',
        'fetchStudentExamResult',
    ],
    data() {
        return {
            teacherExams: [], // { assigner, assignee, id, examid }
        };
    },
    async created() {
        const teacherUserID = +this.$route.params.userid;

        // Get all exams the teacher has assigned
        this.teacherExams = await this.fetchTeacherExams(teacherUserID);

        for (const teacherExam of this.teacherExams) {
            const studentUserID = +teacherExam.assigneeid;
            const studentUser = await this.fetchUser(studentUserID);
            // Store user's username
            teacherExam.assignee = studentUser.username;

            // Find out if the exam is completed
            let studentCompleted = await this.fetchStudentExams(teacherExam.assigneeid);
            studentCompleted = studentCompleted.completed;
            teacherExam.completed = studentCompleted.indexOf(teacherExam.examid) !== -1;

            if (teacherExam.completed) {
                // Find out if the exam has been autograded
                let studentExamResult = await this.fetchStudentExamResult(studentUserID, teacherExam.examid);
                teacherExam.autograded = studentExamResult.autograded;
            }

            // Attach exam information to those assigned exams (name, questions)
            teacherExam.exam = await this.fetchExam(teacherExam.examid);
        }
    },
    methods: {
        postAutoGrade() {
            // make post with { qID and the code } to middle

            // wait for { }
        }
    }
}
</script>

<style scoped>
.exam-item {
    border: 1px solid black;
    margin: 5px;
}
</style>
