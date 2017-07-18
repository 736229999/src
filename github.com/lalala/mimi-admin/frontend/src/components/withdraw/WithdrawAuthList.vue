<template>
    <div class="withdraw-list">
        <div class="query-tools">
    
            <el-form label-position="right" :model="queryForm" label-width="60px">
                <el-row>
                    <el-col :span="2">
                        <!--<el-form-item>-->
                        <el-button type="primary" @click="handleAddOrEdit(0)" style="margin-left:50px">添加</el-button>
                        <!--</el-form-item>-->
                    </el-col>
    
                    <!--<el-col :span="4">
                            <el-form-item label="审核人" style="margin-left:10px">
                                <el-radio-group v-model="queryForm.isMine" @change="handleisMineChange">
                                    <el-radio-button :label="1">全部</el-radio-button>
                                    <el-radio-button :label="2">只看我的</el-radio-button>
                                </el-radio-group>
                            </el-form-item>
                        </el-col>-->
                    <el-col :span="3">
                        <el-form-item label="审核人" prop="userId">
                            <el-select v-model="queryForm.userId" filterable placeholder="请选择">
                                <el-option v-for="auditor in auditors" :key="auditor.userId" :label="auditor.name" :value="auditor.userId">
                                </el-option>
                            </el-select>
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
    
        <el-table border :data="withdrawAuditAuthList" style="width: 100%" stripe>
            <el-table-column prop="userInfo.username" label="姓名" show-overflow-tooltip>
            </el-table-column>
            <el-table-column prop="type" label="操作类型" show-overflow-tooltip :formatter="typeFormat">
            </el-table-column>
            <!--<el-table-column prop="idcardNo" label="身份证号" show-overflow-tooltip>
                                    </el-table-column>
                                    <el-table-column prop="phone" label="手机号码" show-overflow-tooltip>
                                    </el-table-column>-->
            <el-table-column prop="amount" label="操作额度" show-overflow-tooltip>
                <template scope="scope">
                    <span>{{ scope.row.minAmount + "—" + scope.row.maxAmount}}</span>
                </template>
            </el-table-column>
            <el-table-column prop="createTime" label="创建时间" :formatter="timestampFormat">
            </el-table-column>
    
            <el-table-column prop="creatorUserInfo.username" label="创建人" show-overflow-tooltip>
            </el-table-column>
    
            <el-table-column prop="status" label="状态" width="120">
                <template scope="scope">
                    <el-tag :type="scope.row.valid == true ? 'primary' : 'danger'" close-transition>{{ validFormat(scope.row.valid) }}</el-tag>
                </template>
            </el-table-column>
    
            <el-table-column prop="address" label="操作">
                <template scope="scope">
                    <el-button size="small" :type="'primary'" @click="handleAddOrEdit(scope.row.id)">编辑</el-button>
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

import usercenterApi from '../../api/usercenter';

import Vue from 'vue';

export default {

    mounted() {
        this.getWithdrawAuditAuth()
    },
    data() {
        return {
            queryForm: {
                userId: null,
                status: null,
                realname: null,
                page: null,
                pagesize: null
            },
            statusOptions: [
                {
                    value: 1,
                    label: '待审核'
                }, {
                    value: 2,
                    label: '审核通过'
                }, {
                    value: 3,
                    label: '审核未通过'
                },
            ],
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
            withdrawAuditAuthList: [],
            pager: {
                tableListTotal: 20,
                pageSize: 10,
                currentPage: 1
            }
        }
    },
    methods: {
        buttonTextFormat(row) {
            if (row.status > 1) {
                return '查看'
            } else if (row.auditor != '') {
                return '审核'
            }
            return '认领'
        },
        submitQuery() {
            this.getWithdrawAuditAuth()
        },
        getWithdrawAuditAuth() {
            if (this.queryForm.status == '') {
                this.queryForm.status = null
            }
            usercenterApi.getWithdrawAuditAuthList(this.queryForm).then((res) => {
                console.log('getWithdrawAuditAuthList response', res);
                this.withdrawAuditAuthList = res.data.msg.list
                this.pager.tableListTotal = res.data.msg.total
            }).catch((err) => {
                this.$message.error(err.data.msg)
            })
        },
        handleCurrentChange(val) {
            this.pager.currentPage = val;
            this.queryForm.page = this.pager.currentPage
            this.queryForm.pagesize = this.pager.pageSize
            this.getWithdrawApplyList();
        },
        timestampFormat(row, column) {
            let format = Vue.filter('timeStampFormat')
            let msg
            switch (column.property) {
                case 'createTime':
                    msg = format(row.createTime)
                    break;
                case 'auditTime':
                    if (row.status > 1) {
                        msg = format(row.auditTime)
                    } else {
                        msg = ''
                    }
                    break;
                default:
                    break;
            }
            return msg
        },
        validFormat(val) {
            if (val) {
                return '有效'
            } else {
                return '无效'
            }
        },
        typeFormat(row, column) {
            switch (row.type) {
                case 1:
                    return '申请'
                    break;
                case 2:
                    return '转账'
                    break;
                default:
                    break;
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
        handleisMineChange(val) {
            console.log('handleisMineChange', val);
        },
        handleAddOrEdit(id = 0) {
            let topath
            if (id == 0) {
                topath = '/usercenter/withdraw/auth/add'
            } else {
                topath = '/usercenter/withdraw/auth/edit/' + id
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