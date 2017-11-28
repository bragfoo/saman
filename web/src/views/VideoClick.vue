<template>
  <el-card class="box-card">
    <plat-play-bar-chart></plat-play-bar-chart>
    <plat-video-data-line-chart></plat-video-data-line-chart>
  </el-card>
</template>

<script>
  import PlatVideoDataLineChart from './data_views/PlatVideoPlayLineChart.vue'
  import PlatPlayBarChart from './data_views/PlatPlayBarChart.vue'

  export default {
    name: 'Main',
    components: {
      PlatVideoDataLineChart,
      PlatPlayBarChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
