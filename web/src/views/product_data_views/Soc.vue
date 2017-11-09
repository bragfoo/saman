<template>
  <el-card class="box-card">
    <soc-data-line-chart></soc-data-line-chart>
  </el-card>
</template>

<script>
  import SocDataLineChart from '../data_views/SocDataLineChart.vue'

  export default {
    name: 'Soc',
    components: {
      SocDataLineChart
    },
    data () {
      return {}
    },
    created () {},
    methods: {}
  }
</script>
