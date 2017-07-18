<template>
    <div class="withdraw-transfer-list">
        <div class="query-tools">
    
            <el-form label-position="right" :model="queryForm" label-width="60px">
                <el-row>
                    <!--<el-col :span="2">
                                                                                                                                <el-form-item>
                                                                                                                                <el-button type="primary" @click="handleAddOrEdit(0)" style="margin-left:50px">添加</el-button>
                                                                                                                                </el-form-item>
                                                                                                                            </el-col>-->
    
                    <el-col :span="4">
                        <el-form-item label="审核人" style="margin-left:10px">
                            <el-radio-group v-model="queryForm.isMine" @change="handleisMineChange">
                                <el-radio-button :label="null">全部</el-radio-button>
                                <el-radio-button :label="true">只看我的</el-radio-button>
                            </el-radio-group>
                        </el-form-item>
                    </el-col>
                    <el-col :span="3">
                        <el-form-item label="状态">
                            <el-select v-model="status" placeholder="请选择" clearable @change="handleStatusChange">
                                <el-option v-for="item in applyStatusOptions" :key="item.value" :label="item.label" :value="item.value">
                                </el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
    
                    <el-col :span="3">
                        <el-form-item label="姓名">
                            <el-input v-model="queryForm.realname">
    
                            </el-input>
                        </el-form-item>
                    </el-col>
    
                    <el-col :span="2">
                        <el-form-item>
                            <el-button type="primary" @click="submitQuery()">查询</el-button>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
        </div>
    
        <el-table border :data="WithdrawTransferList" style="width: 100%" stripe>
            <el-table-column prop="withdraw_apply.account_id" label="账户ID" show-overflow-tooltip>
            </el-table-column>
            <el-table-column prop="withdraw_apply.realname" label="姓名" show-overflow-tooltip>
            </el-table-column>
            <!--<el-table-column prop="idcardNo" label="身份证号" show-overflow-tooltip>
                                                                                                        </el-table-column>
                                                                                                        <el-table-column prop="phone" label="手机号码" show-overflow-tooltip>
                                                                                                        </el-table-column>-->
            <el-table-column prop="withdraw_apply.amount" label="金额" show-overflow-tooltip>
            </el-table-column>
            <el-table-column prop="withdraw_apply.create_time" label="用户申请时间" :formatter="timestampFormat">
            </el-table-column>
    
            <el-table-column prop="withdraw_apply.step" label="审核状态">
                <template scope="scope">
                    <el-tag :type="scope.row.withdraw_apply.step == 1 ? '' : (scope.row.withdraw_apply.is_success ? 'success' : 'danger')" close-transition>{{ applyStatusFormat(scope.row) }}</el-tag>
                </template>
            </el-table-column>
    
            <el-table-column prop="withdraw_apply.auditor" label="审核人" show-overflow-tooltip>
            </el-table-column>
    
            <el-table-column prop="withdraw_apply.auditComment" label="审核备注" align="center" width="120" :formatter="auditFormat">
                <template scope="scope">
                    <el-popover trigger="hover" placement="top" width="400">
                        <p>{{ scope.row.withdraw_apply.audit_comment }}</p>
                        <div slot="reference" class="name-wrapper">
                            <el-icon name="document" v-if="scope.row.withdraw_apply.step > 1"></el-icon>
                        </div>
                    </el-popover>
                </template>
            </el-table-column>
    
            <el-table-column prop="withdraw_apply.audit_time" label="审核时间" :formatter="timestampFormat">
            </el-table-column>
    
            <el-table-column prop="withdraw_apply.step" label="转账状态">
                <template scope="scope">
                    <el-tag :type="scope.row.step == 1 ? '' : (scope.row.is_success ? 'success' : 'danger')" close-transition>{{ transferStatusFormat(scope.row) }}</el-tag>
                </template>
            </el-table-column>
    
            <el-table-column prop="address" label="操作">
                <template scope="scope">
                    <el-button size="small" :type="scope.row.step > 1 ? '' : 'primary'" @click="confirmOpen(scope.$index, scope.row)" v-show="scope.row.can_operate"> {{ buttonTextFormat(scope.row) }}</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-col :span="24" class="toolbar">
            <el-pagination layout="total, prev, pager, next" @current-change="handleCurrentChange" :page-size="pager.pageSize" :total="pager.tableListTotal" style="float:right;">
            </el-pagination>
        </el-col>
    
    </div>
</template>

<script>

import optionsApi from '../../api/options';

import usercenterApi from '../../api/usercenter';

import Vue from 'vue';

export default {

    mounted() {
        this.getWithdrawTransferList()
    },
    data() {
        return {
            status: null,
            queryForm: {
                isMine: null,
                step: null,
                is_success: null,
                realname: null,
                page: null,
                pagesize: null
            },
            applyStatusOptions: [        // 后台显示状态对应的数值
                { value: 1, label: '未审核', step: 1, is_success: true },
                { value: 2, label: '审核通过', step: 2, is_success: true },
                { value: 3, label: '审核未通过', step: 2, is_success: false },
                { value: 4, label: '银行处理通过', step: 3, is_success: true },
                { value: 5, label: '银行处理未通过', step: 3, is_success: false }
            ],
            transferStatusOptions: [
                { value: 1, label: '等待中', step: 1, is_success: true },
                { value: 2, label: '已转账', step: 2, is_success: true },
                { value: 3, label: '银行处理通过', step: 3, is_success: true },
                { value: 4, label: '银行处理未通过', step: 3, is_success: false },
            ],
            WithdrawTransferList: [],
            pager: {
                tableListTotal: 20,
                pageSize: 10,
                currentPage: 1
            }
        }
    },
    computed: {
        queryArg() {
            _.find(this.applyStatusOptions)
        }
    },
    methods: {
        buttonTextFormat(row) {
            if (row.step > 1) {
                return '查看'
            } else if (row.auditor != '') {
                return '确认转账'
            }
            return '认领'
        },
        handleStatusChange(val) {
            console.log('handleStatusChange', val);
            if (val === '') {
                this.queryForm.step = this.queryForm.is_success = null
            } else {
                let status = _.find(this.applyStatusOptions, { value: val })
                this.queryForm.step = status.step
                this.queryForm.is_success = status.is_success
            }
        },
        submitQuery() {
            this.getWithdrawTransferList()
        },
        getWithdrawTransferList() {
            if (this.queryForm.step == '') {
                this.queryForm.step = null
            }
            usercenterApi.getWithdrawTransferList(this.queryForm).then((res) => {
                console.log('getWithdrawTransferList response', res);
                this.WithdrawTransferList = res.data.msg.list
                this.pager.tableListTotal = res.data.msg.total
            }).catch((err) => {
                this.$message.error(err.data.msg)
            })
        },
        handleCurrentChange(val) {
            this.pager.currentPage = val;
            this.queryForm.page = this.pager.currentPage
            this.queryForm.pagesize = this.pager.pageSize
            this.getWithdrawTransferList();
        },
        timestampFormat(row, column) {
            let format = Vue.filter('timeStampFormat')
            let msg
            switch (column.property) {
                case 'withdraw_apply.create_time':
                    msg = format(row.withdraw_apply.create_time)
                    break;
                case 'withdraw_apply.audit_time':
                    if (row.withdraw_apply.step > 1) {
                        msg = format(row.withdraw_apply.audit_time)
                    } else {
                        msg = ''
                    }
                    break;
                default:
                    break;
            }
            return msg
        },
        applyStatusFormat(val) {
            let status = _.find(this.applyStatusOptions, { step: val.withdraw_apply.step, is_success: val.withdraw_apply.is_success })
            if (status.label !== undefined) {
                return status.label
            } else {
                return ''
            }
        },
        transferStatusFormat(val) {
            let status = _.find(this.transferStatusOptions, { step: val.step, is_success: val.is_success })
            if (status.label !== undefined) {
                return status.label
            } else {
                return ''
            }
        },
        auditFormat(row, column) {
            let msg
            switch (column.property) {
                case '':

                    break;

                default:
                    break;
            }
        },
        confirmOpen(index, row) {
            console.log('confirmOpen');
            if (row.can_operate) {
                // console.log('can Operate');
                if (row.step > 1 || row.auditor != '') {
                    // console.log('自己的未审核');
                    this.handleAddOrEdit(row.id)
                } else {
                    // console.log('初次认领');
                    this.$confirm('确认认领吗？此操作将不可撤销, 是否继续?', '提示', {
                        confirmButtonText: '确定',
                        cancelButtonText: '取消',
                        type: 'warning'
                    }).then(() => {
                        this.handleAddOrEdit(row.id)
                    }).catch(() => {
                        this.$message({
                            type: 'info',
                            message: '已取消认领'
                        });
                    });
                }
            } else {
                this.$message.error('无权操作')
            }
        },
        handleisMineChange(val) {
            console.log('handleisMineChange', val);
            this.getWithdrawTransferList()
        },
        handleAddOrEdit(id = 0) {
            let topath
            if (id == 0) {
                topath = '/usercenter/withdraw/add'
            } else {
                topath = '/usercenter/withdraw/transfer/detail/' + id
            }
            this.$router.push({
                path: topath
            })
        },
    }
}
</script>

<style lang="">
.query-tools {
    margin-top: 50px;
}
</style>