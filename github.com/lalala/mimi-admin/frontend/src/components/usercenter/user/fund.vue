<template>
  <div>
    <el-row :gutter="24">
      <el-col :span="24"><h3 style="margin-left: 10px;">资金统计</h3></el-col>
      <el-col :span="24">
        <el-col :span="5" class="box">
          <div class="black cj">充值</div>
          <div class="black_right">¥ {{form.recharge}}</div>
        </el-col>
        <el-col :span="6" class="box">
          <div class="black tx">提现</div>
          <div class="black_right">¥ {{form.withdraw}}</div>
        </el-col>
        <el-col :span="6" class="box">
          <div class="black gc">购彩</div>
          <div class="black_right">¥ {{form.buycai}}</div>
        </el-col>
        <el-col :span="6" class="box">
          <div class="black zj">中奖</div>
          <div class="black_right">¥ {{form.winning}}</div>
        </el-col>
      </el-col>
      <el-col :gutter="24" style="margin-top: 20px;margin-left: 10px;padding-right: 20px;">
        <el-tabs v-model="data" type="card" @tab-click="handleClick">
          <el-tab-pane label="购彩" name="buycai" style="padding-bottom:15px;">
            <el-table :data="tableData" border style="width: 100%;">
              <el-table-column prop="id" label="序号" width="180"></el-table-column>
              <el-table-column prop="name" label="姓名" width="180"></el-table-column>
              <el-table-column prop="address" label="地址"></el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="奖金" name="winning" style="padding-bottom:15px;">
            <el-table :data="tableData" border style="width: 100%">
              <el-table-column prop="date" label="日期" width="180"></el-table-column>
              <el-table-column prop="name" label="姓名" width="180"></el-table-column>
              <el-table-column prop="address" label="地址"></el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="充值" editable name="recharge" style="padding-bottom:15px;">
            <el-table :data="form.recharge_list" border style="width: 100%">
              <el-table-column prop="id" label="序号" width="180"></el-table-column>
              <el-table-column prop="money" label="充值金额" width="180"></el-table-column>
              <el-table-column label="充值时间" width="300">
                <template scope="scope">
                  <span>{{new Date(scope.row.recharge_time * 1000).toLocaleString()}}</span>
                </template>
              </el-table-column>
              <el-table-column prop="order_id" label="充值订单号"></el-table-column>
              <el-table-column prop="source" label="充值来源"></el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="提现" name="withdraw" style="padding-bottom:15px;">
            <el-table :data="form.withdraw_list" border style="width: 100%">
              <el-table-column prop="id" label="序号" width="180"></el-table-column>
              <el-table-column prop="realname" label="姓名" width="180"></el-table-column>
              <el-table-column prop="idcard_no" label="身份证号码" width="180"></el-table-column>
              <el-table-column prop="phone" label="手机号" width="180"></el-table-column>
              <el-table-column label="提现时间" width="180">
                <template scope="scope">
                  <span>{{new Date(scope.row.create_time * 1000).toLocaleString()}}</span>
                </template>
              </el-table-column>
              <el-table-column prop="amount" label="提现金额" width="180"></el-table-column>
              <el-table-column prop="in_no" label="转入账号" width="180"></el-table-column>
              <el-table-column prop="status" label="提款状态" width="180">
                <template scope="scope">
                  <span v-if="scope.row.status == 0">待审核</span>
                  <span v-if="scope.row.status == 1">审核通过</span>
                  <span v-if="scope.row.status == 2">审核未通过</span>
                </template>
              </el-table-column>
              <el-table-column prop="withdraw_type" label="提现方式" width="180">
                <template scope="scope">
                  <span v-if="scope.row.withdraw_type == 0">支付宝</span>
                  <span v-if="scope.row.withdraw_type == 1">银行卡</span>
                </template>
              </el-table-column>
              <el-table-column prop="auditor" label="审核人" width="180"></el-table-column>
              <el-table-column label="审核时间" width="180">
                <template scope="scope">
                  <span v-if="scope.row.audit_time != 0">{{new Date(scope.row.audit_time * 1000).toLocaleString()}}</span>
                </template>
              </el-table-column>
              <el-table-column prop="audit_comment" label="审核备注" width="180"></el-table-column>
              <el-table-column prop="out_no" label="转出账号" width="180"></el-table-column>
              <el-table-column prop="out_sn" label="流水号" width="180"></el-table-column>
              <el-table-column label="转出时间" width="180">
                <template scope="scope">
                  <span v-if="scope.row.audit_time != 0">{{new Date(scope.row.trans_time * 1000).toLocaleString()}}</span>
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import ucApi from "../../../api/usercenter"
  import Vue from 'vue'
  export default {
    data() {
      return {
        data: "recharge",
        tableData:[],
        form: {
          recharge: "200.00",
          withdraw: "2,040,398.00",
          buycai: "2.00",
          winning: "2040600.00",
          recharge_list:[],
          withdraw_list:[],
        }
      };
    },
    methods: {
      handleClick(tab, event) {
        switch (tab.name) {
          case "buycai":
              this.GetBuycaiList();
              break;
          case "recharge":
              this.GetFund();
              break;
          case "withdraw":
              this.GetWithdraw();
              break;
          case "winning":
              this.GetWinning();
              break;
        }
      },
      GetBuycaiList:function () {
        console.log("获取购彩")
      },
      GetWithdraw:function () {
        ucApi.getWithdrawList({id:this.$route.params.id}).then((res) => {
          console.log("es.data.msg.list:", res.data.msg.list);
          Vue.set(this.form, "withdraw_list", res.data.msg.list);
        })
      },
      GetWinning:function () {
        console.log("获取中奖")
      },
      GetFund: function () {
        ucApi.getFund({id: this.$route.params.id}).then((res) => {
          this.form = res.data.msg
        })
      }
    },
    mounted() {
      this.GetFund();
    }
  };
</script>

<style>

  div {
    color: #475669;
    font-family: Helvetica Neue, Helvetica, PingFang SC, Hiragino Sans GB, Microsoft YaHei, SimSun, sans-serif;
    font-size: 14px;
  }

  .box {
    box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.2);
    -webkit-box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.2);
    margin-left: 10px;
    /*padding-right: 50px;*/
  }

  .black {
    width: 80px;
    height: 80px;
    float: left;
    border-radius: 50%;
    text-align: center;
    line-height: 80px;
    color: #ffffff;
    font-size: 18px;
    margin: 10px;
  }

  .cj {
    background: #ef553a;
  }

  .tx {
    background: #26c281;
  }

  .gc {
    background: #57889c;
  }

  .zj {
    background: #FF4949;
  }

  .black_right {
    float: left;
    margin-left: 30px;
    height: 100px;
    line-height: 100px;
    font-size: 36px;
  }

  .el-tabs__header {
    padding-bottom: 0px;
  }
</style>
