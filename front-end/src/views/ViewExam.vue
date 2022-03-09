<template>
    <h3>View Exams</h3>
    <div v-if="noexams">
        <p>No exams assigned...</p>
    </div>
    <div v-else v-for="exam in teacherExams" class="exam-item">
        <p>Student: {{ exam.assignee }}</p>
        <p>Exam: {{ exam.exam?.name }}</p>
        <div>
            <p>Status: {{ exam.completed ? 'Completed' : 'Not yet completed' }}</p>
            <div v-if="exam.completed">
                <button @click="postAutoGrade(exam.id)" 
                        :disabled="exam.autograded">AutoGrade</button>
                <button @click="reviewAndSubmitExam(exam.studentExamResult)"
                        :disabled="!exam.autograded && !exam.reviewed">Review and Submit</button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'ViewExams',
    inject: [
        'postToAutograder',
        'fetchUser',
        'fetchTeacherExams', 
        'fetchExam',
        'fetchStudentExams',
        'fetchStudentExamResult',
    ],
    data() {
        return {
            teacherExams: [], // { assigner, assignee, id, examid }
            noexams: false,
        };
    },
    async created() {
        const teacherUserID = +this.$route.params.userid;

        // Get all exams the teacher has assigned
        this.teacherExams = await this.fetchTeacherExams(teacherUserID);
        if (Object.keys(this.teacherExams).length === 0) {
            this.noexams = true;
            return;
        }

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
                teacherExam.reviewed = studentExamResult.reviewed;
                teacherExam.studentExamResult = studentExamResult;
            }

            // Attach exam information to those assigned exams (name, questions)
            teacherExam.exam = await this.fetchExam(teacherExam.examid);
        }
    },
    methods: {
        async postAutoGrade(exam) {
            console.log("Sending to autograder", exam);

            const resp = await this.postToAutograder(20, 'some code here!')
            //exam.autograded = true;
            console.log(resp);
        },
        reviewAndSubmitExam(studentExamResult) {
            const teacherUserID = this.$route.params.userid;
            this.$router.push(`/teacher/${teacherUserID}/reviewexam/${studentExamResult.id}`);
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
