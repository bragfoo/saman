<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix">
      <span>Dashboard</span>
      <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
    </div>
    <div class="echarts">
      <IEcharts :option="line" :loading="loading" @ready="onReady" @click="onClick"></IEcharts>
    </div>
    <div class="echarts">
      <IEcharts :option="line" :loading="loading" @ready="onReady" @click="onClick"></IEcharts>
    </div>
    <div class="echarts">
      <IEcharts :option="line" :loading="loading" @ready="onReady" @click="onClick"></IEcharts>
      <button @click="doRandom">Random</button>
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
    data: () => ({
      loading: true,
      bar: {
        title: {
          text: '测试'
        },
        tooltip: {},
        xAxis: {
          data: ['Shirt', 'Sweater', 'Chiffon Shirt', 'Pants', 'High Heels', 'Socks']
        },
        yAxis: {},
        series: [{
          name: 'Sales',
          type: 'bar',
          data: [5, 20, 36, 10, 10, 20]
        }]
      }
    }),
    methods: {
      doRandom () {
        const that = this
        let data = []
        for (let i = 0, min = 5, max = 99; i < 6; i++) {
          data.push(Math.floor(Math.random() * (max + 1 - min) + min))
        }
        that.loading = !that.loading
        that.bar.series[0].data = data
      },
      onReady (instance) {
        console.log(instance)
      },
      onClick (event, instance, echarts) {
        console.log(arguments)
      }
    }
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
