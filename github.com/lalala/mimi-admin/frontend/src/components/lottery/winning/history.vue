<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="期号">
          <el-col :span="23">
            <el-input v-model="form.issue" placeholder="期号"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="开始时间">
          <el-date-picker v-model="form.start_time" type="datetime" placeholder="选择日期时间">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="结束时间">
          <el-date-picker v-model="form.end_time" type="datetime" placeholder="选择日期时间">
          </el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="search">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!--列表-->
    <el-col>
      <el-table :data="form.list" border>
        <el-table-column label="彩种" prop="lottery_name"></el-table-column>
        <el-table-column label="期号" prop="issue"></el-table-column>
        <el-table-column label="开始时间">
          <template scope="scope">
            <span>{{new Date(scope.row.start_time * 1000).toLocaleString()}}</span>
          </template>
        </el-table-column>
        <el-table-column label="结束时间">
          <template scope="scope">
            <span>{{new Date(scope.row.end_time * 1000).toLocaleString()}}</span>
          </template>
        </el-table-column>
        <el-table-column label="开奖时间">
          <template scope="scope">
            <span>{{new Date(scope.row.open_time * 1000).toLocaleString()}}</span>
          </template>
        </el-table-column>
        <el-table-column label="中奖方案" prop="open_balls"></el-table-column>
      </el-table>
    </el-col>
  </div>
</template>

<script>
  import winApi from '../../../api/winning';
  import lotteryApi from "../../../api/lottery"
  import router from "../../../router"
  export default {
    data() {
      return {
        lotteryType: [],
        total: 0,
        form: {
          lottery: '',
          issue: "",
          start_time: "",
          end_time: "",
          size: 100,
          page: 1,
          total: 0,
          list: [],
        },

      }
    }, methods: {
      GetLotteryList: function () {
        lotteryApi.lotteryList().then((res) => {
          let obj = res.data;
          this.lotteryType = obj.msg;
          this.lottery = obj.msg[0].Code;
          this.form.lottery = obj.msg[0].Code;

          this.GetBuycaiOptions()

        }).catch(err => {

        });
      },
      handleClick: function () {
        console.log(1);
      },
      GetWinning() {
        if (this.form.start_time == "") {
          this.form.start_time = 0;
        }
        if (this.form.end_time == "") {
          this.form.end_time = 0;
        }
        this.form.lottery = this.$route.params.lottery;

        winApi.winningHistoryList(this.form).then((res) => {
//          this.form = res.data.msg
          let obj = res.data.msg;
          if (obj.start_time == 0) {
            obj.start_time = "";
          }
          if (obj.end_time == 0) {
            obj.end_time = "";
          }

          this.form = obj

        }).catch(err => {

        });
      },
      handleCurrentChange(val) {
        this.form.page = val;
        this.GetLog();
      },
      changeLottery: function (val) {
        this.form.lottery = val;
      },
      search: function () {

        let moment = require("moment");

        if (typeof(this.form.start_time) == "undefined") {
          this.form.start_time = 0;
        } else {
          this.form.start_time = moment(this.form.start_time).unix()
        }
        if (typeof(this.form.end_time) == "undefined") {
          this.form.end_time = 0;
        } else {
          this.form.end_time = moment(this.form.end_time).unix()
        }
        winApi.winningHistoryList(this.form).then((res) => {
          let obj = res.data.msg;
          if (obj.start_time == 0) {
            obj.start_time = "";
          } else {
              obj.start_time = moment(obj.start_time * 1000).toDate();
          }
          if (obj.end_time == 0) {
            obj.end_time = "";
          } else {
              obj.end_time = moment(obj.end_time * 1000).toDate()
          }
          this.form = obj;
        })
      },

    },
    mounted() {
      this.GetWinning();
    }
  }
</script>

