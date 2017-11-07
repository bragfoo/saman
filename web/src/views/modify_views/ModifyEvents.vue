<template>
  <div>
    <el-row>
      <el-col :span="24" style="padding-bottom: 15px">
        <el-button size="mini" type="primary" plain @click="createRow">添加</el-button>
        <el-button size="mini" type="primary" plain @click="reload">刷新</el-button>
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
            label="开始日期"
            width="150">
            <template slot-scope="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.StartDate | timeToDay}}</span>
            </template>
          </el-table-column>
          <el-table-column
            fixed
            label="结束日期"
            width="150">
            <template slot-scope="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.EndDate | timeToDay}}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="Name"
            label="活动名"
            width="120">
          </el-table-column>
          <el-table-column
            prop="TotalPeople"
            label="全部人数"
            width="120">
          </el-table-column>
          <el-table-column
            prop="TotalWork"
            label="全部工作人员"
            width="120">
          </el-table-column>
          <el-table-column
            prop="UploadPeople"
            label="上传视频人数"
            width="120">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <el-button @click="editRow(scope.row.Ids)" type="text" size="small">编辑</el-button>
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
        <el-form-item label="开始时间">
          <div class="block">
            <el-date-picker
              v-model="rowData.StartDate"
              type="date"
              placeholder="选择日期"
              :picker-options="timeOption">
            </el-date-picker>
          </div>
        </el-form-item>
        <el-form-item label="结束时间">
          <div class="block">
            <el-date-picker
              v-model="rowData.EndDate"
              type="date"
              placeholder="选择日期"
              :picker-options="timeOption">
            </el-date-picker>
          </div>
        </el-form-item>
        <el-form-item label="活动名">
          <el-input v-model="rowData.Name"></el-input>
        </el-form-item>
        <el-form-item label="全部人数">
          <el-input-number v-model="rowData.TotalPeople"></el-input-number>
        </el-form-item>
        <el-form-item label="全部工作人员">
          <el-input-number v-model="rowData.TotalWork"></el-input-number>
        </el-form-item>
        <el-form-item label="上传视频人数">
          <el-input-number v-model="rowData.UploadPeople"></el-input-number>
        </el-form-item>
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
    name: 'ModifyEvents',
    data () {
      return {
        url: 'events',
        tableData: [],
        rowData: {},
        dialogVisible: false,
        playType: '59fae20cef2d1314e0ea2a55',
        isCreate: true,
        timeOption: {
          disabledDate (time) {
            return time.getTime() > Date.now()
          }
        }
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
      saveData (func) {
        if (this.isCreate) {
          this.$http.post(this.url, this.rowData).then(func)
        } else {
          this.$http.put(this.url, this.rowData).then(func)
        }
      },
      deleteData (id) {
        this.$http.delete(this.url).then(this.reload())
      },
      getRowData (id) {
        if (id) {
          return JSON.parse(JSON.stringify(this.tableData.find((row) => {
            return row.Ids === id
          })))
        }
        return {
          StartDate: (new Date()).getTime(),
          EndDate: (new Date()).getTime(),
          totalPeople: 0,
          totalWork: 0,
          uploadPeople: 0,
          name: ''
        }
      },
      reload () {
        this.fetchList()
      },

      // dialog control
      showDialog (id) {
        if (this.isCreate) {
          this.rowData = this.getRowData()
        } else {
          this.rowData = this.getRowData(id)
          this.rowData.StartDate = new Date(this.rowData.StartDate * 1000)
          this.rowData.EntTime = new Date(this.rowData.EndDate * 1000)
        }
        this.dialogVisible = true
      },
      closeDialog () {
        this.dialogVisible = false
        this.isCreate = false
//        this.rowData = this.getRowData()
      },
      createRow () {
        this.isCreate = true
        this.showDialog()
      },
      editRow (id) {
        this.isCreate = false
        this.showDialog(id)
      },
      saveRow () {
        this.rowData.StartDate = Math.floor(new Date(this.rowData.StartDate).getTime() / 1000)
        this.rowData.EndDate = Math.floor(new Date(this.rowData.EndDate).getTime() / 1000)
        this.saveData((response) => {
          this.closeDialog()
          this.reload()
        })
      },
      handleClick (row) {
        console.log(row)
      }
    }
  }
</script>
