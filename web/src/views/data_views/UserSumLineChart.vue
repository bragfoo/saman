<template>
  <el-row>
    <ve-line :data="chartData" :settings="chartSettings"></ve-line>
  </el-row>
</template>

<script>
  export default {
    name: 'AppTalent',
    data () {
      return {
        loading: false,
        url: 'appData',
        chartData: {
          columns: ['CreateTime', 'ActiveUser'],
          rows: [
            {
              CreateTime: (new Date()).getTime(),
              Like: 0,
              ActiveUser: '用户总数'
            }
          ]
        },
        chartSettings: {
          labelMap: {
            CreateTime: '事件',
            ActiveUser: '用户总数'
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
          response.data.forEach((row) => {
            let time = new Date(row.CreateTime)
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
