<template>
  <el-card class="box-card">
    <UserSumLineChart></UserSumLineChart>
  </el-card>
</template>

<script>
  import UserSumLineChart from '../data_views/UserSumLineChart.vue'

  export default {
    name: 'Newly',
    components: {
      UserSumLineChart
    },
    data () {
      return {}
    },
    created () {},
    methods: {}
  }
</script>
