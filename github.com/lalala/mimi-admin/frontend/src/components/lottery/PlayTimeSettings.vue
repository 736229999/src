<template>
  <!--工具条-->
  <div>
    <el-col :span="24" class="toolbar" style="padding-bottom: 20px;">
      <el-col :span="24">
        <el-button type="primary" @click="showDialogFunc">新增彩票</el-button>
      </el-col>
    </el-col>
    <!--列表-->
    <el-col>
      <el-table border :data="playTimeList" style="width: 100%" :row-class-name="tableRowClassName">
        <el-table-column label="彩种">
          <template scope="scope">
            <span>{{scope.row.name}}</span>
          </template>
        </el-table-column>
        <el-table-column label="开始时间(秒)">
          <template scope="scope">
            <el-tooltip class="item" effect="dark" content="双击内容修改数据" placement="top-start">
            <span class="time_span"
                  @dblclick="changeData(scope.$index, 1)">{{scope.row.start_time}}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column prop="end_time" label="结束时间(秒)">
          <template scope="scope">
            <el-tooltip class="item" effect="dark" content="双击内容修改数据" placement="top-start">
          <span class="time_span"
                @dblclick="changeData(scope.$index, 2)">{{scope.row.end_time}}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column label="追号时间(秒)">
          <template scope="scope">
            <el-tooltip class="item" effect="dark" content="双击内容修改数据" placement="top-start">
              <span class="time_span" @dblclick="changeData(scope.$index, 3)">{{scope.row.chase_start_time}}</span>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog title="设置间隔时间单位为秒" :visible.sync="dialogVisible" size="tiny" :before-close="handleClose">
        <el-col :span="10">
          <el-input v-model.number="time" type="number" placeholder="设置间隔时间单位为秒"></el-input>
        </el-col>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="submit">确 定</el-button>
        </span>
      </el-dialog>
    </el-col>

    <el-dialog title="选择你要添加的彩种" :visible.sync="showDialog" size="tiny" :before-close="handleClose">
      <el-col :span="24">
        <el-form :model="form" :rules="formRules" ref="form">
          <el-form-item label="选择彩种 " prop="lottery_id">
            <el-select v-model.number="form.lottery_id" placeholder="请选择" style="width: 100%">
              <el-option v-for="item in lotteryType" :key="item.Id" :disabled="item.Disabled != false"
                         :label="item.Name" :value="item.Id"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="开始时间(秒)" prop="start_time">
            <el-input v-model.number="form.start_time" auto-complete="off"></el-input>
          </el-form-item>
          <el-form-item label="结束时间(秒)" prop="end_time">
            <el-input v-model.number="form.end_time" auto-complete="off"></el-input>
          </el-form-item>
          <el-form-item label="追号时间(秒)" prop="chase_start_time">
            <el-input v-model.number="form.chase_start_time" auto-complete="off"></el-input>
          </el-form-item>
        </el-form>
      </el-col>
      <span slot="footer" class="dialog-footer">
          <el-button @click="showDialog = false">取 消</el-button>
          <el-button type="primary" @click="onSubmit">确 定</el-button>
        </span>
    </el-dialog>
  </div>
</template>

<script>

  import moment from 'moment'
  import ElInput from "../../../node_modules/element-ui/packages/input/src/input";
  import lotteryApi from "../../api/lottery"
  export default {
    components: {ElInput},
    data() {
      return {
        playTimeList: [],
        form: {
          lottery_id: '',
          start_time: '',
          chase_start_time: '',
          end_time: '',
        },
        formRules: {
          lottery_id: [{required: true, message: '请选择彩种'}],
          start_time: [
            {required: true, message: '开始的间隔时间不能为空'},
            {type: 'number', message: '开始的间隔时间必须为整数'}
          ],
          end_time: [
            {required: true, message: '结束的间隔时间不能为空'},
            {type: 'number', message: '结束的间隔时间必须为整数'}
          ],
          chase_start_time: [
            {required: true, message: '追号的间隔时间不能为空'},
            {type: 'number', message: '追号的间隔时间必须为整数'}
          ],
        },
        lotteryType: [],
        lottery: 0,
        dialogValue: "",
        dialogVisible: false,
        time: 1496387988,
        index: 0,
        type: 0,
        showDialog: false,
      };
    },
    methods: {
      tableRowClassName(row, index) {
        if (index === 1) {
          return 'info-row';
        } else if (index === 3) {
          return 'positive-row';
        }
        return '';
      },
      handleClick: function () {
        console.log(1);
      },
      handleClose: function () {

      },
      changeData: function (index, param) {
        //获取当前时间，填充到this.time中.
        let unixTime = 0;
        switch (param) {
          case 1:
            unixTime = this.playTimeList[index].start_time;
            break;
          case 2:
            unixTime = this.playTimeList[index].end_time;
            break;
          case 3:
            unixTime = this.playTimeList[index].chase_start_time;
            break;
        }
        this.index = index;
        this.type = param;

        this.time = unixTime;

        this.dialogVisible = true;

      },
      onSubmit: function () {

        switch (this.type) {
          case 1:
            this.playTimeList[this.index].start_time = this.time;
            break;
          case 2:
            this.playTimeList[this.index].end_time = this.time;
            break;
          case 3:
            this.playTimeList[this.index].chase_start_time = this.time;
            break;
        }

        this.$refs["form"].validate((valid) => {
          if (valid) {
            this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {

              //根据索引获取对应的id.
              this.showDialog = false;

              lotteryApi.addPlayTime(this.form).then((res) => {
                this.$message({type: 'success', message: '已经修改了! 不怕，不怕'});
                this.GetPlayTimeList();
              }).catch(err => {
                this.$message.error("提交了，但是失败了，赶快检查下，是不是哪儿填错了")
              });

            })
          }
        })
      },
      GetPlayTimeList: function () {
        lotteryApi.playTiemSettingGetLottery().then((res) => {
          this.playTimeList = res.data.msg;
          console.log("play_time_list:", this.playTimeList)
        })
      },
      GetLotteryType: function () {
        //获取彩种列表.
        lotteryApi.lotteryList().then((res) => {
          let obj = res.data.msg;
          let pt = this.playTimeList;
          console.log(pt);

          if (pt != null) {
            for (let i = 0; i < obj.length; i++) {
              obj[i]["Disabled"] = false;
              for (let j = 0; j < pt.length; j++) {
                if (obj[i].Id == pt[j].lottery_id) {
                  console.log(obj[i].Id, obj[i].Name);
                  obj[i]["Disabled"] = true;
                }
              }
            }
          } else {
            for (let i = 0; i < obj.length; i++) {
              obj[i]["Disabled"] = false;
            }
          }

          this.lotteryType = obj;
          console.log("obj:", this.lotteryType)
        });
      },
      showDialogFunc: function () {
        this.showDialog = true;
        this.GetLotteryType();
      },
      submit: function () {

        let type = this.type;
        switch (type) {
          case 1:
            this.playTimeList[this.index].start_time = this.time;
            break;
          case 2:
            this.playTimeList[this.index].end_time = this.time;
            break;
          case 3:
            this.playTimeList[this.index].chase_start_time = this.time;
            break;
        }

        lotteryApi.updatePlayTimeSetting(this.playTimeList[this.index]).then((res) => {
          this.$message.success("更新成功")
          this.dialogVisible = false;
        }).catch(err => {
          this.$message.error('更新失败')
        })
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    },
    mounted() {
      this.GetPlayTimeList();
    }
  }
</script>

<style>
  .el-table .info-row {
    background: #c9e5f5;
  }

  .el-table .positive-row {
    background: #e2f0e4;
  }

  .time_span {
    cursor: pointer;
  }
</style>
