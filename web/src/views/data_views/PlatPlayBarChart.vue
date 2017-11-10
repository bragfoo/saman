<template>
  <el-row>
    <PlatTypeSelect v-model="platType" @change="fetchList"></PlatTypeSelect>
    <ve-histogram :data="chartData" :settings="chartSettings" height="300px"></ve-histogram>
  </el-row>
</template>

<script>
  import 'echarts/lib/component/title'
  import PlatTypeSelect from '../../components/PlatTypeSelect.vue'

  export default {
    components: {
      PlatTypeSelect
    },
    name: 'AppUGCChart',
    data () {
      return {
        loading: false,
        url: 'playAmount',
        platType: '59fae20cef2d1314e0ea2a55',
        videoType: '',
        chartData: {
          columns: ['Title', 'Sum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            Title: '视频',
            Sum: '总数'
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
                Sum: 0
              }
              this.chartData.rows.push(newRow)
            }
            if (newRow.Sum < row.Sum) {
              newRow.Sum = row.Sum
            }
          })
          this.chartData.rows.sort((a, b) => {
            return a.Sum === b.Sum ? 0 : (a.Sum < b.Sum ? 1 : -1)
          })
        }).then()
      },
      reload () {
        this.fetchList()
      }
    }
  }
</script>
