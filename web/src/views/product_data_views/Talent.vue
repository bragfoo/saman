<template>
  <el-card class="box-card">
    <talent-sum-line-chart></talent-sum-line-chart>
  </el-card>
</template>

<script>
  import TalentSumLineChart from '../data_views/TalentSumLineChart.vue'

  export default {
    name: 'Talent',
    components: {
      TalentSumLineChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
