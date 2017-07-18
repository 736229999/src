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
        <el-form-item label="选择类型" prop="type">
          <template>
            <el-select v-model="form.type" placeholder="请选择">
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </template>
        </el-form-item>
        <el-form-item label="达到金额" prop="money">
          <el-col :span="10">
            <el-input v-model.number="form.money"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-col :span="24">
            <el-button type="primary" @click="onSubmit">确定添加</el-button>
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
        options: [{
          value: '新人',
          label: '新人'
        }, {
          value: '充值',
          label: '充值'
        }, {
          value: '购彩',
          label: '购彩'
        }],
        form: {
          name: "",
          des : "",
          type : "",
          money : 0,
        },

        disabled: false,
        formRules: {
          name: [{required: true, message: '标题不能为空', trigger: 'blur'}],
          des: [{required: true, message: '描述不能为空', trigger: 'blur'}],
          type: [{required: true, message: '类型不能为空', trigger: 'blur'}],
          money: [
            {required: true, message: '需要达到的金额不能为空'},
            {type: 'number', message: '需要达到的金额为整数'}
          ],
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
                acApi.addTask(this.form).then((res) => {
                this.$message.success("添加成功");
                console.info(this.form);
                router.push({path:"/activity/task/taskList"})
              }).catch(err => {
                this.$message.error("添加失败");
              })
            })
          }
        });
      },


    },
    mounted() {
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

