<template>
	<section>
		<!--工具条-->
		<el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
			<el-form :inline="true" :model="filters">
				<el-form-item>
					<el-input v-model="filters.name" placeholder="名称"></el-input>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" v-on:click="getPrivileges">查询</el-button>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" @click="handleAdd">新增</el-button>
				</el-form-item>
			</el-form>
		</el-col>

		<!--列表-->
		<el-table :data="privileges" highlight-current-row v-loading="listLoading" @selection-change="selsChange" style="width: 100%;">
			<el-table-column type="selection">
			</el-table-column>
			<el-table-column prop="id" label="ID" sortable>
			</el-table-column>
			<el-table-column prop="name" label="名称" sortable>
			</el-table-column>
			<el-table-column prop="path" label="路径" sortable>
			</el-table-column>
			<el-table-column prop="creator" label="创建人" sortable>
			</el-table-column>
			<el-table-column prop="create_time" label="创建时间" sortable>
			</el-table-column>
			<el-table-column label="操作" width="150">
				<template scope="scope">
					<el-button size="small" @click="handleEdit(scope.row.id, scope.row)">编辑</el-button>
					<el-button type="danger" size="small" @click="handleDel(scope.row.id)">删除</el-button>
				</template>
			</el-table-column>
		</el-table>

		<!--工具条-->
		<el-col :span="24" class="toolbar">
			<el-pagination layout="prev, pager, next" @current-change="handleCurrentChange" :page-size="20" :total="total" style="float:right;">
			</el-pagination>
		</el-col>

	<!--新增界面-->
	<el-dialog title="新增" v-model="addFormVisible" :close-on-click-modal="false">
			<el-form :model="addForm" label-width="80px" :rules="addFormRules" ref="addForm">
				<el-form-item label="权限名称" prop="name">
					<el-input v-model="addForm.name" auto-complete="off" type="name"></el-input>
				</el-form-item>
				<el-form-item label="权限key" prop="key">
					<el-input v-model="addForm.key" auto-complete="off" type="key"></el-input>
				</el-form-item>
				<el-form-item label="权限路径" prop="path">
					<el-input v-model="addForm.path" auto-complete="off" type="path"></el-input>
				</el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button @click.native="addFormVisible = false">取消</el-button>
				<el-button type="primary" @click.native="addSubmit" :loading="addLoading">提交</el-button>
			</div>
	</el-dialog>

	<!--编辑界面-->
	<el-dialog title="编辑" v-model="editFormVisible" :close-on-click-modal="false">
		<el-form :model="editForm" label-width="80px" :rules="editFormRules" ref="editForm">
			<el-form-item label="权限名称" prop="name">
				<el-input v-model="editForm.name" auto-complete="off" type="name"></el-input>
			</el-form-item>
			<el-input v-model="editForm.id" auto-complete="off" type="hidden"></el-input>
			<el-form-item label="权限key" prop="key">
				<el-input v-model="editForm.key" auto-complete="off" type="key"></el-input>
			</el-form-item>
			<el-form-item label="权限路径" prop="path">
				<el-input v-model="editForm.path" auto-complete="off" type="path"></el-input>
			</el-form-item>
		</el-form>
		<div slot="footer" class="dialog-footer">
			<el-button @click.native="editFormVisible = false">取消</el-button>
			<el-button type="primary" @click.native="editSubmit" :loading="editLoading">提交</el-button>
		</div>
	</el-dialog>

	</section>
</template>

<script>
  import util from '../../common/js/util'
  import NProgress from 'nprogress'
  import {
    privilegeList,
    addPrivilege,
    editPrivileges,
    deletePrivilege,
  } from '../../api/api';

  export default {
  	data() {
        return {
            filters: {
                name: ''
            },
            privileges: [],
            total: 0,
            page: 1,
            listLoading: false,
            addFormVisible: false, //新增界面是否显示
            editFormVisible: false, //编辑界面是否显示
            addLoading: false,
            editLoading: false,
            //新增界面数据
            addForm: {
                id: '',
                privilege_name: '',
                creator: '',
                privi_ids: []
            },
            addFormRules: {
                name: [{required: true, message: '请输入权限名称', trigger: 'blur'}],
                key: [{required: true, message: '请输入权限的key', trigger: 'blur'}],
                path: [{required: true, message: '请输入权限的路径', trigger: 'blur'}],
            },
            editFormRules: {
                id: [{required: true, message: '系统错误', trigger: 'blur'}],
                name: [{required: true, message: '请输入权限名称', trigger: 'blur'}],
                key: [{required: true, message: '请输入权限key', trigger: 'blur'}],
                path: [{required: true, message: '请输入权限路径', trigger: 'blur'}],
            },
            editForm: {
                id: '',
                name: '',
                key: '',
                path: '',
            },
        }
    },
    methods: {
    	getPrivileges() {
    		let params = new URLSearchParams();
            this.listLoading = true;
            NProgress.start();
    		privilegeList(params).then(( res ) => {
    		    var obj = JSON.parse(res.data);
    		    obj = obj.msg;
    		    console.log(obj);
    			this.listLoading = false;
    			NProgress.done();
    			this.total = obj.length;
    			this.privileges = obj;
    		});
    	},
    	handleCurrentChange(val) {
          this.page = val;
          this.getPrivileges();
        },
        selsChange: function(sels) {
          this.sels = sels;
        },
        //显示新增界面
        handleAdd: function() {
            this.addFormVisible = true;
            this.addForm = {
                name: '',
                key: '',
                path: '',
            };
        },
        //新增
        addSubmit: function() {
            this.$refs.addForm.validate((valid) => {
                if (valid) {
					this.addLoading = true;
					NProgress.start();

					let params = new URLSearchParams();
					params.append('id', this.addForm.id);
					params.append('name', this.addForm.name);
					params.append('key', this.addForm.key);
					params.append('path', this.addForm.path);

					addPrivilege(params).then((res) => {
						this.addLoading = false;
						NProgress.done();
						var obj = JSON.parse(res.data);
						obj = obj.msg;
						if ( res.status !== 200 ) {
							this.$message.error( obj );
						} else {
							this.$message.success('提交成功！')

							this.addFormVisible = false;
                            this.getPrivileges()
						}
					});
                }
            });
        },
        handleEdit:function(index, row) {
            this.editFormVisible = true;
            this.editForm = Object.assign({}, row);
		},
        //删除
        handleDel: function(index) {
            this.$confirm('确认删除该记录吗?', '提示', { type: 'warning' }).then(() => {
                this.listLoading = true;
                NProgress.start();

                let params = new URLSearchParams();
                params.append('id', index);
                deletePrivilege(params).then((res) => {
                    var obj = JSON.parse(res.data);
                    this.listLoading = false;
                    NProgress.done();
                    if ( res.status !== 200 ) {
                        this.$message.error( obj.msg );
                    }else {
                        this.$message.success('删除成功！');
                        this.getPrivileges();
                    }
                });
            }).catch(() => {

            });
        },
        //编辑
        editSubmit: function() {
            this.$refs.editForm.validate((valid) => {
                if (valid) {
					this.editLoading = true;
					NProgress.start();

					let params = new URLSearchParams();
					params.append('id', this.editForm.id);
					params.append('name', this.editForm.name);
					params.append('key', this.editForm.key);
					params.append('path', this.editForm.path);
					editPrivileges(params).then((res) => {
						this.editLoading = false;
						NProgress.done();
						var obj = JSON.parse(res.data)
						obj = obj.msg;
						if ( res.status !== 200 ) {
							this.$message.error( obj );
						} else {
							this.$message.success('提交成功！')
							this.$refs['editForm'].resetFields();
							this.editFormVisible = false;
							this.getPrivileges()
						}
					});
                }
            });
        },
    },
    mounted() {
      this.getPrivileges();
    }
  }
</script>

<style>
	.contain {
		margin:50px;
	}
</style>
