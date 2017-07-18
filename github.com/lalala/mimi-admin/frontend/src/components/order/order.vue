<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label=" 下单时间">
          <el-col :span="23">
            <el-date-picker type="date" placeholder="选择日期" v-model="form.date" style="width: 100%;"></el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="手机号">
        <el-col :span="23">
            <el-input v-model="form.account" placeholder="手机号"></el-input>
            </el-col>
        </el-form-item>
        <el-form-item label="订单号">
            <el-col :span="23">
                <el-input v-model="form.order_num" placeholder="订单号"></el-input>
            </el-col>
        </el-form-item>
        <el-form-item label="订单状态">
             <el-select v-model="form.order_status" placeholder="请选择订单状态">
                  <el-option label="待出票" value="shanghai"></el-option>
                  <el-option label="待开奖" value="beijing"></el-option>
                  <el-option label="已中奖" value="beijing"></el-option>
                  <el-option label="未中奖" value="beijing"></el-option>
                </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" >查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!--列表-->
    <el-table :data="data" border>
      <el-table-column label="订单编号" width="200" prop="order_number"></el-table-column>
      <el-table-column label="彩票类型" width="100" prop="lottery_type"></el-table-column>
      <el-table-column label="彩票期号" width="100" prop="lottery_issue_number"></el-table-column>
      <el-table-column label="下单时间" width="120" prop="order_create_time"></el-table-column>
      <el-table-column label="投注方案" width="300">
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
            <span class="ball_1">06</span>
            <span class="ball_1">08</span>
            <span class="ball_1">06</span>
            <span class="ball_1">10</span>
            <span class="ball_1">00</span>
            <span class="ball_1">26</span>
            <span class="ball_2">95</span>
            <span style="width:100px;color:#FF4949;font-size:12px;">(投注20倍)</span>
          </div>
        </template>
        </el-table-column>
        <el-table-column label="所属账户" width="150" prop="account"></el-table-column>
        <el-table-column label="支付金额" width="150" prop="pay_amount"></el-table-column>
        <el-table-column label="实际支付金额" width="200" prop="actual_pay_amount"></el-table-column>
        <el-table-column label="购彩券抵扣金额" width="200" prop="offset_pay_amount"></el-table-column>
        <el-table-column label="彩金抵扣金额" width="200" prop="lottery_offset_pay_amount"></el-table-column>
        <el-table-column label="订单状态" width="">
            <template scope="scope">
                <el-tag type="primary">{{scope.row.order_type}}</el-tag>
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
                    order_number : "2017040511510124570",
                    lottery_type : "双色球",
                    lottery_issue_number : "20170410",
                    order_create_time : "2017-04-02",
                    lottery_number: "05 06 07 10 19 26 29 21",
                    account : "151****3574",
                    pay_amount : "¥ 123.00",
                    actual_pay_amount : "¥ 100.00",
                    offset_pay_amount : "¥ 20.00",
                    lottery_offset_pay_amount:"¥ 3.00",
                    order_type : "待出票",
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

