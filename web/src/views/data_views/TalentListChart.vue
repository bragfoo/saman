<template>
  <el-row>
    <el-table
      :data="tableData"
      height="300"
      border
      style="width: 100%">
      <el-table-column
        prop="User"
        label="用户"
        width="120">
      </el-table-column>
      <el-table-column
        prop="Submitted"
        label="状态"
        width="120">
        <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ scope.row.Submitted | talentStatus}}</span>
        </template>
      </el-table-column>
      <el-table-column
        fixed
        label="日期"
        width="180">
        <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ scope.row.CreateTime | timeToDay}}</span>
        </template>
      </el-table-column>
      <el-table-column
        prop="SkillName"
        label="分类"
        width="100">
      </el-table-column>
    </el-table>
  </el-row>
</template>

<script>
  export default {
    name: 'TalentListChart',
    data () {
      return {
        url: 'talent',
        tableData: []
      }
    },
    created () {
      this.fetchList()
    },
    methods: {
      // net io
      fetchList (value) {
        this.$http.get(this.url).then((response) => {
          this.tableData = response.data
        })
      },
      reload () {
        this.fetchList()
      }
    }
  }
</script>
