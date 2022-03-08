<template>
    <h1>Student Dashboard</h1>
    <ul>
    <li><router-link :to="{ path: `/student/${$route.params.userid}/examlist` }">View Exams</router-link></li>
    </ul>
    <router-view></router-view>
</template>

<style scoped>
body {
    margin: 0;
}
</style>
