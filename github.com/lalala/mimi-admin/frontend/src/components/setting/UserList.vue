<template>
  <section>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true" :model="filters">
        <el-form-item>
          <el-input v-model="filters.email" placeholder="邮箱"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" v-on:click="getUsers">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleAdd">新增</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!--列表-->
    <el-table :data="users" style="width: 100%">
      <el-table-column type="selection">
      </el-table-column>
      <el-table-column label="账号" prop="email">
      </el-table-column>
      <el-table-column label="昵称" prop="username">
      </el-table-column>
      <el-table-column label="联系方式" prop="mobile">
      </el-table-column>
      <el-table-column label="创建人" prop="creator">
      </el-table-column>
      <el-table-column label="状态" prop="status" width="70">
      </el-table-column>
      <el-table-column label="操作" width="150">
        <template scope="scope">
          <el-button size="small" @click="handleEdit(scope.$email, scope.row)">编辑</el-button>
          <el-button type="danger" size="small" @click="handleDel(scope.row.email)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--工具条-->
    <el-col :span="24" class="toolbar">
      <el-pagination layout="prev, pager, next" @current-change="handleCurrentChange" :page-size="20" :total="total" style="float:right;">
      </el-pagination>
    </el-col>
    <!--编辑界面-->
    <el-dialog title="编辑" v-model="editFormVisible" :close-on-click-modal="false">
      <el-form :model="editForm" label-width="80px" :rules="editFormRules" ref="editForm">
        <el-form-item label="账号" prop="email">
          {{editForm.email}}
        </el-form-item>
        <el-form-item label="密码" prop="password" type="password">
          <el-input v-model="editForm.password" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="昵称" prop="name">
          <el-input v-model="editForm.name" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="是否激活">
          <el-radio-group v-model="editForm.status">
            <el-radio class="radio" :label="1">激活</el-radio>
            <el-radio class="radio" :label="0">未激活</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="联系方式">
          <el-input v-model="editForm.phone" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="其他信息">
          <el-input v-model="editForm.other" auto-complete="off" type="textarea"></el-input>
        </el-form-item>
        <el-form-item label="用户组" prop="roles">
          <el-select v-model="editForm.role_ids" multiple placeholder="请选择/搜索" filterable>
            <el-option v-for="item in roles" :label="item.name" :value="item.id" :key="item.id">
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click.native="editFormVisible = false">取消</el-button>
        <el-button type="primary" @click.native="editSubmit" :loading="editLoading">提交</el-button>
      </div>
    </el-dialog>
    <!--新增界面-->
    <el-dialog title="新增" v-model="addFormVisible" :close-on-click-modal="false">
      <el-form :model="addForm" label-width="80px" :rules="addFormRules" ref="addForm">
        <el-form-item label="账号" prop="email">
          <el-input v-model="addForm.email" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password" type="password">
          <el-input v-model="addForm.password" auto-complete="off" type="password"></el-input>
        </el-form-item>
        <el-form-item label="昵称" prop="name">
          <el-input v-model="addForm.name" auto-complete="off"></el-input>
        </el-form-item>
<!--         <el-form-item label="是否激活">
          <el-radio-group v-model="addForm.status">
            <el-radio class="radio" :label="1">激活</el-radio>
            <el-radio class="radio" :label="0">未激活</el-radio>
          </el-radio-group>
        </el-form-item> -->
        <el-form-item label="联系方式">
          <el-input v-model="addForm.phone" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="其他信息">
          <el-input v-model="addForm.other" auto-complete="off" type="textarea"></el-input>
        </el-form-item>
        <el-form-item label="用户组" prop="roles">
          <el-select v-model="addForm.role_ids" multiple placeholder="请选择/搜索" filterable>
            <el-option v-for="item in roles" :label="item.name" :value="item.id" :key="item.id">
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click.native="addFormVisible = false">取消</el-button>
        <el-button type="primary" @click.native="addSubmit" :loading="addLoading">提交</el-button>
      </div>
    </el-dialog>
  </section>
</template>

<script>
  import util from '../../common/js/util'
  import NProgress from 'nprogress'
  import {
    userList,
    addUser,
    deleteUser,
    editUser,
    roleListSimple
  } from '../../api/api';

  export default {
    data() {
        return {
          filters: {
            email: ''
          },
          users: [],
          total: 0,
          page: 1,
          listLoading: false,
          sels: [], //列表选中列

          //编辑界面数据
          editFormVisible: false, //编辑界面是否显示
          editLoading: false,
          editFormRules: {
            name: [{ required: true,message: '请输入姓名',trigger: 'blur' }]
          },
          editForm: {
            email: '',
            password: '',
            name: '',
            status: 0,
            phone: '',
            other: '',
            role_ids: []
          },

          addFormVisible: false, //新增界面是否显示
          addLoading: false,
          addFormRules: {
            email: [
            { required: true, message: '请输入邮箱', trigger: 'blur' }],
            password: [
            { required: true, message: '请输入密码', trigger: 'blur' }],
            name: [
            { required: true, message: '请输入昵称', trigger: 'blur' }]
          },
          //新增界面数据
          addForm: {
            email: '',
            password: '',
            name: '',
            status: 0,
            phone: '',
            other: '',
            role_ids: []
          },
          roles: []
        }
      },
      methods: {
        // 获取列表
        getUsers() {

            var params = new URLSearchParams();
            params.append('page', this.page);
            params.append('size', 20);

            this.listLoading = true;
            NProgress.start();
            userList(params).then((res) => {
                var user = JSON.parse(res.data);
                let data = user.msg;
                console.log(data);
                this.total = data.length;
                this.users = data;
//                this.listLoading = false;
//                NProgress.done();
            });
        },
        // 获取列表
        getRoles() {
          let params = new URLSearchParams();
          this.listLoading = true;
          NProgress.start();
          roleListSimple(params).then((res) => {
            this.roles = res.data.msg;
            this.listLoading = false;
            NProgress.done();
          });
        },
        //显示转换
//        formatActivate: function(row, column) {
//          return row.status > 0 ? '激活' : '未激活';
//        },
        handleCurrentChange(val) {
          this.page = val;
          this.getUsers();
        },
        //删除
        handleDel: function(index) {
          this.$confirm('确认删除该记录吗?', '提示', { type: 'warning' }).then(() => {
            this.listLoading = true;
            NProgress.start();

            let params = new URLSearchParams();
            params.append('email', index);
            deleteUser(params).then((res) => {
              this.listLoading = false;
              NProgress.done();
              if ( res.status !== 200 ) {
                this.$notify.error( res.data.msg );
              }else {
                this.$notify.success('删除成功！');
                this.getUsers();
              }
            });
          }).catch(() => {

          });
        },
        //显示编辑界面
        handleEdit: function(index, row) {
          this.editFormVisible = true;
          this.editForm = Object.assign({}, row);
        },
        //显示新增界面
        handleAdd: function() {
          this.addFormVisible = true;
          this.addForm = {
            email: '',
            password: '',
            name: '',
            status: 0,
            phone: '',
            other: '',
            role_ids: []
          };
        },
        //编辑
        editSubmit: function() {
          this.$refs.editForm.validate((valid) => {
            if (valid) {
              this.$confirm('确认提交吗？', '提示', {}).then(() => {
                this.editLoading = true;
                NProgress.start();

                let params = new URLSearchParams();
                params.append('email', this.editForm.email);
                params.append('password', this.editForm.password);
                params.append('name', this.editForm.name);
                params.append('status', this.editForm.status);
                params.append('phone', this.editForm.phone);
                params.append('other', this.editForm.other);
                params.append('role_ids', this.editForm.role_ids);
                editUser(params).then((res) => {
                  this.editLoading = false;
                  NProgress.done();

                  if ( res.status !== 200 ) {
                    this.$notify.error( res.data.msg );
                  } else {
                    this.$notify.success('提交成功！')
                    this.$refs['editForm'].resetFields();
                    this.editFormVisible = false;
                    this.getUsers();
                  }
                });
              });
            }
          });
        },
        //新增
        addSubmit: function() {
          this.$refs.addForm.validate((valid) => {
            if (valid) {
              this.$confirm('确认提交吗？', '提示', {}).then(() => {
                this.addLoading = true;
                NProgress.start();

                let params = new URLSearchParams();
                params.append('email', this.addForm.email);
                params.append('password', this.addForm.password);
                params.append('name', this.addForm.name);
                params.append('phone', this.addForm.phone);
                params.append('other', this.addForm.other);
                params.append('role_ids', this.addForm.role_ids);
                addUser(params).then((res) => {
                  this.addLoading = false;
                  NProgress.done();

                  if ( res.status !== 200 ) {
                    this.$notify.error( res.data.msg );
                  } else {
                    this.$notify.success('提交成功！')
                    this.$refs['addForm'].resetFields();
                    this.addFormVisible = false;
                    this.getUsers(); 
                  }
                });
              });
            }
          });
        },
        selsChange: function(sels) {
          this.sels = sels;
        }
      },
      mounted() {
        //this.getRoles();
        this.getUsers();
      }
  }
</script>

<style scoped>
  .table-expand {
    font-size: 0;
  }
  .table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 33%;
  }
</style>

