import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '../views/LoginPage.vue'

const routes = [
    {
        path: '/',
        redirect: '/loginpage'
    },
    {
        path: '/loginpage',
        component: LoginPage
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
  })

  export default router