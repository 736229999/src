<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item>
          <router-link :to="{path:'/user/role/add'}"><el-button type="primary">添加角色</el-button></router-link>
        </el-form-item>
      </el-form>
    </el-col>
    <!--列表-->
    <el-col>
      <el-table :data="form.list" border>
        <el-table-column label="序号" prop="id"></el-table-column>
        <el-table-column label="角色名称" prop="role_name"></el-table-column>
        <el-table-column label="备注" prop="remarks"></el-table-column>
        <el-table-column label="创建人" prop="creator_name"></el-table-column>
        <el-table-column label="创建时间" width="300">
          <template scope="scope">
            <span>{{new Date(scope.row.create_time * 1000).toLocaleString()}}</span>
          </template>
        </el-table-column>

        <el-table-column label="操作" width="100">
          <template scope="scope">
            <router-link :to="{path:'/user/role/edit/'+scope.row.id}">
              <el-button type="text" size="mini" icon="edit">编辑</el-button>
            </router-link>
            <el-button type="danger" size="mini" icon="delete">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
  </div>
</template>

<script>
  import roleApi from '../../../api/role';
  export default {
    data() {
      return {
        form: {
          list: [],
        },

      }
    }, methods: {
      GetRole() {
        roleApi.getRoleList(this.form).then((res) => {
          this.form = res.data.msg;
        }).catch(err => {

        });
      },

    },
    mounted() {
      this.GetRole();
    }
  }
</script>

