<template>
  <el-row>
    <ve-line :data="chartData" :settings="chartSettings"></ve-line>
  </el-row>
</template>

<script>
  export default {
    name: 'UGCPic',
    data () {
      return {
        loading: false,
        url: 'appUGC',
        chartData: {
          columns: ['CreateTime', 'PicSum'],
          rows: [
            {
              CreateTime: (new Date()).getTime(),
              PicSum: 0
            }
          ]
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
