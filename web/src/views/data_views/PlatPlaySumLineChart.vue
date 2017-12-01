<template>
  <el-row style="padding:10px">
    <PlatTypeSelect v-model="platType" @change="fetchList"></PlatTypeSelect>
    <ve-line :data="chartData" :settings="chartSettings" height="300px" :title="title"></ve-line>
    <el-col :span="24">
      <el-table
        :data="chartData.rows"
        :height="300"
        border
        style="width: 100%">
        <el-table-column
          fixed
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
  import PlatTypeSelect from '../../components/PlatTypeSelect.vue'

  export default {
    components: {
      PlatTypeSelect
    },
    name: 'PlatPlaySumLineChart',
    data () {
      return {
        loading: false,
        url: 'liner/playAmount',
        platType: '5a1690eeef2d1345d21327e9',
        videoType: '',
        chartData: {
          columns: ['CreateTime', 'Grow'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '日期',
            Grow: '增长数'
          }
        },
        title: {
          text: '平台所有视频播放总数'
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
            platIds: this.platType
          }
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
      }
    }
  }
</script>
