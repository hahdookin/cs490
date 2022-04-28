<template>
    <div @dragstart="onDragStart($event)" 
         @dragend="onDragEnd($event)" 
         :draggable="draggable">

        <div>
            <h3 style="display: inline">{{ question.title }}</h3>
            <h3 style="display: inline"> 
                [<span :style="{ color: difficultyColor(question.difficulty) }">{{ question.difficulty[0].toUpperCase() }}</span>]
            </h3>
        </div>

        <p>{{ question.desc }}</p>

        <p v-if="question.constraint !== 'none'">
        <strong>Constraint:</strong> {{ constraintFmt(question.constraint) }}
        </p>

        <!--<div class="test-cases">-->
            <!--<code v-for="test in question.tests" v-html="highlight(question.functionname, test)"></code>-->
        <!--</div>-->

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
        showPointsInput: {
            type: Boolean,
            default: false,
        },
        draggable: {
            type: Boolean,
            default: false
        }
    },
    methods: {
        // Map the stored constants to nice text
        constraintFmt(constraint) {
            switch (constraint.toLowerCase()) {
                case 'forloop':
                    return 'For Loop';
                case 'whileloop':
                    return 'While Loop';
                case 'recursion':
                    return 'Recursion';
                default:
                    return constraint;
            }
        }, 
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

.tooltip {
    position: relative;
    display: inline-block;
    border-bottom: 1px dotted black;
}
</style>
