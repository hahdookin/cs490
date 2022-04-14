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
            userid: this.$route.params.userid
        };
    },
    computed: {
        examCreatorPath() {
            return `/teacher/${this.userid}/examcreator`;
        },
        questionCreatorPath() {
            return `/teacher/${this.userid}/questioncreator`;
        },
        assignExamPath() {
            return `/teacher/${this.userid}/assignexam`;
        },
        viewExamsPath() {
            return `/teacher/${this.userid}/viewexams`;
        },
    },
}
</script>

<style>
ul {
    list-style-type: none;
}
</style>
