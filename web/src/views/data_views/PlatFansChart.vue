<template>
  <el-row>
    <PlatTypeSelect v-model="playType" @change="fetchList"></PlatTypeSelect>
    <ve-line :data="chartData" :settings="chartSettings"></ve-line>
  </el-row>
</template>

<script>
  import PlatTypeSelect from '../../components/PlatTypeSelect.vue'

  export default {
    name: 'PlatFansChart',
    components: {PlatTypeSelect},
    data () {
      return {
        loading: false,
        url: 'platformFans',
        playType: '59fae20cef2d1314e0ea2a55',
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
          data.sort((r1, r2) => (r1.CreateTime === r2.CreateTime ? 0 : r1.CreateTime > r2.CreateTime ? 1 : -1)
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
