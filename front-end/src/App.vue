<template>
    <Header :user-info="userInfo" 
            @user-logged-in="handleUserLogin" 
            @user-logged-out="handleUserLogout"
            title="Exam Central"/>

    <!--<button @click="janeLogin">Jane</button>-->
    <!--<button @click="johnLogin">John</button>-->

    <div v-if="userIsLoggedIn">
        <router-view></router-view>
        <Footer :user-info="userInfo" />
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
    inject: ['serialize'],
    data() { 
        return {
            userInfo: {
                username: '',
                type: '',
                userid: -1,
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
        janeLogin() {
            this.handleUserLogin({
                username: 'jane',
                type: 'teacher',
                userid: 1
            });
        },
        johnLogin() {
            this.handleUserLogin({
                username: 'john',
                type: 'student',
                userid: 2
            });
        },
        handleUserLogin(info) {
            this.userInfo = info;
            this.$router.push(`/${this.userInfo.type}/${this.userInfo.userid}`);
        },
        handleUserLogout() {
            this.userInfo.username = '';
            this.userInfo.type = '';
            this.userInfo.userid = -1;
            this.$router.push('/');
        }
    },
    created() {},
}
</script>

<style>
/**
 * Variables
 */
:root {
    --hl-bg: #1E1E1E;
    --hl-fg: #D4D4D4;
    --hl-function: #D7BA7D;
    --hl-string: #CE9178;
    --hl-number: #B5CEA8;
    --hl-keyword: #C586C0;
    /* --hl-function: #444; /1* color07 *1/ */
    /* --hl-string: #5f8700; /1* color03 *1/ */
    /* --hl-number: #d75f00; /1* color13 *1/ */
    /* --hl-keyword: #8700af; /1* color11 *1/ */
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

/**
 * Styling for two-column displays
 */
/* Two column */
.two-column-container {
    display: flex;
    justify-content: safe center;
    width: 75%;
    margin: 0 auto;
    max-width: 800px;
}
/* Single column */
.single-column-container {
    background-color: #344;
    padding: 1rem;
    margin-top: 1rem;
    flex: 1;
}

/* Single item within column */
.single-column-item {
    padding: 1rem;
    background-color: white;
    border: 1px solid blue;
}

/* Input elements are often reused and disabled. */
/* Although this introduces unintuitive UX, */
/* leave disabled input text black so its */
/* easier to read. */
textarea:disabled {
    color: black;
}

/* Minimap */
.minimap {
    border: 1px solid black;
    width: 300px;
    /* position: absolute; */
    /* position: sticky; */
    margin-right: 15%;
    margin-left: auto;
}
.minimap-inner {
    text-align: left;
}

</style>
