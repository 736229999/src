<template>
  <div>
    <el-table :data="form.list">
      <el-table-column label="序号" prop="id" width="100"></el-table-column>
      <el-table-column label="邮箱" prop="email"></el-table-column>
      <el-table-column label="姓名" prop="name"></el-table-column>
      <el-table-column label="反馈内容" prop="content"></el-table-column>
      <el-table-column label="进度">
        <template scope="scope">
          <span v-if="scope.row.status == 0"><el-tag type="warning">待处理</el-tag></span>
          <span v-if="scope.row.status == 1"><el-tag type="primary">已查阅</el-tag></span>
          <span v-if="scope.row.status == 2"><el-tag type="success">已处理</el-tag></span>
          <span v-if="scope.row.status == 3"><el-tag type="gray">已忽略</el-tag></span>
        </template>
      </el-table-column>
      <el-table-column label="反馈时间" prop="create_time" :formatter="formatter"></el-table-column>
      <el-table-column label="操作">
        <template scope="scope">
          <router-link :to="{path:'/options/feedback/detail/'+scope.row.id}">
            <el-button type="primary" size="mini" icon="more"></el-button>
          </router-link>
          <el-button type="danger" size="mini" icon="delete" @click="del(scope.row.id)"></el-button>
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

  import optionsApi from "../../../api/options"
  import ElButton from "../../../../node_modules/element-ui/packages/button/src/button";

  export default {
    components: {ElButton},
    data() {
      return {
        form: {
          total: 0,
          size: 100,
          page: 1,
          list: [],
        },
      }
    },
    methods: {
      GetFeedbackList: function () {
        optionsApi.getFeedbackList(this.form).then((res) => {
          this.form = res.data.msg
        })
      },
      formatter: function (row, column) {
        let moment = require("moment");
        return moment(row.create_time * 1000).format('YYYY-MM-DD HH:mm:ss');
      },
      handleCurrentChange(val) {
        this.form.page = val;
        this.GetFeedbackList();
      },
      del: function (id) {
        optionsApi.delFeedbackById({id: parseInt(id)}).then((res) => {
          this.$message.success("删除成功");
          this.GetFeedbackList();
        })
      }
    },
    mounted() {
      this.GetFeedbackList();
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
