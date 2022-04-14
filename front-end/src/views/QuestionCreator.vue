<template>

    <h3>Question Creator</h3>
    <div class="two-column-container" style="max-width: 1000px">

        <!-- Question creation form -->
        <div class="single-column-container">
            <div class="single-column-item">
                <QuestionCreationForm @question-submitted="onSubmit" :form="form" />
            </div>
        </div>

        <!-- Find questions in bank form -->
        <div class="single-column-container">
            <form @change="onFilterFormChanged" @submit.prevent>
                <div class="single-column-item">
                    <div class="row">
                        <div class="column">
                            <h5>Difficulty: </h5>
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
                            <h5>Category: </h5>
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
                            <h5>Keyword: </h5>
                            <input @input="onFilterFormChanged" 
                                   type="text" 
                                   v-model="filters.keyword">
                        </div>
                    </div>
                </div>
            </form>

            <!-- Results of query -->
            <ExamCreatorItem class="single-column-item" 
                             :question="question"
                             v-for="question in displayQuestions"
                             :key="question.id"/>
        </div>

    </div>
    <p :style="{ color: success ? 'green' : 'red' }">{{ errorMessage }}</p>

</template>

<script>
import QuestionCreationForm from '../components/QuestionCreationForm';
import ExamCreatorItem from '../components/ExamCreatorItem.vue';

export default {
    name: 'QuestionCreator',
    components: {
        QuestionCreationForm,
        ExamCreatorItem
    },
    inject: [
        'postQuestion',
        'fetchQuestionsLike',
        'fetchQuestions'
    ],
    data() {
        return {
            errorMessage: '',
            success: false,

            filters: {
                difficulty: {
                    easy: false,
                    medium: false,
                    hard: false,
                },
                category: {
                    forloop: false,
                    whileloop: false,
                    ifstmt: false,
                    recursion: false,
                    misc: false,
                },
                keyword: '',
            },

            query: '',
            questions: [],
            displayQuestions: [],

            form: {
                title: '',
                desc: '',
                category: 'misc',
                constraint: 'none',
                difficulty: 'easy',
                functionname: '',
                parameters: [],
                tests: [],
            }
        }
    },
    async created() {
        this.questions = await this.fetchQuestions();
    },
    methods: {
        setError(msg) {
            this.success = false;
            this.errorMessage = msg;
            return false;
        },
        setSuccess(msg) {
            this.success = true;
            this.errorMessage = msg;
        },
        updateResults() {
            this.displayQuestions = this.questions;

            const filters = this.filters;
            for (const section in filters) {
                if (section === 'keyword' && filters.keyword !== '') {
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
        },
        onFilterFormChanged() {
            this.updateResults();
        },
        async onSubmit() {
            if (!this.validateForm())
                return;
            const payload = this.form;

            // Post question to DATABASE
            const res = await this.postQuestion(payload);

            this.questions = await this.fetchQuestions();

            this.setSuccess(`Successfully created question: ${this.form.title}`);
        },
        validateForm() {
            const errRequired = 'Missing required field: ';

            if (this.form.title === '')
                return this.setError(errRequired + 'Title');

            if (this.form.desc === '')
                return this.setError(errRequired + 'Description');

            if (this.form.functionname === '')
                return this.setError(errRequired + 'Function Name');

            for (const param of this.form.parameters) {
                if (param === '') 
                    return this.setError(errRequired + 'Param');
            }
            for (const test of this.form.tests) {
                for (const arg of test.arguments)
                    if (arg === '')
                        return this.setError(errRequired + 'Test Argument');
                if (test.output === '')
                    return this.setError(errRequired + 'Test Output');
            }

            // Check patterns
            const validPythonIdent = function(ident) {
                return ident.match(/[A-Za-z_][A-Za-z0-9_]*/) !== null;
            };
            if (!validPythonIdent(this.form.functionname))
                return this.setError('Incorrect format: Function Name');
            for (const param of this.form.parameters) {
                if (!validPythonIdent(param))
                    return this.setError(`Incorrect format: Parameter (${param})`);
            };

            // Check valid parameter count
            if (this.form.parameters.length === 0)
                return this.setError('Need at least 1 parameter.');

            // Check for duplicate parameter names
            const duplicates = this.form.parameters.filter((p, i) => {
                return this.form.parameters.indexOf(p) !== i;
            });
            if (duplicates.length !== 0)
                return this.setError(`Duplicate parameters: ${duplicates.toString()}`);

            // Check for atleast two test cases
            if (this.form.tests.length < 2)
                return this.setError('Need at least 2 test cases.');

            return true;
        }
    },
}
</script>

<style scoped>
.row {
    display: flex;
    justify-content: safe center;
    width: 100%;
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
</style>
