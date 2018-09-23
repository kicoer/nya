import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            component: resolve => require(['@/page/index.vue'], resolve)
        },
        {
            path: '/login/:url',
            component: resolve => require(['@/page/login.vue'], resolve)
        },
        {
            path: '/chat/:name/:img/:f',
            component: resolve => require(['@/page/chat.vue'], resolve)
        }
    ]
})

