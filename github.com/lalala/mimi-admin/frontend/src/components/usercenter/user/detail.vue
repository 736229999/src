<template>
  <div>
    <el-col :span="12" class="contain">
      <el-form :model="form" label-width="80px" ref="form">
        <el-form-item label="昵称:">
          <el-col :span="24">
            <span><el-tag type="text">{{form.nickname}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="手机:">
          <el-col :span="24">
            <span><el-tag type="gray">{{form.phone}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="性别:">
          <el-col :span="24">
            <span>
              <el-tag type="gray" v-if="form.sex == 0">未知</el-tag>
              <el-tag type="gray" v-if="form.sex == 1">男</el-tag>
              <el-tag type="gray" v-if="form.sex == 2">女</el-tag>
            </span>
          </el-col>
        </el-form-item>
        <el-form-item label="邀请码:">
          <el-col :span="24">
            <span><el-tag type="gray">{{form.invitation_code}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="积分:">
          <el-col :span="24">
            <span><el-tag type="gray">{{form.credits}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="开心豆">
          <span><el-tag type="gray">{{form.kxd}}</el-tag></span>
        </el-form-item>
        <el-form-item label="绑定手机">
          <span>
            <el-tag type="gray" v-if="form.bind_phone == true">已绑定</el-tag>
            <el-tag type="gray" v-if="form.bind_phone == false">未绑定</el-tag>
          </span>
        </el-form-item>
        <el-form-item label="绑定QQ">
          <span>
            <el-tag type="gray" v-if="form.bind_qq == true">已绑定</el-tag>
            <el-tag type="gray" v-if="form.bind_qq == false">未绑定</el-tag>
          </span>
        </el-form-item>
        <el-form-item label="绑定微信">
          <span>
            <el-tag type="gray" v-if="form.bind_wechat == true">已绑定</el-tag>
            <el-tag type="gray" v-if="form.bind_wechat == false">未绑定</el-tag>
          </span>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">确定</el-button>
        </el-form-item>

      </el-form>
    </el-col>
  </div>
</template>


<script>
  import ucApi from "../../../api/usercenter"
  import router from "../../../router"
  export default {
    data() {
      return {
        form: {},
      }
    },
    methods: {
      //获取反馈的详细.
      GetUserDetail: function () {
        ucApi.getUserDetail({id: this.$route.params.id}).then((res) => {
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
      this.GetUserDetail();
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

