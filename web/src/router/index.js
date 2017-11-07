import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/Main.vue'
// Acquisition
import Acquisition from '@/views/Acquisition.vue'
import AcqSummary from '@/views/acquisition_views/Summary.vue'
// Activation
import Activation from '@/views/Activation.vue'
// Retention
import Retention from '@/views/Retention.vue'
import RetSummary from '@/views/retention_views/Summary.vue'
// Referral
import Referral from '@/views/Referral.vue'
import RefSummary from '@/views/referral_views/Summary.vue'
// Revenue
import Revenue from '@/views/Revenue.vue'
// ModifyData
/* eslint-disable */
import ModifyData from '@/views/ModifyData.vue'
import ModifyPlatFans from '@/views/modify_views/ModifyPlatFans.vue'
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
      // dashboard
      path: '/',
      name: 'dashboard',
      component: Main
    }, {
      // 运营周报
      path: '/operation_weekly',
      name: 'operation_weekly',
      component: Main
    }, {
      // 用户画像
      path: '/users_portrayal',
      name: 'users_portrayal',
      component: Main
    }, {
      // 视频点击
      path: '/users_click',
      name: 'users_click',
      component: Main
    }, {
      // 投放渠道
      path: '/put_plat',
      name: 'put_plat',
      component: Main
    }, {
      // 产品数据
      path: '/product_data',
      name: 'product_data',
      component: Main,
      children: [
        {
          // 图片发布
          path: '/prd/pic_posted',
          name: 'prd_pic_posted',
          component: Main
        }, {
          // 视频发布
          path: '/prd/video_posted',
          name: 'prd_video_posted',
          component: Main
        }, {
          // 社会化
          path: '/prd/socialization',
          name: 'prd_socialization',
          component: Main
        }, {
          // 达人
          path: '/prd/talent',
          name: 'prd_talent',
          component: Main
        }, {
          // 新增
          path: '/prd/newly',
          name: 'prd_newly',
          component: Main
        }
      ]
    }, {
      // 下载量
      path: '/download_count',
      name: 'download_count',
      component: Main
    }, {
      // 用户转化
      path: '/user_referral',
      name: 'user_referral',
      component: Main,
      children: [
        {
          // 获取总结
          path: '/useref/download',
          name: 'useref_download',
          component: Main
        }
      ]
    }, {
      // 活动检测
      path: '/event_check',
      name: 'event_check',
      component: Main
    }, {
      // 周报
      path: '/weekly',
      name: 'weekly',
      component: Main
    }, {
      // 获取
      path: '/acquisition',
      name: 'acquisition',
      component: Acquisition,
      children: [
        {
          // 获取总结
          path: '/acq/summary',
          name: 'acq_aummary',
          component: AcqSummary
        }
      ]
    }, {
      // 激活
      path: '/activation',
      name: 'activation',
      component: Activation
    }, {
      // 留存
      path: '/Retention',
      name: 'retention',
      component: Retention,
      children: [
        {
          path: '/ret/summary',
          name: 'ret_aummary',
          component: RetSummary
        }
      ]
    }, {
      // 转化
      path: '/Referral',
      name: 'referral',
      component: Referral,
      children: [
        {
          path: '/ref/summary',
          name: 'ref_summary',
          component: RefSummary
        }
      ]
    }, {
      // 付费
      path: '/Revenue',
      name: 'revenue',
      component: Revenue
    }, {
      // 数据修改
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
