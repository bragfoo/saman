<template>
  <el-card class="box-card">
    <PlatTypeSelect v-model="platType"></PlatTypeSelect>
    <date-type-select v-model="dateArr"></date-type-select>
    <plat-fans-chart :platType="platType" :start="start" :end="end"></plat-fans-chart>
  </el-card>
</template>

<script>
  import PlatFansChart from './data_views/PlatFansChart.vue'
  import PlatTypeSelect from '../components/PlatTypeSelect.vue'
  import DateTypeSelect from '../components/DateTypeSelect.vue'

  export default {
    name: 'Main',
    components: {
      PlatFansChart,
      PlatTypeSelect,
      DateTypeSelect
    },
    data () {
      return {
        platType: '5a1690eeef2d1345d21327e9',
        dateArr: [new Date(new Date().getTime() - 3600 * 1000 * 24 * 7), new Date()]
      }
    },
    created () {
    },
    methods: {},
    computed: {
      start () {
        let date = this.dateArr[0]
        return Math.floor(date.getTime() / 1000)
      },
      end () {
        let date = this.dateArr[1]
        return Math.floor(date.getTime() / 1000)
      }
    }
  }
</script>
