import { createRouter, createWebHistory } from 'vue-router'
import Home from './components/Home.vue'
// import {  } from "";
const webHistory = createWebHistory();

const router = createRouter({
    history: webHistory,
    routes: [
        {
            path: "/home",
            name: "Home",
            // component: () => import('./components/Home.vue')
            component: Home
        },
        {
            path: "/",
            name: "Index",
            component: () => import('./components/Index.vue')
        },
        {
            path: "/books",
            name: "Books",
            component: () => import('./components/Books.vue')
        },
        {
            path: "/book_detail/:id",
            name: "book_detail",
            component: () => import('./components/BookDetail.vue')
        },
    ]
})

export default router;
