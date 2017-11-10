<template>
  <div>
    <el-row>
      <el-col :span="24" style="padding-bottom: 15px">
        <PlatTypeSelect v-model="platType" @change="fetchList"></PlatTypeSelect>
        <VideoTypeSelect v-model="videoType" @change="fetchList" :platIds="platType"></VideoTypeSelect>
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
            label="日期"
            width="150">
            <template slot-scope="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.CreateTime | timeToDay}}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="Sum"
            label="总数"
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
        <el-form-item label="时间">
          <div class="block">
            <el-date-picker
              v-model="rowData.TCreateTime"
              type="date"
              placeholder="选择日期"
              :picker-options="timeOption">
            </el-date-picker>
          </div>
        </el-form-item>
        <el-form-item label="播放总数">
          <el-input-number v-model="rowData.Sum"></el-input-number>
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
  import PlatTypeSelect from '../../components/PlatTypeSelect.vue'
  import VideoTypeSelect from '../../components/VideoTypeSelect.vue'

  export default {
    name: 'ModifyPlayAmount',
    components: {VideoTypeSelect, PlatTypeSelect},
    data () {
      return {
        url: 'playAmount',
        tableData: [],
        rowData: {},
        dialogVisible: false,
        videoType: '',
        platType: '59fae20cef2d1314e0ea2a55',
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
        if (this.videoType) {
          this.$http.get(this.url, {
            params: {
              'videoIds': this.videoType,
              'platIds': this.platType
            }
          }).then((response) => {
            this.tableData = response.data
          })
        }
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
          TCreateTime: (new Date()).getTime(),
          CreateTime: (new Date()).getTime(),
          Sum: 0,
          VideoIds: this.videoType
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
          this.rowData.TCreateTime = new Date(this.rowData.CreateTime * 1000)
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
        this.rowData.CreateTime = Math.floor(new Date(this.rowData.TCreateTime).getTime() / 1000)
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
