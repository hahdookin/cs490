<template>
    <h3>View Exams</h3>
    <div v-for="exam in teacherExams">
        <p>Student: {{ exam.assignee }} -> {{ exam.exam?.name }}</p>
        <button v-if="exam.completed">AutoGrade</button>
        <p v-else>Not yet completed</p>
    </div>
</template>

<script>
export default {
    name: 'ViewExams',
    inject: [
        'fetchTeacher', 
        'fetchTeacherExams', 
        'fetchExam',
        'fetchStudent'
    ],
    data() {
        return {
            teacherExams: [], // { assigner, assignee, id, examid }
        };
    },
    async created() {
        const teacherUser = this.$route.params.username;

        // Get all exams the teacher has assigned
        this.teacherExams = await this.fetchTeacherExams(teacherUser);

        for (const teacherExam of this.teacherExams) {
            // Find out if the exam is completed
            let studentCompleted = await this.fetchStudent(teacherExam.assignee);
            studentCompleted = studentCompleted.exams.completed;
            teacherExam.completed = studentCompleted.indexOf(teacherExam.examid) !== -1;
            /* console.log(studentCompleted); */

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
