<template>
  <el-row style="padding:10px">
    <ve-histogram :data="chartData" :settings="chartSettings" height="300px"></ve-histogram>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'

  export default {
    name: 'TalentSumLineChart',
    data () {
      return {
        loading: false,
        url: 'talent',
        chartData: {
          columns: ['SkillName', 'Sub', 'UnSub'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            SkillName: '分类',
            Sub: '推荐人数',
            UnSub: '未推荐人数'
          }
        },
        title: {
          text: '达人统计数据'
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
          if (data !== []) {
            data.sort((r1, r2) => (r1.CreateTime === r2.CreateTime ? 0 : r1.CreateTime > r2.CreateTime ? 1 : -1))
          }
          data.forEach((row) => {
            let newRow = this.chartData.rows.find((r) => (r.SkillName === row.SkillName))
            if (!newRow) {
              newRow = {
                SkillName: row.SkillName,
                Sub: 0,
                UnSub: 0
              }
              this.chartData.rows.push(newRow)
            }
            if (row.Submitted) {
              newRow.Sub++
            } else {
              newRow.UnSub++
            }
          })
        }).then()
      },
      reload () {
        this.fetchList()
      }
    }
  }
</script>
