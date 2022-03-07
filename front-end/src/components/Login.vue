<template>
    <div v-if="userInfo.username === ''">
        <form @submit="onSubmit" method="post">
            Username: <input v-model="username" type="text" name="username" required><br>
            Password: <input v-model="password" type="text" name="password" required><br>
            <input type="submit" value="Login">
            <p style="color:red">{{ errorMessage }}</p>
        </form>
    </div>
    <div v-else>
        <h3>Welcome {{ userInfo.username }}, {{ userInfo.type }}</h3>
        <button @click="onLogout">Logout</button>
    </div>
</template>

<script>
export default {
    name: 'Login',
    inject: ['fetchLoginCredentials'],
    components: {},
    props: {
        userInfo: Object,
    },
    emits: ['user-logged-in', 'user-logged-out'],
    data() {
        return {
            username: '',
            password: '',
            errorMessage: '',
        }
    },
    methods: {
        async onSubmit(e) {
            e.preventDefault();
            // Post username and password to database
            const credientials = await this.fetchLoginCredentials(this.username, this.password);

            // On bad credientials
            if (credientials.length === 0) {
                this.errorMessage = 'Incorrect username or password';
                return;
            }
            // Successful, clear error msg and emit up to app
            this.errorMessage = '';
            this.$emit('user-logged-in', {
                username: this.username,
                type: credientials[0].type,
                //type: 'teacher'
            });
        },
        onLogout() {
            this.$emit('user-logged-out');
        }
    }
}
</script>
