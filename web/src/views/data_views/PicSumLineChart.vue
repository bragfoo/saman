<template>
  <el-row>
    <ve-line :data="chartData" :settings="chartSettings"></ve-line>
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
          if (response.data === null) {
            response.data = []
          }
          response.data.sort((r1, r2) => {
            if (r1.CreateTime === r2.CreateTime) {
              return 0
            }
            return r1.CreateTime > r2.CreateTime ? 1 : -1
          })
          response.data.forEach((row) => {
            let time = new Date(row.CreateTime * 1000)
            row.CreateTime = (time.getMonth() + 1) + '月' + time.getDate() + '日'
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
