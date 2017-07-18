<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="彩种">
          <el-select v-model="form.lottery" @change="changeLottery" placeholder="请选择">
            <el-option v-for="item in lotteryType" :key="item.Code" :label="item.Name" :value="item.Code"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="期号">
          <el-col :span="23">
            <el-input v-model="form.issue" placeholder="期号"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="search">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!--列表-->
    <el-col>
      <el-table :data="data" border>
        <el-table-column label="彩种" prop="lottery_name"></el-table-column>
        <el-table-column label="期号" prop="issue"></el-table-column>
        <el-table-column label="开奖时间">
          <template scope="scope">
            <span>{{new Date(scope.row.open_time * 1000).toLocaleString()}}</span>
          </template>
        </el-table-column>
        <el-table-column label="中奖方案" prop="open_balls"></el-table-column>
        <el-table-column label="操作" width="100">
          <template scope="scope">
            <router-link :to="{path:'/lottery/open/history/'+scope.row.lottery}">
              <el-button @click="handleClick" type="primary" size="mini" icon="more"></el-button>
            </router-link>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
  </div>
</template>

<script>
  import winApi from '../../../api/winning';
  import lotteryApi from "../../../api/lottery"
  import ElFormItem from "../../../../node_modules/element-ui/packages/form/src/form-item";
  export default {
    components: {ElFormItem},
    data() {
      return {
        data: [],
        lotteryType: [],
        total: 0,
        form: {
          lottery: '',
          issue: "",
          size: 100,
          page: 1,
          total: 0,
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
        winApi.winningList(this.form).then((res) => {
          this.data = res.data.msg.list;
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
        if (this.form.lottery == "") {
          this.$message.error("请选择彩种");
          return
        }
        if (this.form.issue == "") {
          this.$message.error("请填写期号");
          return
        }

        winApi.getWinningByIssue(this.form).then((res) => {
          this.data = res.data.msg.list;
          console.log(this.data);
        })
      },

    },
    mounted() {
      this.GetLotteryList();
      this.GetWinning();
    }
  }
</script>

