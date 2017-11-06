import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/Main.vue'
// Acquisition
import Acquisition from '@/views/Acquisition.vue'
// Activation
import Activation from '@/views/Activation.vue'
// Retention
import Retention from '@/views/Retention.vue'
// Referral
import Referral from '@/views/Referral.vue'
// Revenue
import Revenue from '@/views/Revenue.vue'
// ModifyData
/* eslint-disable */
import ModifyData from '@/views/ModifyData.vue'
import ModifyPlatFans from '@/views/modify_views/ModifyAppUGC.vue'
import ModifyVideoPlat from '@/views/modify_views/ModifyAppUGC.vue'
import ModifyVideoPlatAmount from '@/views/modify_views/ModifyAppUGC.vue'
import ModifyMobileData from '@/views/modify_views/ModifyAppUGC.vue'
import ModifyAppUGC from '@/views/modify_views/ModifyAppUGC.vue'
import ModifyTalent from '@/views/modify_views/ModifyAppUGC.vue'
import ModifyAppData from '@/views/modify_views/ModifyAppUGC.vue'
/* eslint-enable */

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: Main
    }, {
      path: '/acquisition',
      name: 'acquisition',
      component: Acquisition
    }, {
      path: '/activation',
      name: 'activation',
      component: Activation
    }, {
      path: '/Retention',
      name: 'Retention',
      component: Retention
    }, {
      path: '/Referral',
      name: 'Referral',
      component: Referral
    }, {
      path: '/Revenue',
      name: 'Revenue',
      component: Revenue
    }, {
      path: '/modify',
      name: 'modify',
      component: ModifyData,
      children: [
        {
          path: '/modify/plat_fans',
          name: 'modify_plat_fans',
          component: ModifyPlatFans
        }, {
          path: 'video_plat',
          name: 'modify_video_plat',
          component: ModifyVideoPlat
        }, {
          path: 'video_plat_amount',
          name: 'modify_video_plat_amount',
          component: ModifyVideoPlatAmount
        }, {
          path: 'mobile_data',
          name: 'modify_mobile_data',
          component: ModifyMobileData
        }, {
          path: 'app_ugc',
          name: 'modify_app_ugc',
          component: ModifyAppUGC
        }, {
          path: 'talent',
          name: 'modify_talent',
          component: ModifyTalent
        }, {
          path: 'app_data',
          name: 'modify_app_data',
          component: ModifyAppData
        }
      ]
    }
  ]
})
