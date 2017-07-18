<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 20px;">
      <el-col :span="24">
        <el-select v-model="lottery" @change="changeLottery" placeholder="请选择">
          <el-option v-for="item in lotteryType" :key="item.Code" :label="item.Name" :value="item.Code"></el-option>
        </el-select>
        <router-link :to="{path:'/lottery/buycai/options/add'}">
          <el-button type="primary">新增配置</el-button>
        </router-link>
      </el-col>
    </el-col>
    <!--列表-->

    <el-table :data="data">
      <el-table-column label="id" prop="id"></el-table-column>
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
      <el-table-column label="操作">
        <template scope="scope">
          <router-link :to="{path:'/lottery/buycai/options/edit/'+scope.row.id+'/lottery/'+scope.row.lottery}">
            <el-button type="primary" size="mini" icon="edit"></el-button>
          </router-link>
          <el-button icon="delete" size="mini" type="danger"
                     @click="delIssue(scope.row.id)"></el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="block">
      <el-pagination @current-change="handleCurrentChange" :current-page="form.page" :page-size="form.size"
                     layout="total, prev, pager, next" :total="form.total"></el-pagination>
    </div>
  </div>
</template>

<script>
  import lotteryApi from "../../api/lottery"

  export default {
    data() {
      return {
        data: [],
        lotteryType: [],
        lottery: "",
        form: {
          page: 1,
          size: 100,
          lottery: "",
          buycai: [],
          total: 1,
        }
      }
    },
    methods: {
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
      GetBuycaiOptions: function () {
        lotteryApi.buycaiOptions(this.form).then((res) => {
          this.data = res.data.msg.buycai;
          for (let i = 0; i < this.data.length; i++) {
            this.data[i].lottery = this.form.lottery
          }
          this.form.total = res.data.msg.total;
          this.form.page = res.data.msg.page;
          this.form.size = res.data.msg.size;
        }).catch(res => {

        })
      },
      handleSizeChange(val) {
        console.log(`每页 ${val} 条`);
      },
      handleCurrentChange(val) {
        this.form.page = val;
        this.GetBuycaiOptions();
      },
      changeLottery: function (val) {
        this.form.lottery = val;
        this.GetBuycaiOptions();
      },
      delIssue: function (id) {
        this.$confirm('确定删除吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          lotteryApi.delIssueById({id: id, lottery: this.form.lottery}).then((res) => {
            this.$message.success("删除成功");
            this.GetLotteryList();
          }).catch(err => {
              this.$message.error(err.data.msg)
          })
        });
      },
    },
    mounted() {
      this.GetLotteryList();
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
