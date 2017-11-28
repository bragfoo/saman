<template>
  <el-row>
    <PlatTypeSelect v-model="platType" @change="fetchList"></PlatTypeSelect>
    <VideoTypeSelect v-model="videoType" @change="fetchList" :platIds="platType"></VideoTypeSelect>
    <ve-line :data="chartData" :settings="chartSettings" :title="title"></ve-line>
    <el-col :span="24">
      <el-table
        :data="chartData.rows"
        height="300px"
        border
        style="width: 100%">
        <el-table-column
          fixed
          prop="CreateTime"
          label="日期">
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
  import 'echarts/lib/component/title'
  import PlatTypeSelect from '../../components/PlatTypeSelect.vue'
  import VideoTypeSelect from '../../components/VideoTypeSelect.vue'

  export default {
    components: {
      PlatTypeSelect,
      VideoTypeSelect
    },
    name: 'AppUGCChart',
    data () {
      return {
        loading: false,
        url: 'playAmount',
        platType: '59fae20cef2d1314e0ea2a55',
        videoType: '',
        chartData: {
          columns: ['CreateTime', 'Sum'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            CreateTime: '时间',
            Sum: '总数'
          }
        },
        title: {
          text: '平台视频播放'
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
            platIds: this.playType,
            videoIds: this.videoType
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
