<template>
  <div>
    <el-row style="padding:10px;border-bottom: 1px solid #e6ebf5;">
      <el-col :span="24">
        <el-table
          :data="chartData.rows"
          :height="300"
          border
          style="width: 100%">
          <el-table-column
            fixed
            prop="CreateTime"
            label="日期">
          </el-table-column>
          <el-table-column
            prop="Like"
            label="点赞数">
          </el-table-column>
          <el-table-column
            prop="CommentSum"
            label="评论数">
          </el-table-column>
          <el-table-column
            prop="ShareSum"
            label="分享数">
          </el-table-column>
          <el-table-column
            prop="PicSum"
            label="图片数">
          </el-table-column>
          <el-table-column
            prop="VideoSum"
            label="视频点击数">
          </el-table-column>
          <el-table-column
            prop="VideoStay"
            label="视频播放时长">
          </el-table-column>
          <el-table-column
            prop="VideoUpload"
            label="视频上传数">
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
    <el-row style="padding:10px;border-bottom: 1px solid #e6ebf5;">
      <el-col :span="12">
        <ve-line :data="chartLikeGrowData" :settings="chartSettings" height="200px" :title="{
          text: '点赞增长数 本周-上周'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartLikeData" :settings="chartSettings" height="200px" :title="{
          text: '点赞数'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartCommentGrowData" :settings="chartSettings" height="200px" :title="{
          text: '评论增长数 本周-上周'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartCommentData" :settings="chartSettings" height="200px" :title="{
          text: '评论数'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartShareGrowData" :settings="chartSettings" height="200px" :title="{
          text: '分享增长数 本周-上周'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartShareData" :settings="chartSettings" height="200px" :title="{
          text: '分享数'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartPicGrowData" :settings="chartSettings" height="200px" :title="{
          text: '图片上传增长数 本周-上周'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartPicData" :settings="chartSettings" height="200px" :title="{
          text: '图片上传数'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartVideoUploadGrowData" :settings="chartSettings" height="200px" :title="{
          text: '视频上传数增长 本周-上周'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartVideoUploadData" :settings="chartSettings" height="200px" :title="{
          text: '视频上传数'
        }"></ve-line>
      </el-col>
    </el-row>
    <el-row style="padding:10px;border-bottom: 1px solid #e6ebf5;">
      <el-col :span="12">
        <ve-line :data="chartVideoGrowData" :settings="chartSettings" height="200px" :title="{
          text: '视频播放增长数'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartVideoData" :settings="chartSettings" height="200px" :title="{
          text: '视频播放数'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartVideoStayGrowData" :settings="chartSettings" height="200px" :title="{
          text: '视频播放时长增长'
        }"></ve-line>
      </el-col>
      <el-col :span="12">
        <ve-line :data="chartVideoStayData" :settings="chartSettings" height="200px" :title="{
          text: '视频播放总时长'
        }"></ve-line>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import 'echarts/lib/component/title'

  export default {
    name: 'AppUGCChart',
    data () {
      return {
        loading: false,
        url: 'appUGCTotal',
        chartData: {
          columns: ['CreateTime', 'Like'],
          rows: []
        },
        chartLikeData: {
          columns: ['CreateTime', 'Like'],
          rows: []
        },
        chartLikeGrowData: {
          columns: ['CreateTime', 'LikeGrow'],
          rows: []
        },
        chartCommentData: {
          columns: ['CreateTime', 'CommentSum'],
          rows: []
        },
        chartCommentGrowData: {
          columns: ['CreateTime', 'CommentGrow'],
          rows: []
        },
        chartShareData: {
          columns: ['CreateTime', 'ShareSum'],
          rows: []
        },
        chartShareGrowData: {
          columns: ['CreateTime', 'ShareGrow'],
          rows: []
        },
        chartPicData: {
          columns: ['CreateTime', 'PicSum'],
          rows: []
        },
        chartPicGrowData: {
          columns: ['CreateTime', 'PicGrow'],
          rows: []
        },
        chartVideoData: {
          columns: ['CreateTime', 'VideoSum'],
          rows: []
        },
        chartVideoGrowData: {
          columns: ['CreateTime', 'VideoGrow'],
          rows: []
        },
        chartVideoStayData: {
          columns: ['CreateTime', 'VideoStay'],
          rows: []
        },
        chartVideoStayGrowData: {
          columns: ['CreateTime', 'VideoGrow'],
          rows: []
        },
        chartVideoUploadData: {
          columns: ['CreateTime', 'VideoUpload'],
          rows: []
        },
        chartVideoUploadGrowData: {
          columns: ['CreateTime', 'VideoUploadGrow'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            Like: '点赞',
            LikeGrow: '点赞增长',
            CommentSum: '评论',
            CommentGrow: '评论增长',
            ShareSum: '分享',
            ShareGrow: '分享增长',
            PicSum: '图片',
            PicGrow: '图片增长',
            VideoSum: '视频',
            VideoGrow: '视频增长',
            VideoStay: '视频播放总时长',
            VideoStayGrow: '视频播放时长增长',
            VideoUpload: '视频上传数',
            VideoUploadGrow: '视频上传数增长'
          },
          label: {
            normal: {
              show: true
            }
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
          if (data !== []) {
            data.sort((r1, r2) => (r1.CreateTime === r2.CreateTime ? 0 : r1.CreateTime > r2.CreateTime ? 1 : -1))
          }
          let grows = [
            {
              'grow': 'LikeGrow',
              'sum': 'Like',
              'growTitle': 'chartLikeGrowData',
              'title': 'chartLikeData',
              'week': 0
            },
            {
              'grow': 'CommentGrow',
              'sum': 'CommentSum',
              'growTitle': 'chartCommentGrowData',
              'title': 'chartCommentData',
              'week': 0
            },
            {
              'grow': 'ShareGrow',
              'sum': 'ShareSum',
              'growTitle': 'chartShareGrowData',
              'title': 'chartShareData',
              'week': 0
            },
            {
              'grow': 'PicGrow',
              'sum': 'PicSum',
              'growTitle': 'chartPicGrowData',
              'title': 'chartPicData',
              'week': 0
            },
            {
              'grow': 'VideoGrow',
              'sum': 'VideoSum',
              'growTitle': 'chartVideoGrowData',
              'title': 'chartVideoData',
              'week': 0
            },
            {
              'grow': 'VideoStayGrow',
              'sum': 'VideoStay',
              'growTitle': 'chartVideoStayGrowData',
              'title': 'chartVideoStayData',
              'week': 0
            },
            {
              'grow': 'VideoUploadGrow',
              'sum': 'VideoUpload',
              'growTitle': 'chartVideoUploadGrowData',
              'title': 'chartVideoUploadData',
              'week': 0
            }
          ]
          data.forEach(row => {
            let time = new Date(row.CreateTime * 1000)
            if (time.getDay() === 0) {
              grows.forEach(grow => {
                row[grow.grow] = row[grow.sum] - grow.week
                grow.week = row[grow.sum]
                this[grow.growTitle].rows.push(row)
              })
            }
            row.CreateTime = (time.getMonth() + 1) + '月' + time.getDate() + '日'
            console.log(row)
            this.chartData.rows.push(row)
          })
          grows.forEach(grow => {
            this[grow.title].rows = this.chartData.rows
          })
        }).then()
      },
      reload () {
        this.fetchList()
      }
    }
  }
</script>
