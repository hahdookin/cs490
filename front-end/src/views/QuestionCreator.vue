<template>

    <h3>Question Creator</h3>
    <div class="two-column-container">

        <!-- Question creation form -->
        <div class="single-column-container">
            <div class="single-column-item">
                <QuestionCreationForm @question-submitted="onSubmit" :form="form" />
            </div>
        </div>

        <!-- Find questions in bank form -->
        <div class="single-column-container">
            <div class="single-column-item">
                <div style="display: inline">
                    <p style="display: inline">Query: </p>
                    <input @input="onQueryChange" 
                           type="text" 
                           v-model="query"
                           size="10">
                </div>
                <!--<button @click="onQueryChange">Submit</button>-->
            </div>
            <!-- Results of query -->
            <div class="single-column-item"
                 :key="question.id" 
                 v-for="question in questions">
                <p><strong>{{ question.title }}</strong></p>
                <p>{{ question.desc }}</p>
            </div>
        </div>

    </div>
    <p :style="{ color: success ? 'green' : 'red' }">{{ errorMessage }}</p>

</template>

<script>
import QuestionCreationForm from '../components/QuestionCreationForm';

export default {
    name: 'QuestionCreator',
    components: {
        QuestionCreationForm,
    },
    inject: [
        'postQuestion',
        'fetchQuestionsLike'
    ],
    data() {
        return {
            errorMessage: '',
            success: false,

            query: '',
            questions: [],

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
        async onQueryChange() {
            // When query is empty, don't show anything
            if (this.query === '')
                this.questions = [];
            else {
                // Catch errors if user entered incorrect regex syntax
                try {
                    this.questions = await this.fetchQuestionsLike(encodeURIComponent(this.query));
                } catch (e) {
                    this.questions = [];
                }
            }
        },
        async onSubmit() {
            if (!this.validateForm())
                return;
            const payload = this.form;

            // Post question to DATABASE
            const res = await this.postQuestion(payload);

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
