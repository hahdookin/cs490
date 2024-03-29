<template>
    <h3>Exam Creator</h3>
    <label>Exam Name: </label>
    <input v-model="examName" type="text">

    <p>Total Points: {{ totalPoints }}</p>

    <!-- Bank filter form -->
    <h4>Bank Filter</h4>

    <!--<form @change="onFilterFormChanged" @submit.prevent>-->
    <form @submit.prevent>
    <div class="row" style="padding-bottom: 5px;border: 1px solid black">
        <div class="column">
            <h5>Difficulty</h5>
            <div class="filter-selection">
                <label>Easy: </label>
                <input type="checkbox" v-model="filters.difficulty.easy">
                <br>
                <label>Medium: </label>
                <input type="checkbox" v-model="filters.difficulty.medium">
                <br>
                <label>Hard: </label>
                <input type="checkbox" v-model="filters.difficulty.hard">
                <br>
            </div>
        </div>

        <div class="column">
            <h5>Category</h5>
            <div class="filter-selection">
                <label>For Loops: </label>
                <input type="checkbox" v-model="filters.category.forloop">
                <br>
                <label>While Loops: </label>
                <input type="checkbox" v-model="filters.category.whileloop">
                <br>
                <label>If Statement: </label>
                <input type="checkbox" v-model="filters.category.ifstmt">
                <br>
                <label>Recursion: </label>
                <input type="checkbox" v-model="filters.category.recursion">
                <br>
                <label>Misc: </label>
                <input type="checkbox" v-model="filters.category.misc">
                <br>
            </div>
        </div>

        <div class="column">
            <h5>Constraint</h5>
            <div class="filter-selection">
                <label>For Loops: </label>
                <input type="checkbox" v-model="filters.constraint.forloop">
                <br>
                <label>While Loops: </label>
                <input type="checkbox" v-model="filters.constraint.whileloop">
                <br>
                <label>Recursion: </label>
                <input type="checkbox" v-model="filters.constraint.recursion">
                <br>
            </div>
        </div>

        <div class="column">
            <h5>Keyword Search</h5>
            <input @input="onFilterFormChanged" type="text" v-model="filters.keyword">
            <button :disabled="this.searching" @click.prevent="onFilterFormChanged" style="margin-top: 30px">Search</button>
        </div>

    </div>
    </form>

    <!-- Here is where the drag-and-drop exam creation happens -->
    <div class="two-column-container">

        <!-- Questions to be added to exam column -->
        <div @drop="onDrop($event, 2)" 
             @dragover.prevent
             @dragenter.prevent
             class="single-column-container">

             <div>
                 <h3 class="column-title">EXAM</h3>
             </div>

            <ExamCreatorItem v-for="question in examList" 
                             :question="question"
                             class="single-column-item moveable"
                             draggable
                             @dragstart="dragStart($event, question)"
                             show-points-input/>

         </div>
    
        <!-- Question bank column -->
        <div @drop="onDrop($event, 1)" 
             @dragover.prevent
             @dragenter.prevent
             class="single-column-container">

             <div>
                 <h3 class="column-title">BANK</h3>
             </div>

            <ExamCreatorItem v-for="question in bankList()" 
                             :question="question"
                             draggable
                             @dragstart="dragStart($event, question)"
                             class="single-column-item moveable"/>
        </div>

    </div>

    <!-- Submit button -->
    <button @click="onSubmit">Create Exam</button>
    <p :style="{ color: success ? 'green' : 'red' }">{{ errorMessage }}</p>
</template>

<script>
import ExamCreatorItem from "../components/ExamCreatorItem";

const bankColumn = 1;
const examColumn = 2;

export default {
    name: 'ExamCreator',
    inject: ['fetchQuestions', 'postExam'],
    components: {
        ExamCreatorItem
    },
    data() {
        return {
            filters: {
                difficulty: {
                    easy: true,
                    medium: true,
                    hard: true,
                },
                category: {
                    forloop: true,
                    whileloop: true,
                    ifstmt: true,
                    recursion: true,
                    misc: true,
                },
                constraint: {
                    forloop: true,
                    whileloop: true,
                    recursion: true,
                },
                keyword: '',
            },
            examName: '',
            questions: [],
            displayQuestions: [],

            errorMessage: '',
            success: false,

            searching: false
        };
    },
    async created() {
        this.questions = await this.fetchQuestions();
        this.resetPoints();
        // Start each question in the bank column
        this.moveAllToColumn(bankColumn);
        this.displayQuestions = this.questions;
    },
    methods: {
        updateResults() {
            this.displayQuestions = this.questions;

            const filters = this.filters;
            for (const section in filters) {
                if (section === "keyword") {
                    // Apply keyword filter
                    this.displayQuestions = this.displayQuestions.filter(q => {
                        const kw = filters.keyword.trim();
                        try {
                            const regex = new RegExp(kw, 'i');
                            return q.desc.match(regex) !== null;
                        } catch (e) {
                            return false;
                        }
                    });
                    continue;
                }
                // Apply flag filter
                for (const [key, val] of Object.entries(filters[section]))
                    if (!val) 
                        this.displayQuestions = this.displayQuestions.filter( 
                            q => q[section] !== key
                        )
            }
            this.searching = false;
        },
        dragStart(event, question) {
            event.dataTransfer.dropEffect = 'move';
            event.dataTransfer.effectAllowed = 'move';
            event.dataTransfer.setData('questionID', question.id);
        },
        onDrop(event, list) {
            // Convert data transfers string id to int
            const questionID = +event.dataTransfer.getData('questionID');
            const question = this.questions.find(q => q.id === questionID);
            question.list = list;
        },
        setError(msg) {
            this.success = false;
            this.errorMessage = msg;
        },
        setSuccess(msg) {
            this.success = true;
            this.errorMessage = msg;
        },
        onFilterFormChanged() {
            this.searching = true;
            setTimeout(this.updateResults, 600 + Math.floor(Math.random() * 300));
        },
        async onSubmit() {
            let examQuestions = this.questions.filter(q => q.list === examColumn);
            examQuestions = examQuestions.map(q => {
                return {
                    id: q.id,
                    points: q.points,
                };
            });
            if (examQuestions.length === 0) {
                this.setError("Please add questions from the questions bank");
                return;
            }
            this.examName = this.examName.trim();
            if (this.examName.length === 0) {
                this.setError("Please enter an exam name");
                return;
            }

            // Post newly created exam to DB
            const payload = {
                name: this.examName,
                questions: examQuestions,
            };
            const res = await this.postExam(payload);

            // TODO:Check response
            this.setSuccess(`Successfully created exam: ${this.examName}`);
            this.examName = '';
            this.resetPoints();
            this.moveAllToColumn(bankColumn);
        },
        moveAllToColumn(col) {
            this.questions.forEach(question => {
                question.list = col;
            });
        },
        resetPoints() {
            this.questions.forEach(question => {
                question.points = 0;
            });
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
        bankList() {
            return this.displayQuestions.filter(q => q.list === bankColumn);
        }
    },
    computed: {
        /* bankList() { */
        /*     return this.questions.filter(q => q.list === bankColumn); */
        /* }, */
        examList() {
            return this.questions.filter(q => q.list === examColumn);
        },
        totalPoints() {
            const points = this.examList.map(q => q.points);
            if (points.length === 0)
                return 0;
            return points.reduce((p, c) => p + c);
        }
    }
}
</script>

<style scoped>
.row {
    display: flex;
    justify-content: safe center;
    width: 75%;
    max-width: 800px;
    margin: 0 auto;
}
.column {
    flex: 50%;
    text-align: center;
}
.column > input[type="text"]  {
    width: 75%;
}


.filter-selection {
    text-align: right;
    display: inline-block;
}

body {
    margin: 0;
}
.column-title {
    color: white;
}
.moveable {
    cursor: move;
}

</style>
