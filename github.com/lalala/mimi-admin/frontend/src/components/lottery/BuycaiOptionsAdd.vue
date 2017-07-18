<template>
  <div>
    <el-col :span="14" class="contain">
      <el-form :model="AddForm" label-width="120px" ref="form">
        <el-form-item label="彩种" prop="lotteryType">
          <el-col :span="24">
            <el-select v-model="selectLottery" @change="changeLottery" placeholder="请选择" style="width: 100%">
              <!--<el-option v-for="item in lotteryType" :key="item.Code" :disabled="item.DayMaxNo < 1 && complete == true" :label="item.Name" :value="item.Code">-->
              <el-option v-for="item in lotteryType" :key="item.Code" :label="item.Name" :value="item.Code">
              </el-option>
            </el-select>
          </el-col>
        </el-form-item>
        <!--<el-form-item label="是否增设整天" prop="">-->
        <!--<el-col ::span="24">-->
        <!--<el-switch v-model="complete" on-text="是" off-text="否" on-color="#20A0FF" @change="completes"-->
        <!--off-color="#8492A6"></el-switch>-->
        <!--</el-col>-->
        <!--</el-form-item>-->
        <div id="day">
          <el-form-item :label="title" prop="day">
            <el-col :span="24">
              <el-input-number v-model="num" @change="changeDay" style="width: 100%" :min="1"></el-input-number>
            </el-col>
            <el-col :span="6">
              &nbsp;<el-tag type="gary" v-model="changeDayIssue">{{changeDayIssue}}</el-tag>
            </el-col>
          </el-form-item>
        </div>
        <div id="bjpk10" style="display:none">
          <el-form-item label="开始期号">
            <el-col :span="24">
              <el-input v-model="issue"></el-input>
            </el-col>
          </el-form-item>
        </div>
        <div id="setting" style="display: none;">
          <el-form-item label="期号数据">
            <el-table :data="issueInfo" border style="width: 100%">
              <el-table-column label="序号" width="80">
                <template scope="scope">
                  <span>{{scope.$index + 1}}</span>
                </template>
              </el-table-column>
              <el-table-column label="期号" width="150">
                <template scope="scope">
                  <el-tooltip class="item" effect="dark" content="双击内容修改数据" placement="top-start">
                    <span class="issueClass" alt="点击修改数据"
                          @dblclick="issueClick(scope.$index, 1)">{{scope.row.issue}}</span>
                  </el-tooltip>
                </template>
              </el-table-column>
              <el-table-column label="开始时间" width="180">
                <template scope="scope">
                  <el-tooltip class="item" effect="dark" content="双击内容修改数据" placement="top-start">
                    <span class="issueClass"
                          @dblclick="issueClick(scope.$index, 2)">{{new Date(scope.row.start_time * 1000).toLocaleString()}}</span>
                  </el-tooltip>
                </template>
              </el-table-column>
              <el-table-column label="结束时间">
                <template scope="scope">
                  <el-tooltip class="item" effect="dark" content="双击内容修改数据" placement="top-start">
                    <span class="issueClass"
                          @dblclick="issueClick(scope.$index, 3)">{{new Date(scope.row.end_time * 1000).toLocaleString()}}</span>
                  </el-tooltip>
                </template>
              </el-table-column>
              <el-table-column label="开奖时间">
                <template scope="scope">
                  <el-tooltip class="item" effect="dark" content="双击内容修改数据" placement="top-start">
                    <span class="issueClass"
                          @dblclick="issueClick(scope.$index, 4)">{{new Date(scope.row.open_time * 1000).toLocaleString()}}</span>
                  </el-tooltip>
                </template>
              </el-table-column>
            </el-table>
            <!--<div class="block">-->
            <!--<el-pagination @current-change="handleCurrentChange" layout="total, prev, pager, next"-->
            <!--:total="issueInfo.length" :page-size="20"></el-pagination>-->
            <!--</div>-->
          </el-form-item>
        </div>
        <el-form-item>
          <el-button type="primary" @click="SettingOk" :disabled="disabledSettingBut">增设</el-button>
          <el-button type="primary" id="submitBtn" @click="submitBtn" style="display: none"
                     :disabled="disabledSettingBut">提交
          </el-button>
        </el-form-item>

      </el-form>
      <el-dialog title="填入你要修改的参数" :visible.sync="dialogVisible" size="tiny" :before-close="handleClose">
        <el-col :span="10">
          <div v-if="type == 1">
            <el-input v-model="dialogValue" style="width: 100%"></el-input>
          </div>
          <div v-else>
            <el-date-picker v-model="time" type="datetime" placeholder="选择时间" style="width: 100%"></el-date-picker>
          </div>
        </el-col>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="onSubmit">确 定</el-button>
        </span>
      </el-dialog>

      <el-dialog title="填写初始期号" :visible.sync="dialogFormInit" :show-close="false" :close-on-click-modal="false">
        <el-form :model="initForm" :rules="initFormRules" ref="initForm">
          <el-form-item label="期号" prop="issue">
            <el-input v-model="initForm.issue" type="text" placeholder="期号"></el-input>
          </el-form-item>
          <el-form-item label="开始时间" prop="start_time">
            <el-date-picker style="width: 100%" v-model="initForm.start_time" type="datetime"
                            placeholder="开奖时间"></el-date-picker>
          </el-form-item>
          <el-form-item label="结束时间" prop="end_time">
            <el-date-picker style="width: 100%" v-model="initForm.end_time" type="datetime"
                            placeholder="开奖时间"></el-date-picker>
          </el-form-item>
          <el-form-item label="开奖时间" prop="open_time">
            <el-date-picker style="width: 100%" v-model="initForm.open_time" type="datetime"
                            placeholder="开奖时间"></el-date-picker>
          </el-form-item>
          <el-form-item label="中奖方案">
            <el-input v-model="initForm.open_balls" type="text" placeholder="中奖方案"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click="initIssueData">确 定</el-button>
        </div>
      </el-dialog>
    </el-col>
  </div>
</template>

<script>

  import lotteryApi from "../../api/lottery"
  import moment from 'moment';
  import router from "../../router"
  export default {
    data() {
      return {
        title: "增设天数",
        page: 20,
        pageData: [],
        date: '',
        changeDayIssue: "",
        dialogVisible: false,
        dialogValue: '',
        dialogTime: '',
        AddForm: {
          lottery: '',
          buycai: [],
        },
        lottery: '',
        complete: true,//是否设置完整的天数.
        demo: 'demo',
        lotteryHistory: [],
        lotteryType: [],
        AddFormRules: [],
        issueInfo: [],
        settingNum: 0,
        disabledSettingBut: false,
        index: '',
        type: 1,
        time: '',
        selectLottery: "", //选中的彩票.
        num: 0, //增设期数或者天数.
        day_max_no: "每天开奖期数",
        issue: "", //期号.
        dialogFormInit: false,
        initForm: { //初始化数据.
          issue: "",
          start_time: "",
          end_time: "",
          open_time: "",
          open_balls: "",
          lottery: "",
          buycai: {},
        },
        initFormRules: {
          issue: [
            {required: true, message: '期号不能为空', trigger: 'blur'}
          ],
          start_time: [
            {type: 'date', required: true, message: '开始时间不能为空', trigger: 'change'}
          ],
          end_time: [
            {type: 'date', required: true, message: '结束时间不能为空', trigger: 'change'}
          ],
          open_time: [
            {type: 'date', required: true, message: '开奖时间不能为空', trigger: 'change'}
          ],
        }
      }

    },
    methods: {
      //初始化数据.
      initIssueData: function () {
        this.$refs["initForm"].validate((valid) => {
          if (valid) {
            this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'error'
            }).then(() => {
              let moment = require("moment");
              this.initForm.buycai.start_time = moment(this.initForm.start_time).unix();
              this.initForm.buycai.end_time = moment(this.initForm.end_time).unix();
              this.initForm.buycai.open_time = moment(this.initForm.open_time).unix();
              this.initForm.buycai.issue = this.initForm.issue;
              this.initForm.lottery = this.selectLottery;
              if (this.initForm.start_time >= this.initForm.end_time) {
                this.$message.error("开始时间不能大于结束时间");
                return
              }

              lotteryApi.initBuycaiOptions(this.initForm).then((res) => {
                this.$message.success("添加成功");
                this.dialogFormInit = false;
                router.push({path: "/lottery/buycai/options"})

              }).catch(err => {
                this.$message.error(err.data.msg)
              })

            }).catch(() => {
              this.$message.info("还是在检查下比较好");
            });
          }
        })
      },
      GetLotteryList: function () {
        //获取彩种列表.
        lotteryApi.lotteryList().then((res) => {
          let obj = res.data;
          this.lotteryType = obj.msg;
          this.selectLottery = obj.msg[0].Code;
          this.num = obj.msg[0].Code;

          this.initIssue();
        });
      },

      //.
      GetLotteryIssue: function () {
        let form = {
          lottery: this.selectLottery,
          num: this.num,
          issue: this.issue,
        };
        this.getIssue(form);
      },

      //确定增设.
      SettingOk: function () {
        this.GetLotteryIssue();
        document.getElementById("submitBtn").style = "display:black"
      },
      issueClick: function (index, type) {
        this.dialogVisible = true;
        this.index = index;
        this.type = type;
      },
      handleClose(done) {
        this.$confirm('确认关闭？')
          .then(_ => {
            done();
          })
          .catch(_ => {
          });
      },
      changeDay: function (day) {
        //判断是高频彩，还是低频彩.
        for (let i = 0; i < this.lotteryType.length; i++) {
          if (this.lotteryType[i].Code == this.selectLottery) {
            if (this.lotteryType[i].DayMaxNo < 1) {
              //低频彩.
              this.changeDayIssue = "增设" + day + "期"
            } else {
              this.num = day;
              let msg = day + " 天 共计 " + day * this.day_max_no + " 期";
              this.changeDayIssue = msg
            }

          }
        }

      },
      initIssue: function () {
        //根据现在选中的彩种来获取每天开奖的期数.
        for (let i = 0; i < this.lotteryType.length; i++) {
          if (this.selectLottery == this.lotteryType[i].Code) {
            this.day_max_no = this.lotteryType[i].DayMaxNo;
          }
        }

        let day = this.num;
        let msg = day + " 天 共计 " + day * this.day_max_no + " 期";

        this.changeDayIssue = msg;
      },
      onSubmit: function () {

        if (this.type == 1) {
          this.issueInfo[this.index].issue = this.dialogValue;
        } else if (this.type == 2) {

          this.issueInfo[this.index].start_time = moment(this.time).unix();
        } else if (this.type == 3) {
          this.issueInfo[this.index].end_time = moment(this.time).unix();
        } else if (this.type == 4) {
          this.issueInfo[this.index].open_time = moment(this.time).unix();
        }
        this.dialogVisible = false
      },

      //根据天数和彩种来获取期号.
      getIssue: function (params) {
        //检查彩种.
        if (params.lottery == "") {
          this.$message.error("请选择彩种");
          return false
        }
        //检查天数 或者期数.
        if (params.num < 1) {
          this.$message.error("请填写天数");
          return false
        }

        lotteryApi.buycaioptionsIssue(params).then((res) => {
          console.log(res);
          this.issueInfo = res.data.msg;
          let sett = document.getElementById("setting");
          sett.style = "display:black";
        })
      },
      changeLottery: function (val) {
        this.num = 1;
        //判断是高频彩，还是低频彩.
        this.selectLottery = val;
        for (let i = 0; i < this.lotteryType.length; i++) {
          if (this.lotteryType[i].Code == val) {
            if (this.lotteryType[i].DayMaxNo < 1) {
              this.title = "增设期数";
              this.changeDayIssue = "增设1期";
            } else {
              document.getElementById("setting").style = "display:none";
              document.getElementById("submitBtn").style = "display:none";
              this.selectLottery = val;

              for (let i = 0; i < this.lotteryType.length; i++) {
                if (this.selectLottery == this.lotteryType[i].Code) {
                  this.day_max_no = this.lotteryType[i].DayMaxNo
                }
              }
              //将天数设为1.
              this.num = 1;

              this.initIssue();
              this.title = "增设天数"
            }
          }
        }

        //获取最后的一期.
        lotteryApi.buycaiOptionsGetIssueByLottery({lottery: val}).then((res) => {

          //判断查询出来的数据是否为空.
          //如果为空则需要手动设置第一期数据.
          if (res.data.msg.id < 1 && val != "bjpk10") {
            document.getElementById("bjpk10").style = "display:none";
            this.dialogFormInit = true;
            return
          }

          //判断最后的一期是否为开奖的最后一期.
          let t = new Date(res.data.msg.start_time * 1000).toLocaleString();
          let arr = t.split(" ");
          if (val == "bjpk10" && arr[1] != "23:52:00") {
            document.getElementById("bjpk10").style = "display:black";
//            alert("数据错误，需要手动设置期号")
          }
          this.issue = res.data.msg.issue
        });
      },
      handleCurrentChange(val) {
        console.log(val)

      },
      submitBtn: function () {
        this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          let form = {
            lottery: this.selectLottery,
            num: this.num,
            issue: this.issue,
            buycai: this.issueInfo,
          };
          lotteryApi.buycaiOptionsAddIssue(form).then((res) => {
            this.$message.success("已经提交了! 不怕，不怕");
            document.getElementById("submitBtn").style = "disbled:true";
            router.push({path: "/lottery/buycai/options"})
          }).catch(err => {
            this.$message.error("已经提交了! 但是失败了，快检查下，看看哪儿是不是填错了？");
            document.getElementById("submitBtn").style = "disbled:false";
          });
        }).catch(() => {
          this.$message({type: 'info', message: '还是在检查下比较好'});
        });
      }
    },
    mounted() {
      this.GetLotteryList();
    }
  }
</script>

<style>
  .contain {
    margin: 50px;
  }

  .form_item {
    margin-top: 15px;
    margin-bottom: 15px;
  }

  .ri {
    text-align: right;
  }

  .specifyTypeShow {
    display: none;
  }

  .specifyTypeNotShow {
    display: block;
  }

  .issueClass {
    cursor: pointer;
  }
</style>

