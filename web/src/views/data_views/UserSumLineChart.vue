<template>
  <el-row>
    <ve-line :data="chartData" :settings="chartSettings" height="300px" :title="title"></ve-line>
    <el-col :span="24">
      <el-table
        :data="chartData.rows"
        heigth="300px"
        border
        style="width: 100%">
        <el-table-column
          fixed
          prop="CreateTime"
          label="日期">
        </el-table-column>
        <el-table-column
          prop="ActiveUser"
          label="活跃用户总数">
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'

  export default {
    name: 'UserSumLineChart',
    data () {
      return {
        loading: false,
        url: 'appData',
        chartData: {
          columns: ['CreateTime', 'ActiveUser'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            ActiveUser: '活跃用户总数'
          }
        },
        title: {
          text: '每日活跃用户'
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
