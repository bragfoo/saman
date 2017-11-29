<template>
  <div>
    <el-row>
      <el-upload
        class="upload-demo"
        drag
        action="/api/v1/upload/fans"
        multiple>
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <div class="el-upload__tip">fans</div>
      </el-upload>
    </el-row>
    <el-row>
      <el-upload
        class="upload-demo"
        drag
        action="/api/v1/upload/video"
        multiple>
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <div class="el-upload__tip">video</div>
      </el-upload>
    </el-row>
    <el-row>
      <el-upload
        class="upload-demo"
        drag
        action="/api/v1/upload/newMedia"
        multiple>
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <div class="el-upload__tip">newMedia</div>
      </el-upload>
    </el-row>
  </div>
</template>

<script>
  export default {
    name: 'AppData',
    data () {
      return {}
    },
    created () {
    },
    methods: {}
  }
</script>
