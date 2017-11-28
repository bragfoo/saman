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
    name: 'VideoTypeSelect',
    props: ['value', 'platIds'],
    data () {
      return {
        url: 'video',
        propValue: '59fad3eeef2d1314e0ea2a4e',
        options: [
          {value: '59fad3eeef2d1314e0ea2a4e', label: 'NoVideo'}
        ]
      }
    },
    created () {
      this.fetchList()
    },
    methods: {
      updateValue (value) {
        this.$emit('input', value)
        this.$emit('change', value)
      },
      fetchList () {
        this.$http.get(this.url, {
          params: {
            platIds: this.platIds
          }
        }).then((response) => {
          this.options = []
          response.data.forEach((row) => {
            this.options.push({value: row.Ids, label: row.Title})
          })
          if (response.data.length > 0) {
            this.propValue = this.options[0].value
          }
          this.updateValue(this.propValue)
        })
      }
    },
    watch: {
      'platIds': function (val, oldval) {
        this.fetchList()
      }
    }
  }
</script>
