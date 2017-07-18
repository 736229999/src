<template>
  <div class="banner-form">
    <el-form :model="form" label-width="80px" :rules="formRules" ref="form">


      <el-row :span="24">
        <el-col :span="4">
          <el-form-item label="球队名称" prop="name">
            <el-input v-model="form.name" placeholder="请输入球队名称"></el-input>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24">
        <el-col :span="5">
          <el-tooltip class="item" effect="dark" content="点击可添加或修改" placement="bottom">
            <el-form-item label="球队队徽" prop="url">
              <el-upload class="avatar-uploader" :action="uploadUrl" :show-file-list="false"
                         :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
                <img v-if="form.url" :src="avatarUrl" class="avatar-team">
                <i v-else class="el-icon-plus avatar-uploader-plus-icon-team" style="line-height: 178px;"></i>
              </el-upload>
            </el-form-item>
          </el-tooltip>
        </el-col>
      </el-row>

      <el-form-item>
        <el-button type="primary" @click="onSubmit()">立即{{ buttonText }}</el-button>
        <el-button @click="goBack">返回</el-button>
      </el-form-item>
    </el-form>

  </div>
</template>
<script>

  import footbApi from '../../api/football';
  import moment from 'moment';
  import apiConst from '../../api/constant';
  import optionsApi from '../../api/options';


  export default {
    data() {
      return {
        teamId: null,
        form: {
          id: null,
          name: null,
          url: '',
        },
        selectQuery: {
          options: [],
          loading: false,
        },
        formRules: {
          name: {message: '队伍名称不能为空', required: true, trigger: 'blur'},
          url: {required: true, message: 'Banner图不能为空', trigger: 'blur'},
        },
      };
    },
    computed: {
      uploadUrl() {
        return apiConst.ASSETS_API + '/assets/backend/upload/teamicon'
      },
      avatarUrl() {
        return apiConst.ASSETS_API + this.form.url
      },
      buttonText() {
        if (this.teamId) {
          return '修改'
        } else {
          return '添加'
        }
      }
    },
    mounted() {
      this.teamId = this.$route.params.id;
      if (this.teamId) {     //更新时需要加载的数据
        console.log(this.teamId)
        this.getTeamById();
      }
    },
    methods: {
      goBack() {
        this.$router.go(-1)
      },
      leagueRemoteMethod(query) {
        if (query !== '') {
          this.selectQuery.loading = true;
          footbApi.queryLeaguesOfSelect({keyword: query}).then((response) => {
            console.log('queryLeagueOfSelect', response);
            this.selectQuery.loading = false;
            this.selectQuery.options = response.data.msg.list
          }).catch((error) => {
          })
        }
      },
      teamRemoteMethod(query) {
        if (query !== '') {
          this.selectQuery.loading = true;
          footbApi.queryTeamsOfSelect({keyword: query}).then((response) => {
            console.log('queryTeamsOfSelect', response);
            this.selectQuery.loading = false;
            this.selectQuery.options = response.data.msg.list
          }).catch((error) => {
          })
        }
      },
      onSubmit() {
        console.log('onSubmit');
        this.$refs.form.validate((valid) => {
          if (valid) {
            this.$confirm('确认提交吗?', '提示', {type: 'warning'}).then(() => {
              console.log('this.form', this.form);
              if (this.teamId) {
                // 更新Game
                this.putTeam()
              } else {
                // 添加Game
                this.postTeam()
              }
            }).catch((err) => {
              console.log(err)
              this.$message.info('取消提交')
            });
          } else {
            return
          }
        })
      },
      getTeamById() {
        const vm = this
        if (this.teamId) {
          footbApi.getTeamById({id: this.teamId}).then((res) => {
            console.log('getTeamById', res);
            if (res.data.msg !== undefined) {
              this.form = res.data.msg
            }
          }).catch((err) => {
            console.log('getTeamById error', error);
          })
        }
      },
      postTeam() {
        this.form.id = 0
        console.log('post Team', this.form);
        footbApi.addTeam(this.form).then((res) => {
          this.$message.success('创建成功')
          console.log('post Team res', res);
          this.$router.push({path: '/football/team'})
        }).catch((error) => {
          console.log('postTeam error', error);
        })
      },
      putTeam() {
        this.form.id = Number(this.teamId)
        console.log('put Team', this.form);
        footbApi.updateTeam(this.form).then((res) => {
          console.log('put Team', res);
          this.$message.success('修改成功')
          this.$router.push({path: '/football/team'})
        }).catch((err) => {
          console.log('put Team', err);
        })
      },
      handleAvatarSuccess(res, file) {
        this.form.url = res.result
      },
      beforeAvatarUpload() {
        console.log('beforeAvatarUpload');
      }
    },
  }
</script>

<style>
  .avatar-uploader-plus-icon-team {
    font-size: 28px;
    color: #8c939d;
    width: 200px;
    height: 200px;
    text-align: center;
  }

  .avatar-team {
    width: 200px;
    height: 200px;
    display: block;
  }
</style>
