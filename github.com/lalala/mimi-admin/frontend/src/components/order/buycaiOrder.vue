<template>
  <div class="withdraw-list">
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="查询内容">
          <el-col :span="23">
            <el-input v-model="searchCondition.condition" placeholder="订单id,姓名"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="search">查询</el-button>

        </el-form-item>
      </el-form>
    </el-col>

    <el-table border :data="form.list" style="width: 100%" stripe>


      <el-table-column label="订单 ID" prop="id" ></el-table-column>
      <el-table-column label="姓名" prop="user_name"></el-table-column>
      <el-table-column label="下单时间" prop="order_time"></el-table-column>
      <el-table-column label="中奖金额" prop="total_money"></el-table-column>
      <el-table-column label="彩种" prop="cai"></el-table-column>

      <el-table-column label="购彩方案"  prop="scheme_list" >
        <template scope="scope">
          <el-popover trigger="hover" placement="top" width="400" ref="scheme_list">
            <el-table :data="scheme_list">
              <el-table-column width="100" property="name" label="姓名"></el-table-column>
              <el-table-column width="300" property="address" label="地址"></el-table-column>
            </el-table>
            <div slot="reference" class="name-wrapper">
              <el-icon name="document" ></el-icon>
            </div>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column label="不抵扣总额" prop="sum_money"></el-table-column>
      <el-table-column label="花费彩金" prop="cai"></el-table-column>
      <el-table-column label="花费余额" prop="balance"></el-table-column>
      <el-table-column label="购彩花费总金额" prop="cost_cai"></el-table-column>
      <el-table-column label="余额总额" prop="cost_balance"></el-table-column>
      <el-table-column label="追期总数" prop="issue_num"></el-table-column>
      <el-table-column label="当前期数" prop="chase_no"></el-table-column>
      <el-table-column label="追期号及倍数" prop="issues">
        <template scope="scope">
          <el-popover trigger="hover" placement="top" width="400" ref="issues">
            <el-table :data="issues">
              <el-table-column width="100" property="name" label="姓名"></el-table-column>
              <el-table-column width="300" property="address" label="地址"></el-table-column>
            </el-table>
            <div slot="reference" class="name-wrapper">
              <el-icon name="document" ></el-icon>
            </div>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column prop="step" label="状态">
        <template scope="scope">
          <el-tag :type="scope.row.step == 1 ? '' : (scope.row.is_success ? 'success' : 'danger')" close-transition>{{ statusFormat(scope.row) }}</el-tag>
        </template>
      </el-table-column>





    </el-table>
    <el-col :span="24" class="toolbar">
      <el-pagination @current-change="handleCurrentChange" :current-page="form.page" :page-size="form.size"
                     layout="total, prev, pager, next" :total="form.total"></el-pagination>
    </el-col>

  </div>
</template>

<script>

  import optionsApi from '../../api/options';

  import usercenterApi from '../../api/usercenter';
  import orderApi from '../../api/order';

  import Vue from 'vue';

  export default {

    mounted() {
      this.getUserOrderList()
    },
    data() {
      return {
        title:"",
        searchCondition:{
          total: 11,
          size: 10,
          page: 1,
          condition:"",
        },
        form:{
          total: 11,
          size: 10,
          page: 1,
          list: [],
        },
        scheme_list:[],
        issues:[],
        status: null,
        queryForm: {
          isMine: null,
          step: null,
          is_success: null,
          page: null,
          pagesize: null
        },
        statusOptions: [
          { value: 1, label: '未审核', step: 1, is_success: true },
          { value: 2, label: '审核通过', step: 2, is_success: true },
          { value: 3, label: '审核未通过', step: 2, is_success: false },
          { value: 4, label: '银行处理通过', step: 3, is_success: true },
          { value: 5, label: '银行处理未通过', step: 3, is_success: false }
        ],
        withdrawApplyList: [],
        pager: {
          tableListTotal: 20,
          pageSize: 10,
          currentPage: 1
        }
      }
    },
      computed: {
        queryArg() {
          _.find(this.statusOptions)
        }
      },
      methods: {
        handleCurrentChange(val) {
          this.form.page = val;
          this.GetTaskList();
        },
        buttonTextFormat(row) {
          if (row.step > 1) {
            return '查看'
          } else if (row.auditor != '') {
            return '审核'
          }
          return '认领'
        },
        handleStatusChange(val) {
          console.log('handleStatusChange', val);
          if (val === '') {
            this.queryForm.step = this.queryForm.is_success = null
          } else {
            let status = _.find(this.statusOptions, { value: val })
            this.queryForm.step = status.step
            this.queryForm.is_success = status.is_success
          }
        },
        submitQuery() {
          this.getWithdrawApplyList()
        },
        search: function () {
          orderApi.searchUserOrder(this.searchCondition).then((res) => {
            console.log(' response', res.data.msg);
            this.form = res.data.msg;

          }).catch((err) => {
            this.$message.error(err.data.msg)
          })
        },
        getUserOrderList() {
          orderApi.getUserOrderList(this.form).then((res) => {
//            console.log(' response', res.data.msg.list[0].issues);
            this.form = res.data.msg;
            //这个地方暂保留，以后看了数据库存的字段格式再处理
//            for (var i=0;i<res.data.msg.length;i++){
//
//            }

//            this.issues = [{name:"啦啦啦",age:213},{name:"123"}];
            this.issues = [res.data.msg.list[0].issues];
          }).catch((err) => {
            this.$message.error(err.data.msg)
          })
        },
        handleCurrentChange(val) {
          this.pager.currentPage = val;
          this.queryForm.page = this.pager.currentPage
          this.queryForm.pagesize = this.pager.pageSize
          this.getWithdrawApplyList();
        },
        timestampFormat(row, column) {
          let format = Vue.filter('timeStampFormat')
          let msg
          switch (column.property) {
            case 'create_time':
              msg = format(row.create_time)
              break;
            case 'audit_time':
              if (row.step > 1) {
                msg = format(row.audit_time)
              } else {
                msg = ''
              }
              break;
            default:
              break;
          }
          return msg
        },
        statusFormat(val) {
          let step = '', stepDesc = '未通过'
          if (val.is_success) {
            stepDesc = '通过'
          }
          switch (val.step) {
            case 1:
              step = '未审核'
              stepDesc = ''
              break;
            case 2:
              step = '审核'
              break;
            case 3:
              step = '银行处理'
            default:
              break;
          }
          return step + stepDesc
        },
        auditFormat(row, column) {
          let msg
          switch (column.property) {
            case '':

              break;

            default:
              break;
          }
        },
        confirmOpen(index, row) {
          console.log('confirmOpen');
          if (row.can_operate) {
            console.log('can Operate');
            if (row.step > 1 || row.auditor != '') {
              console.log('自己的未审核');
              this.handleAddOrEdit(row.id)
            } else {
              console.log('初次认领');
              this.$confirm('确认认领吗？此操作将不可撤销, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(() => {
                this.handleAddOrEdit(row.id)
              }).catch(() => {
                this.$message({
                  type: 'info',
                  message: '已取消认领'
                });
              });
            }
          } else {
            this.$message.error('无权操作')
          }
        },
        handleisMineChange(val) {
          console.log('handleisMineChange', val);
          this.getWithdrawApplyList()
        },
        handleAddOrEdit(id = 0) {
          let topath
          if (id == 0) {
            topath = '/usercenter/withdraw/add'
          } else {
            topath = '/usercenter/withdraw/audit/' + id
          }
          this.$router.push({
            path: topath
          })
        },
      }
  }
</script>

<style lang="">
  .query-tools {
    margin-top: 50px;
  }
</style>
