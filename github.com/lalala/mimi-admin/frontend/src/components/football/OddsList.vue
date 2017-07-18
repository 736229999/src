<template>
  <div>
    <el-row>
      <el-col :span="2">
        <el-button type="primary" style="margin-top:15px;margin-left:50px" @click="addodds()" :disabled=this.addflag>添加
        </el-button>
      </el-col>
      <el-col :span="2">
        <el-button type="primary" style="margin-top:15px" @click="editodds()" :disabled=this.editflag>修改</el-button>
      </el-col>
    </el-row>
    <el-row>
      <el-col style="margin-top: 20px;margin-left: 10px;padding-right: 20px;">
        <el-tabs v-model="data" type="card" value="spf">
          <el-tab-pane label="胜平负" name="spf" style="padding-bottom:15px">
            <el-table border :data="dataspf" style="width: 100%">
              <el-table-column prop="win" label="胜"></el-table-column>
              <el-table-column prop="draw" label="平"></el-table-column>
              <el-table-column prop="lose" label="负"></el-table-column>
            </el-table>
            <el-row style="padding-top: 15px">
              <el-tag :type="spfdg === true ? 'primary' : 'danger'" close-transition>{{spfdg | switchPlayTypeFormat}}
              </el-tag>
            </el-row>
          </el-tab-pane>

          <el-tab-pane label="让球胜平负" name="rqspf" style="padding-bottom:15px">
            <el-table border :data="datarqspf" style="width: 100%">
              <el-table-column prop="win" label="胜"></el-table-column>
              <el-table-column prop="draw" label="平"></el-table-column>
              <el-table-column prop="lose" label="负"></el-table-column>
            </el-table>
            <el-row style="padding-top: 15px">
              <el-tag :type="rqspfdg === true ? 'primary' : 'danger'" close-transition>{{rqspfdg | switchPlayTypeFormat}}</el-tag>
            </el-row>
          </el-tab-pane>

          <el-tab-pane label="总进球数" name="zjqs" style="padding-bottom:15px">
            <el-table border :data="datazjqs" style="width: 100%">
              <el-table-column prop="zero" label="0球"></el-table-column>
              <el-table-column prop="one" label="1球"></el-table-column>
              <el-table-column prop="two" label="2球"></el-table-column>
              <el-table-column prop="three" label="3球"></el-table-column>
              <el-table-column prop="four" label="4球"></el-table-column>
              <el-table-column prop="five" label="5球"></el-table-column>
              <el-table-column prop="six" label="6球"></el-table-column>
              <el-table-column prop="seven" label="7球及以上"></el-table-column>
            </el-table>
            <el-row style="padding-top: 15px">
              <el-tag :type="zjqsdg === true ? 'primary' : 'danger'" close-transition>
                {{zjqsdg | switchPlayTypeFormat}}
              </el-tag>
            </el-row>
          </el-tab-pane>

          <el-tab-pane label="比分" name="bf" style="padding-bottom:15px">
            <el-table border :data="databf" style="width: 100%">
              <el-table-column prop="wo" label="胜其他"></el-table-column>
              <el-table-column prop="w10" label="1:0"></el-table-column>
              <el-table-column prop="w20" label="2:0"></el-table-column>
              <el-table-column prop="w21" label="2:1"></el-table-column>
              <el-table-column prop="w30" label="3:0"></el-table-column>
              <el-table-column prop="w31" label="3:1"></el-table-column>
              <el-table-column prop="w32" label="3:2"></el-table-column>
              <el-table-column prop="w40" label="4:0"></el-table-column>
              <el-table-column prop="w41" label="4:1"></el-table-column>
              <el-table-column prop="w42" label="4:2"></el-table-column>
              <el-table-column prop="w50" label="5:0"></el-table-column>
              <el-table-column prop="w51" label="5:1"></el-table-column>
              <el-table-column prop="w52" label="5:2"></el-table-column>
            </el-table>
            <el-table border :data="databf" style="width: 100%">
              <el-table-column prop="lo" label="胜其他"></el-table-column>
              <el-table-column prop="l01" label="0:1"></el-table-column>
              <el-table-column prop="l02" label="0:2"></el-table-column>
              <el-table-column prop="l12" label="1:2"></el-table-column>
              <el-table-column prop="l03" label="0:3"></el-table-column>
              <el-table-column prop="l13" label="1:3"></el-table-column>
              <el-table-column prop="l23" label="2:3"></el-table-column>
              <el-table-column prop="l04" label="0:4"></el-table-column>
              <el-table-column prop="l14" label="1:4"></el-table-column>
              <el-table-column prop="l24" label="2:4"></el-table-column>
              <el-table-column prop="l05" label="0:5"></el-table-column>
              <el-table-column prop="l15" label="1:5"></el-table-column>
              <el-table-column prop="l25" label="2:5"></el-table-column>
            </el-table>
            <el-table border :data="databf" style="width: 100%">
              <el-table-column prop="do" label="平其他"></el-table-column>
              <el-table-column prop="d00" label="0:0"></el-table-column>
              <el-table-column prop="d11" label="1:1"></el-table-column>
              <el-table-column prop="d22" label="2:2"></el-table-column>
              <el-table-column prop="d33" label="3:3"></el-table-column>
            </el-table>
            <el-row style="padding-top: 15px">
              <el-tag :type="bfdg === true ? 'primary' : 'danger'" close-transition>{{bfdg | switchPlayTypeFormat}}
              </el-tag>
            </el-row>
          </el-tab-pane>

          <el-tab-pane label="半全场" name="bqc" style="padding-bottom:15px">
            <el-table border :data="databqc" style="width: 100%">
              <el-table-column prop="ww" label="胜胜"></el-table-column>
              <el-table-column prop="wd" label="胜平"></el-table-column>
              <el-table-column prop="wl" label="胜负"></el-table-column>
              <el-table-column prop="dw" label="平胜"></el-table-column>
              <el-table-column prop="dd" label="平平"></el-table-column>
              <el-table-column prop="dl" label="平负"></el-table-column>
              <el-table-column prop="lw" label="负胜"></el-table-column>
              <el-table-column prop="ld" label="负平"></el-table-column>
              <el-table-column prop="ll" label="负负"></el-table-column>
            </el-table>
            <el-row style="padding-top: 15px">
              <el-tag :type="bqcdg === true ? 'primary' : 'danger'" close-transition>{{bqcdg | switchPlayTypeFormat}}
              </el-tag>
            </el-row>
          </el-tab-pane>

        </el-tabs>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-button type="primary" style="float:right;margin-bottom:15px;margin-right:50px" @click="goBack()">返回
        </el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import ucApi from "../../api/usercenter"
  import footbApi from '../../api/football';
  import Vue from 'vue'
  export default {
    data() {
      return {
        oddsId: null,
        data: "recharge",
        addflag: null,
        editflag: null,
        dataspf: [],
        datarqspf: [],
        datazjqs: [],
        databqc: [],
        databf: [],
        spfdg: null,
        rqspfdg: null,
        zjqsdg: null,
        bfdg: null,
        bqcdg: null,
      };
    },
    methods: {
      getOddsbyId() {
        footbApi.getOddbyId({id: this.oddsId}).then((res) => {
          if (res.data.msg.id !== 0) {
            this.dataspf = [JSON.parse(res.data.msg.spf)]
            this.datarqspf = [JSON.parse(res.data.msg.rqspf)]
            this.datazjqs = [JSON.parse(res.data.msg.zjqs)]
            this.databf = [JSON.parse(res.data.msg.bf)]
            this.databqc = [JSON.parse(res.data.msg.bqc)]
            this.spfdg = res.data.msg.spfdg
            this.rqspfdg = res.data.msg.rqspfdg
            this.zjqsdg = res.data.msg.zjqsdg
            this.bfdg = res.data.msg.bfdg
            this.bqcdg = res.data.msg.bqcdg
            this.addflag = true
            this.editflag = false
          } else {
            this.addflag = false
            this.editflag = true
          }
        }).catch()
      },
      editodds(){
        this.$router.push({
          path: '/football/odds/edit/' + this.oddsId
        })
      },
      addodds(){
        this.$router.push({
          path: '/football/odds/add/' + this.oddsId
        })
      },
      goBack(){
        this.$router.push({
          path: '/football/game'
        })
      },
    },
    mounted() {
      this.oddsId = this.$route.params.id;
      console.log("oddsid", this.oddsId)
      this.getOddsbyId();
    }
  };
</script>
