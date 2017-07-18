<template>
  <div>
    <el-col :span="12" class="contain">
      <el-form :model="form" :rules="rulesForm">
        <el-form-item label="彩票名称" prop="lottery_name">
          <el-input type="text" v-model="form.lottery_name"></el-input>
        </el-form-item>
        <el-form-item label="开奖简介" prop="info">
          <el-input type="text" v-model="form.info"></el-input>
        </el-form-item>
        <el-form-item label="是否加奖" prop="is_plus_award">
          <br/>
          <el-switch v-model="form.is_plus_award" on-color="#13ce66" off-color="#ff4949" on-text="是"
                     off-text="否"></el-switch>
        </el-form-item>
        <el-form-item label="停止销售" prop="stop_sale">
          <br/>
          <el-switch v-model="form.stop_sale" on-color="#13ce66" off-color="#ff4949" on-text="是"
                     off-text="否"></el-switch>
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
  import ElInput from "../../../node_modules/element-ui/packages/input/src/input";
  import lotteryApi from "../../api/lottery"
  import router from "../../router";
  import Vue from "vue";
  export default {
    components: {ElInput},
    data() {
      return {
        form: {
          id: "",
          lottery_name: '',
          is_plus_award: "",
          info: "",
          stop_sale: "",
        },
        rulesForm: {
          id: [{required: true, message: '请选择彩种', trigger: 'blur'}],
          lottery_name: [{required: true, message: '彩种名字不能为空', trigger: 'blur'}],
          info: [{required: true, message: '开奖简介不能为空', trigger: 'blur'}],
          is_plus_award: [{required: true, message: '是否加奖不能为空', trigger: 'blur'}],
          stop_sale: [{required: true, message: '暂停销售不能为空', trigger: 'blur'}],
        },
        options: [],
      }
    },
    methods: {
      onSubmit() {
        this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          lotteryApi.uploadLotteryOptions(this.form).then((res) => {
            this.$message.success("修改成功");
            router.push({path: '/lottery/home/options'})
          }).catch(err=>{
            this.$message.error("修改失败");
          })
        })
      },
      GetLotteryOptionsById: function () {
        //获取彩种列表.
        let params = {
          id: this.$route.params.id,
        };
        lotteryApi.getLotteryOptionsById(params).then((res) => {
          this.form = res.data.msg;
        });
      },
    },
    mounted() {
      this.GetLotteryOptionsById();
    }
  }
</script>
