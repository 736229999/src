<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="充值时间">
          <el-col :span="23">
            <el-date-picker type="date" placeholder="选择日期" v-model="form.date" style="width: 100%;"></el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="手机号">
        <el-col :span="23">
            <el-input v-model="form.account" placeholder="手机号"></el-input>
            </el-col>
        </el-form-item>
        <el-form-item label="充值金额">
            <el-col :span="23">
                <el-input v-model="form.order_num" placeholder="充值金额"></el-input>
            </el-col>
        </el-form-item>
        <el-form-item label="充值状态">
             <el-select v-model="form.order_status" placeholder="请选择充值状态">
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
      <el-table-column label="流水号"   prop="serial_number"></el-table-column>
      <el-table-column label="充值时间" prop="rechange_money_time"></el-table-column>
      <el-table-column label="充值金额" prop="rechange_money_amount"></el-table-column>
      <el-table-column label="账户余额" prop="balance"></el-table-column>
      <el-table-column label="资金来源" prop="money_source"></el-table-column>
      <el-table-column label="充值账户" prop="rechange_money_account"></el-table-column>
        <el-table-column label="充值状态" >
            <template scope="scope">
                <el-tag type="success">{{scope.row.rechange_money_status}}</el-tag>
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
                    serial_number : "2017040500132156",
                    rechange_money_time : "2017-04-05",
                    rechange_money_amount : "¥ 2000.00",
                    balance : "¥ 2001.00",
                    money_source : "支付宝",
                    rechange_money_account : "151****3574",
                    rechange_money_status : "充值成功",
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

