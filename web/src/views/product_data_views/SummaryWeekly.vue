<template>
  <el-card class="box-card">
    <app-u-g-c-chart></app-u-g-c-chart>
  </el-card>
</template>

<script>
  import AppUGCChart from '../data_views/AppUGCChart.vue'

  export default {
    name: 'SummaryWeekly',
    components: {
      AppUGCChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
