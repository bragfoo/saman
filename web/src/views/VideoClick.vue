<template>
  <el-card class="box-card">
    <plat-play-sum-line-chart></plat-play-sum-line-chart>
    <plat-play-bar-chart></plat-play-bar-chart>
    <plat-video-data-line-chart></plat-video-data-line-chart>
  </el-card>
</template>

<script>
  import PlatVideoDataLineChart from './data_views/PlatVideoPlayLineChart.vue'
  import PlatPlayBarChart from './data_views/PlatPlayBarChart.vue'
  import PlatPlaySumLineChart from './data_views/PlatPlaySumLineChart.vue'

  export default {
    name: 'Main',
    components: {
      PlatVideoDataLineChart,
      PlatPlayBarChart,
      PlatPlaySumLineChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
