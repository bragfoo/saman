// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App.vue'
import router from './router'

import ElementUI from 'element-ui'
import './assets/element/index.css'

import axios from 'axios'
import VueAxios from 'vue-axios'

import VCharts from 'v-charts'

import * as filters from './util/filters'

Vue.config.productionTip = false

Vue.use(ElementUI)

Vue.use(VueAxios, axios)

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
