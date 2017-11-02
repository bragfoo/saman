import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/Main.vue'
import ModifyData from '@/views/ModifyData.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: Main
    }, {
      path: '/modify',
      name: 'modify',
      component: ModifyData,
      children: [
        {
          path: '/modify/plat_fans',
          name: 'modify_plat_fans',
          component: ModifyData
        }, {
          path: 'video_plat',
          name: 'modify_video_plat',
          component: ModifyData
        }, {
          path: 'video_plat_amount',
          name: 'modify_video_plat_amount',
          component: ModifyData
        }, {
          path: 'mobile_data',
          name: 'modify_mobile_data',
          component: ModifyData
        }, {
          path: 'app_ugc',
          name: 'modify_app_ugc',
          component: ModifyData
        }, {
          path: 'talent',
          name: 'modify_talent',
          component: ModifyData
        }, {
          path: 'app_data',
          name: 'modify_app_data',
          component: ModifyData
        }
      ]
    }
  ]
})
