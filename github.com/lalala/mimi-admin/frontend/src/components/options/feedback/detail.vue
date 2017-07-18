<template>
  <div>
    <el-col :span="12" class="contain">
      <el-form :model="form" label-width="80px" ref="form">
        <el-form-item label="称呼:">
          <el-col :span="24">
            <span><el-tag type="text">{{form.name}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="邮箱:">
          <el-col :span="24">
            <span><el-tag type="gray">{{form.email}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="反馈内容:">
          <el-col :span="24">
            <span><el-tag type="gray">{{form.content}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="反馈时间:">
          <el-col :span="24">
            <span><el-tag type="gray">{{new Date(form.create_time * 1000).toLocaleString()}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="状态:">
          <el-col :span="24">
            <span>
              <el-select v-model="form.status" placeholder="请选择">
                <el-option v-for="item in options" :key="item.value" :label="item.label"
                           :value="item.value"></el-option>
              </el-select>
            </span>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">确定</el-button>
        </el-form-item>

      </el-form>
    </el-col>
  </div>
</template>


<script>
  import opApi from "../../../api/options"
  import router from "../../../router"
  export default {
    data() {
      return {
        form: {
          id: "",
          name: "",
          email: "",
          content: "",
          create_time: "",
          status: "",
        },
        options: [{
          value: 0,
          label: "待处理",
        }, {
          value: 1,
          label: "已查看",
        }, {
          value: 2,
          label: "已处理"
        }, {
          value: 3,
          label: "忽略"
        }],
      }
    },
    methods: {
      //获取反馈的详细.
      GetFeedbackDetail: function () {
        opApi.getFeedbackById({id: this.$route.params.id}).then((res) => {
          this.form = res.data.msg
        })
      },
      onSubmit: function () {
          this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'error'
          }).then(() => {
            opApi.updateFeedbackById(this.form).then((res) => {
              this.$message.success("修改成功");
              router.push({path:"/options/feedback/list"})
            }).catch(err => {
              this.$message.error("修改失败");
            })
          })
      }
    },
    mounted() {
      this.GetFeedbackDetail();
    }
  }
</script>

<style>
  .contain {
    margin: 50px;
  }

  .form_item {
    margin-top: 20px;
    margin-bottom: 20px;
    background: #f6f6f6;
    border-radius: 3px;
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

  #showTickets, #edit_ticket_black {
    margin-top: 15px;
    margin-bottom: 15px;
  }

  .el-table .info-row {
    background: #c9e5f5;
  }

  .el-table .positive-row {
    background: #e2f0e4;
  }
</style>

