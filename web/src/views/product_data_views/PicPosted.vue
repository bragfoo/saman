<template>
  <el-card class="box-card">
    <pic-sum-line-chart></pic-sum-line-chart>
  </el-card>
</template>

<script>
  import PicSumLineChart from '../data_views/PicSumLineChart.vue'

  export default {
    name: 'PicPosted',
    components: {
      PicSumLineChart
    },
    data () {
      return {}
    },
    created () {},
    methods: {}
  }
</script>
