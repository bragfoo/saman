<template>
  <div id="app">
    <div class="menu">
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
    </div>
    <div class="main">
      <router-view/>
    </div>
  </div>
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

  .menu {
    float: left;
  }

  .main {
    float: left;
    width: 100%;
    padding: 10px
  }

  #app {
    display: flex;
    flex-direction: row;
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    /*margin-top: 60px;*/
  }
</style>
