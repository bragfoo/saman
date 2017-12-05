<template>
  <el-card class="box-card">
    <el-row>
      <PlatTypeSelect v-model="platType"></PlatTypeSelect>
      <plat-play-sum-line-chart :platType="platType"></plat-play-sum-line-chart>
    </el-row>
    <el-row>
      <el-col :span="10" :offset="2">
        <date-type-select v-model="dateArr"></date-type-select>
      </el-col>
      <el-col :span="10" :offset="2">
        <VideoTypeSelect v-model="videoType" :platIds="platType"></VideoTypeSelect>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <plat-play-bar-chart :platType="platType" :start="start" :end="end"></plat-play-bar-chart>
      </el-col>
      <el-col :span="12">
        <plat-video-data-line-chart :platType="platType" :videoType="videoType"></plat-video-data-line-chart>
      </el-col>
    </el-row>
  </el-card>
</template>

<script>
  import PlatVideoDataLineChart from './data_views/PlatVideoPlayLineChart.vue'
  import PlatPlayBarChart from './data_views/PlatPlayBarChart.vue'
  import PlatPlaySumLineChart from './data_views/PlatPlaySumLineChart.vue'
  import PlatTypeSelect from '../components/PlatTypeSelect.vue'
  import DateTypeSelect from '../components/DateTypeSelect.vue'
  import VideoTypeSelect from '../components/VideoTypeSelect.vue'

  export default {
    name: 'Main',
    components: {
      PlatVideoDataLineChart,
      PlatPlayBarChart,
      PlatPlaySumLineChart,
      PlatTypeSelect,
      VideoTypeSelect,
      DateTypeSelect
    },
    data () {
      return {
        platType: '5a1690eeef2d1345d21327e9',
        videoType: '',
        dateArr: [new Date(new Date().getTime() - 3600 * 1000 * 24 * 7), new Date()]
      }
    },
    created () {
    },
    methods: {},
    computed: {
      start () {
        let date = this.dateArr[0]
        return date.getTime() / 1000
      },
      end () {
        let date = this.dateArr[1]
        return date.getTime() / 1000
      }
    }
  }
</script>
