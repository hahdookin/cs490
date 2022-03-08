import { createRouter, createWebHistory } from 'vue-router';

import Home from '../views/Home';
import NotFound from '../views/NotFound';
import TeacherDashboard from '../views/TeacherDashboard';
import StudentDashboard from '../views/StudentDashboard';
import Exam from '../views/Exam.vue';
import ExamList from '../views/ExamList';
import AssignExam from '../views/AssignExam';
import ExamCreator from '../views/ExamCreator';
import QuestionCreator from '../views/QuestionCreator';
import ViewExam from '../views/ViewExam';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/teacher/:userid',
        name: 'TeacherDashboard',
        component: TeacherDashboard,
        children: [
            {
                path: 'examcreator',
                component: ExamCreator
            },
            {
                path: 'questioncreator',
                component: QuestionCreator
            },
            {
                path: 'assignexam',
                component: AssignExam
            },
            {
                path: 'viewexams',
                component: ViewExam
            }
        ]
    },
    {
        path: '/student/:userid',
        name: 'StudentDashboard',
        component: StudentDashboard,
        children: [
            {
                path: 'exam/:examid',
                component: Exam
            },
            {
                path: 'examlist',
                component: ExamList
            }
        ]
    },
    // 404 catch all
    {
        path: '/:catchAll(.*)',
        name: 'NotFound',
        component: NotFound
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
