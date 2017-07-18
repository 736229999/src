<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="购彩时间">
          <el-col :span="23">
            <el-date-picker type="date" placeholder="选择日期" v-model="form.date" style="width: 100%;"></el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="手机号">
        <el-col :span="23">
            <el-input v-model="form.account" placeholder="手机号"></el-input>
            </el-col>
        </el-form-item>
        <el-form-item label="购买金额">
            <el-col :span="23">
                <el-input v-model="form.order_num" placeholder="充值金额"></el-input>
            </el-col>
        </el-form-item>
        <el-form-item label="购买状态">
             <el-select v-model="form.order_status" placeholder="请选择购买状态">
                  <el-option label="成功" value="shanghai"></el-option>
                  <el-option label="失败" value="beijing"></el-option>
                </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" >查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!--列表-->
    <el-table :data="data" border>
      <el-table-column label="购彩时间" prop="buy_lottery_time"></el-table-column>
      <el-table-column label="彩票名称" prop="lottery_name"></el-table-column>
      <el-table-column label="购彩方案" width="300">
        <template scope="scope">
            <div class="ballbg">
                <span class="ball_1">05</span>
                <span class="ball_1">06</span>
                <span class="ball_1">07</span>
                <span class="ball_1">10</span>
                <span class="ball_1">19</span>
                <span class="ball_1">26</span>
                <span class="ball_2">29</span>
                <span style="width:100px; color:#FF4949;font-size:12px;">(投注10倍)</span>
            </div>
            <div class="ballbg">
                <span class="ball_1">05</span>
                <span class="ball_1">06</span>
                <span class="ball_1">07</span>
                <span class="ball_1">10</span>
                <span class="ball_1">19</span>
                <span class="ball_1">26</span>
                <span class="ball_2">29</span>
                <span style="width:100px; color:#FF4949;font-size:12px;">(投注5倍)</span>
            </div>
        </template>
      </el-table-column>
      <el-table-column label="购买账户" prop="buy_lottery_account"></el-table-column>
      <el-table-column label="购买金额" prop="buy_amount"></el-table-column>
      <el-table-column label="资金来源" prop="money_source"></el-table-column>
        <el-table-column label="购买状态" >
            <template scope="scope">
                <el-tag type="success">{{scope.row.buy_lottery_status}}</el-tag>
            </template>
        </el-table-column>
    </el-table>
    </el-col>
  </div>
</template>

<script>
    export default {
        data() {
            return {
                data : [{
                    buy_lottery_time : "2017-04-05 15:00",
                    lottery_name : "双色球",
                    buy_lottery_program : [],
                    buy_lottery_account : "151****3574",
                    buy_amount : "¥ 200.00",
                    money_source : "账户余额",
                    buy_lottery_status : "购买成功",
                }],
                form : {
                    date : '',
                    account : '',
                    order_num: '',
                    order_status : '',
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
        float:left;
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

