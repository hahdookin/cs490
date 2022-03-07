// Common functionality for querying the database
export default {
    install: (app, options) => {

        const database = 'http://localhost:5000';

        // Fetch single student table
        const fetchStudent = async function(user) {
            const res = await fetch(`${database}/students?username=${user}`);
            const data = await res.json();
            return data[0]; // return first result
        };
        // Fetch single exam
        const fetchExam = async function (id) {
            const res = await fetch(`${database}/exams?id=${id}`);
            const data = await res.json();
            if (data.length === 0) {
                console.error(`Couldn't fetch exam data: ID ${id}`);
                return;
            }
            return data[0]; // return first result
        }
        // Fetch all exams
        const fetchExams = async function() {
            const res = await fetch(`${database}/exams`);
            const data = await res.json();
            return data;
        };
        // Fetch single teacher table
        const fetchTeacher = async function(user) {
            const res = await fetch(`${database}/teachers?username=${user}`);
            const data = await res.json();
            return data[0]; // return first result
        };
        // Fetch all teacher exams (examID, assignee, assigner)
        const fetchTeacherExams = async function(user) {
            const res = await fetch(`${database}/teacherexams?assigner=${user}`);
            const data = await res.json();
            return data;
        };
        // Fetch all student exams (completed (IDs), incompleted (IDs))
        const fetchStudentExams = async function(user) {
            const res = await fetch(`${database}/studentexams?username=${user}`);
            const data = await res.json();
            return data[0];
        };
        // Fetch all question tables from question IDs in exam questions
        const fetchQuestionsFromExam = async function(examID) {
            const exam = await fetchExam(examID);
            const questions = [];
            for (const { id, points } of exam.questions) {
                const question = await fetchQuestion(id);
                question.points = points;
                questions.push(question);
            }
            return questions;
        };
        // Fetch all questions (in bank)
        const fetchQuestions = async function() {
            const res = await fetch(`${database}/questions`);
            const data = await res.json();
            return data;
        };
        // Fetch single question
        const fetchQuestion = async function(questionID) {
            const res = await fetch(`${database}/questions?id=${questionID}`);
            const data = await res.json();
            return data[0]; // return first result
        };
        // Fetch single user login credentials
        const fetchLoginCredentials = async function(user, pass) {
            const res = await fetch(`${database}/users?username=${user}&password=${pass}`);
            const data = await res.json();
            return data;
        };

        app.provide('fetchStudent', fetchStudent);
        app.provide('fetchTeacher', fetchTeacher);

        app.provide('fetchExam', fetchExam);
        app.provide('fetchExams', fetchExams);

        app.provide('fetchStudentExams', fetchStudentExams);
        app.provide('fetchTeacherExams', fetchTeacherExams);

        app.provide('fetchQuestion', fetchQuestion);
        app.provide('fetchQuestions', fetchQuestions);
        app.provide('fetchQuestionsFromExam', fetchQuestionsFromExam);

        app.provide('fetchLoginCredentials', fetchLoginCredentials);
    }
}
