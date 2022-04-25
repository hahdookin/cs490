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
                <button @click="proceedAutoGrade(exam)" 
                        :disabled="exam.autograded">{{ autogradeStatus }}</button>
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
        'zip',
        'postToAutograder',
        'putStudentExamAnswer',
        'putStudentExamResult',
        'fetchUser',
        'fetchTeacherExams', 
        'fetchExam',
        'fetchStudentExams',
        'fetchStudentExamResult',
        'fetchStudentExamAnswers',
    ],
    data() {
        return {
            teacherExams: [], // { assigner, assignee, id, examid }
            noexams: false,
            autogradeInProgress: false
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
        /*
        {namecorrect: bool, output: [], pass: []}
        */
        async proceedAutoGrade(exam) {
            this.autogradeInProgress = true;
            const studentUserID = exam.assigneeid;
            const ser = await this.fetchStudentExamResult(studentUserID, exam.examid);
            if (!ser) {
                console.log("error fetching student exam result", studentUserID, exam.id);
                return;
            }
            const sea = await this.fetchStudentExamAnswers(ser.id);
            if (!sea) {
                console.log("error fetching student exam answers", ser.id);
                return;
            }
            for (const answer of sea) {
                const res = await this.postToAutograder(answer.questionid, answer.code);
                answer.runs = res.runs;
                answer.namecorrect = res.namecorrect;
                answer.namecorrectpoints = +res.namecorrect; // 1 if true, 0 otherwise
                answer.constraintmet = res.constraint;
                answer.constraintmetpoints = +res.constraint * 5; // 5 points
                for (let i = 0; i < answer.tests.length; i++) {
                    answer.tests[i].pass = res.pass[i];
                    answer.tests[i].studentoutput = res.output[i];
                    answer.tests[i].points *= +res.pass[i];
                }
                const res2 = await this.putStudentExamAnswer(answer);
            }
            ser.autograded = true;
            exam.autograded = true;

            const res = await this.putStudentExamResult(ser);

            this.autogradeInProgress = false;
        },
        reviewAndSubmitExam(studentExamResult) {
            const teacherUserID = this.$route.params.userid;
            this.$router.push(`/teacher/${teacherUserID}/reviewexam/${studentExamResult.id}`);
        }
    },
    computed: {
        autogradeStatus() {
            return this.autogradeInProgress ? 'In Progress' : 'AutoGrade';
        }
    }
}
</script>

<style scoped>
.exam-item {
    border: 1px solid black;
    margin: 5px auto;
    padding: 5px;
    width: 30%;
}
</style>
