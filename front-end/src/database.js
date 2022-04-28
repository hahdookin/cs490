// Common functionality for querying the database
export default {
    install: (app, options) => {

        //const database = 'http://localhost:5000';
        const database = 'http://ec2-3-92-132-35.compute-1.amazonaws.com';
        const middle = 'http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com';

        const serialize = function(obj) {
            var str = [];
            for (var p in obj)
                if (obj.hasOwnProperty(p)) {
                    str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
                }
            return str.join("&");
        }

        const formatArg = function(arg) {
            if (isNaN(parseInt(arg)))
                return `'${arg}'`;
            return arg;
        }
        app.provide('formatArg', formatArg);

        const postToAutograder = async function(qid, code) {
            const res = await fetch(`${middle}/actual`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    qid: "" + qid, // Ensure qid is a string
                    code: code
                })
                // body: JSON.stringify({
                //     qid: qid,
                //     code: code
                // })
            });
            const resp = await res.json();
            return resp;
        };
        app.provide('postToAutograder', postToAutograder);

        // Fetch single user table
        const fetchUser = async function(userID) {
            const res = await fetch(`${database}/users/${userID}`);
            const data = await res.json();
            return data;
        }
        // Fetch single exam table by examID
        const fetchExam = async function(examID) {
            const res = await fetch(`${database}/exams/${examID}`);
            const data = await res.json();
            if (data.length === 0) {
                console.error(`Couldn't fetch exam data: ID ${id}`);
                return;
            }
            return data;
        }
        // Fetch all exams
        const fetchExams = async function() {
            const res = await fetch(`${database}/exams`);
            const data = await res.json();
            return data;
        };
        // Fetch all teacher exams (examID, assignee, assigner)
        const fetchTeacherExams = async function(userID) {
            const res = await fetch(`${database}/teacherexams?assignerid=${userID}`);
            const data = await res.json();
            return data;
        };
        // Fetch all student exams (completed (IDs), incompleted (IDs))
        const fetchStudentExams = async function(userID) {
            const res = await fetch(`${database}/studentexams?userid=${userID}`);
            const data = await res.json();
            return data[0]; // just want the one result
        };
        // Fetch all question tables from question IDs in exam questions
        const fetchQuestionsFromExam = async function(examID) {
            const exam = await fetchExam(examID);
            // Exam doesn't exist or has no questions (??), should throw but just return empty
            if (!exam.questions) return [];
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
        // Fetch all questions with query in desc
        const fetchQuestionsLike = async function(query) {
            const res = await fetch(`${database}/questions?desc_like=${query}`);
            const data = await res.json();
            return data;
        };
        // Fetch all questions matching the filter
        // const fetchQuestionsByFilter = async function(filter) {
        //     const res = await fetch(`${database}/questions?${filter}`);
        //     const data = await res.json();
        //     return data;
        // };
        // Fetch single question
        const fetchQuestion = async function(questionID) {
            const res = await fetch(`${database}/questions/${questionID}`);
            const data = await res.json();
            return data;
        };
        // Fetch single user login credentials
        const fetchLoginCredentials = async function(user, pass) {
            const res = await fetch(`${database}/users?username=${user}&password=${pass}`);
            const data = await res.json();
            return data[0]; // just want the one result
        };
        // Fetch userID by user name
        const fetchUserID = async function(user) {
            const res = await fetch(`${database}/users?username=${user}`);
            const data = await res.json();
            return data[0].id; // just want the one result
        }
        const fetchStudentExamResult = async function(userID, examID) {
            const res = await fetch(`${database}/studentexamresult?userid=${userID}&examid=${examID}`);
            const data = await res.json();
            return data[0]; // just want the one result
        }
        const fetchStudentExamResultByID = async function(ID) {
            const res = await fetch(`${database}/studentexamresult/${ID}`);
            const data = await res.json();
            return data;
        };
        const fetchStudentExamAnswers = async function(serID) {
            const res = await fetch(`${database}/studentexamanswers?studentexamresultid=${serID}`);
            const data = await res.json();
            return data;
        };
        app.provide('fetchStudentExamAnswers', fetchStudentExamAnswers);

        const postQuestion = async function(question) {
            const res = await fetch(`${database}/questions`, {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(question)
            });
        };
        app.provide('postQuestion', postQuestion);

        const postExam = async function(exam) {
            const res = await fetch(`${database}/exams`, {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(exam)
            });
        };
        app.provide('postExam', postExam);

        const postTeacherExam = async function(teacherexam) {
            const res = await fetch(`${database}/teacherexams`, {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(teacherexam)
            });
        };
        app.provide('postTeacherExam', postTeacherExam);

        const updateStudentExams = async function(studentexams) {
            const res = await fetch(`${database}/studentexams/${studentexams.id}`, {
                method: 'put',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(studentexams)
            });
        };
        app.provide('updateStudentExams', updateStudentExams);

        const postStudentExamResult = async function(studentexamresult) {
            const res = await fetch(`${database}/studentexamresult`, {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(studentexamresult)
            });
            return res;
        };
        app.provide('postStudentExamResult', postStudentExamResult);


        const postStudentExamAnswer = async function(sea) {
            const res = await fetch(`${database}/studentexamanswers`, {
                method: 'post',
                headers: {
                    'Content-type': 'application/json',
                },
                body: JSON.stringify(sea)
            });
        };
        app.provide('postStudentExamAnswer', postStudentExamAnswer);

        const putStudentExamResult = async function(ser) {
            const res = await fetch(`${database}/studentexamresult/${ser.id}`, {
                method: 'put',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(ser)
            });
        };
        app.provide('putStudentExamResult', putStudentExamResult);

        const putStudentExamAnswer = async function(sea) {
            const res = await fetch(`${database}/studentexamanswers/${sea.id}`, {
                method: 'put',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(sea)
            });
        };
        app.provide('putStudentExamAnswer', putStudentExamAnswer);

        app.provide('fetchExam', fetchExam);
        app.provide('fetchExams', fetchExams);

        app.provide('fetchStudentExams', fetchStudentExams);
        app.provide('fetchTeacherExams', fetchTeacherExams);

        app.provide('fetchStudentExamResult', fetchStudentExamResult);
        app.provide('fetchStudentExamResultByID', fetchStudentExamResultByID);

        app.provide('fetchQuestion', fetchQuestion);
        app.provide('fetchQuestions', fetchQuestions);
        app.provide('fetchQuestionsLike', fetchQuestionsLike);
        //app.provide('fetchQuestionsByFilter', fetchQuestionsByFilter);
        app.provide('fetchQuestionsFromExam', fetchQuestionsFromExam);

        app.provide('fetchUser', fetchUser);
        app.provide('fetchUserID', fetchUserID);
        app.provide('fetchLoginCredentials', fetchLoginCredentials);

        app.provide('serialize', serialize);

        // Divide points into n close-to equal parts
        const split = function(x, n) {
            if(x < n)
                return 0;

            let res = [];
            if (x % n == 0) {
                for (let i = 0; i < n; i++)
                    res.push(x/n);
            } else {
                let zp = n - (x % n);
                let pp = Math.floor(x/n);
                for(let i = 0; i < n; i++) {
                    if(i>= zp)
                        res.push(pp + 1);
                    else
                        res.push(pp);
                }
            }
            return res;
        };

        const zip = function(a, b) {
            return a.map((e, i) => [e, b[i]]);
        }

        app.provide('split', split);
        app.provide('zip', zip);
    }
}
