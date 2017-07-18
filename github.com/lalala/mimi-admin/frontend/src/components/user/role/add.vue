<template>
  <div>

    <el-col :span="10" style="margin: 20px 0px 20px 20px;">
      <el-form :model="form">
        <el-form-item label="角色名称">
          <el-input v-model="form.role_name" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="form.remarks" auto-complete="off"></el-input>
        </el-form-item>
        <!--<el-form-item>-->
      <el-tree :data="options" 　:default-expand-all	="true"　 :props="defaultProps" show-checkbox node-key="id" :expand-on-click-node="true"
               :render-content="renderContent"></el-tree>
        <!--</el-form-item>-->
      </el-form>
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
        </span>);
      }    },
    mounted() {
      this.GetPrivilegeList();
    }
  };
</script>
