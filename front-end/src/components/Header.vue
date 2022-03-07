<template>
    <header>
        <h1 @click="$router.push('/')">{{ title }}</h1>
        <Login :userInfo="userInfo" 
               @user-logged-in="onUserLogin"
               @user-logged-out="onUserLogout"/>
    </header>
    <hr>
</template>

<script>
import Login from "../components/Login.vue";

export default {
    name: 'Header',
    emits: ['user-logged-in', 'user-logged-out'],
    components: {
        Login
    },
    props: {
        title: String,
        userInfo: Object
    },
    methods: {
        onUserLogin(info) {
            this.$emit('user-logged-in', info);
        },
        onUserLogout() {
            this.$emit('user-logged-out');
        }
    }
}
</script>

<style scoped>
footer {
    display: flex;
    justify-content: space-around;
    align-items: center;
}
</style>
