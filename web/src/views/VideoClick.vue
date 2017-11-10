<template>
  <el-card class="box-card">
    <plat-video-data-line-chart></plat-video-data-line-chart>
  </el-card>
</template>

<script>
  import PlatVideoDataLineChart from './data_views/PlatVideoDataLineChart.vue'

  export default {
    name: 'Main',
    components: {
      PlatVideoDataLineChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
