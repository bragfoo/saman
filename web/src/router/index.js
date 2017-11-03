import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/Main.vue'
import Gine from '@/components/Line.vue'
import ModifyData from '@/views/ModifyData.vue'
import AppUGC from '@/views/modify_views/AppUGC.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: Main
    }, {
      path: '/gain',
      name: 'gain',
      component: Gine
    }, {
      path: '/modify',
      name: 'modify',
      component: ModifyData,
      children: [
        {
          path: '/modify/plat_fans',
          name: 'modify_plat_fans',
          component: AppUGC
        }, {
          path: 'video_plat',
          name: 'modify_video_plat',
          component: AppUGC
        }, {
          path: 'video_plat_amount',
          name: 'modify_video_plat_amount',
          component: AppUGC
        }, {
          path: 'mobile_data',
          name: 'modify_mobile_data',
          component: AppUGC
        }, {
          path: 'app_ugc',
          name: 'modify_app_ugc',
          component: AppUGC
        }, {
          path: 'talent',
          name: 'modify_talent',
          component: AppUGC
        }, {
          path: 'app_data',
          name: 'modify_app_data',
          component: AppUGC
        }
      ]
    }
  ]
})
