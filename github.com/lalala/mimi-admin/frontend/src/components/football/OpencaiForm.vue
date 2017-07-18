<template>
  <div class="banner-form">
    <el-form :model="form" label-width="100px" :rules="formRules" ref="form">

      <el-row :span="24">
        <el-col :span="4">
          <el-form-item label="主队进球" prop="homeball" >
            <el-input-number v-model="form.homeball" :min="0" :max="20"></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="4">
          <el-form-item label="客队进球" prop="guestball">
            <el-input-number v-model="form.guestball" :min="0" :max="20"></el-input-number>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24">
        <el-col :span="4">
          <el-form-item label="主队半场进球" prop="homeball" >
            <el-input-number v-model="form.homeHball" :min="0" :max="20"></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="4">
          <el-form-item label="客队半场进球" prop="guestball">
            <el-input-number v-model="form.guestHball" :min="0" :max="20"></el-input-number>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24">
        <el-col :span="12">
          <el-form-item label="让球" prop="rqspf">
            <el-select v-model="form.rqspf" placeholder="请选择">
              <el-option v-for="item in optionspf" :key="item.value" :label="item.label" :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24">
        <el-col :span="12">
          <el-form-item label="半全场" prop="bqc">
            <el-select v-model="form.bqc" placeholder="请选择">
              <el-option v-for="item in optionbqc" :key="item.value" :label="item.label" :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="4">
          <el-form-item label="是否开奖" prop="ifopen">
            <el-switch v-model="form.ifopen" on-color="#13ce66" off-color="#ff4949">
            </el-switch>
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
        resultId: null,
        form: {
          id: null,
          homeball: 0,
          guestball: 0,
          homeHball:0,
          guestHball:0,
          rqspf: null,
          bqc: null,
          ifopen: false,
        },

        formRules: {
          rqspf: [{message: '赛事编号不能为空', required: true, trigger: 'blur'}],
          bqc: [{message: '联赛不能为空', required: true, trigger: 'blur'}],
        },
        optionspf: [{
          value: '胜',
          label: '胜'
        }, {
          value: '平',
          label: '平'
        }, {
          value: '负',
          label: '负'
        }],
        optionbqc: [{
          value: '胜-胜',
          label: '胜-胜'
        }, {
          value: '胜-平',
          label: '胜-平'
        }, {
          value: '胜-负',
          label: '胜-负'
        }, {
          value: '平-胜',
          label: '平-胜'
        }, {
          value: '平-平',
          label: '平-平'
        }, {
          value: '平-负',
          label: '平-负'
        }, {
          value: '负-胜',
          label: '负-胜'
        }, {
          value: '负-平',
          label: '负-平'
        }, {
          value: '负-负',
          label: '负-负'
        }
        ],
      };
    },
    computed: {
      buttonText() {
        if (this.resultId) {
          return '修改'
        } else {
          return '添加'
        }
      }
    },
    mounted() {
      this.gameId = this.$route.params.id;
      if (this.gameId) {
        this.getOpencaiById();
      }
    },
    methods: {
      goBack() {
        this.$router.go(-1)
      },
      onSubmit() {
        console.log('onSubmit');
        this.$refs.form.validate((valid) => {
          if (valid) {
            this.$confirm('确认提交吗?', '提示', {type: 'warning'}).then(() => {
              console.log('this.form', this.form);
              if (this.resultId) {
                // 更新Game
                this.putOpencai()
              } else {
                // 添加Game
                this.postOpencai()
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
      getOpencaiById() {
        const vm = this
        if (this.gameId) {
          footbApi.getOpencaiById({id: this.gameId}).then((res) => {
            console.log('getopencaiById', res);
            if (res.data.msg !== undefined) {
              this.form = res.data.msg
              this.resultId = res.data.msg.id
            }
          }).catch((err) => {
            console.log('getopencaiById error', error);
          })
        }
      },
      postOpencai() {
        this.form.id = Number(this.gameId)
        console.log('post Opencai', this.form);
        footbApi.addOpencai(this.form).then((res) => {
          this.$message.success('创建成功')
          console.log('post Opencai res', res);
          this.$router.push({path: '/football/opencai'})
        }).catch((error) => {
          console.log('post Opencai error', error);
        })
      },
      putOpencai() {
        this.form.id = Number(this.gameId)
        console.log('put Opencai', this.form);
        footbApi.updateOpencai(this.form).then((res) => {
          console.log('put Opencai', res);
          this.$message.success('修改成功')
          this.$router.push({path: '/football/opencai'})
        }).catch((err) => {
          console.log('put Opencai error', err);
        })
      },
    },
  }
</script>

<style>

</style>
