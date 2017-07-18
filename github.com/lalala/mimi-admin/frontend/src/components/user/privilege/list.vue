<template>
  <div>
    <el-col :span="24">
      <el-button type="primary" @click="addAuth" style="margin: 20px 0px 0px 20px;">添加模块</el-button>
    </el-col>
  <el-col :span="10" style="margin: 20px 0px 20px 20px;">
    <el-tree :data="options" 　:default-expand-all	="true"　 :props="defaultProps" show-checkbox node-key="id" :expand-on-click-node="true"
             :render-content="renderContent"></el-tree>

    <el-dialog title="编辑权限" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="权限名称">
          <el-input v-model="form.name" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="权限标识">
          <el-input v-model="form.key" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="权限路径">
          <el-input v-model="form.path" auto-complete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
      </div>
    </el-dialog>

    <!--新增模块-->
    <el-dialog title="新增权限模块" :visible.sync="authDialog">
      <el-form :model="authForm">
        <el-form-item label="模块名称" >
          <el-input v-model="authForm.name" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="模块标识" >
          <el-input v-model="authForm.key" auto-complete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="authDialog = false">取 消</el-button>
        <el-button type="primary" @click="addAuthOk">确 定</el-button>
      </div>
    </el-dialog>

    <!--新增权限-->
    <el-dialog title="新增权限" :visible.sync="addDialog">
      <el-form :model="addForm">
        <el-form-item label="权限名称" >
          <el-input v-model.number="addForm.p_id" v-show="false"></el-input>
          <el-input v-model="addForm.name" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="权限标识" >
          <el-input v-model="addForm.key" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="权限路径 " >
          <el-input v-model="addForm.path" auto-complete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="addDialog = false">取 消</el-button>
        <el-button type="primary" @click="addOk">确 定</el-button>
      </div>
    </el-dialog>
  </el-col>
  </div>
</template>

<script>

  import authApi from "../../../api/auth"
  export default {
    data() {
      return {
        dialogFormVisible: false,
        authDialog:false,
        addDialog:false,
        authForm:{
          name:"",
          key:"",
          p_id:0,
          path:"/",
        },
        addForm:{
          name:"",
          key:"",
          p_id:"",
          path:"",
        },
        form: {},
        options: [],
        defaultProps: {
          children: 'children',
          label: 'label'
        }
      }
    },

    methods: {
      addAuth:function () {
        this.authDialog = true;
      },
      addAuthOk:function () {
          if (this.authForm.name == "") {
              this.$message.error("模块名不能为空")
              return
          }

          if (this.authForm.key == "") {
              this.$message.error("模块标识不能为空");
              return
          }
        authApi.addAuth(this.authForm).then((res)=>{
          this.$message.success("添加成功");
          this.GetPrivilegeList();
          this.authDialog = false;
        }).catch(err => {
            this.$message.error("添加失败")
        })
      },
      addOk:function () {
          if (this.addForm.name == "") {
              this.$message.error("模块名不能为空")
              return
          }

          if (this.addForm.key == "") {
              this.$message.error("模块标识不能为空");
              return
          }
        authApi.addAuth(this.addForm).then((res)=>{
          this.$message.success("添加成功");
          this.GetPrivilegeList();
          this.addDialog = false;
        }).catch(err => {
            this.$message.error("添加失败")
        })
      },
      //获取权限.
      GetPrivilegeList: function () {
        authApi.getAuthList({}).then((res)=>{
            console.log(res);
          this.options = res.data.msg.children
        })
      },
      append(store, data) {
//        let id = 1000;
//        store.append({id: id++, label: 'testtest', children: []}, data);
        this.addDialog = true;
        this.addForm.p_id = data.id
      },

      remove(store, data) {
        store.remove(data);
      },

      edit(store, data) {
        this.dialogFormVisible = true
      },

      renderContent(h, { node, data, store }) {
        return (
          <span>
          <span>
          <span>{node.label}</span>
        </span>
        <span style="float: right; margin-right: 20px">
          <el-button size="mini" type="primary" icon="plus" on-click={ () => this.append(store, data) }></el-button>
        <el-button size="mini"  type="danger" icon="delete" on-click={ () => this.remove(store, data) }></el-button>
        <el-button size="mini"  type="warning" icon="edit" on-click={ () => this.edit(store, data) }></el-button>
        </span>
        </span>);
      }    },
    mounted() {
      this.GetPrivilegeList();
    }
  };
</script>
