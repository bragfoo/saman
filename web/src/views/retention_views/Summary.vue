<template>
  <el-card class="box-card">
    <app-u-g-c></app-u-g-c>
  </el-card>
</template>

<script>
  import AppUGC from '../data_views/AppUGCChart.vue'

  export default {
    name: 'RetSummary',
    components: {
      AppUGC
    },
    data () {
      return {}
    },
    created () {
      this.fetchList()
    },
    methods: {}
  }
</script>
