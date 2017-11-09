<template>
  <el-card class="box-card">
    <pic-sum-line-chart></pic-sum-line-chart>
    <pic-daily-line-chart></pic-daily-line-chart>
  </el-card>
</template>

<script>
  import PicSumLineChart from '../data_views/PicSumLineChart.vue'
  import PicDailyLineChart from '../data_views/PicDailyLineChart.vue'

  export default {
    name: 'PicPosted',
    components: {
      PicSumLineChart,
      PicDailyLineChart
    },
    data () {
      return {}
    },
    created () {},
    methods: {}
  }
</script>
