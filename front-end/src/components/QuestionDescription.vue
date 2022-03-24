<template>
    <h3>{{ number }}) {{ question.title }}</h3>
    <h5>Points: {{ question.points }}</h5>
    <p>{{ question.desc }}</p>
    <p>Examples:</p>
    <div>
        <p v-for="test in question.tests">
        <code>
        {{ testStr(test, question.functionname) }}
        </code>
        </p>
    </div>
</template>

<script>
export default {
    name: 'QuestionDescription',
    props: {
        question: Object,
        number: Number
    },
    methods: {
        testStr(test, fname) {
            return `${fname}(${test.arguments.join(',')}) -> ${test.output}`;
        }
    }
}
</script>
