<template>
  <el-row>
    <ve-line :data="chartData" :settings="chartSettings"></ve-line>
  </el-row>
</template>

<script>
  export default {
    name: 'AppUGC',
    data () {
      return {
        loading: false,
        url: 'appUGC',
        chartData: {
          columns: ['CreateTime', 'Like', 'CommentSum', 'ShareSum', 'PicSum', 'VideoSum'],
          rows: [
            {
              CreateTime: (new Date()).getTime(),
              Like: 0,
              CommentSum: 0,
              ShareSum: 0,
              PicSum: 0,
              VideoSum: 0
            }
          ]
        },
        chartSettings: {
          labelMap: {
            CreateTime: '事件',
            Like: '点赞',
            CommentSum: '评论',
            ShareSum: '分享',
            PicSum: '图片',
            VideoSum: '视频'
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
