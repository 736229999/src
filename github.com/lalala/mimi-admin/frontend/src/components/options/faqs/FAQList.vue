<template>
    <div class="faq-list">
        <div class="query-tools">
    
            <el-form label-position="right" :model="queryForm" label-width="40px">
                <el-row>
                    <el-col :span="2">
                        <!--<el-form-item>-->
                        <el-button type="primary" @click="handleAddOrEdit(0)" style="margin-left:50px">添加</el-button>
                        <!--</el-form-item>-->
                    </el-col>
    
                    <el-col :span="4">
                        <el-form-item label="标题">
                            <el-input v-model="queryForm.title">
    
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
    
        <el-table border :data="faqList" style="width: 100%" stripe>
            <el-table-column prop="title" label="标题" show-overflow-tooltip>
            </el-table-column>
    
            <el-table-column prop="content" label="内容" align="center" width="120">
                <template scope="scope">
                    <el-popover trigger="hover" placement="top" width="400">
                        <p>{{ scope.row.content }}</p>
                        <div slot="reference" class="name-wrapper">
                            <el-icon name="document"></el-icon>
                        </div>
                    </el-popover>
                </template>
            </el-table-column>
    
            <el-table-column prop="createTime" label="添加时间" :formatter="timestampFormat">
            </el-table-column>
            <el-table-column prop="updateTime" label="最近一次修改" :formatter="timestampFormat">
            </el-table-column>
    
            <el-table-column prop="isVisible" label="状态" width="120">
                <template scope="scope">
                    <el-tag :type="scope.row.isVisible === true ? 'primary' : 'danger'" close-transition>{{ scope.row.isVisible | switchStatusFormat }}</el-tag>
                </template>
            </el-table-column>
    
            <el-table-column prop="address" label="操作">
                <template scope="scope">
                    <el-button size="small" type="primary" @click="handleAddOrEdit(scope.row.id)">编辑</el-button>
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

import optionsApi from '../../../api/options';

import Vue from 'vue';

export default {
    mounted() {
        this.getFaqList()
    },
    data() {
        return {
            queryForm: {
                title: null,
                page: null,
                pagesize: null
            },
            faqList: [],
            pager: {
                tableListTotal: 20,
                pageSize: 10,
                currentPage: 1
            }
        }
    },
    methods: {
        submitQuery() {
            this.getFaqList()
        },
        getFaqList() {
            optionsApi.getFaqList(this.queryForm).then((res) => {
                console.log('getFaqList response', res);
                this.faqList = res.data.msg.list
                this.pager.tableListTotal = res.data.msg.total
            }).catch((err) => {
                this.$message.error(err.data.msg)
            })
        },
        handleCurrentChange(val) {
            this.pager.currentPage = val;
            this.queryForm.page = this.pager.currentPage
            this.queryForm.pagesize = this.pager.pageSize
            this.getFaqList();
        },
        timestampFormat(row, column) {
            let format = Vue.filter('timeStampFormat')
            let msg
            switch (column.property) {
                case 'createTime':
                    msg = row.createTime
                    break;
                case 'updateTime':
                    msg = row.updateTime
                    break;
                default:
                    break;
            }
            return format(msg)
        },
        handleAddOrEdit(id = 0) {
            let topath
            if (id == 0) {
                topath = '/options/faq/add'
            } else {
                topath = '/options/faq/edit/' + id
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
    margin-top: 10px;
}
</style>