<template>
  <el-row>
    <ve-line :data="chartData" :settings="chartSettings"></ve-line>
  </el-row>
</template>

<script>
  export default {
    name: 'Share',
    data () {
      return {
        loading: false,
        url: 'appUGC',
        chartData: {
          columns: ['CreateTime', 'ShareSum'],
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
        this.$http.get(this.url).then((response) => {
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
</style>
