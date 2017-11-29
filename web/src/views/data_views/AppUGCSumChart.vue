<template>
  <el-row>
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
        <el-table-column
          prop="PicSum"
          label="图片">
        </el-table-column>
        <el-table-column
          prop="VideoSum"
          label="视频">
        </el-table-column>
      </el-table>
    </el-col>
    <el-row>
      <el-col :span="8">
        <ve-line :data="chartLikeData" :settings="chartSettings" height="200px" :title="title"></ve-line>
      </el-col>
      <el-col :span="8">
        <ve-line :data="chartCommentData" :settings="chartSettings" height="200px" :title="title"></ve-line>
      </el-col>
      <el-col :span="8">
        <ve-line :data="chartShareData" :settings="chartSettings" height="200px" :title="title"></ve-line>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <ve-line :data="chartPicData" :settings="chartSettings" height="200px" :title="title"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartVideoData" :settings="chartSettings" height="200px" :title="title"></ve-line>
      </el-col>
    </el-row>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'

  export default {
    name: 'AppUGCChart',
    data () {
      return {
        loading: false,
        url: 'appUGCTotal',
        chartLikeData: {
          columns: ['CreateTime', 'Like'],
          rows: []
        },
        chartCommentData: {
          columns: ['CreateTime', 'CommentSum'],
          rows: []
        },
        chartShareData: {
          columns: ['CreateTime', 'ShareSum'],
          rows: []
        },
        chartPicData: {
          columns: ['CreateTime', 'PicSum'],
          rows: []
        },
        chartVideoData: {
          columns: ['CreateTime', 'VideoSum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            Like: '点赞',
            CommentSum: '评论',
            ShareSum: '分享',
            PicSum: '图片',
            VideoSum: '视频'
          }
        },
        title: {
          text: 'test'
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
