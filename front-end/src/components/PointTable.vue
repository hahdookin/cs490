<template>
     <!--Instructors point column table -->
    <div>
        <table>
            <!-- Namecorrect -->
            <thead>
                <th colspan="2">Correct Name?</th>
                <td :style="{ backgroundColor: passColor(question.namecorrect) }" colspan="1"></td>
                <td>{{ question.namecorrectpoints }} / 1</td>
                <td v-if="!disabled">
                <input type="text" size="4" v-model="question.override">
                </td>
            </thead>

            <!-- Constraint if applicable -->
            <thead v-if="question.constraint !== 'none'">
                <th colspan="2">Constraint met?</th>
                <td :style="{ backgroundColor: passColor(question.constraintmet) }"></td>
                <td></td>
                <td v-if="!disabled">
                <input type="text" size="4" v-model="question.constraint_override">
                </td>
            </thead>

            <!--Labels for tests -->
            <thead>
                <th>Expected</th>
                <th>Run</th>
                <th>Pass?</th>
                <th>Points</th>
                <th v-if="!disabled">Override</th>
            </thead>

            <!-- Tests and points input -->
            <tbody v-for="test in question.tests">
                <td v-html="testStr(test, question.functionname)"></td>
                <td>{{ test.studentoutput }}</td>
                <td :style="{ backgroundColor: passColor(test.pass) }"></td>
                <td>{{ test.points }} / {{ test.maxpoints }}</td>
                <td v-if="!disabled">
                <input type="text" size="4" v-model="test.override">
                </td>
            </tbody>

            <!-- Totals -->
            <tfoot>
                <tr>
                    <th colspan="3">Total Points: </th>
                    <td>{{ totalPoints(question) }} / {{ question.points }}</td>
                </tr>
            </tfoot>
        </table>
    </div>
</template>

<script>
export default {
    name: 'PointTable',
    inject: ['zip', 'formatArg'],
    props: {
        question: Object,
        disabled: {
            type: Boolean,
            default: false
        }
    },
    created() {
        if (this.question.override === null || this.question.override === undefined)
            this.question.override = '';
        if (this.question.constraint_override === null || this.question.constraint_override === undefined)
            this.question.constraint_override = '';
        for (const test of this.question.tests)
            if (test.override === null || test.override === undefined)
                test.override = '';
    },
    methods: {
        testStr(test, fname) {
            const fmtArgs = test.arguments.map(arg => this.formatArg(arg));
            return `${fname}(${fmtArgs.join(',')}) &#8594; ${test.output}`;
        },
        passColor(bool) {
            return bool ? 'green' : 'red';
        },
        totalPoints(question) {
            let nc_override = question.override.trim();
            let total = nc_override === '' ? question.namecorrectpoints : Number(nc_override);
                             
            for (const test of question.tests) {
                let test_override = test.override.trim();
                total += test_override === '' ? test.points : Number(test_override);
            };
            return isNaN(total) ? '-' : total;
        }
    }
}
</script>

<style scoped>
    
.comment-box {
    background-color: white;
    padding: 10px;
}

table {
    background-color: white;
    width: 100%;
    border: 1px solid black;
}
th, tr, td {
    border: 1px solid black;
}
</style>
