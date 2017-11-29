<template>
  <el-row :style="padding=10">
    <ve-radar :data="chartData" :settings="chartSettings" height="300px" :title="title"></ve-radar>
    <el-col :span="24">
      <el-table
        :data="chartData.rows"
        :height="300"
        border
        style="width: 100%">
        <el-table-column
          fixed
          prop="名称"
          label="名称">
        </el-table-column>
        <el-table-column
          prop="参加人数"
          label="参加人数">
        </el-table-column>
        <el-table-column
          prop="工作人数"
          label="工作人数">
        </el-table-column>
        <el-table-column
          prop="上传人数"
          label="上传人数">
        </el-table-column>
        <el-table-column
          prop="开始"
          label="开始">
        </el-table-column>
        <el-table-column
          prop="结束"
          label="结束">
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'

  export default {
    name: 'AppUGCChart',
    data () {
      return {
        loading: false,
        url: 'event',
        chartData: {
          columns: ['名称', '参加人数', '工作人数', '上传人数', '开始', '结束'],
          rows: []
        },
        chartSettings: {
          dimension: ['名称'],
          metrics: ['参加人数', '工作人数', '上传人数'],
//          dataType: {'参加人数': 'normal'},
          label: {
            normal: {
              position: 'top',
              show: false,
              formatter: function (params) {
                console.log(params)
              }
            }
          }
        },
        dataZoom: [
          {
            type: 'slider',
            start: 0,
            end: 20
          }
        ],
        title: {
          text: '活动检测'
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
          let data = response.data === null ? [] : response.data
          data.sort((r1, r2) => (r1.CreateTime === r2.CreateTime ? 0 : r1.CreateTime > r2.CreateTime ? 1 : -1))
          data.forEach((row) => {
            let time = new Date(row.StartDate * 1000)
            row.StartDate = (time.getMonth() + 1) + '月' + time.getDate() + '日'
            time = new Date(row.EndDate * 1000)
            row.EndDate = (time.getMonth() + 1) + '月' + time.getDate() + '日'
            this.chartData.rows.push({
              // CreateTime: '时间',
              '名称': row.Name,
              '参加人数': row.TotalPeople,
              '工作人数': row.TotalWork,
              '上传人数': row.UploadPeople,
              '开始': row.StartDate,
              '结束': row.EndDate
            })
          })
        }).then()
      },
      reload () {
        this.fetchList()
      }
    }
  }
</script>
