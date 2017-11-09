<template>
  <el-card class="box-card">
    <video-sum-line-chart></video-sum-line-chart>
  </el-card>
</template>

<script>
  import VideoSumLineChart from '../data_views/VideoSumLineChart.vue'

  export default {
    name: 'VideoPosted',
    components: {
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
