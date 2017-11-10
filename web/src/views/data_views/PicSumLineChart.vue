<template>
  <el-row>
    <ve-line :data="chartData" :settings="chartSettings" height="300px"></ve-line>
  </el-row>
</template>

<script>
  export default {
    name: 'PicSumLineChart',
    data () {
      return {
        loading: false,
        url: 'appUGC',
        chartData: {
          columns: ['CreateTime', 'PicSum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            PicSum: '图片'
          }
        }
      }
    },
    created () {
      this.fetchList()
    },
    methods: {
      // net io
      fetchList () {
        this.$http.get(this.url, {
          params: {
            platIds: this.playType
          }
        }).then((response) => {
          this.chartData.rows = []
          let data = response.data === null ? response.data = [] : response.data
          data.sort((r1, r2) => (r1.CreateTime === r2.CreateTime ? 0 : r1.CreateTime > r2.CreateTime ? 1 : -1)
          let sum = 0
          data.forEach((row) => {
            let time = new Date(row.CreateTime * 1000)
            row.CreateTime = (time.getMonth() + 1) + '月' + time.getDate() + '日'
            sum += row.PicSum
            row.PicSum = sum
            this.chartData.rows.push(row)
          })
        }).then()
      },
      reload () {
        this.fetchList()
      }
    }
  }
</script>
