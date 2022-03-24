<template>
    <div @dragstart="onDragStart($event)" 
         @dragend="onDragEnd($event)" 
         draggable="true">

        <h3>{{ question.title }}</h3>
        <h5 :style="{ color: difficultyColor(question.difficulty) }">{{ question.difficulty }}</h5>
        <p>{{ question.desc }}</p>
        <div class="test-cases">
            <code v-for="test in question.tests" v-html="highlight(question.functionname, test)"></code>
        </div>

        <div v-if="showPointsInput">
            <label>Points: </label>
            <input v-model="question.points" type="number" min="0" size="4">
        </div>

    </div>
</template>

<script>
export default {
    name: 'ExamCreatorItem',
    props: {
        question: Object,
        showPointsInput: Boolean
    },
    methods: {
        onDragStart(e) {
            setTimeout(() => e.target.classList.add('dragging'), 0);
        },
        onDragEnd(e) {
            e.target.classList.remove('dragging');
        },
        difficultyColor(difficulty) {
            switch (difficulty) {
                case 'easy':
                    return 'green';
                case 'medium':
                    return 'orange';
                case 'hard':
                    return 'red';
                default:
                    return 'black';
            }
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
.dragging {
    opacity: 0.1;
}

code {
    display: block;
    padding: 5px;
}

.test-cases {
    border: 1px solid #aaa;
    border-radius: 10px;
    background-color: var(--hl-bg);
    color: var(--hl-fg);
    justify-content: safe center;
    width: 80%;
    margin: 0 auto;
}
</style>
