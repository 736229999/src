<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 20px;">
      <el-col :span="24">
        <router-link :to="{path:'/lottery/home/options/add'}">
          <el-button type="primary">新增彩种</el-button>
        </router-link>
      </el-col>
    </el-col>
    <!--列表-->

    <el-table :data="data">
      <el-table-column label="id" prop="id"></el-table-column>
      <el-table-column label="名称" prop="lottery_name"></el-table-column>
      <el-table-column label="是否加奖">
        <template scope="scope">
          <span v-if="scope.row.is_plus_award == true">是</span>
          <span v-else-if="scope.row.is_plus_award == false">否</span>
        </template>
      </el-table-column>
      <el-table-column label="开奖简介" prop="info"></el-table-column>
      <el-table-column label="停止销售">
        <template scope="scope">
          <span v-if="scope.row.stop_sale == true">是</span>
          <span v-else>否</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" prop="create_time" :formatter="formatter"></el-table-column>
      <el-table-column label="更新时间" prop="update_time" :formatter="formatter"></el-table-column>
      <el-table-column label="操作">
        <template scope="scope">
          <router-link :to="{path:'/lottery/home/options/edit/'+scope.row.id}"><el-button type="warning" size="mini" icon="edit"></el-button></router-link>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>

  import lotteryApi from "../../api/lottery"
  import ElButton from "../../../node_modules/element-ui/packages/button/src/button";

  export default {
    components: {ElButton},
    data() {
      return {
        data: [],
      }
    },
    methods: {
      GetLotteryOptionsList: function () {
        lotteryApi.lotteryOptionsList().then((res) => {
          console.log(res);
          this.data = res.data.msg
        })
      },
      formatter: function (row, column) {
        let moment = require("moment");
        if (column.property == "update_time") {
          return moment(row.update_time * 1000).format('YYYY-MM-DD HH:mm:ss');
        } else if (column.property == "create_time") {
          return moment(row.create_time * 1000).format('YYYY-MM-DD HH:mm:ss');
        }
      }
      },
      mounted() {
        this.GetLotteryOptionsList();
      }
    }
</script>

<style>
  .block {
    float: right;
    margin-top: 20px;
  }

  .list {
    margin-top: 20px;
  }
</style>
