import Vue from 'vue'
import VueRouter from "vue-router";

// 载入页面
import Index from "./pages/index.vue"
import Login from "./pages/login.vue";
import Chat from "./pages/chat.vue";

Vue.use(VueRouter);

const routes = [
    {
        path:"/login/:url",
        component: Login
    },
    {
        path: "/chat/:name/:img",
        component: Chat
    },
    {
        path: "/",
        component: Index
    }
]

var router =  new VueRouter({
    mode: 'history',
    routes
})
export default router;