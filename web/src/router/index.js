import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/Main.vue'
// 用户点击
import VideoClick from '@/views/VideoClick.vue'
// 投放渠道
import PutPlat from '@/views/PutPlat.vue'
// 产品数据
import ProductData from '@/views/ProductData.vue'
import PrdPicPosted from '@/views/product_data_views/PicPosted.vue'
import PrdVideoPosted from '@/views/product_data_views/VideoPosted.vue'
import PrdSoc from '@/views/product_data_views/Soc.vue'
import PrdTalent from '@/views/product_data_views/Talent.vue'
import PrdNewly from '@/views/product_data_views/Newly.vue'
// 下载量
import DownloadCount from '@/views/DownloadCount.vue'
// 用户转化
import UserReferral from '@/views/UserReferral.vue'
import UserefDownload from '@/views/user_referral_views/Download.vue'
import UserefRetention from '@/views/user_referral_views/Retention.vue'
// 活动检测
import EventCheck from '@/views/EventCheck.vue'
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
import ModifyVideoPlat from '@/views/modify_views/ModifyVideo.vue'
import ModifyVideoPlayAmount from '@/views/modify_views/ModifyVideoPlayAmount.vue'
import ModifyMobileData from '@/views/modify_views/ModifyMobileData.vue'
import ModifyAppUGC from '@/views/modify_views/ModifyAppUGC.vue'
import ModifyTalent from '@/views/modify_views/ModifyAppUGC.vue'
import ModifyAppData from '@/views/modify_views/ModifyAppData.vue'
import ModifyEvents from '@/views/modify_views/ModifyEvents.vue'
/* eslint-enable */

Vue.use(Router)

export default new Router({
  routes: [
    {
      // dashboard
      path: '/',
      name: 'dashboard',
      component: Main
      // }, {
      // 运营周报
      // path: '/operation_weekly',
      // name: 'operation_weekly',
      // component: Main
      // }, {
      // 用户画像
      // path: '/users_portrayal',
      // name: 'users_portrayal',
      // component: Main
    }, {
      // 视频点击
      path: '/video_click',
      name: 'video_click',
      component: VideoClick
    }, {
      // 投放渠道
      path: '/put_plat',
      name: 'put_plat',
      component: PutPlat
    }, {
      // 产品数据
      path: '/product_data',
      name: 'product_data',
      component: ProductData,
      children: [
        {
          // 图片发布
          path: '/prd/pic_posted',
          name: 'prd_pic_posted',
          component: PrdPicPosted
        }, {
          // 视频发布
          path: '/prd/video_posted',
          name: 'prd_video_posted',
          component: PrdVideoPosted
        }, {
          // 社会化
          path: '/prd/socialization',
          name: 'prd_socialization',
          component: PrdSoc
        }, {
          // 达人
          path: '/prd/talent',
          name: 'prd_talent',
          component: PrdTalent
        }, {
          // 新增
          path: '/prd/newly',
          name: 'prd_newly',
          component: PrdNewly
        }
      ]
    }, {
      // 下载量
      path: '/download_count',
      name: 'download_count',
      component: DownloadCount
    }, {
      // 用户转化
      path: '/user_referral',
      name: 'user_referral',
      component: UserReferral,
      children: [
        {
          // 下载
          path: '/useref/download',
          name: 'useref_download',
          component: UserefDownload
        }, {
          // 留存
          path: '/useref/retention',
          name: 'useref_retention',
          component: UserefRetention
        }
      ]
    }, {
      // 活动检测
      path: '/event_check',
      name: 'event_check',
      component: EventCheck
      // }, {
      // 周报
      // path: '/weekly',
      // name: 'weekly',
      // component: Main
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
          path: '/modify/video_plat',
          name: 'modify_video_plat',
          component: ModifyVideoPlat
        }, {
          path: '/modify/video_plat_amount',
          name: 'modify_video_plat_amount',
          component: ModifyVideoPlayAmount
        }, {
          path: '/modify/mobile_data',
          name: 'modify_mobile_data',
          component: ModifyMobileData
        }, {
          path: '/modify/app_ugc',
          name: 'modify_app_ugc',
          component: ModifyAppUGC
        }, {
          path: '/modify/events',
          name: 'modify_events',
          component: ModifyEvents
        }, {
          path: '/modify/talent',
          name: 'modify_talent',
          component: ModifyTalent
        }, {
          path: '/modify/app_data',
          name: 'modify_app_data',
          component: ModifyAppData
        }
      ]
    }
  ]
})
