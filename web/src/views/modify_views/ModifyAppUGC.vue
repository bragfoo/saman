<template>
  <div>
    <el-row>
      <el-col :span="24" style="padding-bottom: 15px">
        <el-button size="mini" type="primary" plain @click="createRow">添加</el-button>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-table
          :data="tableData"
          border
          style="width: 100%">
          <el-table-column
            fixed
            prop="date"
            label="日期"
            width="150">
          </el-table-column>
          <el-table-column
            prop="name"
            label="姓名"
            width="120">
          </el-table-column>
          <el-table-column
            prop="province"
            label="省份"
            width="120">
          </el-table-column>
          <el-table-column
            prop="city"
            label="市区"
            width="120">
          </el-table-column>
          <el-table-column
            prop="address"
            label="地址"
            width="300">
          </el-table-column>
          <el-table-column
            prop="zip"
            label="邮编"
            width="120">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <el-button @click="editRow(scope.row)" type="text" size="small">查看</el-button>
              <el-button type="text" size="small">编辑</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
    <el-dialog
      title="更新数据"
      :visible.sync="dialogVisible"
      width="30%"
      center>
      <el-form ref="form" :model="rowData" label-width="80px">
        <el-form-item label="活动名称">
          <el-input v-model="rowData.name"></el-input>
        </el-form-item>
        <!--<el-form-item label="活动区域">-->
        <!--<el-select v-model="form.region" placeholder="请选择活动区域">-->
        <!--<el-option label="区域一" value="shanghai"></el-option>-->
        <!--<el-option label="区域二" value="beijing"></el-option>-->
        <!--</el-select>-->
        <!--</el-form-item>-->
        <!--<el-form-item label="活动时间">-->
        <!--<el-col :span="11">-->
        <!--<el-date-picker type="date" placeholder="选择日期" v-model="form.date1" style="width: 100%;"></el-date-picker>-->
        <!--</el-col>-->
        <!--<el-col class="line" :span="2">-</el-col>-->
        <!--<el-col :span="11">-->
        <!--<el-time-picker type="fixed-time" placeholder="选择时间" v-model="form.date2"-->
        <!--style="width: 100%;"></el-time-picker>-->
        <!--</el-col>-->
        <!--</el-form-item>-->
        <!--<el-form-item label="即时配送">-->
        <!--<el-switch v-model="form.delivery"></el-switch>-->
        <!--</el-form-item>-->
        <!--<el-form-item label="活动性质">-->
        <!--<el-checkbox-group v-model="form.type">-->
        <!--<el-checkbox label="美食/餐厅线上活动" name="type"></el-checkbox>-->
        <!--<el-checkbox label="地推活动" name="type"></el-checkbox>-->
        <!--<el-checkbox label="线下主题活动" name="type"></el-checkbox>-->
        <!--<el-checkbox label="单纯品牌曝光" name="type"></el-checkbox>-->
        <!--</el-checkbox-group>-->
        <!--</el-form-item>-->
        <!--<el-form-item label="特殊资源">-->
        <!--<el-radio-group v-model="form.resource">-->
        <!--<el-radio label="线上品牌商赞助"></el-radio>-->
        <!--<el-radio label="线下场地免费"></el-radio>-->
        <!--</el-radio-group>-->
        <!--</el-form-item>-->
        <!--<el-form-item label="活动形式">-->
        <!--<el-input type="textarea" v-model="form.desc"></el-input>-->
        <!--</el-form-item>-->
        <el-form-item>
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="saveRow">确 定</el-button>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
      </span>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    name: 'Main',
    data () {
      return {
        url: '/api/ugc',
        tableData: [],
        rowData: {},
        dialogVisible: false,
        isCreate: true
      }
    },
    created () {

    },
    methods: {
      // net io
      fetchList () {
        this.$http.get(this.url, {
          params: {
            ID: 12345
          }
        }).then((response) => {
          this.tableData = response.data
        })
      },
      fetchData (id) {
        this.$http.get(this.url, {
          params: {
            ID: 12345
          }
        }).then((response) => {
          this.rowData = response.data
        })
      },
      saveData (func) {
        this.$http.post(this.url, this.rowData).then(func)
      },
      deleteData (id) {
        this.$http.delete(this.url).then(this.reload())
      },
      getRowData () {
        return {
          name: ''
        }
      },
      reload () {
        this.fetchList()
      },

      // dialog control
      showDialog () {
        if (this.isCreate) {
          this.rowData = this.getRowData()
        }
        this.fetchData()
        this.dialogVisible = true
      },
      closeDialog () {
        this.dialogVisible = false
        this.isCreate = false
        this.rowData = this.getRowData()
      },
      createRow () {
        this.isCreate = true
        this.showDialog()
      },
      editRow () {
        this.isCreate = false
        this.showDialog()
      },
      saveRow () {
        this.saveData((response) => {
          this.closeDialog()
        })
      },
      handleClick (row) {
        console.log(row)
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h1, h2 {
    font-weight: normal;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #42b983;
  }
</style>
