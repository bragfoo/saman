<template>
  <div>
    <el-row>
      <el-upload
        class="upload-demo"
        drag
        action="/upload/fans"
        multiple>
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      </el-upload>
    </el-row>
    <el-row>
      <el-upload
        class="upload-demo"
        drag
        action="/upload/video"
        multiple>
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      </el-upload>
    </el-row>
    <el-row>
      <el-upload
        class="upload-demo"
        drag
        action="/upload/newMedia"
        multiple>
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      </el-upload>
    </el-row>
  </div>
</template>

<script>
  export default {
    name: 'AppData',
    data () {
    },
    created () {
    },
    methods: {}
  }
</script>
