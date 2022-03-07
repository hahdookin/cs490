<template>
    <Header :userInfo="userInfo" 
            @user-logged-in="handleUserLogin" 
            @user-logged-out="handleUserLogout"
            title="Exam Central"/>

    <div v-if="userIsLoggedIn">
        <router-view></router-view>
        <Footer :userInfo="userInfo" />
    </div>
</template>

<script>
import Header from "./components/Header";
import Footer from "./components/Footer";

export default {
    name: 'App',
    components: {
        Header,
        Footer,
    },
    data() { 
        return {
            userInfo: {
                username: '',
                type: '',
            },
            examActive: false
        }; 
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
        }
    },
    methods: {
        handleUserLogin(info) {
            this.userInfo = info;
            this.$router.push(`/${this.userInfo.type}/${this.userInfo.username}`);
        },
        handleUserLogout() {
            this.userInfo.username = '';
            this.userInfo.type = '';
            this.$router.push('/');
        },
    },
    created() {},
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
