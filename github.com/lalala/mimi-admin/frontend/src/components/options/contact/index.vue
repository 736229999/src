<template>
  <div>
    <el-col :span="12" class="contain">
      <el-form :model="form">
        <el-form-item label="客服QQ" prop="qq">
          <el-input type="text" v-model="form.qq"></el-input>
        </el-form-item>
        <el-form-item label="客服微信" prop="wechat">
          <el-input type="text" v-model="form.wechat"></el-input>
        </el-form-item>
        <el-form-item label="官方邮箱" prop="email">
          <el-input type="email" v-model="form.email"></el-input>
        </el-form-item>
        <el-form-item label="官方电话" prop="telphone">
          <el-input type="email" v-model="form.telphone"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">立即修改</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </div>
</template>
<script>
  import ElInput from "../../../../node_modules/element-ui/packages/input/src/input";
  import lotteryApi from "../../../api/lottery"
  import router from "../../../router/index";
  import Vue from "vue";
  export default {
    components: {ElInput},
    data() {
      return {
        form: {
          qq: "",
          wechat: '',
          email: "",
          telphone: "",
        },
      }
    },
    methods: {
      onSubmit() {
        this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          lotteryApi.updateContactInfo(this.form).then((res) => {
            this.$message.success("修改成功");
          }).catch(err => {
            this.$message.error("修改失败");
          })
        })
      },

      GetContact: function () {
        lotteryApi.getContact({}).then((res) => {
          this.form = res.data.msg
        }).catch(err => {
          this.$message.error("获取失败");
        })
      }
    },
    mounted() {
      this.GetContact()
    }
  }
</script>
