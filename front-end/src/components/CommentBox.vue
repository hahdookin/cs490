<template>
<div class="comment-box">
    <p>Comment:</p>
    <textarea :disabled="disabled" v-model="question.comment"></textarea>
</div>
</template>

<script>
export default {
    name: 'CommentBox',
    props: {
        question: Object,
        disabled: {
            type: Boolean,
            default: false
        }
    },
}
</script>

<style>
.comment-box {
    background-color: white;
    padding: 10px;
}
</style>
