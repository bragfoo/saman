<template>
  <div>
    <el-dialog
      title="请先登录"
      :visible.sync="dialogVisible"
      width="30%">
      <el-form ref="form" :model="form" label-width="80px">
        <el-form-item label="用户名">
          <el-input v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input type="password" v-model="form.passwd"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">登录</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-container id="app">
      <el-aside width="65px">
        <el-card style="border:0" :body-style="{ padding: '0px' }" v-show="auth">
          <el-menu default-active="1-4-1" :router=true :collapse="true">
            <el-menu-item index="0" :route="{'name':'dashboard'}">
              <i class="el-icon-menu"></i>
              <span slot="title">Dashboard</span>
            </el-menu-item>

            <!--new menu -->
            <!--<el-menu-item index="1" :route="{'name':'operation_weekly'}">-->
            <!--<i class="el-icon-date"></i>-->
            <!--<span slot="title">运营周报</span>-->
            <!--</el-menu-item>-->
            <!--<el-menu-item index="2" :route="{'name':'users_portrayal'}">-->
            <!--<i class="el-icon-edit-outline"></i>-->
            <!--<span slot="title">用户画像</span>-->
            <!--</el-menu-item>-->
            <el-menu-item index="3" :route="{'name':'video_click'}">
              <i class="el-icon-plus"></i>
              <span slot="title">视频点击</span>
            </el-menu-item>
            <el-menu-item index="4" :route="{'name':'put_plat'}">
              <i class="el-icon-upload"></i>
              <span slot="title">投放渠道</span>
            </el-menu-item>
            <el-menu-item index="5" :route="{'name':'product_data'}">
              <i class="el-icon-mobile-phone"></i>
              <span slot="title">产品数据</span>
            </el-menu-item>
            <el-menu-item index="6" :route="{'name':'download_count'}">
              <i class="el-icon-download"></i>
              <span slot="title">下载量</span>
            </el-menu-item>
            <el-menu-item index="7" :route="{'name':'user_referral'}">
              <i class="el-icon-sort"></i>
              <span slot="title">用户转化</span>
            </el-menu-item>
            <el-menu-item index="8" :route="{'name':'event_check'}">
              <i class="el-icon-circle-check"></i>
              <span slot="title">活动检测</span>
            </el-menu-item>
            <!--<el-menu-item index="9" :route="{'name':'weekly'}">-->
            <!--<i class="el-icon-document"></i>-->
            <!--<span slot="title">周报</span>-->
            <!--</el-menu-item>-->
            <!--new menu end-->

            <!--old menu-->
            <!--
            <el-menu-item index="10" :route="{'name':'acquisition'}">
              <i class="el-icon-news"></i>
              <span slot="title">获取</span>
            </el-menu-item>
            <el-menu-item index="11" :route="{'name':'activation'}">
              <i class="el-icon-mobile-phone"></i>
              <span slot="title">激活</span>
            </el-menu-item>
            <el-menu-item index="12" :route="{'name':'retention'}">
              <i class="el-icon-view"></i>
              <span slot="title">留存</span>
            </el-menu-item>
            <el-menu-item index="13" :route="{'name':'referral'}">
              <i class="el-icon-share"></i>
              <span slot="title">转化</span>
            </el-menu-item>
            <el-menu-item index="14" :route="{'name':'revenue'}">
              <i class="el-icon-sold-out"></i>
              <span slot="title">付费</span>
            </el-menu-item>
            -->
            <el-menu-item index="15" :route="{'name':'modify'}">
              <i class="el-icon-setting"></i>
              <span slot="title">数据修改</span>
            </el-menu-item>
            <!--old menu end-->
          </el-menu>
        </el-card>
      </el-aside>
      <el-container v-show="auth">
        <!--<el-header>Header</el-header>-->
        <el-main>
          <router-view/>
        </el-main>
        <!--<el-footer>Footer</el-footer>-->
      </el-container>
    </el-container>
  </div>
</template>

<script>
  export default {
    name: 'app',
    data () {
      return {
        dialogVisible: true,
        auth: false,
        form: {
          username: '',
          passwd: ''
        }
      }
    },
    methods: {
      onSubmit () {
        if (this.form.username === 'admin' && this.form.passwd === 'woaixuexi') {
          this.dialogVisible = false
          this.auth = true
        } else {

        }
      }
    }
  }
</script>

<style>
  body {
    margin: 0
  }
</style>
