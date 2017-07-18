<template>
  <div class="news-list">

    <div class="query-tools">
      <el-form label-position="right" :model="queryForm" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="2">
            <el-button type="primary" @click="handleAddOrEdit(0)" style="margin-left:50px">添加</el-button>
          </el-col>

          <el-col :span="4">
            <el-form-item label="球队">
              <el-select v-model="queryForm.teamname" clearable filterable remote placeholder="输入球队名称搜索"
                         :remote-method="teamRemoteMethod" :loading="selectQuery.loading" style="width:100%">
                <el-option v-for="item in selectQuery.options" :key="item.id" :label="item.name"
                           :value="item.id"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="5">
            <el-form-item label="开始时间">
              <el-date-picker v-model="dateRange" type="datetimerange" placeholder="选择日期范围"
                              :picker-options="pickerOptions">
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

    <el-table border :data="gamesList" style="width: 100%" stripe>
      <el-table-column prop="id" label="比赛id" width="150">
      </el-table-column>
      <el-table-column prop="gameNo" label="赛事编号" width="150">
      </el-table-column>
      <el-table-column prop="gameType" label="赛事类型">
      </el-table-column>
      <el-table-column prop="openTime" label="开赛时间" :formatter="openFormat">
      </el-table-column>
      <el-table-column prop="homeTeam" label="主队">
      </el-table-column>
      <el-table-column prop="guestTeam" label="客队">
      </el-table-column>
      <el-table-column prop="giveball" label="让球数">
      </el-table-column>
      <el-table-column prop="startTime" label="开售时间" :formatter="startFormat">
      </el-table-column>
      <el-table-column prop="endTime" label="停售时间" :formatter="endFormat">
      </el-table-column>
      <el-table-column prop="address" label="操作" width="300">
        <template scope="scope">
          <el-button size="small" type="primary" icon="edit" @click="handleAddOrEdit(scope.row.id)"></el-button>
          <el-button size="small" type="primary" icon="search" @click="handleviewOdds(scope.row.id)">赔率</el-button>
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
      this.getGamesList()
    },
    data() {
      return {
        queryForm: {
          teamname: null,
          start: null,
          end: null,
          page: null,
          pagesize: null
        },
        dateRange: [],
        gamesList: [],
        selectQuery: {
          options: [],
          loading: false,
        },
        pager: {
          tableListTotal: 10,
          pageSize: 10,
          currentPage: 1
        },
        pickerOptions: {
          shortcuts: [{
            text: '未来3天',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              end.setTime(end.getTime() + 3600 * 1000 * 24 * 3);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近一周',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
              picker.$emit('pick', [start, end]);
            }
          }, {
            text: '最近一个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
              picker.$emit('pick', [start, end]);
            }
          }]
        },
      }
    },
    methods: {
      submitQuery() {
        if (this.dateRange.length === 2) {
          if (this.dateRange[0] !== null || this.dateRange[1] !== null) {
            this.queryForm.start = moment(this.dateRange[0]).unix()
            this.queryForm.end = moment(this.dateRange[1]).unix()
          } else {
            this.queryForm.start = this.queryForm.end = null
          }
        }
        this.getGamesList()
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
      getGamesList() {
        footbApi.getGamesList(this.queryForm).then((res) => {
          console.log('getGameList response', res);
          this.gamesList = res.data.msg.list
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
        return Vue.filter('timeStampFormat')(row.openTime)
      },
      startFormat(row, column) {
        return Vue.filter('timeStampFormat')(row.startTime)
      },
      endFormat(row, column) {
        return Vue.filter('timeStampFormat')(row.endTime)
      },
      handleAddOrEdit(id = 0) {
        let topath
        if (id === 0) {
          topath = '/football/game/add'
        } else {
          topath = '/football/game/edit/' + id
        }
        this.$router.push({
          path: topath
        })
      },
      handleviewOdds(id){
        let topath
        topath = '/football/odds/' + id
        this.$router.push({
          path: topath
        })
      }
    }
  }
</script>

<style lang="">
  .query-tools {
    margin-top: 10px;
  }
</style>
