<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="昵称">
          <el-col :span="23">
            <el-input v-model="form.nickname" placeholder="昵称"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="手机号">
          <el-col :span="23">
            <el-input v-model="form.phone" placeholder="手机号"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="search">查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <el-table :data="form.list">
      <el-table-column label="序号" prop="id" width="100"></el-table-column>
      <el-table-column label="昵称" prop="nickname"></el-table-column>
      <el-table-column label="性别">
        <template scope="scope">
          <span v-if="scope.row.sex == 0">未知</span>
          <span v-if="scope.row.sex == 1">男</span>
          <span v-if="scope.row.sex == 1">女</span>
        </template>
      </el-table-column>
      <el-table-column label="手机" prop="phone"></el-table-column>
      <el-table-column label="邀请码" prop="invitation_code"></el-table-column>
      <el-table-column label="积分" prop="credits"></el-table-column>
      <el-table-column label="开心豆" prop="kxd"></el-table-column>
      <el-table-column label="是否被邀请">
        <template scope="scope">
          <el-tag type="primary" v-if="scope.row.is_invited == false">未被邀请</el-tag>
          <el-tag type="success" v-if="scope.row.is_invited == true">已被邀请</el-tag>
        </template>
      </el-table-column>
      <!--<el-table-column label="登录时间" prop="login_time"></el-table-column>-->
      <!--<el-table-column label="注册时间" prop="create_time"></el-table-column>-->
      <!--<el-table-column label="注册ip地址" prop=""></el-table-column>-->
      <!--<el-table-column label="登录ip地址" prop=""></el-table-column>-->
      <el-table-column label="操作">
        <template scope="scope">
          <router-link :to="{path:'/usercenter/user/detail/'+scope.row.id}">
            <el-button type="primary" size="mini"><icon name="ellipsis-h" scale="0.6"></icon></el-button>
          </router-link>
          <router-link :to="{path:'/usercenter/user/fund/'+scope.row.id}">
            <el-button type="success" size="mini"><icon name="money" scale="0.6"></icon></el-button>
          </router-link>
          <router-link :to="{path:'/usercenter/user/detail/'+scope.row.id}">
            <el-button type="warning" size="mini"><icon name="shopping-cart" scale="0.6"></icon></el-button>
          </router-link>
          <router-link :to="{path:'/usercenter/user/detail/'+scope.row.id}">
            <el-button type="info" size="mini"><icon name="qrcode" scale="0.6" label="资金明细"></icon></el-button>
          </router-link>
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

  import userApi from "../../../api/usercenter"
  var Icon = require('vue-awesome');
  export default {
    components: {
      Icon
    },
    data() {
      return {
        form: {
          total: 0,
          size: 100,
          page: 1,
          nickname: "",
          phone: "",
          list: [],
        },
      }
    },
    methods: {
      GetUserList: function () {
        userApi.getUserList(this.form).then((res) => {
          this.form = res.data.msg;
        })
      },
      formatter: function (row, column) {
        let moment = require("moment");
        return moment(row.create_time * 1000).format('YYYY-MM-DD HH:mm:ss');
      },
      handleCurrentChange(val) {
        this.form.page = val;
        this.GetUserList();
      },
      search:function () {
        this.GetUserList();
      }
    },
    mounted() {
      this.GetUserList();
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
