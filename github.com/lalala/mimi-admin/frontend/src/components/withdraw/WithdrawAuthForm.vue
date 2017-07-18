<template>
    <div class="withdraw-auth-form">
        <el-form :model="form" label-width="90px" :rules="formRules" ref="form">
            <el-row :span="24">
                <el-col :span="12">
                    <el-form-item label="审核人" prop="userId">
                        <el-select v-model="form.userId" filterable placeholder="请选择" :disabled="WithdrawAuditAuthId ? true : false">
                            <el-option v-for="auditor in auditors" :key="auditor.userId" :label="auditor.name" :value="auditor.userId">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row :span="24">
                <el-col :span="6">
                    <el-form-item label="最小值">
                        <el-input-number v-model="form.minAmount" @change="handleValueChange" :min="1" :disabled="form.unlimited ? true : false"></el-input-number>
                        <span>大于且等于该值</span>
                    </el-form-item>
                </el-col>
    
            </el-row>
    
            <el-row :span="24">
                <el-col :span="6">
                    <el-form-item label="最大值">
                        <el-input-number v-model="form.maxAmount" @change="handleValueChange" :min="1" :disabled="form.unlimited ? true : false"></el-input-number>
                        <span>小于该值</span>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-row>
                <el-col :span="4">
                    <el-form-item label="是否无限制" prop="unlimited">
                        <el-switch v-model="form.unlimited" on-color="#13ce66" off-color="#ff4949" on-text="是" off-text="否">
                        </el-switch>
                    </el-form-item>
                </el-col>
            </el-row>

            <el-row :span="24">
                <el-col :span="12">
                    <el-form-item label="操作类型" prop="type">
                        <el-radio-group v-model="form.type" @change="handleLocationChange">
                            <el-radio-button :label="1" :disabled="WithdrawAuditAuthId ? (form.type == 2 ? true : false) : false">审核</el-radio-button>
                            <el-radio-button :label="2" :disabled="WithdrawAuditAuthId ? (form.type == 1 ? true : false) : false">转账</el-radio-button>
                        </el-radio-group>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-row>
                <el-col :span="4">
                    <el-form-item label="当前状态" prop="valid">
                        <el-switch v-model="form.valid" on-color="#13ce66" off-color="#ff4949" on-text="允许" off-text="禁止">
                        </el-switch>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-form-item>
                <el-button type="primary" @click="onSubmit()">立即{{ buttonText }}</el-button>
                <el-button @click="goBack">返回</el-button>
            </el-form-item>
        </el-form>
    
    </div>
</template>
<script>
import discoverApi from '../../api/discover';
import apiConst from '../../api/constant';
import optionsApi from '../../api/options';
import usercenterApi from '../../api/usercenter';
export default {
    data() {
        return {
            WithdrawAuditAuthId: 0,
            form: {
                id: null,
                userId: null,
                minAmount: null,
                maxAmount: null,
                unlimited: false,
                vaild:false,
                type: 0,
            },
            auditors: [
                {
                    name: '侯忠建',
                    userId: 1,

                },
                {
                    name: '张三',
                    userId: 2,

                },
                {
                    name: '杨泽丰',
                    userId: 3,
                },
            ],
            selectQuery: {
                options: [],
                loading: false,
            },
            formRules: {
                userId: [{ type: 'number', required: true, message: '审核人不能为空', trigger: 'blur' }],
            },

            src: 'http://img1.vued.vanthink.cn/vued0a233185b6027244f9d43e653227439a.png',
            imageUrl: '',
            dialogTableVisible: false,
        };
    },
    computed: {
        uploadUrl() {
            return apiConst.ASSETS_API + '/assets/backend/upload/news'
        },
        avatarUrl() {
            return apiConst.ASSETS_API + this.form.url
        },
        buttonText() {
            if (this.WithdrawAuditAuthId) {
                return '修改'
            } else {
                return '添加'
            }
        }
    },
    mounted() {
        this.WithdrawAuditAuthId = this.$route.params.id;
        if (this.WithdrawAuditAuthId) {     //修改时加载的数据
            this.getWithdrawAuditAuthById();
        }
    },
    methods: {
        goBack() {
            this.$router.go(-1)
        },
        newsRemoteMethod(query) {
            console.log('newsRemoteMethod keyword', query);
            if (query !== '') {
                this.selectQuery.loading = true;
                discoverApi.queryNewsOfSelect({ keyword: query }).then((response) => {
                    console.log('queryNewsOfSelect', response);
                    this.selectQuery.loading = false;
                    this.selectQuery.options = response.data.msg.list
                }).catch((error) => {

                })
            }
        },
        handleValueChange(val) {
            console.log('handleSortChange', val);
            // if (this.form.maxAmount < this.form.minAmount) {
            //     this.$message('最大值不能小于最小值')
            //     this.form.maxAmount = this.form.minAmount + 1
            // }
        },
        handleLocationChange(val) {
            console.log('handleLocationChange', val);
        },
        handleTargetTypeChange(val) {
            console.log('handleTargetTypeChange', val);
        },
        handlePreview(file) {
            console.log(file);
            console.log('fileList:', this.fileList);
        },
        onSubmit() {
            console.log('onSubmit');
            this.$refs.form.validate((valid) => {
                if (valid) {
                    this.$confirm('确认提交吗?', '提示', { type: 'warning' }).then(() => {
                        console.log('this.form', this.form);
                        if (this.WithdrawAuditAuthId) {
                            // 更新Banner
                            this.putWithdrawAuditAuth()
                        } else {
                            // 添加审核信息
                            this.postWithdrawAuditAuth()
                        }
                    }).catch(() => {
                        this.$message.info('取消提交')
                    });
                } else {
                    return
                }
            })

        },
        getWithdrawAuditAuthById() {
            const vm = this
            if (this.WithdrawAuditAuthId) {
                usercenterApi.getWithdrawAuditAuthDetail({ id: this.WithdrawAuditAuthId }).then((res) => {
                    console.log('getWithdrawAuditAuthDetail', res);
                    if (res.data.msg !== undefined) {
                        this.form = res.data.msg
                    }
                }).catch((err) => {

                })
            }
        },
        postWithdrawAuditAuth() {
            console.log('post WithdrawAuditAuth', this.form);
            usercenterApi.addWithdrawAuditAuth(this.form).then((res) => {
                this.$message.success('创建成功')
                console.log('post WithdrawAuditAuth res', res);
                this.$router.push({ path: '/usercenter/withdraw/auth' })
            }).catch((error) => {
                console.log('post WithdrawAuditAuth error', error);
                this.$message.error(error.data.msg)
            })
        },
        putWithdrawAuditAuth() { 
            this.form.id = this.WithdrawAuditAuthId
            this.form.id = Number(this.form.id)
            console.log('put WithdrawAuditAuth', this.form);
            usercenterApi.updateWithdrawAuditAuth(this.form).then((res) => {
                console.log('put WithdrawAuditAuth', res);
                this.$message.success('修改成功')
                console.log('put WithdrawAuditAuth res', res);
                this.$router.push({ path: '/usercenter/withdraw/auth' })
            }).catch((err) => {
                console.log('put WithdrawAuditAuth', err);
            })
        },
        handleAvatarSuccess(res, file) {
            this.form.url = res.result
        },
        beforeAvatarUpload() {
            console.log('beforeAvatarUpload');
        }
    },
}
</script>

<style>
.banner-form-select {
    width: 150%;
}

.md-editor {
    margin: 20px;
}

.withdraw-auth-form {
    margin: 50px;
}

.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
}

.avatar-uploader .el-upload:hover {
    border-color: #20a0ff;
}

.avatar-uploader-plus-icon-banner {
    font-size: 28px;
    color: #8c939d;
    width: 375px;
    height: 199px;
    text-align: center;
}

.avatar-banner {
    width: 375px;
    height: 199px;
    display: block;
}
</style>