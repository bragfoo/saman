<template>
  <el-row style="padding:10px">
    <ve-line :data="chartData" :settings="chartSettings" :title="title"></ve-line>
    <ve-line :data="netIncreaseData" :settings="netIncreaseSettings" :title="title"></ve-line>
    <el-col :span="24">
      <el-table
        :data="chartData.rows"
        :height="300"
        border
        show-summary
        style="width: 100%">
        <el-table-column
          fixed
          prop="CreateTime"
          label="日期">
        </el-table-column>
        <el-table-column
          prop="Decrease"
          label="减少">
        </el-table-column>
        <el-table-column
          prop="Increase"
          label="增加">
        </el-table-column>
        <el-table-column
          prop="Sum"
          label="总数">
        </el-table-column>
        <el-table-column
          prop="NetIncrease"
          label="净增">
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'

  export default {
    name: 'PlatFansChart',
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
    data () {
      return {
        loading: false,
        url: 'platformFans',
        chartData: {
          columns: ['CreateTime', 'Sum'],
          rows: []
        },
        netIncreaseData: {
          columns: ['CreateTime', 'NetIncrease'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            Decrease: '减少',
            Increase: '增加',
            NetIncrease: '净增',
            Sum: '总数'
          },
          label: {
            normal: {
              show: true
            }
          }
        },
        netIncreaseSettings: {
          labelMap: {
            CreateTime: '时间',
            NetIncrease: '净增'
          },
          label: {
            normal: {
              show: true
            }
          }
        },
        title: {
          text: '平台粉丝数'
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
          this.netIncreaseData.rows = []
          let data = response.data === null ? response.data = [] : response.data
          if (data !== []) {
            data.sort((r1, r2) => (r1.CreateTime === r2.CreateTime ? 0 : r1.CreateTime > r2.CreateTime ? 1 : -1))
          }
          data.forEach((row) => {
            let time = new Date(row.CreateTime * 1000)
            row.CreateTime = (time.getMonth() + 1) + '月' + time.getDate() + '日'
            row.NetIncrease = row.Increase - row.Decrease
            this.chartData.rows.push(row)
            this.netIncreaseData.rows.push(row)
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
