<template>
  <el-card class="box-card">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="图片发布" name="prd_pic_posted"></el-tab-pane>
      <el-tab-pane label="视频发布" name="prd_video_posted"></el-tab-pane>
      <el-tab-pane label="社会化" name="prd_socialization"></el-tab-pane>
      <el-tab-pane label="达人" name="prd_talent"></el-tab-pane>
      <el-tab-pane label="新增" name="prd_newly"></el-tab-pane>
    </el-tabs>
    <router-view></router-view>
  </el-card>
</template>

<script>
  export default {
    name: 'Acquisition',
    data () {
      return {
        activeName: 'prd_pic_posted'
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
          this.activeName = 'prd_pic_posted'
          this.$router.replace({'name': 'prd_pic_posted'})
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
