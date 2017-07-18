<template>
    <div class="banner-form">
        <el-form :model="form" label-width="80px" :rules="formRules" ref="form">
            <el-row :span="24">
                <el-col :span="12">
                    <el-form-item label="描述" prop="description">
                        <el-input v-model="form.description" auto-complete="off"></el-input>
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row :span="24">
                <el-col :span="12">
                    <el-form-item label="位置" prop="location">
                        <el-radio-group v-model="form.location" @change="handleLocationChange">
                            <el-radio-button :label="1">首页</el-radio-button>
                            <el-radio-button :label="2">发现</el-radio-button>
                        </el-radio-group>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-row :span="24">
                <el-col :span="12">
                    <el-form-item label="跳转类别" prop="targetType">
                        <el-radio-group v-model="form.targetType" @change="handleTargetTypeChange">
                            <el-radio-button :label="1">新闻</el-radio-button>
                            <el-radio-button :label="2">活动</el-radio-button>
                            <el-radio-button :label="3">链接</el-radio-button>
                        </el-radio-group>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-row :span="24" v-if="this.form.targetType === 1">
                <el-col :span="12">
                    <el-form-item label="跳转新闻" prop="targetId">
                        <el-select v-model="form.targetId" filterable remote placeholder="输入关键词搜索，支持新闻名和作者" :remote-method="newsRemoteMethod" :loading="selectQuery.loading" style="width:100%">
                            <el-option v-for="item in selectQuery.options" :key="item.id" :label="item.title + ' -- ' + item.author" :value="item.id">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-row :span="24" v-if="this.form.targetType === 2">
                <el-col :span="12">
                    <el-form-item label="跳转活动" prop="targetId">
                        <el-select v-model="form.targetId" filterable remote placeholder="输入关键词搜索，支持活动名" :remote-method="newsRemoteMethod" :loading="selectQuery.loading" style="width:100%">
                            <el-option v-for="item in selectQuery.options" :key="item.id" :label="item.title + ' -- ' + item.author" :value="item.id">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-row :span="24" v-if="this.form.targetType === 3">
                <el-col :span="12">
                    <el-form-item label="跳转链接" prop="targetId">
                        <el-input v-model="form.targetLink" auto-complete="off"></el-input>
                    </el-form-item>
                </el-col>
            </el-row>
   
            <el-row :span="24">
                <el-col :span="5">
    
                    <el-tooltip class="item" effect="dark" content="点击可添加或修改" placement="bottom">
                        <el-form-item label="banner图" prop="url">
                            <el-upload class="avatar-uploader" :action="uploadUrl" :show-file-list="false" :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
                                <img v-if="form.url" :src="avatarUrl" class="avatar-banner">
                                <i v-else class="el-icon-plus avatar-uploader-plus-icon-banner" style="line-height: 178px;"></i>
                            </el-upload>
                        </el-form-item>
                    </el-tooltip>
                </el-col>
            </el-row>
    
            <el-row>
                <el-col :span="4">
                    <el-form-item label="当前状态" prop="isVisible">
                        <el-switch v-model="form.isVisible" on-color="#13ce66" off-color="#ff4949" on-text="打开" off-text="关闭">
                        </el-switch>
                    </el-form-item>
    
                </el-col>
            </el-row>
    
            <el-row :span="24">
                <el-col :span="12">
                    <el-form-item label="排序">
                        <el-input-number v-model="form.sort" @change="handleSortChange" :min="1" :max="100"></el-input-number>
                        <span>数字越大，排序越靠前</span>
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
export default {
    data() {
        return {
            bannerId: 0,
            form: {
                id: null,
                url: '',
                targetType: 1,
                targetLink: '',
                // location 1为首页，2为发现
                location: 1,
                description: '',
                targetId: null,
                isVisible: false,
                sort: 50,
            },
            selectQuery: {
                options: [],
                loading: false,
            },
            formRules: {
                title: [{ required: true, message: '标题不能为空', trigger: 'blur' }],
                location: [{ type: 'number', required: true, message: '位置不能为空', trigger: 'blur' }],
                targetType: [{ type: 'number', required: true, message: '跳转类别不能为空', trigger: 'blur' }],
                url: [{ required: true, message: 'Banner图不能为空', trigger: 'blur' }],
                description: [{ required: true, message: 'Banner描述不能为空', trigger: 'blur' }],
                content: [{ required: true, message: '正文内容不能为空', trigger: 'blur' }],
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
            if (this.bannerId) {
                return '修改'
            }else {
                return '添加'
            }
        }
    },
    mounted() {
        this.bannerId = this.$route.params.id;
        if (this.bannerId) {     //修改时加载的数据
            this.getBannerById();
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
        onSubmit() {
            console.log('onSubmit');
            this.$refs.form.validate((valid) => {
                if (valid) {
                    this.$confirm('确认提交吗?', '提示', { type: 'warning' }).then(() => {
                        console.log('this.form', this.form);
                        switch (this.form.targetType) {
                            // 新闻
                            case 1:
                                this.form.targetLink = '/options/discover/news/detail?id=' + this.form.targetId
                                break;
                            // 其他。。。
                            case 3:
                                this, form.targetId = 0
                                break
                            default:
                                break;
                        }
                        if (this.bannerId) {
                            // 更新Banner
                            this.putBanner()
                        } else {
                            // 添加Banner
                            this.postBanner()
                        }
                    }).catch(() => {
                        this.$message.info('取消提交')
                    });
                } else {
                    return
                }
            })

        },
        getBannerById() {
            const vm = this
            if (this.bannerId) {
                optionsApi.getBannerById({ id: this.bannerId }).then((res) => {
                    console.log('getBannerById', res);
                    if (res.data.msg !== undefined) {
                        this.form = res.data.msg
                        switch (res.data.msg.targetType) {
                            case 1:
                                this.getTargetNewsById(res.data.msg.targetId)
                                break;
                            default:
                                break;
                        }
                    }
                }).catch((err) => {

                })
            }
        },
        getTargetNewsById(id) {
            console.log('getTargetNewsById', id);
            discoverApi.getNewsById({ id: id }).then((newsRes) => {
                // this.form.
                console.log('getNewsById', newsRes);
                this.newsRemoteMethod(newsRes.data.msg.title)
                // this.form.targetId = newsRes.data.msg.title //+ '--' + newsRes.data.msg.author
            }).catch((err) => {
                this.$message.warning('跳转新闻获取错误')
            })
        },
        postBanner() {
            console.log('post Banner', this.form);
            optionsApi.addBanner(this.form).then((res) => {
                this.$message.success('创建成功')
                console.log('post news res', res);
                this.$router.push({ path: '/discover/banner' })
            }).catch((error) => {
                console.log('postBanner error', error);
            })
        },
        putBanner() {
            this.form.id = this.bannerId
            this.form.id = Number(this.form.id)
            console.log('put Banner', this.form);
            optionsApi.updateBanner(this.form).then((res) => {
                console.log('put Banner', res);
                this.$message.success('修改成功')
                console.log('put news res', res);
                this.$router.push({ path: '/discover/banner' })
            }).catch((err) => {
                console.log('put Banner', err);
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

.banner-form {
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