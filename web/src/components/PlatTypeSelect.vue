<template>
  <el-select v-model='propValue' @change="updateValue" placeholder="请选择">
    <el-option
      v-for="item in options"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    >
    </el-option>
  </el-select>
</template>


<script>
  export default {
    name: 'PlatTypeSelect',
    props: ['value'],
    data () {
      return {
        propValue: '59fae20cef2d1314e0ea2a55',
        options: [
          {value: '59fae20cef2d1314e0ea2a55', label: '头条音乐'},
          {value: '59fae276ef2d1314e0ea2a56', label: '头条体育'},
          {value: '59fae294ef2d1314e0ea2a57', label: '腾讯音乐'},
          {value: '59fae2a8ef2d1314e0ea2a58', label: '腾讯体育'},
          {value: '59fae2beef2d1314e0ea2a59', label: '公众空间'},
          {value: '59fae2d4ef2d1314e0ea2a5a', label: 'A站'},
          {value: '59fae2ecef2d1314e0ea2a5b', label: 'B站'},
          {value: '59fae313ef2d1314e0ea2a5c', label: '爱奇艺'},
          {value: '59fae32cef2d1314e0ea2a5d', label: '优酷'},
          {value: '59fae33fef2d1314e0ea2a5e', label: '斗鱼'},
          {value: '59fae351ef2d1314e0ea2a5f', label: '人人'}
        ]
      }
    },
    methods: {
      updateValue (value) {
        this.$emit('input', value)
        this.$emit('change', value)
      }
    }
  }
</script>
