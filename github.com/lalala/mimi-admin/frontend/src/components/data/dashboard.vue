<template>
  <div class="cont">
    <el-row :gutter="24">
      <el-col :span="12">
        <el-col :span="24"><h3>收入</h3></el-col>
        <el-col :span="12">
          <div class="dashboard_income_order dashboard_box">
            <div class="top income_top_bg">¥</div>
            <div class="dashboard_total">
              <div class="top_amount">¥ {{top.new_income}}</div>
              <div class="top_desc">今日收入</div>
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="dashboard_income_order dashboard_box">
            <div class="top total_income_top_bg">¥</div>
            <div class="dashboard_total">
              <div class="top_amount">¥ {{top.income_total}}</div>
              <div class="top_desc">历史总收入</div>
            </div>
          </div>
        </el-col>
      </el-col>
      <el-col :span="12">
        <el-col :span="24"><h3>用户订单</h3></el-col>
        <el-col :span="12">
          <div class="dashboard_income_order dashboard_box">
            <div class="top user_order_top_bg">订</div>
            <div class="dashboard_total">
              <div class="top_amount">{{top.new_user_order}}</div>
              <div class="top_desc">今日新增用户订单</div>
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="dashboard_income_order dashboard_box">
            <div class="top total_user_order_top_bg">订</div>
            <div class="dashboard_total">
              <div class="top_amount">{{top.user_order_total}}</div>
              <div class="top_desc">历史用户总订单</div>
            </div>
          </div>
        </el-col>
      </el-col>
      <el-col :span="12">
        <el-col :span="24"><h3>购彩订单</h3></el-col>
        <el-col :span="12">
          <div class="dashboard_income_order dashboard_box">
            <div class="top vendor_order_top_bg">订</div>
            <div class="dashboard_total">
              <div class="top_amount">{{top.new_buycai_order}}</div>
              <div class="top_desc">今日新增购彩订单</div>
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="dashboard_income_order dashboard_box">
            <div class="top total_vendor_order_top_bg">订</div>
            <div class="dashboard_total">
              <div class="top_amount">{{top.buycai_order_total}}</div>
              <div class="top_desc">历史购彩总订单</div>
            </div>
          </div>
        </el-col>
      </el-col>
      <el-col :span="12">
        <el-col :span="24"><h3>用户数</h3></el-col>
        <el-col :span="12">
          <div class="dashboard_income_order dashboard_box">
            <div class="top user_top_bg">人</div>
            <div class="dashboard_total">
              <div class="top_amount">{{user.new_user_num}}</div>
              <div class="top_desc">今日新增用户</div>
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="dashboard_income_order dashboard_box">
            <div class="top total_user_top_bg">人</div>
            <div class="dashboard_total">
              <div class="top_amount">{{user.total_user_num}}</div>
              <div class="top_desc">总用户数</div>
            </div>
          </div>
        </el-col>
      </el-col>
      <el-col>

      </el-col>
    </el-row>
  </div>
</template>

<style>
  .dashboard_income_order {
    width: 100%;
    height: 100px;
  }

  .dashboard_income_order .top {
    width: 80px;
    height: 80px;
    margin: 10px 0px 10px 10px;
    font-size: 3em;
    line-height: 80px;
    text-align: center;
    -webkit-border-radius: 50%;
    -moz-border-radius: 50%;
    border-radius: 50%;
    float: left;
    color: #ffffff;
  }

  .dashboard_box {
    box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.2);
    -webkit-box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.2);
  }

  .dashboard_box:hover {
    box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.5);
    -webkit-box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.5);
    transition: box-shadow 0.8s;
  }

  .dashboard_total {
    width: 260px;
    height: 80px;
    float: left;
    margin: 10px 0px 10px 20px;
    /*margin-left: 20px;*/
  }

  .dashboard_total .top_amount {
    width: 260px;
    height: 50px;
    float: left;
    line-height: 50px;
    font-size: 2em;
    font-weight: bold;
  }

  .dashboard_total .top_desc {
    width: 260px;
    height: 20px;
    float: left;
    line-height: 20px;
    font-size: 16px;
    color: #bbbbbb;
    font-weight: bold;
  }

  .income_top_bg {
    background: #26c281;
  }

  .total_income_top_bg {
    background: #ef553a;
  }

  .user_order_top_bg {
    background: #57889c;
  }

  .total_user_order_top_bg {
    background: #FF4949;
  }

  .vendor_order_top_bg {
    background: #20A0FF;
  }

  .total_vendor_order_top_bg {
    background: #475669;
  }
  .user_top_bg {
    background: #CC6666;
  }

  .total_user_top_bg {
    background: #66CCFF;
  }

</style>

<script>
  import dataApi from "../../api/data"
  export default {
    data() {
      return {
        top: {},
        user:{
            new_user_num:"获取失败",
            total_user_num:"获取失败",
        },
      }
    }, methods: {
      //获取收入和订单的情况.
      getOrderAndIncome: function () {
        dataApi.getOrderAndIncome({}).then((res) => {
          this.top = res.data.msg
          this.top.new_income = this.moneyFormat(this.top.new_income);
          this.top.income_total = this.moneyFormat(this.top.income_total)
        });
      },
      getUserStatistics:function () {
        dataApi.getUserStatistics({}).then((res)=>{
            console.log(res.data.msg)
          this.user = res.data.msg
        })
      },
      moneyFormat: function (value) {
        value = value.toString();
        let arrStr = value.split('.'),
          floatPart = arrStr[1],
          step = 3,
          len = 0;
        value = arrStr[0];
        len = value.length;

        if (len > step) {
          let c1 = len % step,
            c2 = parseInt(len / step),
            arr = [],
            first = value.substr(0, c1);
          if (first != '') {
            arr.push(first);
          }

          for (let i = 0; i < c2; i++) {
            arr.push(value.substr(c1 + i * step, step));
          }

          value = arr.join(',');
        }
        return value + '.' + floatPart;
      },
    },
    mounted() {
      //获取订单和收入的情况.
      this.getOrderAndIncome();

      //获取用户数.
      this.getUserStatistics();
    }
  }

</script>
