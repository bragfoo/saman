<template>
  <el-card class="box-card">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="图片发布" name="useref_download"></el-tab-pane>
      <el-tab-pane label="视频发布" name="useref_retention"></el-tab-pane>
    </el-tabs>
    <router-view></router-view>
  </el-card>
</template>

<script>
  export default {
    name: 'UserReferral',
    data () {
      return {
        activeName: 'useref_download'
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
        if (name === 'user_referral') {
          this.activeName = 'useref_download'
          this.$router.replace({'name': 'useref_download'})
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
