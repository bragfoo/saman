<template>
  <el-card class="box-card">
    <div class="echarts">
      <ve-line :data="chartData" :settings="chartSettings"></ve-line>
    </div>
  </el-card>
</template>


<script>
  import IEcharts from 'vue-echarts-v3/src/full.vue'

  export default {
    name: 'Main',
    components: {
      IEcharts
    },
    props: {},
    data: () => (
      {
        chartData: {
          columns: ['日期', '成本', '利润', '占比', '其他'],
          rows: [
            {'成本': 1523, '日期': '1月1日', '利润': 1523, '占比': 0.12, '其他': 100},
            {'成本': 1223, '日期': '1月2日', '利润': 1523, '占比': 0.345, '其他': 100},
            {'成本': 2123, '日期': '1月3日', '利润': 1523, '占比': 0.7, '其他': 100},
            {'成本': 4123, '日期': '1月4日', '利润': 1523, '占比': 0.31, '其他': 100},
            {'成本': 3123, '日期': '1月5日', '利润': 1523, '占比': 0.12, '其他': 100},
            {'成本': 7123, '日期': '1月6日', '利润': 1523, '占比': 0.65, '其他': 100}
          ]
        },
        chartSettings: {
          metrics: ['成本', '利润'],
          dimension: ['日期']
        }
      }),
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

  .echarts {
    width: 1000px;
    height: 200px;
  }
</style>
