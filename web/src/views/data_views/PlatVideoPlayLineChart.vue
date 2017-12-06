<template>
  <el-row style="padding:10px">
    <ve-line :data="chartData" :settings="chartSettings" :title="title"></ve-line>
    <el-col :span="24">
      <el-table
        :data="chartData.rows"
        :height="300"
        border
        style="width: 100%">
        <el-table-column
          prop="CreateTime"
          label="日期">
        </el-table-column>
        <el-table-column
          prop="Sum"
          label="增长数">
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'

  export default {
    props: {
      start: {
        type: Number,
        default: -1
      },
      end: {
        type: Number,
        default: 0
      },
      platType: {
        type: String,
        default: '5a1690eeef2d1345d21327e9'
      },
      videoType: {
        type: String,
        default: ''
      }
    },
    name: 'PlatVideoPlayLineChart',
    data () {
      return {
        loading: false,
        url: 'playAmount',
        chartData: {
          columns: ['CreateTime', 'Sum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            Sum: '增长数'
          },
          label: {
            normal: {
              show: true
            }
          }
        },
        title: {
          text: '平台视频播放'
        }
      }
    },
    created () {
      this.fetchList()
    },
    methods: {
      // net io
      fetchList () {
        let params = {
          platIds: this.platType,
          videoIds: this.videoType
        }
        if (this.start >= 0) {
          params.start = this.start
          params.end = this.end
        }
        this.$http.get(this.url, {
          params: params
        }).then((response) => {
          this.chartData.rows = []
          let data = response.data === null ? response.data = [] : response.data
          if (data !== []) {
            data.sort((r1, r2) => (r1.CreateTime === r2.CreateTime ? 0 : r1.CreateTime > r2.CreateTime ? 1 : -1))
          }
          data.forEach((row) => {
            let time = new Date(row.CreateTime * 1000)
            row.CreateTime = (time.getMonth() + 1) + '月' + time.getDate() + '日'
            this.chartData.rows.push(row)
          })
        }).then()
      },
      reload () {
        this.fetchList()
      },
      watch_reload (val, oldval) {
        if (val !== oldval) {
          this.reload()
        }
      }
    },
    watch: {
      start: 'watch_reload',
      // end: 'watch_reload',
      // platType: 'watch_reload',
      videoType: 'watch_reload'
    }
  }
</script>
