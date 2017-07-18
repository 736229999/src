<template>
  <div>
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="手机号">
          <el-col :span="23">
            <el-input v-model="form.mobile" placeholder="标题"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="search">查询</el-button>
          <router-link :to="{path:'/activity/cdkey/add'}">
            <el-button type="primary" icon="plus">添加</el-button>
          </router-link>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-table :data="form.list" style="width: 100%">
        <el-table-column prop="id" label="序号" width="180"></el-table-column>
        <el-table-column prop="username" label="用户昵称" width="180"></el-table-column>
        <el-table-column prop="info" label="中奖信息"></el-table-column>
        <el-table-column label="创建时间">
          <template scope="scope">
            <span>{{new Date(scope.row.valid_start * 1000).toLocaleString()}}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template scope="scope">
            <router-link :to="{path:'/activity/cdkey/detail/'+scope.row.id}">
              <el-button type="primary" size="mini" icon="more"></el-button>
            </router-link>
            <router-link :to="{path:'/activity/cdkey/edit/'+scope.row.id}">
              <el-button type="warning" size="mini" icon="edit"></el-button>
            </router-link>
            <el-button type="danger" size="mini" icon="delete" @click="del(scope.row.id)"></el-button>
            <el-button type="success" size="mini" icon="upload" @click="exportCsv(scope.row.id)"></el-button>
            <a href="" id="downloadLink" download="asd.csv" style="display: none"></a>
          </template>
        </el-table-column>
      </el-table>
      <div class="block">
        <el-pagination @current-change="handleCurrentChange" :current-page="form.page" :page-size="form.size"
                       layout="total, prev, pager, next" :total="form.total"></el-pagination>
      </div>
    </el-col>
  </div>
</template>

<script>

  import ElFormItem from "../../../node_modules/element-ui/packages/form/src/form-item";
  import ElCol from "element-ui/packages/col/src/col";
  import ElButton from "../../../node_modules/element-ui/packages/button/src/button";
  import winningApi from "../../api/winning"

  export default {
    components: {
      ElButton,
      ElCol,
      ElFormItem
    },
    data() {
      return {
        form: {
          mobile: "",
          title: "",
          total: 0,
          size: 50,
          page: 1,
          list: [],
        },
        disabled: false,
      }
    },
    methods: {
      //获取中奖列表.
      GetWinningList: function () {
        winningApi.getWinningList(this.form).then((res) => {
          this.form = res.data.msg
        }).catch(err => {
          this.$message.error("获取失败")
        })
      },
      handleCurrentChange(val) {
        this.form.page = val;
        this.GetCdkeyList();
      },
      search: function () {
        this.GetCdkeyList();
      },
      del: function (id) {
        this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {

          acApi.delCdkeyById({id: id}).then((res) => {
            this.$message.success("删除成功");
            this.GetCdkeyList();
          }).catch(err => {
            this.$message.error("删除失败");
          })
        })
      },
    },
    mounted() {
      this.GetWinningList();
    }
  }
</script>

<style>

  .form_item {
    margin-top: 15px;
    margin-bottom: 15px;
  }

  .ri {
    text-align: right;
  }

  .specifyTypeShow {
    display: none;
  }

  .specifyTypeNotShow {
    display: block;
  }
</style>

