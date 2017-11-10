<template>
  <el-row>
    <ve-line :data="chartData" :settings="chartSettings"></ve-line>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'

  export default {
    name: 'TalentSumLineChart',
    data () {
      return {
        loading: false,
        url: 'appData',
        chartData: {
          columns: ['CreateTime', 'TalentSum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            TalentSum: '达人'
          }
        },
        title: {
          text: '每日数据'
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
          data.sort((r1, r2) => (r1.CreateTime === r2.CreateTime ? 0 : r1.CreateTime > r2.CreateTime ? 1 : -1))
          data.forEach((row) => {
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
