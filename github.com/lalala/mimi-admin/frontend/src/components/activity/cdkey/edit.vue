<template>
  <div>
    <el-col :span="8" class="contain">
      <el-form :model="form" label-width="80px" :rules="formRules" ref="form">
        <el-form-item label="活动标题" prop="title">
          <el-col :span="24">
            <el-input v-model="form.title" auto-complete="off"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="活动描述" prop="cdkey_desc">
          <el-col :span="24">
            <el-input v-model="form.cdkey_desc" auto-complete="off"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="可兑换数" prop="max_exchange">
          <el-col :span="24">
            <el-input v-model.number="form.max_exchange"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="选择礼包" prop="gift_package_id">
          <template>
            <el-select v-model.number="form.gift_template_id" placeholder="请选择" style="width:100%;">
              <el-option v-for="item in gift.gift_list" :key="item.id" :label="item.title" :value="item.id">
              </el-option>
            </el-select>
          </template>
        </el-form-item>
        <el-form-item label="开始时间" prop="valid_start">
          <el-col :span="24">
            <el-date-picker style="width: 100%" v-model="form.valid_start" type="datetime" placeholder="选择日期时间">
            </el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="结束时间" prop="valid_start">
          <el-col :span="24">
            <el-date-picker style="width:100%" v-model="form.valid_end" type="datetime" placeholder="选择日期时间">
            </el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-col :span="24">
            <el-button type="primary" @click="onEdit">修改</el-button>
          </el-col>
        </el-form-item>
      </el-form>
    </el-col>
  </div>
</template>


<script>

  import acApi from "../../../api/activity"
  import router from "../../../router"

  export default {
    data() {
      return {
        form: {
          title: "",
          valid_start: "",
          valid_end: "",
          max_exchange: "",
          gift_template_id: "",
        },
        gift: {
          type: 0,
          gift_list: [],
        },
        disabled: false,
        formRules: {
          title: [{required: true, message: '标题不能为空', trigger: 'blur'}],
          cdkey_desc: [{required: true, message: '描述不能为空', trigger: 'blur'}],
          max_exchange: [
            {required: true, message: '可兑换数不能为空'},
            {type: 'number', message: '可兑换数为整数'}
          ],
          gift_template_id: [
            {required: true, message: '礼包不能为空'},
            {type: 'number', message: '礼包为整型'}
          ],
          valid_start: [{type: 'date', required: true, message: '请选择开始时间', trigger: 'blur'}],
          valid_end: [{type: 'date', required: true, message: '请选择结束时间', trigger: 'blur'}],
        }
      }
    },
    methods: {
      onEdit() {
        this.$refs.form.validate((valid) => {
          if (valid) {
            this.$confirm('确定提交了吗? 填错了会被老板骂哟！要不要在确认一下？', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'error'
            }).then(() => {
              let moment = require("moment");
              this.form.valid_start = moment(this.form.valid_start).unix();
              this.form.valid_end = moment(this.form.valid_end).unix();
              this.form.id = parseInt(this.$route.params.id);

              acApi.editCdkeyBatch(this.form).then((res) => {
                this.$message.success("更新成功");
                router.push({path: "/activity/cdkey/list"})
              }).catch(err => {
                this.$message.error("更新失败");
              })
            })
          }
        });
      },

      //获取礼包的列表.
      GetGiftTemplateList: function () {
        acApi.getGiftTemplateList(this.gift).then((res) => {
          this.gift.gift_list = res.data.msg.list;
        }).catch(err => {

        })
      },
      GetCdkeyById: function () {
        acApi.getCdkeyDetail({id: this.$route.params.id}).then((res) => {
          this.form = res.data.msg;
          let moment = require("moment");
          this.form.valid_start = moment(this.form.valid_start * 1000).toDate();
          this.form.valid_end = moment(this.form.valid_end * 1000).toDate();

          console.log(this.form)
        }).catch(err => {

        })
      }
    },
    mounted() {
      this.GetGiftTemplateList();
      this.GetCdkeyById();
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

