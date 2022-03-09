<template>
    <h3>Exam Creator</h3>
    <label>Exam Name: </label>
    <input v-model="examName" type="text">

    <p>Total Points: {{ totalPoints }}</p>

    <!-- Here is where the drag-and-drop exam creation happens -->
    <div class="exam-creator-container">

        <!-- Questions to be added to exam column -->
        <div @drop="onDrop($event, 2)" 
             @dragover.prevent
             @dragenter.prevent
             class="questions-container">

             <div>
             <h3 class="column-title">EXAM</h3>
             </div>

            <div :key="question.id" 
                 v-for="question in examList"
                 draggable="true"
                 class="question-item" 
                 :id="'question' + question.id"
                 @dragstart="dragStart($event, question)">
                <h3>{{ question.title }}({{question.id}})</h3>
                <h5 :style="{ color: difficultyColor(question.difficulty) }">{{ question.difficulty }}</h5>
                <p>{{ question.desc }}</p>
                <p v-for="test in question.tests">
                    <code>
                    {{ question.functionname }}
                    ({{ test.arguments.join(", ") }}) -&gt; {{ test.output }}
                    </code>
                </p>
                <label>Points: </label>
                <!-- TODO: Handle invalid numeric input -->
                <input v-model="question.points" type="number" min="0">
            </div>

         </div>
    
        <!-- Question bank column -->
        <div @drop="onDrop($event, 1)" 
             @dragover.prevent
             @dragenter.prevent
             class="questions-container">

             <div>
             <h3 class="column-title">BANK</h3>
             </div>

            <div :key="question.id" 
                 v-for="question in bankList" 
                 draggable="true"
                 class="question-item" 
                 :id="'question' + question.id"
                 @dragstart="dragStart($event, question)">
                <h3>{{ question.title }}({{question.id}})</h3>
                <h5 :style="{ color: difficultyColor(question.difficulty) }">{{ question.difficulty }}</h5>
                <p>{{ question.desc }}</p>
                <p v-for="test of question.tests">
                    <code>
                    {{ question.functionname }}
                    ({{ test.arguments.join(", ") }}) -&gt; {{ test.output }}
                    </code>
                </p>
            </div>

        </div>

    </div>

    <!-- Submit button -->
    <button @click="onSubmit">Create Exam</button>
    <p :style="{ color: success ? 'green' : 'red' }">{{ errorMessage }}</p>
</template>

<script>
const bankColumn = 1;
const examColumn = 2;

export default {
    name: 'ExamCreator',
    inject: ['fetchQuestions'],
    components: {},
    props: {},
    data() {
        return {
            examName: '',
            questions: [],
            errorMessage: "",
            success: false
        };
    },
    methods: {
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
            const res = await fetch('http://localhost:5000/exams', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(payload)
            });
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
        }
    },
    async created() {
        this.questions = await this.fetchQuestions();
        this.resetPoints();
        // Start each question in the bank column
        // 1 -> bank
        // 2 -> exam
        this.moveAllToColumn(bankColumn);
    },
    computed: {
        bankList() {
            return this.questions.filter(q => q.list === bankColumn);
        },
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
body {
    margin: 0;
}
.column-title {
    color: white;
}

.questions-container {
    background-color: #344;
    padding: 1rem;
    margin-top: 1rem;
    flex: 1;
}

.question-item {
    padding: 1rem;
    background-color: white;
    cursor: move;
    border: 1px solid blue;
}

.exam-creator-container {
    display: flex;
    justify-content: safe center;
    width: 75%;
    margin: 0 auto;
    max-width: 800px;
}

.invisible {
    display: none;
}
</style>
