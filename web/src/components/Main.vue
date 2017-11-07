<template>
  <el-card class="box-card">
    <plat-fans-chart></plat-fans-chart>
  </el-card>
</template>


<script>
  import PlatFansChart from '../views/data_views/PlatFansChart.vue'

  export default {
    name: 'Main',
    components: {
      PlatFansChart
    },
    props: {},
    data: () => ({}),
    methods: {}
  }
</script>
