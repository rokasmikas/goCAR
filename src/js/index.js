import '../css/main.scss'

import _ from 'lodash'
import "babel-polyfill"
import Vue from "vue"
import VueRouter from 'vue-router'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'


Vue.use(VueRouter)
Vue.use(Antd)

// partials
import NavBar from "./partials/nav.vue"
import Footer from "./partials/footer.vue"


Vue.component('navbar', NavBar)
Vue.component('footbar', Footer)

//components
import Home from "./home.vue"
import AddLog from "./addCarLog.vue"
import ShowLog from "./showLog.vue"

// TODO: Export router to different file
const routes = [
  { path: '/', component: Home },
  { path: '/addlog', component: AddLog },
  { path: '/carlog/:id', name: 'showlog', component: ShowLog },
]

const router = new VueRouter({
  mode: 'hash',
  routes
})


const app = new Vue({
  router
}).$mount('#app')
