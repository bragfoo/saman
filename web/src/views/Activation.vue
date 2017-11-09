<template>
  <el-card class="box-card">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="用户激活" name="modify_plat_fans"></el-tab-pane>
      <el-tab-pane label="移动数据" name="modify_mobile_data"></el-tab-pane>
      <el-tab-pane label="UGC" name="modify_app_ugc"></el-tab-pane>
    </el-tabs>
    <router-view></router-view>
  </el-card>
</template>

<script>
  export default {
    name: 'Activation',
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
