<template>
    <hr>
    <footer>
        <div v-if="userIsLoggedIn">
            <h3>Navigate</h3>
            <router-link :to="userDashboard">Dashboard</router-link><br>
        </div>
    </footer>
</template>

<script>
export default {
    name: 'Footer',
    components: {},
    props: {
        userInfo: Object,
    },
    computed: {
        userIsLoggedIn() {
            return this.userInfo.username !== '';
        },
        userIsTeacher() {
            return this.userInfo.type.toLowerCase() === 'teacher';
        },
        userIsStudent() {
            return this.userInfo.type.toLowerCase() === 'student';
        },
        userDashboard() {
            const { userid, type } = this.userInfo;
            return '/' + type + '/' + userid;
        }
    }
}
</script>

<style scoped>
footer {
    /* margin-top: 30px; */
    margin-bottom: 30px;
    text-align: center;
}
</style>
