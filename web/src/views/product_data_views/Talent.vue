<template>
  <el-card class="box-card">
    <talent-sum-line-chart></talent-sum-line-chart>
    <talent-list-chart></talent-list-chart>
  </el-card>
</template>

<script>
  import TalentSumLineChart from '../data_views/TalentSumLineChart.vue'
  import TalentListChart from '../data_views/TalentListChart.vue'

  export default {
    name: 'Talent',
    components: {
      TalentSumLineChart,
      TalentListChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
