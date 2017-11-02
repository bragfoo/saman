<template>
    <el-card class="box-card">
    <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="平台粉丝数" name="modify_plat_fans"></el-tab-pane>
        <el-tab-pane label="平台视频" name="modify_video_plat"></el-tab-pane>
        <el-tab-pane label="视频播放量" name="modify_video_plat_amount"></el-tab-pane>
        <el-tab-pane label="移动数据" name="modify_mobile_data"></el-tab-pane>
        <el-tab-pane label="UGC" name="modify_app_ugc"></el-tab-pane>
        <el-tab-pane label="达人" name="modify_talent"></el-tab-pane>
        <el-tab-pane label="移动数据" name="modify_app_data"></el-tab-pane>
    </el-tabs>
    <router-view></router-view>
    </el-card>
</template>

<script>
export default {
  name: 'Main',
  data () {
    return {
      activeName: 'modify_plat_fans'
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
      if (name === 'modify') {
        this.activeName = 'modify_plat_fans'
        this.$router.replace({'name': 'modify_plat_fans'})
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
    font-weight: normal;
}

ul {
    list-style-type: none;
padding: 0;
}

li {
display: inline-block;
margin: 0 10px;
}

a {
color: #42b983;
}
</style>
