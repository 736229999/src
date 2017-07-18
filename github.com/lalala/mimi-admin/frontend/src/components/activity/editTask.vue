<template>
  <div>
    <el-col :span="8" class="contain">
      <el-form :model="form" label-width="80px" :rules="formRules" ref="form">
        <el-form-item label="任务名" prop="name">
          <el-col :span="10">
            <el-input v-model="form.name" auto-complete="off"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="任务描述" prop="des">
          <el-col :span="24">
            <el-input v-model="form.des" auto-complete="off"></el-input>
          </el-col>
        </el-form-item>

        <el-form-item>
          <el-col :span="24">
            <el-button type="primary" @click="onSubmit">确定修改</el-button>
          </el-col>
        </el-form-item>
      </el-form>
    </el-col>
  </div>
</template>


<script>
  import acApi from "../../api/activity"
  import router from "../../router"
  export default {
    data() {
      return {
        form: {
          id:0,
          name: "",
          des : "",
        },

        disabled: false,
        formRules: {
          name: [{required: true, message: '标题不能为空', trigger: 'blur'}],
          des: [{required: true, message: '描述不能为空', trigger: 'blur'}],
        }
      }
    },
    methods: {
      onSubmit() {
        this.disabled = true;
        this.$refs.form.validate((valid) => {
          if (valid) {
            this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'error'
            }).then(() => {
                let moment = require("moment");
                acApi.updateTask(this.form).then((res) => {
                this.$message.success("修改成功");
                console.info(this.form);
                router.push({path:"/activity/task/taskList"})
              }).catch(err => {
                this.$message.error("添加失败");
              })
            })
          }
        });
      },
      QueryTaskById:function () {
        acApi.queryTaskById({id: this.$route.params.id}).then((res) => {

          this.form = res.data.msg;
          console.log(this.form.id);
        }).catch(err => {

        })

      },

    },
    mounted() {
      this.QueryTaskById();
    }
  }
</script>

<style>
  .contain {
    margin: 50px;
  }

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

