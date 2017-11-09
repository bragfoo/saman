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
          rows: [
            {'CreateTime': '1月1日', 'Decrease': 123, 'Increase': 3, 'Sum': 0, 'Type': ''}
          ]
        },
        chartSettings: {
          labelMap: {
            CreateTime: '事件',
            Decrease: '减少',
            Increase: '增加',
            Sum: '总数',
            Type: '平台'
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
          response.data.forEach((row) => {
            let time = new Date(row.CreateTime)
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
