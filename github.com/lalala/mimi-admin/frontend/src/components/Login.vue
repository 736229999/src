<template>
  <el-form :model="form" :rules="rule" ref="form" label-position="left" label-width="0px" class="login-container">
    <h3 class="title">系统登录</h3>
    <el-form-item prop="email">
      <el-input type="text" v-model="form.email" auto-complete="off" placeholder="账号"></el-input>
    </el-form-item>
    <el-form-item prop="checkPass">
      <el-input type="password" v-model="form.checkPass" auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-form-item style="width:100%;">
      <el-button type="primary" style="width:100%;" @click.native.prevent="loginSubmit" :loading="logining">登录</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
  import userApi  from '../api/user';
  import NProgress from 'nprogress'
  import Vue from 'vue'
  export default {
    data() {
      return {
        logining: false,
        form: {
          email: '1556469348@qq.com',
          checkPass: 'a123456'
        },
        rule: {
          email: [
            { required: true, message: '请输入账号', trigger: 'blur' },
          ],
          checkPass: [
            { required: true, message: '请输入密码', trigger: 'blur' },
          ]
        },
        checked: true
      };
    },
    methods: {
      loginSubmit(ev) {
        var _this = this;
        this.$refs.form.validate((valid) => {
          if (valid) {
            this.logining = true;
            NProgress.start();
            let params = new URLSearchParams();
            params.append('email', this.form.email);
            params.append('password', this.form.checkPass);
            userApi.userLogin(params).then(res => {

                sessionStorage.setItem('user', res.data.msg);
                this.$router.push({ path: '/dashboard' });

            }).catch(err => {

                console.log(err)
               this.$message.error(err.data.msg)

            })

          }
    });
  },
      }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  .login-container {
    -webkit-border-radius: 5px;
    border-radius: 5px;
    -moz-border-radius: 5px;
    background-clip: padding-box;
    margin-bottom: 20px;
    background-color: #F9FAFC;
    margin: 180px auto;
    border: 2px solid #8492A6;
    width: 350px;
    padding: 35px 35px 15px 35px;
    .title {
      margin: 0px auto 30px auto;
      text-align: center;
      color: #505458;
    }
  }
</style>
