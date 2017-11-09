<template>
  <el-card class="box-card">
    <video-sum-line-chart></video-sum-line-chart>
    <video-daily-line-chart></video-daily-line-chart>
  </el-card>
</template>

<script>
  import VideoSumLineChart from '../data_views/VideoSumLineChart.vue'
  import VideoDailyLineChart from '../data_views/VideoDailyLineChart.vue'

  export default {
    name: 'VideoPosted',
    components: {
      VideoSumLineChart,
      VideoDailyLineChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
