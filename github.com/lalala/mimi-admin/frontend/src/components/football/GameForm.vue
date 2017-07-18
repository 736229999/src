<template>
  <div class="banner-form">
    <el-form :model="form" label-width="80px" :rules="formRules" ref="form">

      <el-row :span="24">
        <el-col :span="4">
          <el-form-item label="赛事编号" prop="gameNo">
            <el-input v-model="form.gameNo" auto-complete="off"></el-input>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24">
        <el-col :span="4">
          <el-form-item label="联赛" prop="gameType">
            <el-select v-model="form.gameType" clearable filterable remote placeholder="输入关键词搜索"
                       :remote-method="leagueRemoteMethod" :loading="selectQuery.loading" style="width:100%">
              <el-option v-for="item in selectQuery.options" :key="item.id" :label="item.name" :value="item.name">
              </el-option>
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24">
        <el-col :span="4">
          <el-form-item label="开赛时间" prop="openTime">
            <el-date-picker v-model="form.openTime" type="datetime" placeholder="选择日期时间">
            </el-date-picker>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24">
        <el-col :span="4">
          <el-form-item label="主队" prop="homeTeam">
            <el-select v-model="form.homeTeam" clearable filterable remote placeholder="输入关键词搜索"
                       :remote-method="teamRemoteMethod" :loading="selectQuery.loading" style="width:100%">
              <el-option v-for="item in selectQuery.options" :key="item.id" :label="item.name" :value="item.name">
              </el-option>
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24">
        <el-col :span="4">
          <el-form-item label="客队" prop="guestTeam">
            <el-select v-model="form.guestTeam" clearable filterable remote placeholder="输入关键词搜索"
                       :remote-method="teamRemoteMethod" :loading="selectQuery.loading" style="width:100%">
              <el-option v-for="item in selectQuery.options" :key="item.id" :label="item.name" :value="item.name">
              </el-option>
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24">
        <el-col :span="12">
          <el-form-item label="让球">
            <el-input-number v-model="form.giveball" :min="-5" :max="10"></el-input-number>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="4">
          <el-form-item label="开售时间" prop="startTime">
            <el-date-picker v-model="form.startTime" type="datetime" placeholder="选择日期时间">
            </el-date-picker>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="4">
          <el-form-item label="停售时间" prop="endTime">
            <el-date-picker v-model="form.endTime" type="datetime" placeholder="选择日期时间">
            </el-date-picker>
          </el-form-item>
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
        gameId: null,
        form: {
          id: null,
          gameNo: '',
          gameType: '',
          openTime: null,
          homeTeam: null,
          guestTeam: null,
          giveball: 0,
          startTime: null,
          endTime: null,
        },
        selectQuery: {
          options: [],
          loading: false,
        },
        formRules: {
          gameNo: [{message: '赛事编号不能为空', required: true, trigger: 'blur'},
            {min: 3, message: '最少为3个字符', trigger: 'blur'}],
          gameType: [{message: '联赛不能为空', required: true, trigger: 'blur'}],
          openTime: [{type: 'date', message: '开赛时间不能为空', required: true, trigger: 'blur'}],
          homeTeam: [{message: '主队不能为空', required: true, trigger: 'blur'}],
          guestTeam: [{message: '客队不能为空', required: true, trigger: 'blur'}],
          startTime: [{type: 'date', message: '开售时间不能为空', required: true, trigger: 'blur'}],
          endTime: [{type: 'date', message: '停售时间不能为空', required: true, trigger: 'blur'}],
        },
      };
    },
    computed: {
      buttonText() {
        if (this.gameId) {
          return '修改'
        } else {
          return '添加'
        }
      }
    },
    mounted() {
      this.gameId = this.$route.params.id;
      if (this.gameId) {     //更新时需要加载的数据
        this.getGameById();
      }
    },
    methods: {
      goBack() {
        this.$router.go(-1)
      },
      leagueRemoteMethod(query) {
        if (query !== '') {
          this.selectQuery.loading = true;
          footbApi.queryLeaguesOfSelect({name: query}).then((response) => {
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
          footbApi.queryTeamsOfSelect({name: query}).then((response) => {
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
              this.form.openTime = moment(this.form.openTime).unix()
              this.form.startTime = moment(this.form.startTime).unix()
              this.form.endTime = moment(this.form.endTime).unix()
              if (this.gameId) {
                // 更新Game
                this.putGame()
              } else {
                // 添加Game
                this.postGame()
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
      getGameById() {
        const vm = this
        if (this.gameId) {
          footbApi.getGameById({id: this.gameId}).then((res) => {
            console.log('getGameById', res);
            if (res.data.msg !== undefined) {
              this.form = res.data.msg
              this.form.openTime = moment.unix(res.data.msg.openTime).toDate()
              this.form.startTime = moment.unix(res.data.msg.startTime).toDate()
              this.form.endTime = moment.unix(res.data.msg.endTime).toDate()
            }
          }).catch((err) => {
            console.log('getGameById error', error);
          })
        }
      },
      postGame() {
        this.form.id = 0
        console.log('post Game', this.form);
        footbApi.addGame(this.form).then((res) => {
          this.$message.success('创建成功')
          console.log('post game res', res);
          this.$router.push({path: '/football/game'})
        }).catch((error) => {
          console.log('postGame error', error);
        })
      },
      putGame() {
        this.form.id = this.gameId
        this.form.id = Number(this.form.id)
        footbApi.updateGame(this.form).then((res) => {
          console.log('put Game result', res);
          this.$message.success('修改成功')
          this.$router.push({path: '/football/game'})
        }).catch((err) => {
          console.log('put Game', err);
        })
      },
    },
  }
</script>

<style>

</style>
