<template>
  <div class="news-list">
    <div class="query-tools">

      <el-form label-position="right" :model="queryForm" label-width="80px">
        <el-row>
          <el-col :span="2">
            <!--<el-form-item>-->
            <el-button type="primary" @click="handleAddOrEdit(0)" style="margin-left:50px">添加</el-button>
            <!--</el-form-item>-->
          </el-col>

          <el-col :span="5">
            <el-form-item label="新闻标题">
              <el-input v-model="queryForm.title">

              </el-input>
            </el-form-item>
          </el-col>

          <el-col :span="4">
            <el-form-item label="作者">
              <el-input v-model="queryForm.author">

              </el-input>
            </el-form-item>
          </el-col>

          <el-col :span="4">
            <el-form-item label="发布时间">
              <el-date-picker v-model="dateRange" type="daterange" align="right" placeholder="选择日期范围"
                              :picker-options="pickerOptions">
              </el-date-picker>
            </el-form-item>
          </el-col>

          <el-col :span="4">
            <el-form-item>
              <el-button type="primary" @click="submitQuery()">查询</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>

    <el-table border :data="newsList" style="width: 100%" stripe>
      <el-table-column prop="title" label="标题" show-overflow-tooltip>
      </el-table-column>
      <el-table-column prop="description" label="封面描述" show-overflow-tooltip>
      </el-table-column>
      <el-table-column prop="author" label="作者">
      </el-table-column>
      <el-table-column prop="created" label="发布时间" :formatter="timestampFormat">
      </el-table-column>
      <el-table-column prop="pageViews" label="浏览量" width="120">
      </el-table-column>
      <el-table-column prop="cover" label="封面">
        <template scope="scope">
          <div slot="reference" class="name-wrapper">
            <el-button type="text" @click="preview(scope.row.cover,scope.row.title)">
              <el-icon name="picture"></el-icon>
            </el-button>
          </div>
        </template>
      </el-table-column>

      <el-table-column prop="isVisible" label="状态" width="120">
        <template scope="scope">
          <el-tag :type="scope.row.isVisible === true ? 'primary' : 'danger'" close-transition>
            {{ scope.row.isVisible | switchStatusFormat }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column prop="address" label="操作">
        <template scope="scope">
          <el-button size="small" type="primary" @click="handleAddOrEdit(scope.row.id)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-col :span="24" class="toolbar">
      <el-pagination layout="total, prev, pager, next" @current-change="handleCurrentChange" :page-size="pager.pageSize"
                     :total="pager.tableListTotal" style="float:right;">
      </el-pagination>
    </el-col>

    <pic-view :imgUrl="picRreview.imgUrl" :title="picRreview.title" v-model="picRreview.visible"></pic-view>

  </div>
</template>

<script>

  import newsApi from '../../api/discover';

  import PicPreview from '../picPreview/PicPreview';

  import apiConst from '../../api/constant';

  import moment from 'moment';

  export default {
    components: {
      'pic-view': PicPreview
    },
    mounted() {
      this.getNewsList()
    },
    data() {
      return {
        queryForm: {
          title: null,
          author: null,
          end: null,
          class: null,
          page: null,
          pagesize: null
        },
        dateRange: [],
        newsList: [
          // {
          //     title: '震惊13亿人的消息-测试',
          //     author: '无节操小编-测试',
          //     created: '2017-05-17-测试',
          //     description: '震惊13亿人的消息，男默女泪-测试',
          //     pageViews: 999,
          // }
        ],
        pager: {
          tableListTotal: 20,
          pageSize: 10,
          currentPage: 1
        },
        pickerOptions: {
          shortcuts: [{
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
          }, {
            text: '最近三个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
              picker.$emit('pick', [start, end]);
            }
          }]
        },
        picRreview: {
          imgUrl: null,
          visible: false,
          title: '预览图片',
        }
      }
    },
    computed: {},
    methods: {
      submitQuery() {
        if (this.dateRange.length === 2) {
          if (this.dateRange[0] !== null || this.dateRange[0] !== null) {
            this.queryForm.start = moment(this.dateRange[0]).unix()
            this.queryForm.end = moment(this.dateRange[1]).unix()
          } else {
            this.queryForm.start = this.queryForm.end = null
          }
        }
        this.getNewsList()
      },
      getNewsList() {
        newsApi.getNewsList(this.queryForm).then((res) => {
          console.log('getNewsList response', res);
          this.newsList = res.data.list
          this.pager.tableListTotal = res.data.total
        }).catch()
      },
      handleCurrentChange(val) {
        this.pager.currentPage = val;
        this.queryForm.page = this.pager.currentPage
        this.queryForm.pagesize = this.pager.pageSize
        this.getNewsList();
      },
      preview(url, title) {
        this.picRreview = {
          imgUrl: apiConst.ASSETS_API + url,
          visible: true,
          title: title
        }
      },
      timestampFormat(row, column) {
        var moment = require('moment');
        return moment.unix(row.created).format('YYYY-MM-DD HH:mm:ss');
      },
      handleAddOrEdit(id = 0) {
        let topath
        if (id == 0) {
          topath = '/news/add'
        } else {
          topath = '/news/edit/' + id
        }
        this.$router.push({
          path: topath
        })
      },
      switchChange(valSta, valId) {
        console.log('val', valSta, valId);
        // console.log('row', row);
      }
    }
  }
</script>

<style lang="">
  .query-tools {
    margin-top: 10px;
  }
</style>
