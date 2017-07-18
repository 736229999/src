<template>
  <div>
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true">
        <el-form-item label="活动标题">
          <el-col :span="23">
            <el-input v-model="form.title" placeholder="标题"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="search">查询</el-button>
          <router-link :to="{path:'/activity/activity/add'}">
            <el-button type="primary" icon="plus">添加</el-button>
          </router-link>
        </el-form-item>
      </el-form>
    </el-col>
    <el-col :span="24">
      <el-table :data="form.list" style="width: 100%">
        <el-table-column prop="id" label="序号" width="180"></el-table-column>
        <el-table-column prop="title" label="活动标题" width="180"></el-table-column>
        <el-table-column prop="des" label="活动描述"></el-table-column>
        <el-table-column prop="logo" label="活动logo">
          <template scope="scope">
            <div slot="reference" class="name-wrapper">
              <el-button type="text" @click="preview(scope.row.logo, scope.row)">
                <el-icon name="picture"></el-icon>
              </el-button>
            </div>

          </template>
        </el-table-column>
        <el-table-column prop="num" label="限制人数"></el-table-column>
        <el-table-column prop="package_name" label="对应礼包"></el-table-column>

        <el-table-column prop="starttime"label="活动开始时间">
          <template scope="scope">
            <span>{{ scope.row.starttime | timeStampFormat }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="endtime"label="活动结束时间">
          <template scope="scope">
            <span>{{ scope.row.endtime | timeStampFormat }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="create_time"label="活动创建时间">
          <template scope="scope">
            <span>{{ scope.row.create_time | timeStampFormat }}</span>
          </template>
        </el-table-column>
        <!--<el-table-column prop="create_admin" label="活动创建人" ></el-table-column>-->


        <el-table-column label="操作">
          <template scope="scope">
            <router-link :to="{path:'/activity/activity/edit/'+scope.row.id}">
              <el-button type="primary" size="mini" icon="more"></el-button>
            </router-link>
            <router-link :to="{path:'/activity/activity/edit/'+scope.row.id}">
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
    <pic-view :imgUrl="picRreview.imgUrl" :title="picRreview.title" v-model="picRreview.visible"></pic-view>
  </div>
</template>

<script>

  import acApi from "../../api/activity"
  import PicPreview from '../picPreview/PicPreview';
  import apiConst from '../../api/constant';
  export default {
    components: {
      'pic-view': PicPreview
    },
    data() {
      return {
        picRreview: {
          imgUrl: null,
          visible: false,
          title: '预览图片',
        },
        form: {
          total: 0,
          size: 3,
          page: 1,
          list: [],

        },
        disabled: false,
      }
    },
    methods: {
      preview(url, row) {
          console.log(url)
        this.picRreview = {
//          imgUrl: "http://dev.jiditv.com/index.jpg",
          imgUrl: apiConst.ASSETS_API + url,
          visible: true,
          title: "活动标题"
        }
      },
      GetActivityList: function () {
        acApi.activityList(this.form).then((res) => {
//          console.log(this.form)
          console.log(res.data.msg)
          this.form = res.data.msg;
        }).catch(err => {
          this.$message.error("获取失败")

        })
      },
      handleCurrentChange(val) {
        this.form.page = val;
        this.GetActivityList();
      },
      search: function () {
        this.GetActivityList();
      },
      del: function (id) {
        this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {

          acApi.deleteActivity({id: id}).then((res) => {
            this.$message.success("删除成功");
            this.GetActivityList();
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

          let context = "活动标题:,"+obj.title+"\n序号,兑换码\n";
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
      this.GetActivityList();
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

