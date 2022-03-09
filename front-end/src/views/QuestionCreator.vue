<template>
    <h3>Question Creator</h3>
    <!-- A Question has the following required elements: -->
    <!-- - Title -> Title of the question -->
    <!-- - Desc -> A short description of the problem -->
    <!-- - Difficulty -> The problems difficulty: one of 3 options -->
    <!-- - Function Name -> Name of the function -->
    <!-- - Parameters -> Up to X parameters -->
    <!-- - Examples -> Up to X examples -->
    <form name="questioncreatorform" 
          id="questioncreatorform" 
          @submit="onSubmit">

        <!-- Title -->
        <label for="title">Title</label>
        <input v-model="title"
               type="text" 
               name="title" 
               required><br>

        <!-- Description -->
        <label for="desc">Description</label><br>
        <textarea v-model="desc"
                  name="desc" 
                  rows="3" 
                  cols="50" 
                  required></textarea><br>

        <!-- Difficulty -->
        <label for="difficulty">Difficulty</label>
        <select v-model="difficulty" name="difficulty" required>
            <option value="easy">Easy</option>
            <option value="medium">Medium</option>
            <option value="hard">Hard</option>
        </select><br>

        <!-- Function Name -->
        <label for="functionname">Function Name</label>
        <input pattern="[A-Za-z_][A-Za-z0-9_]*" 
               v-model="functionname" 
               type="text" 
               name="functionname" 
               required><br>

        <label>Parameters: </label>
        <div>
            <input :key="index" 
                   pattern="[A-Za-z_][A-Za-z0-9_]*"
                   v-for="(param, index) in parameters"
                   v-model="parameters[index]"
                   type="text" 
                   :placeholder="`param${index}`"
                   required>
        </div>
        <button @click.prevent="addParam">Add Parameter</button>
        <button @click.prevent="removeParam">Remove Parameter</button>
        <br>

        <label>Test Cases: </label>
        <div>
            <div :key="index"
                 v-for="(test, index) in tests">
                <label>Test {{ index }}: </label>
                <input :key="idx"
                       v-for="(param, idx) in parameters"
                       v-model="test.arguments[idx]"
                       type="text"
                       :placeholder="`arg${idx}`"
                       required>
                <input placeholder="output"
                       v-model="test.output">
            </div>
        </div>
        <button @click.prevent="addTest">Add Test</button>
        <button @click.prevent="removeTest">Remove Test</button>

        <!-- DEBUG REMOVE ME -->
        <div>
            <p>Params: {{ parameters.toString() }}</p>
            <p>Tests: </p>
            <div v-for="(test, i) in tests">
                <p>({{ test.arguments.toString() }}) -> {{ test.output }}</p>
            </div>
        </div>

        <div>
            <input type="submit" value="Create Question">
            <input type="reset">
            <p :style="{ color: success ? 'green' : 'red' }">{{ errorMessage }}</p>
        </div>
    </form>
</template>

<script>

export default {
    name: 'QuestionCreator',
    data() {
        return {
            errorMessage: '',
            success: false,
            title: '',
            desc: '',
            difficulty: 'easy',
            functionname: '',
            parameters: [],
            tests: [],
        }
    },
    methods: {
        setError(msg) {
            this.success = false;
            this.errorMessage = msg;
        },
        setSuccess(msg) {
            this.success = true;
            this.errorMessage = msg;
        },
        async onSubmit(e) {
            e.preventDefault();
            if (!this.validateForm())
                return;
            console.log("Submitting form!");
            const payload = {
                title: this.title,
                desc: this.desc,
                difficulty: this.difficulty,
                functionname: this.functionname,
                parameters: this.parameters,
                tests: this.tests
            };
            // Post question to DATABASE
            const res = await fetch('http://localhost:5000/questions', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(payload)
            });
            this.setSuccess(`Successfully created question: ${this.title}`);
        },
        validateForm() {
            // Validate form (the parts which haven't already been validated)
            // Check for duplicate parameters
            const duplicates = this.parameters.filter((p, i) => {
                return this.parameters.indexOf(p) !== i;
            });
            if (duplicates.length !== 0) {
                this.setError(`Duplicate parameters: ${duplicates.toString()}`);
                return false;
            }
            // Check for no test cases
            if (this.tests.length < 2) {
                this.setError('Need at least 2 test cases.');
                return false;
            }
            return true;
        },
        addParam() {
            this.parameters.push('');
        },
        removeParam() {
            this.parameters.pop();
            // Remove parameter from all tests!
            this.tests.forEach(test => {
                test.arguments.pop();
            });
        },
        addTest() {
            this.tests.push({
                arguments: [],
                output: '',
            });
        },
        removeTest() {
            this.tests.pop();
        }
    },
}
</script>