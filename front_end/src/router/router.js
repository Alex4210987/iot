import { createRouter, createWebHistory } from 'vue-router'
import Index from "@/page/IndexPage.vue";
import History from "@/page/HistoryPage.vue";
import AllDevice from "@/page/AllDevicePage.vue";
import AllUsers from "@/page/AllUsersPage.vue";
import Login from "@/page/LoginPage.vue";
import Register from "@/page/RegisterPage.vue";
import FaceIdentify from '@/page/FaceIdentifyPage.vue'

const routes = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/index',
        name: 'Index',
        component: Index
    },
    {
        path: '/history',
        name: 'History',
        component: History
    },
    {
        path: '/allDevice',
        name: 'AllDevice',
        component: AllDevice
    },
    {
        path: '/allUsers',
        name: 'AllUsers',
        component: AllUsers
    },
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/register',
        name:'Register',
        component: Register
    },
    {
        path: '/faceidentify',
        name:'FaceIdentify',
        component: FaceIdentify
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;