import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import database from './database';

const app = createApp(App);

app.use(router);
app.use(database);

app.mount('#app');

//createApp(App).use(router).mount('#app')
