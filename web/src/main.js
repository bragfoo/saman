// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App.vue'
import router from './router'

import ElementUI from 'element-ui'
import './assets/element/index.css'

import axios from 'axios'

import VCharts from 'v-charts'

import * as filters from './util/filters'

Vue.config.productionTip = false

Vue.use(ElementUI)

// Vue.use(VueAxios, axios)
Vue.prototype.$http = axios.create({
  baseURL: 'http://localhost:8081/api/v1/',
  timeout: 1000,
  headers: {'X-Custom-Header': 'foobar'}
})

Vue.use(VCharts)

Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {App}
})
