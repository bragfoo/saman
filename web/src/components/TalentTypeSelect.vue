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
    name: 'SystemTypeSelect',
    props: ['value'],
    data () {
      return {
        propValue: '5a055e14ef2d13f0054063ea',
        options: [
          {value: '5a055e14ef2d13f0054063ea', label: 'Swag girl'},
          {value: '5a055e22ef2d13f0054063eb', label: '潮人'},
          {value: '5a055e2def2d13f0054063ec', label: '街舞'},
          {value: '5a055e36ef2d13f0054063ed', label: '滑板'}
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
