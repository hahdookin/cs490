<template>
    <h1>Teacher Dashboard</h1>
    <ul>
        <li><router-link :to="{ path: examCreatorPath }">Exam Creator</router-link></li>
        <li><router-link :to="{ path: questionCreatorPath }">Question Creator</router-link></li>
        <li><router-link :to="{ path: assignExamPath }">Assign Exam</router-link></li>
        <li><router-link :to="{ path: viewExamsPath }">View Exams</router-link></li>
    </ul>
    <router-view></router-view>
</template>

<script>
export default {
    name: 'TeacherDashboard',
    data() {
        return {
            username: this.$route.params.username
        };
    },
    computed: {
        examCreatorPath() {
            return `/teacher/${this.username}/examcreator`;
        },
        questionCreatorPath() {
            return `/teacher/${this.username}/questioncreator`;
        },
        assignExamPath() {
            return `/teacher/${this.username}/assignexam`;
        },
        viewExamsPath() {
            return `/teacher/${this.username}/viewexams`;
        },
    },
}
</script>
