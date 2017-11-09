<template>
  <el-card class="box-card">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="付费" name="modify_plat_fans"></el-tab-pane>
    </el-tabs>
    <router-view></router-view>
  </el-card>
</template>

<script>
  export default {
    name: 'Revenue',
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
