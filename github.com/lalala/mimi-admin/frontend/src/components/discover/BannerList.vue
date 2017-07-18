<template>
    <div>
        <div class="discover-banner" align="center">
            <span class="demonstration">
                <span style="color:red">{{ locationText }}</span> 效果预览</span>
                <p v-if="carousels.total == 0"> 无</p>
            <el-carousel indicator-position="outside" height="200px" v-if="carousels.total > 0">
                <el-carousel-item v-for="item in bannerListFormat.list" :key="item" :label="item.id">
                    <!--<h3>{{ item }}</h3>-->
                    <img :src="item.url" alt="" style="width: 100%; height: 100%">
                    <div>
                        <p>{{item.description}}</p>
                    </div>
                </el-carousel-item>
            </el-carousel>
        </div>
        <div>
            <div class="query-tools">

                <el-form label-position="right" :model="queryForm" label-width="80px">
                    <el-row>
                        <el-col :span="2">
                            <!--<el-form-item>-->
                            <el-button type="primary" @click="handleAddOrEdit(0)" style="margin-left:50px">添加</el-button>
                            <!--</el-form-item>-->
                        </el-col>

                        <el-col :span="5">
                            <el-form-item label="位置">
                                <el-radio-group v-model="queryForm.location" @change="handleLocationChange">
                                    <el-radio-button :label="1">首页</el-radio-button>
                                    <el-radio-button :label="2">发现</el-radio-button>
                                </el-radio-group>
                            </el-form-item>
                        </el-col>

                        <!--<el-col :span="4">
                                            <el-form-item>
                                                <el-button type="primary" @click="submitQuery()">查询</el-button>
                                            </el-form-item>
                                        </el-col>-->
                    </el-row>
                </el-form>
            </div>
            <el-table border :data="bannerList" style="width: 100%" stripe>
                <el-table-column prop="id" label="ID" show-overflow-tooltip>
                </el-table-column>
                <el-table-column prop="cover" label="预览">
                    <template scope="scope">
                        <div slot="reference" class="name-wrapper">
                            <el-button type="text" @click="preview(scope.row.url, scope.row)">
                                <el-icon name="picture"></el-icon>
                            </el-button>
                        </div>

                    </template>
                </el-table-column>
                <el-table-column prop="description" label="描述" show-overflow-tooltip>
                </el-table-column>
                <el-table-column prop="location" label="位置" show-overflow-tooltip :formatter="locationFormat">
                </el-table-column>
                <el-table-column prop="targetType" label="跳转类型" :formatter="targetTypeFormat">
                </el-table-column>
                <el-table-column prop="targetId" label="跳转目标ID">
                </el-table-column>
                <el-table-column prop="created" label="创建时间" :formatter="createdFormat">
                </el-table-column>

                <el-table-column prop="isVisible" label="状态" width="120">
                    <template scope="scope">
                        <el-tag :type="scope.row.isVisible === true ? 'primary' : 'danger'" close-transition>{{ scope.row.isVisible | switchStatusFormat }}</el-tag>
                    </template>
                </el-table-column>

                <el-table-column prop="" label="操作">
                    <template scope="scope">
                        <el-button size="small" type="primary" @click="handleAddOrEdit(scope.row.id)">编辑</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-col :span="24" class="toolbar">
                <el-pagination layout="total, prev, pager, next" @current-change="handlePagerChange" :page-size="pager.pageSize" :total="pager.tableListTotal" style="float:right;">
                </el-pagination>
            </el-col>

            <pic-view :imgUrl="picRreview.imgUrl" :title="picRreview.title" v-model="picRreview.visible"></pic-view>
        </div>
    </div>
</template>

<script>
import PicPreview from '../picPreview/PicPreview';

import optionsApi from '../../api/options'

import Vue from 'vue';

import apiConst from '../../api/constant';

export default {
    components: {
        'pic-view': PicPreview
    },
    mounted() {
        this.getBannerList()
        this.getPreBanner()
    },
    data() {
        return {
            location: 1,
            currentDate: new Date(),
            testNum: 4,
            queryForm: {
                location: 1,
                page: null,
                pagesize: null
            },
            bannerList: [],
            pager: {
                tableListTotal: 20,
                pageSize: 10,
                currentPage: 1
            },
            carousels: {
                list:[],
                total: 0,
            },
            picRreview: {
                imgUrl: null,
                visible: false,
                title: '预览图片',
            }
        };
    },
    computed: {
        locationText() {
            switch (this.queryForm.location) {
                case 1:
                    return '首页';
                case 2:
                    return '发现';

                default:
                    break;
            }
        },
        bannerListFormat() {
            this.carousels.list.forEach(function(element) {
                element.url = apiConst.ASSETS_API + element.url
            }, this);
            return {
                list: this.carousels.list,
                total: this.carousels.total,
            }
        }
    },
    methods: {
        getBannerList() {
            optionsApi.getBannerList(this.queryForm).then((res) => {
                console.log('getBannerList', res.data);
                this.bannerList = res.data.msg.list
                this.pager.tableListTotal = res.data.msg.total
            }).catch((err) => {
                console.log('getBannerList error', err);
                this.$message.error(err.data.msg)
            })
        },
        locationFormat(row, column) {
            switch (row.location) {
                case 1:
                    return '首页';
                case 2:
                    return '发现';
            }
        },
        targetTypeFormat(row, column) {
            let msg = ''
            switch (row.targetType) {
                case 1:
                    msg = '新闻'
                    break;
                case 2:
                    msg = '活动'
                    break
                case 3:
                    msg = '链接'
                    break
                default:
                    break;
            }
            return msg
        },
        createdFormat(row, column) {
            return Vue.filter('timeStampFormat')(row.created)
        },
        handleAddOrEdit(id = 0) {
            this.testNum = 2;
            let topath
            if (id == 0) {
                topath = '/discover/banner/add'
            } else {
                topath = '/discover/banner/edit/' + id
            }
            this.$router.push({
                path: topath
            })
        },
        handleLocationChange(status) {
            this.getPreBanner();
            this.getBannerList();
        },
        handlePagerChange(val) {
            this.pager.currentPage = val;
            this.queryForm.page = this.pager.currentPage
            this.queryForm.pagesize = this.pager.pageSize
            this.getBannerList();
        },
        getPreBanner() {
            optionsApi.getPreBanner(this.queryForm).then( (res) => {
                console.log('get Preview Banner', res);
                this.carousels = res.data.msg

            }).catch( (err) => {
                console.log('getPreBanner error', err);
                this.$message.error(err.data.msg)
            })
        },
        preview(url, row) {
            this.picRreview = {
                imgUrl: apiConst.ASSETS_API + url,
                visible: true,
                title: this.locationFormat({ location: row.location }) + '-' + row.description
            }
        },
    }
}
</script>

<style>
.el-carousel__item h3 {
    color: #475669;
    font-size: 18px;
    opacity: 0.75;
    line-height: 180px;
    margin: 0;
}

/*.el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
}*/

.discover-banner {
    width: 20%;
    /*height: 20%;*/
    /*height: 200px;*/
    margin: 20px auto;
}

time {
    font-size: 13px;
    color: #999;
}

.bottom {
    margin-top: 13px;
    line-height: 12px;
}

.button {
    padding: 0;
    float: right;
}

.image {
    width: 100%;
    display: block;
}

.clearfix:before,
.clearfix:after {
    display: table;
    content: "";
}

.clearfix:after {
    clear: both
}
</style>
