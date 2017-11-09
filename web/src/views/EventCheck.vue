<template>
  <el-card class="box-card">
    <event-rardar-chart></event-rardar-chart>
  </el-card>
</template>

<script>
  import EventRardarChart from './data_views/EventRardarChart.vue'

  export default {
    name: 'Main',
    components: {
      EventRardarChart
    },
    data () {
      return {}
    },
    created () {},
    methods: {}
  }
</script>
