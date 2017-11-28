<template>
  <el-row>
    <ve-line :data="chartData" :settings="chartSettings" height="300px" :title="title"></ve-line>
    <el-col :span="24">
      <el-table
        :data="chartData.rows"
        height="300px"
        border
        style="width: 100%">
        <el-table-column
          fixed
          prop="CreateTime"
          label="日期">
        </el-table-column>
        <el-table-column
          prop="Like"
          label="点赞">
        </el-table-column>
        <el-table-column
          prop="CommentSum"
          label="评论">
        </el-table-column>
        <el-table-column
          prop="ShareSum"
          label="分享">
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'

  export default {
    name: 'SocDataLineChart',
    data () {
      return {
        loading: false,
        url: 'appUGC',
        chartData: {
          columns: ['CreateTime', 'Like', 'CommentSum', 'ShareSum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            Like: '点赞',
            CommentSum: '评论',
            ShareSum: '分享'
          }
        },
        title: {
          text: '每日点赞数'
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
