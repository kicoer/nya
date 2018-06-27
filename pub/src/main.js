import Vue from 'vue'
import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './app.vue'

Vue.use(Element, { size: 'small' })

// 载入路由
import router from './router.js'

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})