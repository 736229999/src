<template>
  <el-row class="container">
    <!-- Form -->
    <el-dialog title="修改密码" size="small" v-model="password.visible" @close="cancelPassword">
      <el-form :model="password.form" :rules="password.rules" ref="password">
        <el-row :gutter="20">
          <el-col :span="16">
            <el-form-item label="原始密码" prop="oldPassword">
              <el-input type="password" v-model="password.form.oldPassword" auto-complete="off"></el-input>
            </el-form-item>
            <el-form-item label="新密码" prop="newPassword">
              <el-input type="password" v-model="password.form.newPassword" auto-complete="off"></el-input>
            </el-form-item>
            <el-form-item label="再次确认" prop="confirmPassword">
              <el-input type="password" v-model="password.form.confirmPassword" auto-complete="off"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="请注意：">
              <br/>
              <div>1. 密码必须以字母开头。</div>
              <div>2. 长度在6～18位之间</div>
              <div>3. 密码只能出现字母、数字和下划线 _</div>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click.native="password.visible = false">取 消</el-button>
        <el-button type="primary" @click.native="changePassword">确 定</el-button>
      </div>
    </el-dialog>

    <el-col :span="24" class="header">
      <el-col :span="20" class="logo">
        <img src="../assets/logo.png" /> <span>亨通彩管理后台</span>
      </el-col>
      <el-col :span="4" class="userinfo">
        <el-dropdown trigger="click">
          <span class="el-dropdown-link userinfo-inner"><img :src="this.sysUserAvatar" /> {{sysUserName}}</span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item @click.native='redirectProfile'>个人中心</el-dropdown-item>
            <el-dropdown-item @click.native="password.visible = true">修改密码</el-dropdown-item>
            <el-dropdown-item divided @click.native="logout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-col>
    </el-col>

    <el-col :span="24" class="main">
      <aside>
        <el-menu :default-active="$route.path" class="el-menu-vertical-demo" @open="handleopen" @close="handleclose" @select="handleselect" theme="dark" unique-opened router>
          <template v-for="(item,index) in $router.options.routes" v-if="!item.hidden">
            <el-submenu :index="index+''" v-if="!item.leaf">
              <template slot="title"><i :class="item.iconCls"></i>{{item.name}}</template>
              <el-menu-item v-for="child in item.children" :index="child.path" :key="child.path" v-if="!child.hidden">{{child.name}}</el-menu-item>
            </el-submenu>
            <el-menu-item v-if="item.leaf&&item.children.length>0" :index="item.children[0].path"><i :class="item.iconCls"></i>{{item.children[0].name}}</el-menu-item>
          </template>
        </el-menu>
      </aside>
      <section class="content-container">
        <div class="grid-content bg-purple-light">
          <el-col :span="24" class="breadcrumb-container">
            <strong class="title">{{$route.name}}</strong>
            <el-breadcrumb separator="/" class="breadcrumb-inner">
              <el-breadcrumb-item v-for="item in $route.matched" :key="item.name">
                {{ item.name }}
              </el-breadcrumb-item>
            </el-breadcrumb>
          </el-col>
          <el-col :span="24" class="content-wrapper">
            <transition>
              <router-view></router-view>
            </transition>
          </el-col>
        </div>
      </section>
    </el-col>
  </el-row>
</template>
<script>
  import { userLogout } from '../api/api';
  import { userChangepwd } from '../api/api';
  import NProgress from 'nprogress'
  export default {
    data() {
      return {
        sysUserName: '',
        sysUserAvatar: '',

        // 密码相关
        password: {
          visible: false,
          form: {
            oldPassword: '',
            newPassword: '',
            confirmPassword: ''
          },
          formLabelWidth: '120px',
          rules: {
            oldPassword: [{
              required: true, message: '原密码不能为空', trigger: 'blur'
            }],
            newPassword: [{
              required: true, message: '新密码不能为空', trigger: 'blur'
            }],
            confirmPassword: [{
              required: true, message: '请再次确认修改的密码', trigger: 'blur'
            },{
              trigger: 'blur', message: '两次密码输入不一致', validator: (rule, value, callback) => {
                if (value !== this.password.form.newPassword) {
                  callback(new Error(rule.message))
                } else {
                  callback()
                }
              }
            }]
          }
        }

      }
    },
    methods: {
      handleopen() {
      },
      handleclose() {
      },
      handleselect: function(a, b) {},
      redirectProfile() {
        this.$router.push({ path: '/setting/profile' })
      },
      logout: function() {
        this.$confirm('确认退出吗?', '提示', {}).then(() => {
          //sessionStorage.removeItem('user');
          userLogout().then( res => { this.$router.push('/login'); } );
        }).catch(() => {
        });
      },
      cancelPassword() {
        this.$refs.password.resetFields();
      },
      changePassword() {
        this.$refs.password.validate(valid => {
          if (valid) {
            let params = new URLSearchParams();
            params.append('oldpwd', this.password.form.oldPassword);
            params.append('newpwd', this.password.form.newPassword);
            userChangepwd(params).then( res => {
              if (res.status !== 200) {
                this.$notify.error( res.data.msg );
              } else {
                // 修改成功
                this.password.visible = false;
                this.$notify.success('修改成功！');
              }
            });
          } else {
            console.log('error submit!!');
            return false
          }
        });
      }
    },
    mounted() {
      var user = sessionStorage.getItem('user');
        console.debug(user)
        if (user) {

        this.sysUserName = user.name || '';
        this.sysUserAvatar = user.avatar || 'static/default_avatar.png';
      }
    }
  }
</script>
<style scoped lang="scss">
  .container {
    position: absolute;
    top: 0px;
    bottom: 0px;
    width: 100%;
    .header {
      height: 60px;
      line-height: 60px;
      background: #1F2D3D;
      color: #c0ccda;
      .userinfo {
        text-align: right;
        padding-right: 35px;
        .userinfo-inner {
          color: #c0ccda;
          cursor: pointer;
          img {
            width: 40px;
            height: 40px;
            border-radius: 20px;
            margin: 10px 0px 10px 10px;
            float: right;
          }
        }
      }
      .logo {
        font-size: 22px;
        img {
          width: 40px;
          float: left;
          margin: 10px 10px 10px 18px;
        }
        .txt {
          color: #20a0ff
        }
      }
    }
    .main {
      background: #324057;
      position: absolute;
      top: 60px;
      bottom: 0px;
      /*overflow: hidden;*/
      aside {
        width: 230px;
      }
      .content-container {
        background: #f1f2f7;
        position: absolute;
        right: 0px;
        top: 0px;
        bottom: 0px;
        left: 230px;
        overflow-y: scroll;
        padding: 20px;
        .breadcrumb-container {
          margin-bottom: 15px;
          .title {
            width: 200px;
            float: left;
            color: #475669;
          }
          .breadcrumb-inner {
            float: right;
          }
        }
        .content-wrapper {
          background-color: #fff;
          box-sizing: border-box;
          min-width:1670px;
        }
      }
    }
  }
</style>

