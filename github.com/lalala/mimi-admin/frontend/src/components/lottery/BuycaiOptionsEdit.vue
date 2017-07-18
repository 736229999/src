<template>
  <div>
    <el-col :span="10" class="contain">
      <el-form :model="form" label-width="120px" :rules="formRules" ref="form">
        <el-form-item label="期号" prop="issue">
          <el-col>
            <el-input v-model="form.issue"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="开始时间" prop="start_time">
          <el-col>
            <el-date-picker style="width:100%" v-model="form.start_time" type="datetime"
                            placeholder="开始时间"></el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="结束时间" prop="end_time">
          <el-col>
            <el-date-picker style="width:100%" v-model="form.end_time" type="datetime"
                            placeholder="结束时间"></el-date-picker>

          </el-col>
        </el-form-item>
        <el-form-item label="开奖时间" prop="open_time">
          <el-col>
            <el-date-picker style="width: 100%" v-model="form.open_time" type="datetime"
                            placeholder="开奖时间"></el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="中奖方案" prop="open_balls">
          <el-col>
            <el-input v-model="form.open_balls"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
        </el-form-item>
        <el-form-item>
          <el-col>
            <el-button type="primary" @click="submitBtn">修改</el-button>
          </el-col>
        </el-form-item>
      </el-form>
    </el-col>
  </div>
</template>

<script>

  import lotteryApi from "../../api/lottery"
  import moment from 'moment';
  import Vue from "vue"
  import router from "../../router"

  export default {
    data() {
      return {
        form: {
          id: "",
          lottery: "",
          start_time: "",
          end_time: "",
          open_time: "",
          open_balls: "",
          issue:"",
          buycai: {},
        },
        formRules: {
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
          ]
        }
      }
    },
    methods: {
      submitBtn: function () {
        this.$refs["form"].validate((valid) => {
          if (valid) {
            this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'error'
            }).then(() => {
              let moment = require("moment");
              this.form.id = parseInt(this.$route.params.id);
              this.form.lottery = this.$route.params.lid;
              let buycai = {
                  start_time: moment(this.form.start_time).unix(),
                  end_time: moment(this.form.end_time).unix(),
                  open_time: moment(this.form.open_time).unix(),
                  open_balls: this.form.open_balls,
                  issue: this.form.issue,
                };
              Vue.set(this.form, "buycai", buycai);
              lotteryApi.updateBuycaiOptions(this.form).then((res) => {
                this.$message.success("提交成功");
                router.push({path: "/lottery/buycai/options"})
              }).catch(err => {
                this.$message.error(err.data.msg)
              });
            }).catch(err => {
              console.log(err);
              this.$message.info("还是检查下比较好")
            })
          }
        })
      },
      GetBuycaiOptionsById: function () {
        let params = {
          id: this.$route.params.id,
          lottery: this.$route.params.lid,
        };
        lotteryApi.getLotteryBuycaiOptionsById(params).then((res) => {
          let obj = res.data.msg;
          let moment = require("moment");
          obj.end_time = moment(obj.end_time * 1000).toDate();
          obj.start_time = moment(obj.start_time * 1000).toDate();
          obj.open_time = moment(obj.open_time * 1000).toDate();
          this.form = obj;
//          Vue.set(this.form, "buycai", obj)
        })
      }
    },
    mounted() {
      this.GetBuycaiOptionsById();
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

