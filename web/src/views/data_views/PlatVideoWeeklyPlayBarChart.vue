<template>
  <el-row style="padding:10px">
    <ve-histogram :data="chartData" :settings="chartSettings" :title="title"></ve-histogram>
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
      }
    },
    name: 'PlatVideoPlayWeeklyBarChart',
    data () {
      return {
        loading: false,
        url: 'weeklyPlay',
        chartData: {
          columns: ['PlatName', 'Sum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            PlatName: '平台名',
            Sum: '增长数'
          },
          label: {
            normal: {
              show: true
            }
          }
        },
        title: {
          text: '平台视频播放每周增量'
        }
      }
    },
    created () {
      this.fetchList()
    },
    methods: {
      // net io
      fetchList () {
        let params = {}
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
            data.sort((r1, r2) => (r1.PlatIds === r2.PlatIds ? 0 : r1.PlatIds > r2.PlatIds ? 1 : -1))
          }
          data.forEach((row) => {
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
      start: 'watch_reload'
      // end: 'watch_reload',
      // platType: 'watch_reload',
      // videoType: 'watch_reload'
    }
  }
</script>
