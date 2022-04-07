<template>
     <!--Instructors point column table -->
    <div>
        <table>
            <tr>
                <th colspan="2">Runs?</th>
                <td colspan="1">{{ question.runs ? "Yes" : "No" }}</td>
            </tr>
            <tr>
                <th colspan="2">Correct Name?</th>
                <td :style="{ backgroundColor: passColor(question.namecorrect) }" colspan="1"></td>
                <td>
                <p v-if="disabled"><strong>{{ question.namecorrectpoints }}</strong></p>
                <input v-else type="number" size="4" :max="1" v-model="question.namecorrectpoints" min="0">
                </td>
            </tr>
        </table>
         <!--Test cases and point assignment -->
        <table>
            <tr>
                <th>Expected</th>
                <th>Run</th>
                <th>Pass?</th>
                <th>Points</th>
                <th>Override</th>
            </tr>
            <tr v-for="test in question.tests">
                <td>{{ testStr(test, question.functionname) }}</td>
                <td>{{ test.studentoutput }}</td>
                <td :style="{ backgroundColor: passColor(test.pass) }"></td>
                <td>
                <p v-if="disabled"><strong>{{ test.points }}</strong></p>
                <input v-else type="number" size="4" :max="test.maxpoints" v-model="test.points" min="0">
                </td>
                <td>
                <input type="number" size="4" :max="test.maxpoints" min="0">
                </td>
            </tr>
            <tfoot>
                <tr>
                    <th colspan="3">Total Points: </th>
                    <td>{{ totalPoints(question) }}</td>
                </tr>
            </tfoot>
        </table>
    </div>
</template>

<script>
export default {
    name: 'PointTable',
    props: {
        question: Object,
        disabled: {
            type: Boolean,
            default: false
        }
    },
    methods: {
        testStr(test, fname) {
            return `${fname}(${test.arguments.join(',')}) -> ${test.output}`;
        },
        passColor(bool) {
            return bool ? 'green' : 'red';
        },
        totalPoints(question) {
            let total = question.namecorrectpoints;
            for (const test of question.tests) {
                total += test.points;
            };
            return total;
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
