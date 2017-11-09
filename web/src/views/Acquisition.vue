<template>
  <el-card class="box-card">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="用户获取" name="acq_aummary"></el-tab-pane>
      <el-tab-pane label="平台粉丝数量" name="modify_plat_fans"></el-tab-pane>
      <el-tab-pane label="平台" name="modify_video_plat"></el-tab-pane>
      <el-tab-pane label="视频播放量" name="modify_video_plat_amount"></el-tab-pane>
    </el-tabs>
    <router-view></router-view>
  </el-card>
</template>

<script>
  export default {
    name: 'Acquisition',
    data () {
      return {
        activeName: 'acq_aummary'
      }
    },
    created () {
      this.parseRoute()
    },
    methods: {
      handleClick (tab, event) {
        this.$router.replace({'name': tab.name})
      },
      parseRoute () {
        let name = this.$route.name
        if (name === 'acquisition') {
          this.activeName = 'acq_aummary'
          this.$router.replace({'name': 'acq_aummary'})
        } else {
          this.activeName = name
        }
        console.log(this.$route)
      }
    },
    watch: {
      '$route': 'parseRoute'
    }
  }
</script>
