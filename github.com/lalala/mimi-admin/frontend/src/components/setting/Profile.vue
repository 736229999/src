<template>
  <el-form ref="user" :model="user.form" label-width="80px" @submit.prevent="onSubmit" style="margin:20px;width:60%;min-width:600px;">
    <h3>个人信息</h3>
    <el-form-item label="账号" prop="email">
      {{user.form.email}}
    </el-form-item>
    <el-form-item label="昵称" prop="phone">
      <el-input v-model="user.form.username" value=""></el-input>
    </el-form-item>
    <el-form-item label="联系方式" prop="phone">
      <el-input v-model="user.form.mobile"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click.native.prevent="updateProfile">更新信息</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
  import { userUpdate } from '../../api/api';
  export default {
    data() {
      var user = sessionStorage.getItem('user');
      //user = JSON.parse(user).msg;
      return {
        labelPosition: 'left',
        user: {
          form: {
            email: '',
            username: '',
            mobile: '',
          }
        }
      };
    },
    methods: {
      updateProfile() {
        this.$refs.user.validate((valid) => {
          if (valid) {
            var params = new URLSearchParams();
            params.append('username', this.user.form.username);
            params.append('mobile', this.user.form.mobile);
            userUpdate(params).then(res => {
                var u = JSON.parse(res.data);

              if (res.status !== 200) {
                this.$message.error( u.msg );
              } else {
                sessionStorage.setItem('user', JSON.stringify(res.data));
                this.$message.success('更新成功！');
              }
            });
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      }
    },
    mounted() {
      var user = sessionStorage.getItem('user');

        if (user) {
          var u = JSON.parse(user);
          console.log(u)
            var userinfo = u.msg;
        this.user.form.email = userinfo.email;
        this.user.form.username = userinfo.username;
        this.user.form.mobile = userinfo.mobile;
      }
    }
  }
</script>

