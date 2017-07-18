<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="操作人">
          <el-col :span="23">
            <el-input v-model="form.account" placeholder="账号"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" >查询</el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!--列表-->
    <el-col>
    <el-table :data="data" border>
      <el-table-column label="序号" prop="id"></el-table-column>
      <el-table-column label="操作人" prop="username"></el-table-column>
      <el-table-column label="访问路径" prop="path"></el-table-column>
      <el-table-column label="操作类型" width="300">
          <template scope="scope">
            <div v-if="scope.row.operating == 2">
              <span style="color: red">删除</span>
            </div>
            <div v-else-if="scope.row.operating == 3">
              <span style="color: yellow">修改</span>
            </div>
            <div v-else-if="scope.row.operating == 1">
              <span style="color: blue;">添加</span>
            </div>
            <div v-else>
              <span style="color: green">查看</span>
            </div>
          </template>
      </el-table-column>
      <el-table-column label="备注" prop="message"></el-table-column>
      <el-table-column label="操作时间">
        <template scope="scope">
          <span>{{new Date(scope.row.create_time * 1000).toLocaleString()}}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="100">
        <template scope="scope">
          <el-button @click="handleClick" type="text" size="small">查看</el-button>
        </template>
      </el-table-column>
    </el-table>
      <div class="block">
        <el-pagination  @current-change="handleCurrentChange" :current-page="form.page" :page-size="form.size" layout="total, prev, pager, next" :total="form.total"></el-pagination>
      </div>
    </el-col></div>
</template>

<script>
    import logApi from '../../api/log';

    export default {
        data() {
            return {
                data : [],
                total:0,
                form : {
                    account : '',
                    size:100,
                    page:1,
                    total:0,
                },

            }
        },methods : {
             handleClick:function() {
                console.log(1);
            },
            GetLog() {
                logApi.logList(this.form).then((res) => {
                      this.data = res.data.msg.log;
                      this.form.total = res.data.msg.total;
                }).catch(err => {

                });
            },
            handleCurrentChange(val) {
                this.form.page = val;
                this.GetLog();
            },

        },
        mounted() {
            this.GetLog();
        }
    }
</script>

