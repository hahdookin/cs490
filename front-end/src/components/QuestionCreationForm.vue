<template>
    <!-- Title -->
    <label for="title">Title</label>
    <input v-model="form.title"
           type="text" 
           name="title" 
           required><br>

    <!-- Description -->
    <label for="desc">Description</label><br>
    <textarea v-model="form.desc"
              name="desc" 
              rows="3" 
              cols="50" 
              required></textarea><br>

    <!-- Category -->
    <label for="category">Category</label>
    <select v-model="form.category" name="category" required>
        <option value="misc">Miscellaneous</option>
        <option value="for">For Loops</option>
        <option value="while">While Loops</option>
        <option value="recursion">Recursion</option>
        <option value="if">If Statements</option>
    </select><br>

    <!-- Constraint -->
    <label for="constraint">Constraint</label>
    <select v-model="form.constraint" name="constraint">
        <option value="none">None</option>
        <option value="for">For Loops</option>
        <option value="while">While Loops</option>
        <option value="recursion">Recursion</option>
    </select><br>

    <!-- Difficulty -->
    <label for="difficulty">Difficulty</label>
    <select v-model="form.difficulty" name="difficulty" required>
        <option value="easy">Easy</option>
        <option value="medium">Medium</option>
        <option value="hard">Hard</option>
    </select><br>

    <!-- Function Name -->
    <label for="functionname">Function Name</label>
    <input pattern="[A-Za-z_][A-Za-z0-9_]*" 
           v-model="form.functionname" 
           type="text" 
           name="functionname" 
           required><br>

    <label>Parameters: </label>
    <div>
        <input :key="index" 
               pattern="[A-Za-z_][A-Za-z0-9_]*"
               v-for="(param, index) in form.parameters"
               v-model="form.parameters[index]"
               type="text" 
               :placeholder="`param${index}`"
               size="10"
               required>
    <button @click.prevent="addParam">Add Parameter</button>
    <button @click.prevent="removeParam">Remove Parameter</button>
    </div>

    <label>Test Cases: </label>
    <div>
        <div :key="index"
             v-for="(test, index) in form.tests">
            <label>Test {{ index }}: </label>
            <input :key="idx"
                   v-for="(param, idx) in form.parameters"
                   v-model="test.arguments[idx]"
                   type="text"
                   :placeholder="`arg${idx}`"
                   size="10"
                   required>
            <input type="text"
                   placeholder="output"
                   v-model="test.output"
                   size="10"
                   required>
        </div>
        <button @click.prevent="addTest">Add Test</button>
        <button @click.prevent="removeTest">Remove Test</button>

        <!-- Submit or reset -->
        <div>
            <input @click="$emit('question-submitted')" 
                   type="submit" 
                   value="Create Question">
            <input @click.prevent="onReset" type="reset">
        </div>
    </div>
</template>

<script>
export default {
    name: 'QuestionCreationForm',
    emits: ['question-submitted'],
    props: {
        form: Object,
    },
    methods: {
        onReset() {
            this.form.title = '';
            this.form.desc = '';
            this.form.category = 'misc';
            this.form.constraint = 'none';
            this.form.difficulty = 'easy';
            this.form.functionname = '';
            this.form.parameters = [];
            this.form.tests = [];
        },
        addParam() {
            this.form.parameters.push('');
        },
        removeParam() {
            this.form.parameters.pop();
            // Remove parameter from all tests!
            this.form.tests.forEach(test => {
                test.form.arguments.pop();
            });
        },
        addTest() {
            this.form.tests.push({
                arguments: [],
                output: '',
            });
        },
        removeTest() {
            this.form.tests.pop();
        }
    }
}
</script>

<style>
</style>
