<template>
  <el-row>
    <PlatTypeSelect v-model="playType" @change="fetchList"></PlatTypeSelect>
    <ve-line :data="chartData" :settings="chartSettings" :title="title"></ve-line>
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
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
  import PlatTypeSelect from '../../components/PlatTypeSelect.vue'
  import 'echarts/lib/component/title'

  export default {
    name: 'PlatFansChart',
    components: {PlatTypeSelect},
    data () {
      return {
        loading: false,
        url: 'platformFans',
        playType: '5a1690eeef2d1345d21327e9',
        chartData: {
          columns: ['CreateTime', 'Decrease', 'Increase', 'Sum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            Decrease: '减少',
            Increase: '增加',
            Sum: '总数'
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
