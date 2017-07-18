<template>
  <div class="banner-form">
    <el-form :model="form" label-width="80px" ref="form">
      <el-row :span="24">
        <el-tag type="gray">胜平负</el-tag>
      </el-row>
      <el-row :span="24">
        <el-col :span="2">
          <el-form-item label="胜" prop="dataspf.win" :rules="formRules">
            <el-input v-model.number="form.dataspf.win"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="平" prop="dataspf.draw" :rules="formRules">
            <el-input v-model.number="form.dataspf.draw"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="负" prop="dataspf.lose" :rules="formRules">
            <el-input v-model.number="form.dataspf.lose"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="是否单关" prop="ifopen">
            <el-switch v-model="form.spfdg" on-color="#13ce66" off-color="#ff4949"></el-switch>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24" style="margin-top: 10px">
        <el-tag type="gray">让球胜平负</el-tag>
      </el-row>
      <el-row :span="24">
        <el-col :span="2">
          <el-form-item label="胜" prop="datarqspf.win" :rules="formRules">
            <el-input v-model.number="form.datarqspf.win"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="平" prop="datarqspf.draw" :rules="formRules">
            <el-input v-model.number="form.datarqspf.draw"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="负" prop="datarqspf.lose" :rules="formRules">
            <el-input v-model.number="form.datarqspf.lose"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="是否单关" prop="ifopen">
            <el-switch v-model="form.rqspfdg" on-color="#13ce66" off-color="#ff4949"></el-switch>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24" style="margin-top: 10px">
        <el-tag type="gray">总进球数</el-tag>
      </el-row>
      <el-row :span="24">
        <el-col :span="2">
          <el-form-item label="0颗" prop="datazjqs.zero" :rules="formRules">
            <el-input v-model.number="form.datazjqs.zero"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="1颗" prop="datazjqs.one" :rules="formRules">
            <el-input v-model.number="form.datazjqs.one"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="2颗" prop="datazjqs.two" :rules="formRules">
            <el-input v-model.number="form.datazjqs.two"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="3颗" prop="datazjqs.three" :rules="formRules">
            <el-input v-model.number="form.datazjqs.three"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="4颗" prop="datazjqs.four" :rules="formRules">
            <el-input v-model.number="form.datazjqs.four"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="5颗" prop="datazjqs.five" :rules="formRules">
            <el-input v-model.number="form.datazjqs.five"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="6颗" prop="datazjqs.six" :rules="formRules">
            <el-input v-model.number="form.datazjqs.six"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="7颗" prop="datazjqs.seven" :rules="formRules">
            <el-input v-model.number="form.datazjqs.seven"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="是否单关" prop="ifopen">
            <el-switch v-model="form.zjqsdg" on-color="#13ce66" off-color="#ff4949"></el-switch>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24" style="margin-top: 10px">
        <el-tag type="gray">半全场</el-tag>
      </el-row>
      <el-row :span="24">
        <el-col :span="2">
          <el-form-item label="胜胜" prop="databqc.ww" :rules="formRules">
            <el-input v-model.number="form.databqc.ww"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="胜平" prop="databqc.wd" :rules="formRules">
            <el-input v-model.number="form.databqc.wd"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="胜负" prop="databqc.wl" :rules="formRules">
            <el-input v-model.number="form.databqc.wl"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="平胜" prop="databqc.dw" :rules="formRules">
            <el-input v-model.number="form.databqc.dw"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="平平" prop="databqc.dd" :rules="formRules">
            <el-input v-model.number="form.databqc.dd"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="平负" prop="databqc.dl" :rules="formRules">
            <el-input v-model.number="form.databqc.dl"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="负胜" prop="databqc.lw" :rules="formRules">
            <el-input v-model.number="form.databqc.lw"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="负平" prop="databqc.ld" :rules="formRules">
            <el-input v-model.number="form.databqc.ld"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="负负" prop="databqc.ll" :rules="formRules">
            <el-input v-model.number="form.databqc.ll"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="是否单关" prop="ifopen">
            <el-switch v-model="form.bqcdg" on-color="#13ce66" off-color="#ff4949"></el-switch>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24" style="margin-top: 10px">
        <el-tag type="gray">比分</el-tag>
        <el-tag type="success">胜</el-tag>
      </el-row>
      <el-row :span="24">
        <el-col :span="2">
          <el-form-item label="1:0" prop="databf.w10" :rules="formRules">
            <el-input v-model.number="form.databf.w10"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="2:0" prop="databf.w20" :rules="formRules">
            <el-input v-model.number="form.databf.w20"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="3:0" prop="databf.w30" :rules="formRules">
            <el-input v-model.number="form.databf.w30"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="3:1" prop="databf.w31" :rules="formRules">
            <el-input v-model.number="form.databf.w31"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="3:2" prop="databf.w32" :rules="formRules">
            <el-input v-model.number="form.databf.w32"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="4:0" prop="databf.w40" :rules="formRules">
            <el-input v-model.number="form.databf.w40"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="4:1" prop="databf.w41" :rules="formRules">
            <el-input v-model.number="form.databf.w41"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="4:2" prop="databf.w42" :rules="formRules">
            <el-input v-model.number="form.databf.w42"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="5:0" prop="databf.w50" :rules="formRules">
            <el-input v-model.number="form.databf.w50"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="5:1" prop="databf.w51" :rules="formRules">
            <el-input v-model.number="form.databf.w51"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="5:2" prop="databf.w52" :rules="formRules">
            <el-input v-model.number="form.databf.w52"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="胜其他" prop="databf.wo" :rules="formRules">
            <el-input v-model.number="form.databf.wo"></el-input>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24" style="margin-left: 40px">
        <el-tag type="warning">平</el-tag>
      </el-row>
      <el-row :span="24">
        <el-col :span="2">
          <el-form-item label="0:0" prop="databf.d00" :rules="formRules">
            <el-input v-model.number="form.databf.d00"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="1:1" prop="databf.d11" :rules="formRules">
            <el-input v-model.number="form.databf.d11"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="2:2" prop="databf.d22" :rules="formRules">
            <el-input v-model.number="form.databf.d22"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="3:3" prop="databf.d33" :rules="formRules">
            <el-input v-model.number="form.databf.d33"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="平其他" prop="databf.do" :rules="formRules">
            <el-input v-model.number="form.databf.do"></el-input>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row :span="24" style="margin-left: 40px">
        <el-tag type="danger">负</el-tag>
      </el-row>
      <el-row :span="24">
        <el-col :span="2">
          <el-form-item label="0:1" prop="databf.l01" :rules="formRules">
            <el-input v-model.number="form.databf.l01"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="0:2" prop="databf.l02" :rules="formRules">
            <el-input v-model.number="form.databf.l02"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="0:3" prop="databf.l03" :rules="formRules">
            <el-input v-model.number="form.databf.l03"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="1:3" prop="databf.l13" :rules="formRules">
            <el-input v-model.number="form.databf.l13"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="2:3" prop="databf.l23" :rules="formRules">
            <el-input v-model.number="form.databf.l23"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="0:4" prop="databf.l04" :rules="formRules">
            <el-input v-model.number="form.databf.l04"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="1:4" prop="databf.l14" :rules="formRules">
            <el-input v-model.number="form.databf.l14"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="2:4" prop="databf.l24" :rules="formRules">
            <el-input v-model.number="form.databf.l24"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="0:5" prop="databf.l05" :rules="formRules">
            <el-input v-model.number="form.databf.l05"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="1:5" prop="databf.l15" :rules="formRules">
            <el-input v-model.number="form.databf.l15"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="2:5" prop="databf.l25" :rules="formRules">
            <el-input v-model.number="form.databf.l25"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="负其他" prop="databf.lo" :rules="formRules">
            <el-input v-model.number="form.databf.lo"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2">
          <el-form-item label="是否单关" prop="ifopen">
            <el-switch v-model="form.bfdg" on-color="#13ce66" off-color="#ff4949"></el-switch>
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

  export default {
    data() {
      return {
        option: '',
        oddsId: null,
        form: {
          id: null,
          dataspf: {},
          datarqspf: {},
          datazjqs: {},
          databqc: {},
          databf: {},
          spfdg: null,
          rqspfdg: null,
          zjqsdg: null,
          bfdg: null,
          bqcdg: null,
        },
        submitform: {
          id: null,
          dataspf: null,
          datarqspf: null,
          datazjqs: null,
          databqc: null,
          databf: null,
          spfdg: null,
          rqspfdg: null,
          zjqsdg: null,
          bfdg: null,
          bqcdg: null,
        },
        formRules: [{required: true, message: '赔率不能为空'},
          {type: 'number', message: '赔率必须为数字值'}],
      };
    },
    computed: {
      buttonText() {
        if (this.option === 'edit') {
          return '修改'
        } else {
          return '添加'
        }
      }
    },
    mounted() {
      this.oddsId = this.$route.params.id;
      this.option = this.$route.params.op;
      if (this.option === 'edit') {     //更新时需要加载的数据
        this.getOddsById();
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
              this.submitform.id = Number(this.oddsId)
              this.submitform.spf = JSON.stringify(this.form.dataspf)
              this.submitform.rqspf = JSON.stringify(this.form.datarqspf)
              this.submitform.zjqs = JSON.stringify(this.form.datazjqs)
              this.submitform.bf = JSON.stringify(this.form.databf)
              this.submitform.bqc = JSON.stringify(this.form.databqc)
              this.submitform.spfdg = this.form.spfdg
              this.submitform.rqspfdg = this.form.rqspfdg
              this.submitform.zjqsdg = this.form.zjqsdg
              this.submitform.bfdg = this.form.bfdg
              this.submitform.bqcdg = this.form.bqcdg
              if (this.option === 'edit') {
                // 更新Game
                this.putOdds()
              } else {
                // 添加Game
                this.postOdds()
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
      getOddsById() {
        footbApi.getOddbyId({id: this.oddsId}).then((res) => {
          this.form.dataspf = JSON.parse(res.data.msg.spf)
          this.form.datarqspf = JSON.parse(res.data.msg.rqspf)
          this.form.datazjqs = JSON.parse(res.data.msg.zjqs)
          this.form.databf = JSON.parse(res.data.msg.bf)
          this.form.databqc = JSON.parse(res.data.msg.bqc)
          this.form.spfdg = res.data.msg.spfdg
          this.form.rqspfdg = res.data.msg.rqspfdg
          this.form.zjqsdg = res.data.msg.zjqsdg
          this.form.bfdg = res.data.msg.bfdg
          this.form.bqcdg = res.data.msg.bqcdg
        }).catch((error) => {
          console.log('getOddsById error', error);
        })
      },
      postOdds() {
        footbApi.addOdds(this.submitform).then((res) => {
          this.$message.success('创建成功')
          console.log('postOdds res', res);
          this.$router.push({path: '/football/game'})
        }).catch((error) => {
          console.log('postOdds error', error);
        })
      },
      putOdds() {
        footbApi.updateOdds(this.submitform).then((res) => {
          console.log('putOdds', res);
          this.$message.success('修改成功')
          this.$router.push({path: '/football/odds/' + this.oddsId})
        }).catch((err) => {
          console.log('put Odds error', err);
        })
      },
    },
  }
</script>

<style>
  .odds-title {
    background: #20A0FF;
  }
</style>
