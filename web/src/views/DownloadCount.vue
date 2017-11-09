<template>
  <el-card class="box-card">
    <plat-fans-chart></plat-fans-chart>
  </el-card>
</template>

<script>
  import PlatFansChart from './data_views/PlatFansChart.vue'

  export default {
    name: 'Main',
    components: {
      PlatFansChart
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h1, h2 {
    font-weight: normal;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #42b983;
  }
</style>
