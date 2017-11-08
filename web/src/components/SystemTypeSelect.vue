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
        propValue: 'ios',
        options: [
          {value: 'ios', label: 'ios'},
          {value: 'android', label: 'android'},
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
