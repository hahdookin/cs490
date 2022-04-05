<template>
    <div class="wrapper">
        <h3>{{ number }}) {{ question.title }}</h3>
        <h5>Points: {{ question.points }}</h5>
        <p>{{ question.desc }}</p>
        <p>Examples:</p>
        <div class="test-cases">
            <code v-for="test in question.tests" v-html="highlight(question.functionname, test)"></code>
        </div>
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
        },
        highlight(fname, test) {
            let hlargs = [];
            for (const arg of test.arguments)
                hlargs.push(this.highlightToken(arg));
            let res = `<span style="color: var(--hl-function)">${fname}</span>`;
            res += `(${hlargs.join(', ')})`;
            res += ' &#8594; ';
            res += this.highlightToken(test.output);
            return res;
        },
        highlightClass(s) {
            if (s.startsWith('"') && s.endsWith('"'))
                return 'string';
            return isNaN(parseInt(s)) ? 'string' : 'number';
        },
        highlightToken(t) {
            const hlclass = this.highlightClass(t);
            let v;
            switch (hlclass) {
                case 'string':
                    v = '--hl-string';
                    break;
                case 'number':
                    v = '--hl-number';
                    break;
            }
            return `<span style="color: var(${v})">${t}</span>`;
        }
    }
}
</script>

<style scoped>
code {
    display: block;
    padding: 5px;
}

.wrapper {
    border: 1px solid #aaa;
    border-radius: 10px;
    margin: 0 auto;
    margin-top: 5px;
    margin-bottom: 5px;
    padding: 0 5px 0;
    width: fit-content;
}

.test-cases {
    border: 1px solid #aaa;
    border-radius: 10px;
    background-color: var(--hl-bg);
    color: var(--hl-fg);
    justify-content: safe center;
    width: fit-content;
    padding-left: 5px;
    padding-right: 5px;
    margin: 0 auto;
    margin-bottom: 25px;
}
</style>
