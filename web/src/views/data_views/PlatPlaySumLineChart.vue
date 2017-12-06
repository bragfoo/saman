<template>
  <el-row style="padding:10px">
    <ve-line :data="chartData" :settings="chartSettings" height="300px" :title="title"></ve-line>
    <el-col :span="24">
      <el-table
        :data="chartData.rows"
        :height="300"
        border
        style="width: 100%">
        <el-table-column
          prop="CreateTime"
          label="时间">
        </el-table-column>
        <el-table-column
          prop="Grow"
          label="增长数">
        </el-table-column>
        <!--<el-table-column-->
        <!--prop="Sum"-->
        <!--label="总数">-->
        <!--</el-table-column>-->
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
      }
    },
    name: 'PlatPlaySumLineChart',
    data () {
      return {
        loading: false,
        url: 'liner/playAmount',
        chartData: {
          columns: ['CreateTime', 'Grow'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            Grow: '增长数'
          }
        },
        title: {
          text: '每周播放量增长趋势'
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
          platIds: this.platType
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
          if (data === []) {
            return
          }
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
      platType: 'watch_reload'
    }
  }
</script>
