import '../css/main.scss'

import _ from 'lodash'
import "babel-polyfill"
import Vue from "vue"
import VueRouter from 'vue-router'


Vue.use(VueRouter)

import Home from "./home.vue"
import AddLog from "./addCarLog.vue"

const routes = [
  { path: '/', component: Home },
  { path: '/addlog', component: AddLog },
]

const router = new VueRouter({
  mode: 'hash',
  routes
})


const app = new Vue({
  router
}).$mount('#app')
