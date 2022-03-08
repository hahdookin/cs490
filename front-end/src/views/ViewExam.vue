<template>
    <h3>View Exams</h3>
    <div v-for="exam in teacherExams" class="exam-item">
        <p>Student: {{ exam.assignee }}</p>
        <p>Exam: {{ exam.exam?.name }}</p>
        <div>
            <p>Status: {{ exam.completed ? 'Completed' : 'Not yet completed' }}</p>
            <div v-if="exam.completed">
                <button @click="postAutoGrade(exam)" 
                        :disabled="exam.autograded">AutoGrade</button>
                <button @click="reviewAndSubmitExam(exam.studentExamResult)"
                        :disabled="!exam.autograded">Review and Submit</button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'ViewExams',
    inject: [
        'serialize',
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

            // If completed, query the studentExamResult
            if (teacherExam.completed) {
                // Find out if the exam has been autograded
                let studentExamResult = await this.fetchStudentExamResult(studentUserID, teacherExam.examid);
                teacherExam.autograded = studentExamResult.autograded;
                teacherExam.studentExamResult = studentExamResult;
            }

            // Attach exam information to those assigned exams (name, questions)
            teacherExam.exam = await this.fetchExam(teacherExam.examid);
        }
    },
    methods: {
        async postAutoGrade(exam) {
            console.log("Sending to autograder", exam);

            /* const res = await fetch('http://localhost:5000/') */
            // make post with { qID and the code } to middle
            exam.autograded = true;
        },
        reviewAndSubmitExam(studentExamResult) {
            const teacherUserID = this.$route.params.userid;
            this.$router.push(`/teacher/${teacherUserID}/reviewexam/${studentExamResult.id}`);
            /* const payload = { */
            /*     qid: 2, */
            /*     code: 'def my_fn(x, y):\n\treturn x + y' */
            /* }; */
            /* const res = await fetch('http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com/autograde', { */
            /*     method: 'post', */
            /*     headers: { */
            /*         'Content-type': 'application/x-www-form-urlencoded', */
            /*     }, */
            /*     body: serialize(payload) */
            /* }); */
            /* const resp = await res.json(); */
            /* console.log(resp); */

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
