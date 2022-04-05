<template>
<textarea :disabled="disabled"
          @keydown="onKeydown($event)"
          @input="handleInput($event)"
          rows="20" 
          cols="60"
          autocapitalize="off" 
          spellcheck="false" 
          contenteditable="true">{{ question.code }}</textarea>
</template>

<script>
export default {
    name: 'AnswerBox',
    emits: ['student-input'],
    props: {
        question: Object,
        content: String,
        disabled: {
            type: Boolean,
            default: false
        }
    },
    methods: {
        // When a student inputs, emit the current value of
        // the answerbox to the parent component, along with
        // the associated question. (MIMIC OF V-MODEL)
        handleInput(e) {
            this.$emit('student-input', e.target.value, this.question);
        },

        // Allow Tab to function as it would in a text editor
        onKeydown(e) {
            if (e.key == 'Tab') {
                const tar = e.target;
                e.preventDefault();
                var start = tar.selectionStart;
                var end = tar.selectionEnd;

                // set textarea value to: text before caret + tab + text after caret
                tar.value = tar.value.substring(0, start) +
                "\t" + tar.value.substring(end);

                // put caret at right position again
                tar.selectionStart = tar.selectionEnd = start + 1;
            }
        }
    }
}
</script>
