<template>
  <div>
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="任务标题">
          <el-col :span="23">
            <el-input v-model="form.title" placeholder="标题"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="search">查询</el-button>
          <router-link :to="{path:'/activity/task/add'}">
            <el-button type="primary" icon="plus">添加</el-button>
          </router-link>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-table :data="form.list" style="width: 100%">
        <el-table-column prop="id" label="序号" width="180"></el-table-column>
        <el-table-column prop="name" label="任务名" width="180"></el-table-column>
        <el-table-column prop="des" label="任务描述"></el-table-column>
        <el-table-column prop="type" label="所属类型"></el-table-column>
        <el-table-column prop="addtime"label="创建时间">
          <template scope="scope">
            <span>{{ scope.row.addtime | timeStampFormat }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="is_finish" label="后台是否实现"></el-table-column>

        <el-table-column label="操作">
          <template scope="scope">
            <router-link :to="{path:'/activity/cdkey/detail/'+scope.row.id}">
              <el-button type="primary" size="mini" icon="more"></el-button>
            </router-link>
            <router-link :to="{path:'/activity/task/edit/'+scope.row.id}">
              <el-button type="warning" size="mini" icon="edit"></el-button>
            </router-link>
            <el-button type="danger" size="mini" icon="delete" @click="del(scope.row.id)"></el-button>
            <el-button type="success" size="mini" icon="upload" @click="exportCsv(scope.row.id)"></el-button>
            <a href="" id="downloadLink" download="asd.csv" style="display: none"></a>
          </template>
        </el-table-column>
      </el-table>
      <div class="block">
        <el-pagination @current-change="handleCurrentChange" :current-page="form.page" :page-size="form.size"
                       layout="total, prev, pager, next" :total="form.total"></el-pagination>
      </div>
    </el-col>
  </div>
</template>

<script>

  import acApi from "../../api/activity"

  export default {
    data() {
      return {
        form: {
          total: 0,
          size: 10,
          page: 1,
          list: [],
        },
        disabled: false,
      }
    },
    methods: {
      //获取cdkey列表.
      GetTaskList: function () {
        acApi.taskList(this.form).then((res) => {
          this.form = res.data.msg;
          console.log(res.data.msg);
        }).catch(err => {
          this.$message.error("获取失败")

        })
      },
      handleCurrentChange(val) {
        this.form.page = val;
        this.GetTaskList();
      },
      search: function () {
        this.GetTaskList();
      },
      del: function (id) {
        this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {

          acApi.deleteTask({id: id}).then((res) => {
            this.$message.success("删除成功");
            this.GetTaskList();
          }).catch(err => {
            this.$message.error("删除失败");
          })
        })
      },

      //导出兑换码.
      exportCsv: function (id) {
        acApi.exportCsv({id: parseInt(id)}).then((res) => {
          let obj = res.data.msg;

          var downloadLink = document.getElementById('downloadLink');

          let context = "任务标题:,"+obj.title+"\n序号,兑换码\n";
          for (let i = 0; i < obj.cdkey.length; i++) {
            context += i+1+","+obj.cdkey[i]+"\n"
          }

          context = encodeURIComponent(context);
          downloadLink.download = obj.title+".csv";// 下载的文件名称
          downloadLink.href = "data:text/csv;charset=utf-8,\ufeff" + context; //加上 \ufeff BOM 头
          downloadLink.click();
        });
      }
    },
    mounted() {
      this.GetTaskList();
    }
  }
</script>

<style>

  .form_item {
    margin-top: 15px;
    margin-bottom: 15px;
  }

  .ri {
    text-align: right;
  }

  .specifyTypeShow {
    display: none;
  }

  .specifyTypeNotShow {
    display: block;
  }
</style>

