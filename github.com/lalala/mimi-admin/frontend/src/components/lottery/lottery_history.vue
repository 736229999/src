<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="开奖时间">
          <el-col :span="23">
            <el-date-picker type="date" placeholder="选择日期" v-model="form.date" style="width: 100%;"></el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" >查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!--列表-->
    <el-table :data="data" border>
      <el-table-column label="期号" prop="lottery_issue_number"></el-table-column>
      <el-table-column label="开奖时间" prop="lottery_time"></el-table-column>
      <el-table-column label="开奖号码">
        <template scope="scope">
          <div class="ballbg">
            <span class="ball_1">05</span>
            <span class="ball_1">06</span>
            <span class="ball_1">07</span>
            <span class="ball_1">10</span>
            <span class="ball_1">19</span>
            <span class="ball_1">26</span>
            <span class="ball_2">29</span>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
    export default {
        data() {
            return {
                data : [{
                    lottery_name : "双色球",
                    lottery_issue_number : "20170410",
                    lottery_time : "2017-04-02",
                    lottery_number: "05 06 07 10 19 26 29 21"
                },
                {
                    lottery_name : "福彩3D",
                    lottery_issue_number : "20170410",
                    lottery_time : "2017-04-02",
                    lottery_number: "05 06 07 10 19 26 29 21"
                },
                {
                    lottery_name : "七乐彩",
                    lottery_issue_number : "20170410",
                    lottery_time : "2017-04-02",
                    lottery_number: "05 06 07 10 19 26 29 21"
                },
                {
                    lottery_name : "大乐透",
                    lottery_issue_number : "20170410",
                    lottery_time : "2017-04-02",
                    lottery_number: "05 06 07 10 19 26 29 21"
                }],
                form : {
                    date : ''
                }
            }
        },methods : {
            getLotteryListInfo() {

            }
        },
        mounted() {
            this.getLotteryListInfo();
        }
    }
</script>

<style scoped>
  .table-expand {
    font-size: 0;
  }
  .table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 33%;
  }
   .ballbg {
      text-align: left;
      margin: 0 8px;
  }
     .ballbg span.ball_1 {
        background: url(http://static.lecai.com/img/openprize_ball_bg.gif) 0 0 no-repeat;
    }
    .ballbg span {
        width: 20px;
        height: 21px;
        line-height: 21px;
        text-align: center;
        font-weight: bold;
        color: #fff;
        display: inline-block;
    }
    .ballbg span.ball_2 {
        background: url(http://static.lecai.com/img/openprize_ball_bg.gif) -21px 0 no-repeat;
    }
    .history {
        size:12px;
        color:#20a0ff;
        text-decoration:none;
    }
</style>

