<template>
  <el-container id="app">
    <el-aside width="65px">
      <el-card style="border:0" :body-style="{ padding: '0px' }">
        <el-menu default-active="1-4-1" @open="handleOpen" @close="handleClose" :router=true
                 :collapse="true">
          <el-menu-item index="0" :route="{'name':'dashboard'}">
            <i class="el-icon-menu"></i>
            <span slot="title">Dashboard</span>
          </el-menu-item>
          <el-menu-item index="1" :route="{'name':'gain'}">
            <i class="el-icon-news"></i>
            <span slot="title">获取</span>
          </el-menu-item>
          <el-menu-item index="2" :route="{'name':'dashboard'}">
            <i class="el-icon-mobile-phone"></i>
            <span slot="title">激活</span>
          </el-menu-item>
          <el-menu-item index="3" :route="{'name':'dashboard'}">
            <i class="el-icon-view"></i>
            <span slot="title">留存</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-share" :route="{'name':'dashboard'}"></i>
            <span slot="title">转化</span>
          </el-menu-item>
          <el-menu-item index="5" :route="{'name':'dashboard'}">
            <i class="el-icon-sold-out"></i>
            <span slot="title">付费</span>
          </el-menu-item>
          <el-menu-item index="6" :route="{'name':'modify'}">
            <i class="el-icon-setting"></i>
            <span slot="title">数据修改</span>
          </el-menu-item>
        </el-menu>
      </el-card>
    </el-aside>
    <el-container>
      <!--<el-header>Header</el-header>-->
      <el-main>
        <router-view/>
      </el-main>
      <!--<el-footer>Footer</el-footer>-->
    </el-container>
  </el-container>
</template>

<script>
  export default {
    name: 'app',
    methods: {
      handleOpen (key, keyPath) {
        console.log(key, keyPath)
      },
      handleClose (key, keyPath) {
        console.log(key, keyPath)
      }
    }
  }
</script>

<style>
  body {
    margin: 0
  }

</style>
