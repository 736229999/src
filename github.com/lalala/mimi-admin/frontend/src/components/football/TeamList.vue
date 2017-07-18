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
              <el-select v-model="queryForm.name" clearable filterable remote placeholder="输入球队名称搜索" :remote-method="teamRemoteMethod" :loading="selectQuery.loading" style="width:100%">
                <el-option v-for="item in selectQuery.options" :key="item.id" :label="item.name" :value="item.name"></el-option>
              </el-select>
            </el-form-item>
          </el-col>


          <el-col :span="2">
            <el-form-item>
              <el-button type="primary" @click="getTeamList()">查询</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>

    <el-table border :data="teamList" style="width: 100%" stripe>
      <el-table-column prop="id" label="球队ID">
      </el-table-column>
      <el-table-column prop="name" label="球队名称">
      </el-table-column>
      <el-table-column prop="alias" label="别名">
      </el-table-column>
      <el-table-column prop="cover" label="预览队徽">
        <template scope="scope">
          <div slot="reference" class="name-wrapper">
            <el-button type="text" @click="preview(scope.row.url, scope.row)">
              <el-icon name="picture"></el-icon>
            </el-button>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="address" label="操作" width="300">
        <template scope="scope">
          <el-button size="small" type="primary" icon="edit" @click="handleAddOrEdit(scope.row.id)"></el-button>
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
  import PicPreview from '../picPreview/PicPreview';
  import footbApi from '../../api/football';
  import apiConst from '../../api/constant';

  export default {
    components: {
      'pic-view': PicPreview
    },
    mounted() {
      this.getTeamList()
    },

    data() {
      return {
        teamList: [],
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
        picRreview: {
          imgUrl: null,
          visible: false,
          title: '预览图片',
        }
      }
    },
    methods: {
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
      getTeamList() {
        footbApi.getTeamList(this.queryForm).then((res) => {
          console.log('getTeamList response', res);
          this.teamList = res.data.msg.list
          this.pager.tableListTotal = res.data.msg.total
        }).catch()
      },
      handleCurrentChange(val) {
        this.pager.currentPage = val;
        this.queryForm.page = this.pager.currentPage
        this.queryForm.pagesize = this.pager.pageSize
        this.getTeamList();
      },

      handleAddOrEdit(id = 0) {
        let topath
        if (id === 0) {
          topath = '/football/team/add'
        } else {
          topath = '/football/team/edit/' + id
        }
        this.$router.push({
          path: topath
        })
      },
      preview(url, row) {
        this.picRreview = {
          imgUrl: apiConst.ASSETS_API + url,
          visible: true,
        }
      },
    }
  }
</script>

<style>

</style>
