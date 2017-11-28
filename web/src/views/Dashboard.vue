<template>
  <el-row>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>平台数据</span>
      </div>
      <app-u-g-c-chart></app-u-g-c-chart>
      <plat-fans-chart></plat-fans-chart>
      <plat-video-data-line-chart></plat-video-data-line-chart>
    </el-card>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>应用数据</span>
      </div>
      <user-sum-line-chart></user-sum-line-chart>
      <talent-sum-line-chart></talent-sum-line-chart>
      <pic-sum-line-chart></pic-sum-line-chart>
      <video-sum-line-chart></video-sum-line-chart>
      <share-chart></share-chart>
    </el-card>
  </el-row>
</template>

<script>
  import AppUGCChart from './data_views/AppUGCChart.vue'
  import EventRardarChart from './data_views/EventRardarChart.vue'
  import PicDailyLineChart from './data_views/PicDailyLineChart.vue'
  import PicSumLineChart from './data_views/PicSumLineChart.vue'
  import PlatFansChart from './data_views/PlatFansChart.vue'
  import PlatVideoDataLineChart from './data_views/PlatVideoPlayLineChart.vue'
  import ShareChart from './data_views/ShareChart.vue'
  import SocDataLineChart from './data_views/SocDataLineChart.vue'
  import TalentSumLineChart from './data_views/TalentSumLineChart.vue'
  import UserSumLineChart from './data_views/UserSumLineChart.vue'
  import VideoDailyLineChart from './data_views/VideoDailyLineChart.vue'
  import VideoSumLineChart from './data_views/VideoSumLineChart.vue'

  export default {
    name: 'Dashboard',
    components: {
      AppUGCChart,
      EventRardarChart,
      PicDailyLineChart,
      PicSumLineChart,
      PlatFansChart,
      PlatVideoDataLineChart,
      ShareChart,
      SocDataLineChart,
      TalentSumLineChart,
      UserSumLineChart,
      VideoDailyLineChart,
      VideoSumLineChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
