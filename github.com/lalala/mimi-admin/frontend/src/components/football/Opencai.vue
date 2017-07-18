<template>
  <div class="news-list">

    <div class="query-tools">
      <el-form label-position="right" :model="queryForm" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="4">
            <el-form-item label="开售日期">
              <el-date-picker v-model="startTime" type="date" placeholder="选择日期">
              </el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="2">
            <el-form-item>
              <el-button type="primary" @click="submitQuery()">查询</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>

    <el-table border :data="opencaiList" style="width: 100%" stripe>
      <!--<el-table-column prop="game.id" label="比赛id">-->
      <!--</el-table-column>-->
      <el-table-column prop="game.gameNo" label="赛事编号">
      </el-table-column>
      <el-table-column prop="game.openTime" label="开赛时间" :formatter="openFormat" width="300px">
      </el-table-column>
      <el-table-column prop="game.homeTeam" label="主队">
      </el-table-column>
      <el-table-column prop="result.homeball" label="主队进球">
      </el-table-column>
      <el-table-column prop="result.homeHball" label="主队半场进球">
      </el-table-column>
      <el-table-column prop="game.guestTeam" label="客队">
      </el-table-column>
      <el-table-column prop="result.guestball" label="客队进球">
      </el-table-column>
      <el-table-column prop="result.guestHball" label="客队半场进球">
      </el-table-column>
      <el-table-column prop="result.spf" label="胜负平" :formatter="spfFormat">
      </el-table-column>
      <el-table-column prop="result.rqspf" label="让球胜平负">
      </el-table-column>
      <el-table-column prop="result.zjqs" label="总进球数" :formatter="zjqsFormat">
      </el-table-column>
      <el-table-column prop="result.bqc" label="半全场">
      </el-table-column>
      <el-table-column prop="result.ifopen" label="是否开奖">
      </el-table-column>
      <el-table-column prop="address" label="操作">
        <template scope="scope">
          <el-button size="small" type='primary' icon="edit" @click="handleEdit(scope.row.game.id)"></el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-col :span="24" class="toolbar">
      <el-pagination layout="total, prev, pager, next" @current-change="handleCurrentChange" :page-size="pager.pageSize"
                     :total="pager.tableListTotal" style="float:right;">
      </el-pagination>
    </el-col>

  </div>
</template>

<script>

  import footbApi from '../../api/football';
  import Vue from 'vue';
  import moment from 'moment';

  export default {
    mounted() {
      this.getOpencaiList()
    },
    computed: {},
    data() {
      return {
        queryForm: {
          date: null,
          page: null,
          pagesize: null
        },
        startTime: null,
        opencaiList: [],
        pager: {
          tableListTotal: 10,
          pageSize: 10,
          currentPage: 1
        },
      }
    },
    methods: {
      submitQuery() {
        if (this.startTime) {
          this.queryForm.date = moment(this.startTime).unix()
        } else {
          this.queryForm.date = null
        }
        this.getOpencaiList()
      },
      getOpencaiList() {
        console.log('查询日期', this.startTime, this.queryForm.date)
        footbApi.getOpencaiList(this.queryForm).then((res) => {
          console.log('getOpencaiList response', res);
          this.queryForm.date = (moment.unix(this.queryForm.date)).toDate()
          this.opencaiList = res.data.msg.list
          for(var i = 0;i < this.opencaiList.length; i++) {
              if (this.opencaiList[i].result){
                if(this.opencaiList[i].result.ifopen===true){
                  this.opencaiList[i].result.ifopen='是'
                }else{
                  this.opencaiList[i].result.ifopen='否'
                }
              }
          }
          this.pager.tableListTotal = res.data.msg.total
        }).catch()
      },
      handleCurrentChange(val) {
        this.pager.currentPage = val;
        this.queryForm.page = this.pager.currentPage
        this.queryForm.pagesize = this.pager.pageSize
        this.getGamesList();
      },
      openFormat(row, column) {
        return Vue.filter('timeStampFormat')(row.game.openTime)
      },
      zjqsFormat(row, column) {
        if (row.result) {
          return row.result.homeball + row.result.guestball
        }
      },
      spfFormat(row, column) {
        if (row.result) {
          if (row.result.homeball > row.result.guestball) {
            return "胜"
          } else if (row.result.homeball === row.result.guestball) {
            return "平"
          } else {
            return "负"
          }
        }
      },
      handleEdit(id = 0) {
        let topath
        topath = '/football/opencai/edit/' + id
        this.$router.push({
          path: topath
        })
      },
    }
  }
</script>

<style lang="">
  .query-tools {
    margin-top: 10px;
  }
</style>
