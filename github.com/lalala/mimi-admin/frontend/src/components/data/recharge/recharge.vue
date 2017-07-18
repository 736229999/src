<template>
  <div>
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="按天统计" name="day">
        <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
          <el-form :inline="true">
            <el-form-item label="开始时间">
              <el-date-picker v-model="form.start_time" type="datetime" placeholder="选择日期时间">
              </el-date-picker>
            </el-form-item>
            <el-form-item label="结束时间">
              <el-date-picker v-model="form.end_time" type="datetime" placeholder="选择日期时间">
              </el-date-picker>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="search">查询</el-button>
            </el-form-item>
          </el-form>
        </el-col>
        <el-col>
          <el-table :data="form.list" style="width: 100%" :default-sort="{prop: 'date', order: 'descending'}">
            <el-table-column prop="recharge_time" label="充值日期" :formatter="formatter"></el-table-column>
            <el-table-column sortable label="总充值金额/次数">
              <template scope="scope">
                <span>¥ {{scope.row.recharge_total_amount}} <b>/</b> {{scope.row.recharge_num}}次</span>
              </template>
            </el-table-column>
            <el-table-column sortable label="微信充值金额/次数">
              <template scope="scope">
                <span>¥ {{scope.row.wechat_recharge_amount}} <b>/</b> {{scope.row.wechat_recharge_num}}次</span>
              </template>
            </el-table-column>
            <el-table-column prop="" sortable label="支付宝充值金额/次数">
              <template scope="scope">
                <span>¥ {{scope.row.alipay_recharge_amount}} <b>/</b> {{scope.row.alipay_recharge_num}}次</span>
              </template>
            </el-table-column>
            <el-table-column prop="iphone_recharge_amount" sortable label="iOS充值金额/次数">
              <template scope="scope">
                <span>¥ {{scope.row.iphone_recharge_amount}} <b>/</b> {{scope.row.iphone_recharge_num}}次</span>
              </template>
            </el-table-column>
            <el-table-column prop="android_recharge_amount" sortable label="Android充值金额/次数">
              <template scope="scope">
                <span>¥ {{scope.row.android_recharge_amount}} <b>/</b> {{scope.row.android_recharge_num}}次</span>
              </template>
            </el-table-column>
          </el-table>
          <div class="block">
            <el-pagination
              layout="prev, pager, next"
              :total="form.total"
              @current-change="currentChange"
              :page-size="form.size"
              :current-page="form.page">
            </el-pagination>
          </div>

        </el-col>

      </el-tab-pane>
      <el-tab-pane label="按月统计" name="month">
        <el-table :data="form.list" style="width: 100%" :default-sort="{prop: 'date', order: 'descending'}">
          <el-table-column prop="recharge_time" label="充值月份" :formatter="formatter"></el-table-column>
          <el-table-column sortable label="总充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.recharge_total_amount}} <b>/</b> {{scope.row.recharge_num}}次</span>
            </template>
          </el-table-column>
          <el-table-column sortable label="微信充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.wechat_recharge_amount}} <b>/</b> {{scope.row.wechat_recharge_num}}次</span>
            </template>
          </el-table-column>
          <el-table-column prop="" sortable label="支付宝充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.alipay_recharge_amount}} <b>/</b> {{scope.row.alipay_recharge_num}}次</span>
            </template>
          </el-table-column>
          <el-table-column prop="iphone_recharge_amount" sortable label="iOS充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.iphone_recharge_amount}} <b>/</b> {{scope.row.iphone_recharge_num}}次</span>
            </template>
          </el-table-column>
          <el-table-column prop="android_recharge_amount" sortable label="Android充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.android_recharge_amount}} <b>/</b> {{scope.row.android_recharge_num}}次</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="按年度统计" name="year">
        <el-table :data="form.list" style="width: 100%" :default-sort="{prop: 'date', order: 'descending'}">
          <el-table-column prop="recharge_time" label="充值年份" :formatter="formatter"></el-table-column>
          <el-table-column sortable label="总充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.recharge_total_amount}} <b>/</b> {{scope.row.recharge_num}}次</span>
            </template>
          </el-table-column>
          <el-table-column sortable label="微信充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.wechat_recharge_amount}} <b>/</b> {{scope.row.wechat_recharge_num}}次</span>
            </template>
          </el-table-column>
          <el-table-column prop="" sortable label="支付宝充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.alipay_recharge_amount}} <b>/</b> {{scope.row.alipay_recharge_num}}次</span>
            </template>
          </el-table-column>
          <el-table-column prop="iphone_recharge_amount" sortable label="iOS充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.iphone_recharge_amount}} <b>/</b> {{scope.row.iphone_recharge_num}}次</span>
            </template>
          </el-table-column>
          <el-table-column prop="android_recharge_amount" sortable label="Android充值金额/次数">
            <template scope="scope">
              <span>¥ {{scope.row.android_recharge_amount}} <b>/</b> {{scope.row.android_recharge_num}}次</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>

  </div>
</template>

<script>
  import reApi from "../../../api/recharge"
  import Vue from "vue"
  export default {
    data() {
      return {
        activeName: 'day',
        tab: "day",
        form: {
          start_time: 0,
          end_time: 0,
          total: 0,
          page: 1,
          size: 100,
          list: [],
        }
      }
    }, methods: {
      currentChange(val) {
        this.form.page = val;
        this.GetRechargeList();
      },
      search: function () {
        let moment = require("moment");
        this.form.start_time = moment(this.form.start_time).unix();
        this.form.end_time = moment(this.form.end_time).unix();

        this.GetRechargeList()
      },
      GetRechargeList: function () {
        let moment = require("moment");
        reApi.getRechargeList(this.form).then((res) => {
          let obj = res.data.msg;
          if (obj.start_time < 1) {
            obj.start_time = ""
          } else {
            obj.start_time = moment(obj.start_time * 1000).toDate();
          }
          if (obj.end_time < 1) {
            obj.end_time = ""
          } else {
            obj.end_time = moment(obj.end_time * 1000).toDate();
          }
          this.form = obj
        });
      },
      formatter(row, column) {
        let moment = require("moment");
        if (this.tab == "day") {
          return moment.unix(row.recharge_time).format('YYYY-MM-DD');
        } else if (this.tab == "month") {
          return moment.unix(row.recharge_time).format('YYYY-MM');
        } else if (this.tab == "year") {
          return moment.unix(row.recharge_time).format('YYYY');
        }
      },
      handleClick(tab, event) {
        //月份.
        this.tab = tab.name;
        if (tab.name == "day") {
          this.GetRechargeList()
        } else if (tab.name == "month") {
          reApi.getRechargeListByMonth({}).then((res) => {
            let obj = res.data.msg.list;
            Vue.set(this.form, "list", obj)
          })
        } else if (tab.name == "year") {
          reApi.getRechargeListByYear({}).then((res) => {
            let obj = res.data.msg.list;
            Vue.set(this.form, "list", obj)
          })
        }
      },
    },
    mounted() {
      this.GetRechargeList()
    }
  }

</script>
