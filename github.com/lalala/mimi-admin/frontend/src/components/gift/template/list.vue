<template>
  <div>
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="礼包模板标题">
          <el-col :span="23">
            <el-input v-model="form.title" placeholder="礼包模板标题"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="search" icon="search">查询</el-button>
        </el-form-item>
        <el-form-item>
          <router-link :to="{path:'/gift/template/add'}">
            <el-button type="primary" icon="plus">新增礼包模板</el-button>
          </router-link>
        </el-form-item>
      </el-form>
    </el-col>
    <!--列表-->
    <el-col>
      <el-table :data="data" border>
        <el-table-column label="序号" prop="id" width="100"></el-table-column>
        <el-table-column label="标题" prop="title"></el-table-column>
        <el-table-column label="描述" prop="content_desc"></el-table-column>
        <el-table-column label="添加时间">
          <template scope="scope">
            <span>{{new Date(scope.row.add_time * 1000).toLocaleString()}}</span>
          </template>
        </el-table-column>
        <el-table-column label="礼包状态">
          <template scope="scope">
            <span v-if="scope.row.status == false && scope.row.gift_type != 3">未使用</span>
            <span v-else="scope.row.status == true">使用中</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="300">
          <template scope="scope">
            <router-link :to="{path:'/gift/template/detail/'+scope.row.id}">
              <el-button type="primary" size="small" icon="more"></el-button>
            </router-link>
            <router-link :to="{path:'/gift/template/edit/'+scope.row.id}"
                         v-if="scope.row.status == false && scope.row.gift_type != 3">
              <el-button type="warning" size="small" icon="edit"></el-button>
            </router-link>
            <el-button type="warning" size="small" icon="edit" v-else @click="message"></el-button>
            <el-button type="danger" size="small" icon="delete" @click="del(scope.row.id)"
                       v-if="scope.row.status == false && scope.row.gift_type != 3"></el-button>
            <el-button type="danger" size="small" icon="delete" @click="message" v-else
                       style="margin-left: -0px;"></el-button>
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
  import giftApi from '../../../api/gift';

  export default {
    data() {
      return {
        data: [],
        total: 0,
        form: {
          title: '',
          size: 100,
          page: 1,
          total: 0,
        },
        dialogTableVisible: false,
      }
    }, methods: {
      handleClick: function () {
        console.log(1);
      },
      GetGiftList() {
        giftApi.giftTemplateList(this.form).then((res) => {
          this.data = res.data.msg.list;
          this.form.total = res.data.msg.total;
        })
      },
      handleCurrentChange(val) {
        this.form.page = val;
        this.GetGiftList()
      },
      //查询.
      search: function () {
        this.GetGiftList()
      },

      //删除.
      del: function (id) {
        this.$confirm('是否确定删除当前礼包, 是否删除?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          giftApi.delGift({id: id}).then((res) => {
            this.$message.success("删除成功");
            this.GetGiftList()
          }).catch(err => {
            this.$message.error('删除失败')
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });
        });
      },
      message: function () {
        this.$message.error("当前礼包处于使用中，无法执行此操作")
      }

    },
    mounted() {
      this.GetGiftList();
    }
  }
</script>

