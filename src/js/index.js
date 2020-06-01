import '../css/main.scss'

import _ from 'lodash'
import "babel-polyfill"
import Vue from "vue"
import VueRouter from 'vue-router'


Vue.use(VueRouter)

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
