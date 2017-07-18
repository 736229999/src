<template>
  <div>
    <el-col :span="8" class="contain">
      <el-form :model="form" label-width="80px" :rules="formRules" ref="form">
        <el-form-item label="活动标题" prop="title">
          <el-col :span="24">
            <el-input v-model="form.title" auto-complete="off"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="活动描述" prop="des">
          <el-col :span="24">
            <el-input v-model="form.des" auto-complete="off"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="参与人数" prop="num">
          <el-col :span="24">
            <el-input v-model.number="form.num"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="选择礼包" prop="package_id">
          <template>
            <el-select v-model.number="form.package_id" placeholder="请选择" style="width:100%;" >
              <el-option v-for="item in gift.gift_list" :key="item.id" :label="item.title" :value="item.id">
              </el-option>
            </el-select>
          </template>
        </el-form-item>
        <el-form-item label="开始时间" prop="starttime">
          <el-col :span="24">
            <el-date-picker style="width: 100%" v-model="form.starttime" type="datetime" placeholder="选择日期时间">
            </el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="结束时间" prop="endtime">
          <el-col :span="24">
            <el-date-picker style="width:100%" v-model="form.endtime" type="datetime" placeholder="选择日期时间">
            </el-date-picker>
          </el-col>
        </el-form-item>
        <el-form-item label="绑定任务" prop="starttime">
          <el-select v-model="form.taskLists" multiple placeholder="请选择">
            <el-option
              v-for="item in tasks"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>

        <el-row :span="24">
          <el-col :span="5">
            <el-tooltip class="item" effect="dark" content="点击可添加或修改" placement="bottom">
              <el-form-item label="活动封面">
                <el-upload class="avatar-uploader" :action="uploadUrl" :show-file-list="false" :on-success="handleAvatarSuccess">
                  <img v-if="form.logo" :src="avatarUrl" class="avatar">
                  <i v-else class="el-icon-plus avatar-uploader-plus-icon" style="line-height: 178px;"></i>
                </el-upload>
              </el-form-item>
            </el-tooltip>
          </el-col>

        </el-row>


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
  import apiConst from '../../api/constant';
  export default {
    data() {
      return {
        tasks: [],

        form: {
          logo:"",
          title: "",
          des:"",
          num:"",
          package_id:"",
          starttime:"",
          endtime:"",
          taskLists:[],
        },
        gift:{
//            type:0,
            gift_list:[],
        },
        disabled: false,
        formRules: {
          title: [{required: true, message: '标题不能为空', trigger: 'blur'}],
          des: [{required: true, message: '描述不能为空', trigger: 'blur'}],
          num: [
            {required: true, message: '限制人数不能为空'},
            {type: 'number', message: '限制人数为整数'}
          ],
          package_id: [
            {required: true, message: '礼包不能为空'},
            {type: 'number', message: '礼包为整型'}
          ],
          starttime: [{type: 'date', required: true, message: '请选择开始时间', trigger: 'blur'}],
          endtime: [{type: 'date', required: true, message: '请选择结束时间', trigger: 'blur'}],
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
                this.form.starttime = moment(this.form.starttime).unix();
                this.form.endtime = moment(this.form.endtime).unix();
              acApi.addActivity(this.form).then((res) => {
                this.$message.success("添加成功");
                router.push({path:"/activity/activity/activityList"})
              }).catch(err => {
                this.$message.error("添加失败");
              })
            })
          }
        });
      },
      handleAvatarSuccess(res, file) {
        this.form.logo = res.result
      },

      //获取礼包的列表.
      GetGiftList:function () {
        acApi.giftList(this.gift).then((res) => {
            this.gift.gift_list = res.data.msg.list;
        }).catch(err=>{

        });
        //获取任务列表
        acApi.allTask(this.form).then((res) => {
            this.tasks = res.data.msg.list;

        }).catch(err=>{

        });
      },

    },
    computed:{
      simplemde() {
        return this.$refs.markdownEditor.simplemde
      },
      uploadUrl() {
          //路径暂时存到新闻目录下，后面写api的时候再改地址
        return apiConst.ASSETS_API + '/assets/backend/upload/news'
      },
      avatarUrl() {
        return apiConst.ASSETS_API + this.form.logo
//        return "http://img1.vued.vanthink.cn/vued0a233185b6027244f9d43e653227439a.png";
      },
      buttonText() {
        if (this.newsId) {
          return '修改'
        } else {
          return '添加'
        }
      }
    },
    mounted() {
      this.GetGiftList();
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

