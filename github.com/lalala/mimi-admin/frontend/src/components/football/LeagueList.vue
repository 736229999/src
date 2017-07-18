<template>
  <div class="news-list">

    <div class="query-tools">
      <el-form label-position="right" :model="queryForm" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="2">
            <el-button type="primary" @click="handleDialog(null)" style="margin-left:50px">添加</el-button>
          </el-col>

          <el-col :span="4">
            <el-form-item label="联赛">
              <el-select v-model="queryForm.name" clearable filterable remote placeholder="输入联赛名称搜索"
                         :remote-method="leagueRemoteMethod" :loading="selectQuery.loading" style="width:100%">
                <el-option v-for="item in selectQuery.options" :key="item.id" :label="item.name"
                           :value="item.name"></el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="2">
            <el-form-item>
              <el-button type="primary" @click="getLeagueList()">查询</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>

    <el-table border :data="leagueList" style="width: 100%" stripe>
      <el-table-column prop="id" label="联赛id" width="400px">
      </el-table-column>

      <el-table-column prop="name" label="联赛名称">
      </el-table-column>

      <el-table-column prop="address" label="操作" width="400px">
        <template scope="scope">
          <el-button size="small" type="primary" icon="edit" @click="handleDialog(scope.row)"></el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-col :span="24" class="toolbar">
      <el-pagination layout="total, prev, pager, next" @current-change="handleCurrentChange" :page-size="pager.pageSize"
                     :total="pager.tableListTotal" style="float:right;">
      </el-pagination>
    </el-col>

    <el-dialog title="联赛管理" :visible.sync="dialogFormVisible" size="tiny">
      <el-form :model="form" :rules="formRules" ref="form">
        <el-form-item label="联赛名称" prop="name" :label-width="formLabelWidth">
          <el-input v-model="form.name" auto-complete="off" style="width:50%"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible=false">取 消</el-button>
        <el-button type="primary" @click="handleAddorEdit()">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>
<script>
  import footbApi from '../../api/football';

  export default {
    mounted() {
      this.getLeagueList()
    },
    data() {
      return {
        dialogFormVisible: false,
        leagueList: [],
        form: {
          id: null,
          name: null,
        },
        formRules: {
          name: [{message: '联赛不能为空', required: true, trigger: 'blur'}],
        },
        selectQuery: {
          options: [],
          loading: false,
        },
        queryForm: {
          name: null,
          page: null,
          pagesize: null
        },
        pager: {
          tableListTotal: 1,
          pageSize: 10,
          currentPage: 1
        },
        formLabelWidth: '120px'
      }
    },
    methods: {
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
      getLeagueList() {
        footbApi.getLeagueList(this.queryForm).then((res) => {
          console.log('getLeagueList response', res);
          this.leagueList = res.data.msg.list
          this.pager.tableListTotal = res.data.msg.total
        }).catch()
      },
      handleCurrentChange(val) {
        this.pager.currentPage = val;
        this.queryForm.page = this.pager.currentPage
        this.queryForm.pagesize = this.pager.pageSize
        this.getLeagueList();
      },
      handleDialog(row) {
        this.dialogFormVisible = true
        if (row === null) {
          this.form.id = 0
          this.form.name = null
        } else {
          this.form=_.cloneDeep(row)

        }
      },
      handleAddorEdit(){
        this.$refs.form.validate((valid) => {
          if (valid) {
            if (this.form.id === 0) {
              this.postLeague()
            } else {
              this.putLeague()
            }
          } else {
            this.$message.error('参数不正确')
          }
        })
        this.dialogFormVisible = false
      },

      postLeague() {
        console.log('post League', this.form);
        footbApi.addLeague(this.form).then((res) => {
          this.$message.success('创建成功')
          console.log('post League res', res);
          this.getLeagueList()
        }).catch((error) => {
          console.log('post League error', error);
        })
      },

      putLeague() {
        console.log('put League', this.form);
        footbApi.updateLeague(this.form).then((res) => {
          console.log('put League', res);
          this.$message.success('修改成功')
          this.getLeagueList()
        }).catch((err) => {
          console.log('put League', err);
        })
      },
    }
  }

</script>

<style>
  .query-tools {
    margin-top: 15px;
  }
</style>
