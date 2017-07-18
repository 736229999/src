<template>
	<section>
		<!--工具条-->
		<el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
			<el-form :inline="true" :model="filters">
				<el-form-item>
					<el-input v-model="filters.name" placeholder="名称"></el-input>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" v-on:click="getRoles">查询</el-button>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" @click="handleAdd">新增</el-button>
				</el-form-item>
			</el-form>
		</el-col>

		<!--列表-->
		<el-table :data="roles" highlight-current-row v-loading="listLoading" @selection-change="selsChange" style="width: 100%;">
			<el-table-column type="selection">
			</el-table-column>
			<el-table-column prop="id" label="ID" width="120" sortable>
			</el-table-column>
			<el-table-column prop="role_name" label="角色" width="100" sortable>
			</el-table-column>
			<el-table-column prop="creator" label="创建人" min-width="180" sortable>
			</el-table-column>
			<el-table-column prop="create_time" label="创建时间" min-width="180" sortable>
			</el-table-column>
			<el-table-column label="操作" width="150">
				<template scope="scope">
					<el-button size="small" @click="handleEdit(scope.row.id, scope.row)">详细</el-button>
					<el-button type="danger" size="small" @click="handleDel(scope.row.id)">删除</el-button>
				</template>
			</el-table-column>
		</el-table>

		<!--工具条-->
		<el-col :span="24" class="toolbar">
			<el-pagination layout="prev, pager, next" @current-change="handleCurrentChange" :page-size="20" :total="total" style="float:right;">
			</el-pagination>
		</el-col>

		<!--编辑界面-->
		<el-dialog title="编辑" v-model="editFormVisible" :close-on-click-modal="false">
		  <el-form :model="editForm" label-width="80px" :rules="editFormRules" ref="editForm">
        <el-form-item label="ID" prop="id">
          {{editForm.id}}
        </el-form-item>
		  	<el-form-item label="名称" prop="name">
          {{editForm.name}}
        </el-form-item>
        	<el-form-item label="备注" prop="remark">
          	  <el-input v-model="editForm.remark" auto-complete="off" type="remark"></el-input>
        	</el-form-item>
        	<el-form-item label="权限" prop="privi_ids">
    			  <!--<el-select v-model="editForm.privi_ids" multiple placeholder="请选择/搜索" filterable>-->
    	    		<!--<el-option v-for="item in privileges" :label="item.name" :value="item.id" :key="item.id">-->
    				  <!--</el-option>-->
    			  <!--</el-select>-->
        	</el-form-item>
	      </el-form>
	      <div slot="footer" class="dialog-footer">
    			<el-button @click.native="editFormVisible = false">取消</el-button>
    			<el-button type="primary" @click.native="editSubmit" :loading="editLoading">提交</el-button>
		  </div>
		</el-dialog>

		<!--新增界面-->
		<el-dialog title="新增" v-model="addFormVisible" :close-on-click-modal="false">
		  <el-form :model="addForm" label-width="80px" :rules="addFormRules" ref="addForm">
		  	<el-form-item label="角色名称" prop="name">
          	  <el-input v-model="addForm.name" auto-complete="off" type="name"></el-input>
        	</el-form-item>
        	<el-form-item label="权限" prop="privi_ids">
				<el-transfer v-model="value1" :titles="['权限列表', '已选择']"  @change="handleChange" :data="data"></el-transfer>
        	</el-form-item>
	      </el-form>
		  <div slot="footer" class="dialog-footer">
        	<el-button @click.native="addFormVisible = false">取消</el-button>
        	<el-button type="primary" @click.native="addSubmit" :loading="addLoading">提交</el-button>
      	  </div>
		</el-dialog>
	</section>
</template>

<script>
  import jsx from 'babel-plugin-transform-vue-jsx'
  import util from '../../common/js/util'
  import NProgress from 'nprogress'
  import {
    roleList,
    addRole,
    deleteRole,
    editRole,
    detailRole,
    privilegeList,
  } from '../../api/api';

  export default {
  	data() {
//        const generateData = _ => {
//            const data = [];
//            console.debug("data:", data);
////            for (let i = 1; i <= data,length; i++) {
//////                data.push({
//////                    label: `备选项`,
//////                });
////                console.log(data)
////            }
//            return data;
//        };
  	  return {
  		filters: {
          name: ''
	    },
	    roles: [],
	    total: 0,
	    page: 1,
	    listLoading: false,
	    //编辑界面数据
	    editFormVisible: false, //编辑界面是否显示
	    editLoading: false,
	    editFormRules: {
        remark: [{ required: true,message: '请输入备注',trigger: 'blur' }]
      },
	    editForm: {
	      id: '',
	      name: '',
	      remark: '',
	      creator: '',
	      create_time: '',
	      privi_ids: []
	    },
	    addFormVisible: false, //新增界面是否显示
	    addLoading: false,
	    //新增界面数据
	    addForm: {
  		  name : '',
		  privi_ids: []
	    },
          addFormRules: {
              name: [{ required: true,message: '请输入备注',trigger: 'blur' }],
          },
          data: [],
          value1: [],
      };
  	},
  	methods: {
        // 获取列表
        getRoles() {
          let params = new URLSearchParams();
          params.append('page', this.page);
          params.append('name', this.filters.name);
          params.append('size', 20);

          this.listLoading = true;
          NProgress.start();
          roleList(params).then((res) => {
              var obj = JSON.parse(res.data)
			  obj = obj.msg
            this.total = obj.length;
            this.roles = obj;
            this.listLoading = false;
            NProgress.done();
          });
        },
        // 权限
        getPrivileges() {

            NProgress.start();
            let params = new URLSearchParams();
            privilegeList(params).then((res) => {
                this.addLoading = false;
                NProgress.done();
				        var obj = JSON.parse(res.data);
                if ( res.status !== 200 ) {
                    this.$notify.error( obj.msg );
                } else {
                    console.log("obj.msg",obj.msg);
                    for(var i in obj.msg){
                        this.data.push({
                          key: obj.msg[i].key,
                          label: obj.msg[i].name
                        })
                    }
                }
            });
        },
        handleCurrentChange(val) {
          this.page = val;
          this.getRoles();
        },
        //显示新增界面
        handleAdd: function() {

            this.addLoading = true;
            this.addFormVisible = true;
            this.getPrivileges()
        },
        //删除
        handleDel: function(index) {
          this.$confirm('确认删除该记录吗?', '提示', { type: 'warning' }).then(() => {
            this.listLoading = true;
            NProgress.start();

            let params = new URLSearchParams();
            params.append('id', index);
            deleteRole(params).then((res) => {
              this.listLoading = false;
              NProgress.done();
              if ( res.status !== 200 ) {
                this.$notify.error( res.data.msg );
              }else {
                this.$notify.success('删除成功！');
                this.getRoles();
              }
            });
          }).catch(() => {

          });
        },
        //显示编辑界面
        handleEdit: function(index, row) {
          this.editFormVisible = true;
          this.editForm = Object.assign({}, row);
        },
        //编辑
        editSubmit: function() {
          this.$refs.editForm.validate((valid) => {
            if (valid) {
              this.$confirm('确认提交吗？', '提示', {}).then(() => {
                this.editLoading = true;
                NProgress.start();

                let params = new URLSearchParams();
                params.append('id', this.editForm.id);
                params.append('remark', this.editForm.remark);
                params.append('privi_ids', this.editForm.privi_ids);
                editRole(params).then((res) => {
                  this.editLoading = false;
                  NProgress.done();

                  if ( res.status !== 200 ) {
                    this.$notify.error( res.data.msg );
                  } else {
                    this.$notify.success('提交成功！');
                    this.$refs['editForm'].resetFields();
                    this.editFormVisible = false;
                    this.getRoles();
                  }
                });
              });
            }
          });
        },
        //新增
        addSubmit: function() {
            let num = this.addForm.privi_ids.length;
            num = num - 1;
            var privi_ids = this.addForm.privi_ids[num];
          this.$refs.addForm.validate((valid) => {
            if (valid) {
              this.$confirm('确认提交吗？', '提示', {}).then(() => {
                this.addLoading = true;
                NProgress.start();

                let params = new URLSearchParams();
                params.append('name', this.addForm.name);
                params.append('privi_ids', privi_ids);
                addRole(params).then((res) => {
                  this.addLoading = false;
                  NProgress.done();

                  if ( res.status !== 200 ) {
                    this.$notify.error( res.data.msg );
                  } else {
                    this.$notify.success('提交成功！');
                    this.$refs['addForm'].resetFields();
                    this.addFormVisible = false;
                    this.getRoles();
                  }
                });
              });
            }
          });
        },
        selsChange: function(sels) {
            this.sels = sels;
        },
        handleChange(value, direction, movedKeys) {
            this.addForm.privi_ids.push(value)
        }
  	},
  	mounted() {
  		this.getRoles();
  	}
  }
</script>
