<template>
  <el-card class="box-card">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="一周数据分析" name="prd_summary_weekly"></el-tab-pane>
      <el-tab-pane label="总数据变量" name="prd_summary_sum"></el-tab-pane>
      <el-tab-pane label="图片发布" name="prd_pic_posted"></el-tab-pane>
      <el-tab-pane label="视频发布" name="prd_video_posted"></el-tab-pane>
      <el-tab-pane label="社会化" name="prd_socialization"></el-tab-pane>
      <el-tab-pane label="达人" name="prd_talent"></el-tab-pane>
      <el-tab-pane label="新增用户" name="prd_newly"></el-tab-pane>
    </el-tabs>
    <router-view></router-view>
  </el-card>
</template>

<script>
  export default {
    name: 'Acquisition',
    data () {
      return {
        activeName: 'prd_summary_weekly'
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
        if (name === 'product_data') {
          this.activeName = 'prd_summary_weekly'
          this.$router.replace({'name': 'prd_summary_weekly'})
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
