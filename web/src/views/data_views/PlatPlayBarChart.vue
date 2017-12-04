<template>
  <el-row style="padding:10px">
    <PlatTypeSelect v-model="platType" @change="fetchList"></PlatTypeSelect>
    <ve-histogram :data="chartData" :settings="chartSettings" height="300px"></ve-histogram>
    <el-col :span="24">
      <el-table
        :data="chartData.rows"
        :height="300"
        border
        style="width: 100%">
        <el-table-column
          fixed
          label="视频">
          <template slot-scope="scope">
            <a :href="scope.row.Title" target="_blank">
              <span style="margin-left: 10px">{{ scope.row.Title | shortText }}</span>
            </a>
          </template>
        </el-table-column>
        <el-table-column
          prop="Sum"
          label="本周播放量">
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'
  import PlatTypeSelect from '../../components/PlatTypeSelect.vue'
  import * as util from '../../util/filters'

  export default {
    components: {
      PlatTypeSelect
    },
    name: 'AppUGCChart',
    data () {
      return {
        loading: false,
        url: 'playAmount?chart=week',
        platType: '5a1690eeef2d1345d21327e9',
        videoType: '',
        chartData: {
          columns: ['ShortTitle', 'Sum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            ShortTitle: '视频',
            Title: '视频',
            Sum: '本周播放量'
          }
        },
        title: {
          text: '平台视频总数'
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
          data.forEach((row) => {
            let newRow = this.chartData.rows.find((r) => (r.Title === row.Title))
            if (!newRow) {
              newRow = {
                Title: row.Title,
                ShortTitle: util.shortTitle(row.Title),
                Sum: 0
              }
              this.chartData.rows.push(newRow)
            }
            if (newRow.Sum < row.Sum) {
              newRow.Sum = row.Sum
            }
          })

          if (this.chartData.rows !== []) {
            this.chartData.rows.sort((a, b) => a.Sum === b.Sum ? 0 : (a.Sum < b.Sum ? 1 : -1))
          }
        }).then()
      },
      reload () {
        this.fetchList()
      }
    }
  }
</script>
