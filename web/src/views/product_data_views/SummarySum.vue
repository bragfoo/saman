<template>
  <el-card class="box-card">
    <app-u-g-c-sum-chart></app-u-g-c-sum-chart>
  </el-card>
</template>

<script>
  import AppUGCSumChart from '../data_views/AppUGCSumChart.vue'

  export default {
    name: 'SummarySum',
    components: {
      AppUGCSumChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
