<template>
    <div class="withdraw-audit-form">
        <el-row>
            <h3 style="margin-left: 10px;">提现申请详情</h3>
        </el-row>
        <el-form :model="form" label-width="80px" :rules="formRules" ref="form">
            <el-row :span="24">
                <el-col :span="3">
                    <el-form-item label="提现人">
                        <span>
                            <el-tag type="gray" style="font-size:15px"> {{ form.realname }}</el-tag>
                        </span>
                    </el-form-item>
                </el-col :span="4">
            </el-row>
    
            <el-row :span="24">
                <el-col :span="4">
                    <el-form-item label="身份证号">
                        <span>
                            <el-tag type="gray" style="font-size:15px" class="haha"> {{ form.idcard_no }}</el-tag>
                        </span>
                    </el-form-item>
                </el-col>
                <el-col :span="4">
                    <el-form-item label="手机号码">
                        <span>
                            <el-tag type="gray" style="font-size:15px"> {{ form.phone }}</el-tag>
                        </span>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-row :span="24">
                <el-col :span="4">
                    <el-form-item label="提现账户">
                        <span>
                            <el-tag type="primary" style="font-size:15px"> {{ form.in_no }}</el-tag>
                        </span>
                    </el-form-item>
                </el-col>
                <el-col :span="4">
                    <el-form-item label="账户类型">
                        <span>
                            <el-tag type="gray" style="font-size:15px"> {{ withdrawTypeText }}</el-tag>
                        </span>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-row :span="24">
                <el-col :span="3">
                    <el-form-item label="申请时间">
                        <span>
                            <el-tag type="gray" style="font-size:15px" v-if="form.create_time"> {{ form.create_time | timeStampFormat }}</el-tag>
                        </span>
                    </el-form-item>
                </el-col :span="4">
            </el-row>
            <el-row :span="24">
                <el-col :span="3">
                    <el-form-item label="提现金额">
                        <span>
                            <el-tag type="danger" style="font-size:15px">¥ {{ form.amount }}</el-tag>
                        </span>
                    </el-form-item>
                </el-col :span="4">
            </el-row>
    
            <el-row :span="24">
                <el-col :span="3">
                    <el-form-item label="当前状态">
                        <span>
                            <el-steps :space="100" :active="form.step" :finish-status="form.is_success ? 'success' : 'error'">
                                <el-step title="提现申请" :status="form.step > 1 ? 'success' : ''"></el-step>
                                <el-step title="平台审核" :status="form.step > 2 ? 'success' : ''"></el-step>
                                <el-step title="银行处理" :status="form.step < 3 ? 'wait': ''"></el-step>
                                <!--<el-step title="银行处理" :status="'process'"></el-step>-->
                            </el-steps>
                            <!--<el-tag :type="form.status == 1 ? '' : (form.status == 2 ? 'success' : 'danger')" style="font-size:15px">{{ statusFormat(form.status) }}</el-tag>-->
                        </span>
                    </el-form-item>
                </el-col :span="4">
            </el-row>
    
            <el-row :span="24">
                <el-col :span="12">
                    <el-form-item label="审核备注" prop="targetId">
                        <el-input type="textarea" :autosize="{ minRows: 2, maxRows: 4}" placeholder="请输入内容" v-model="form.auditComment">
                        </el-input>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-form-item>
                <el-button type="primary" @click="onSubmit(2, true)" v-if="this.form.step == 1">审核通过</el-button>
                <el-button type="danger" @click="onSubmit(2, false)" v-if="this.form.step == 1">驳回申请</el-button>
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

import userFund from '../../sections/usercenter/Fund';

// =====================后台显示状态对应的数值=====================
// 未审核: step: 1, is_success: true ,
// 审核通过: step: 2, is_success: true ,
// 审核未通过: step: 2, is_success: false ,
// 银行处理通过: step: 3, is_success: true ,
// 银行处理未通过: step: 3, is_success: false ,
export default {
    components: {
        'user-fund': userFund
    },
    beforeMount() {
        this.withdrawApplyId = this.$route.params.id;
        if (this.withdrawApplyId) {     //修改时加载的数据
            this.checkWithdrawApply();
        }

    },
    mounted() {
        // this.withdrawApplyId = this.$route.params.id;
        // if (this.withdrawApplyId) {     //修改时加载的数据
        //     this.getWithdrawApplyById();
        // }
    },
    data() {
        return {
            withdrawApplyId: 0,
            // 检查申请的结果
            checkRes: {
                // 是否存在审核人
                existAuditor: false,
                // 是否为自己审核的申请
                isOwn: false,
            },
            form: {
                id: null,
                account_id: null,
                realname: null,
                idcard_no: null,
                phone: null,
                withdraw_type: null,
                in_no: null,
                amount: null,
                audit_comment: null,
                create_time: null,

            },
            selectQuery: {
                options: [],
                loading: false,
            },
            formRules: {

            },

            imageUrl: '23434',
            dialogVisible: false,
        };
    },
    computed: {
        withdrawTypeText() {
            switch (this.form.withdraw_type) {
                case 1:
                    return '银行卡'
                case 2:
                    return '支付宝'
                default:
                    break;
            }
        },
    },
    methods: {
        goBack() {
            this.$router.go(-1)
        },
        handleClose(down) {
            console.log('down', down);
        },
        checkWithdrawApply() {
            usercenterApi.checkWithdrawApply({ id: this.withdrawApplyId }).then((res) => {
                console.log('checkWithdrawApply res', res);
                this.checkRes = res.data.msg
                this.checkWithdrawApplyRes()
            }).catch((err) => {
                this.$message.error(err.data.msg)
            })
        },
        checkWithdrawApplyRes() {
            // 判断是否可以认领
            if (!this.checkRes.exist_auditor) {
                console.log('no auditor');
                this.claimWithdrawApply();
            } else {
                console.log('exist auditor');
                if (!this.checkRes.is_own) {
                    this.$message.error('无权查看')
                    this.$router.push('/usercenter/withdraw/list')
                    return
                }
                this.getWithdrawApplyById()
            }
        },
        statusFormat(val) {
            let msg
            switch (val) {
                case 1:
                    msg = '未审核'
                    break;
                case 2:
                    msg = '审核通过'
                    break;
                case 3:
                    msg = '审核未通过'
                default:
                    break;
            }
            return msg
        },
        handleSortChange(val) {
            console.log('handleSortChange', val);
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
        onSubmit(step, is_success) {
            console.log('onSubmit step', step, 'is_success', is_success);
            this.$refs.form.validate((valid) => {
                if (valid) {
                    this.$confirm('确认提交吗? \n 请注意核实该用户中奖金额是否大于提现金额', '提示', { type: 'warning' }).then(() => {
                        console.log('this.form', this.form);
                        this.putWithdrawApply(step, is_success)
                    }).catch(() => {
                        this.$message.info('取消提交')
                    });
                } else {
                    return
                }
            })

        },
        getWithdrawApplyById() {
            const vm = this
            if (this.withdrawApplyId) {
                usercenterApi.getWithdrawApplyDetail({ id: this.withdrawApplyId }).then((res) => {
                    console.log('getWithdrawApplyDetail res', res);
                    if (res.data.msg !== undefined) {
                        this.form = res.data.msg

                    }
                }).catch((err) => {
                    this.$message.error(err.data.msg)
                    this.goBack()
                })
            }
        },
        claimWithdrawApply() {
            const vm = this
            if (this.withdrawApplyId) {
                usercenterApi.claimWithdrawApply({ withdrawApplyId: Number(this.withdrawApplyId) }).then((res) => {
                    console.log('claimWithdrawApply res', res);
                    if (res.data.msg !== undefined) {
                        this.form = res.data.msg
                    }
                }).catch((err) => {
                    console.log('err', err);
                    this.$message.error(err.data.msg)
                })
            }
        },
        putWithdrawApply(step, is_success) {
            this.form.id = this.bannerId
            this.form.id = Number(this.form.id)
            console.log('put WithdrawApply', this.form);

            let params = {
                step: step,
                is_success: is_success,
                id: Number(this.withdrawApplyId),
                auditComment: this.form.auditComment,
            }

            usercenterApi.updateWithdrawApplyStatus(params).then((res) => {
                console.log('put WithdrawApply res', res);
                this.$message('修改成功')
                this.$router.push({ path: '/usercenter/withdraw/list' })
            }).catch((err) => {
                this.$message.error(err.msg)
                console.log('put WithdrawApply error', err);
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

.dialog-audit {
    text-align: center;
    margin-bottom: 100px;
}

.md-editor {
    margin: 20px;
}


.banner-form {
    margin: 50px;
}

.withdraw-audit-form el-tag {
    font-size: 20px;
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