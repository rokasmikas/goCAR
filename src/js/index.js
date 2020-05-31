import '../css/main.scss'

import _ from 'lodash'
import "babel-polyfill"
import Vue from "vue"
import VueRouter from 'vue-router'


Vue.use(VueRouter)

import Home from "./home.vue"
import About from "./about.vue"

const routes = [
  { path: '/', component: Home },
  { path: '/about', component: About },
]

const router = new VueRouter({
  mode: 'hash',
  routes
})


const app = new Vue({
  router
}).$mount('#app')
