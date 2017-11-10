<template>
  <el-card class="box-card">
    <talent-sum-line-chart></talent-sum-line-chart>
    <el-row>
      <el-col :span="12">
        <talent-list-chart></talent-list-chart>
      </el-col>
      <el-col :span="12">
        <talent-bar-chart></talent-bar-chart>
      </el-col>
    </el-row>
  </el-card>
</template>

<script>
  import TalentSumLineChart from '../data_views/TalentSumLineChart.vue'
  import TalentListChart from '../data_views/TalentListChart.vue'
  import TalentBarChart from '../data_views/TalentBarChart.vue'

  export default {
    name: 'Talent',
    components: {
      TalentSumLineChart,
      TalentListChart,
      TalentBarChart
    },
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
