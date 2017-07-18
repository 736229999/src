<template>
    <div class="faq-form">
        <el-form :model="form" label-width="80px" :rules="formRules" ref="form">
            <el-row :span="24">
                <el-col :span="12">
                    <el-form-item label="问题标题" prop="title">
                        <el-input v-model="form.title" auto-complete="off"></el-input>
                    </el-form-item>
                </el-col>
            </el-row>
    
            <el-row :span="24">
                <el-col :span="4">
                    <el-form-item label="内容图片" prop="desc">
                        <el-upload class="upload-demo" :action="uploadUrl" :on-preview="handlePreview" :on-success="handleSuccess" :on-remove="handleRemove" :file-list="fileList">
                            <el-button size="small" type="primary">点击上传</el-button>
                            <!--<div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>-->
                        </el-upload>
                    </el-form-item>
                </el-col>
                <el-col :span="4">
                    <el-form-item>
                        <el-button type="primary" @click="dialogTableVisible = true" size="small">查看上传列表</el-button>
                    </el-form-item>
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
    
            <el-row :span="24">
                <div>
                    <el-col :span="18">
                        <el-form-item label="问题解答" prop="content">
                            <markdown-editor v-model="form.content" :configs="configs" ref="markdownEditor" @input="handleInput"></markdown-editor>
                        </el-form-item>
                    </el-col>
                </div>
            </el-row>
            <el-form-item>
                <el-button type="primary" @click="onSubmit()">立即{{ buttonText }}</el-button>
                <el-button @click="setEmpty()">清空</el-button>
                <el-button @click="goBack">返回</el-button>
            </el-form-item>
        </el-form>
    
        <el-dialog title="上传图片列表" :visible.sync="dialogTableVisible">
            <el-table :data="fileList" :border="true">
                <el-table-column property="name" label="文件名"></el-table-column>
                <el-table-column property="url" label="图片地址"></el-table-column>
                <el-table-column property="uploadTime" label="上传时间"></el-table-column>
            </el-table>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="dialogTableVisible = false">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>
<script>
import { markdownEditor } from 'vue-simplemde'
import optionsApi from '../../../api/options';
import apiConst from '../../../api/constant';
import moment from 'moment';
export default {
    components: {
        markdownEditor,
    },
    data() {
        return {
            faqId: 0,
            form: {
                title: '',
                content: '',
                html: '',
                isVisible: false,
                sort: 50,
            },
            formRules: {
                title: [{ required: true, message: '标题不能为空', trigger: 'blur' }],
                content: [{ required: true, message: '正文内容不能为空', trigger: 'blur' }],
            },

            configs: {
                status: true,
                initialValue: '',
                renderingConfig: {
                    codeSyntaxHighlighting: false,
                    highlightingTheme: 'atom-one-light'
                },
                spellChecker: false,
                toolbar: ['bold', 'italic', '|', 'heading', 'heading-smaller', 'heading-bigger', '|', 'unordered-list', 'ordered-list', 'link', 'image', 'table', 'preview', 'side-by-side', 'fullscreen', '|', 'quote', '|', 'guide',
                    {
                        name: 'custom',
                        action: this.imgCenter,
                        className: 'fa fa-file-image-o',
                        title: '居中图片'
                    }],
            },
            fileList: [],
            src: 'http://img1.vued.vanthink.cn/vued0a233185b6027244f9d43e653227439a.png',
            imageUrl: '',
            dialogTableVisible: false,
        };
    },
    mounted() {
        this.faqId = this.$route.params.id;
        if (this.faqId) {     //修改时加载的数据
            this.getFaqById()
        }
    },
    computed: {
        simplemde() {
            return this.$refs.markdownEditor.simplemde
        },
        uploadUrl() {
            return apiConst.ASSETS_API + '/assets/backend/upload/news'
        },
        avatarUrl() {
            return apiConst.ASSETS_API + this.form.cover
        },
        buttonText() {
            if (this.faqId) {
                return '修改'
            } else {
                return '添加'
            }
        }
    },
    methods: {
        goBack() {
            this.$router.go(-1)
        },
        imgCenter(editor) {
            this.form.content += `<div align="center">
                ![]()
</div>\n`
            console.log('editor', editor);
        },
        handleRemove(file, fileList) {
            console.log(file, fileList);
        },
        handlePreview(file) {
            console.log(file);
            console.log('fileList:', this.fileList);
        },
        handleSuccess(res, file) {
            console.log('handleSuccess');
            console.log('res:', res);
            console.log('file:', file);
            var moment = require('moment');

            this.fileList.push({
                name: file.name, url: apiConst.ASSETS_API + res.result,
                uploadTime: moment().format('YYYY-MM-DD HH:mm:ss')
            })
        },
        handleInput() {
            console.log('handleInput')
            // console.log('content ', this.form.content);
            // console.log('html', this.mdCompiler(this.form.content));
        },
        handleSortChange(val) {
            
        },
        onSubmit() {
            console.log('onSubmit');
            console.log('parse html', this.simplemde.markdown(this.form.content));
            this.$refs.form.validate((valid) => {
                if (valid) {
                    this.$confirm('确认提交吗?', '提示', { type: 'warning' }).then(() => {
                        this.pieceHtml()
                    }).catch(() => {
                        this.$message.info('取消提交')
                    });
                } else {

                }
            })

        },
        pieceHtml() {
            this.form.html = this.simplemde.markdown(this.form.content)
            console.log('submit this.form', this.form);
            if (this.faqId) {
                this.putFaq()
            } else {
                this.postFaq();
            }

        },
        setEmpty() {
            this.form.title = this.form.content = ""
        },
        imageuploaded(res) {
            if (res.errcode == 0) {
                this.src = res.data.src;
            }
        },
        postFaq() {
            console.log('post Faq', this.form);
            optionsApi.addFaq(this.form).then((res) => {
                this.$message.success('创建成功')
                console.log('post Faq res', res);
                this.$router.push({ path: '/options/faq/list' })
            }).catch((error) => {
                console.log('post Faq error', error);
            })
        },
        handleAvatarSuccess(res, file) {
            console.log('handleAvatarSuccess');
            console.log('res:', res);
            console.log('res.result:', res.result);
            this.form.cover = res.result
            console.log('file:', file);
        },
        beforeAvatarUpload() {
            console.log('beforeAvatarUpload');
        },
        getFaqById() {
            console.log('getFaqById');
            const vm = this
            if (this.faqId) {
                optionsApi.getFaqById({ id: this.faqId }).then((res) => {
                    console.log('getFaqById', res);
                    if (res.data.msg !== undefined) {
                        this.form = res.data.msg
                    }
                }).catch((err) => {
                    this.$message.error(err.data.msg)
                })
            }
        },
        putFaq() {
            this.form.id = Number(this.faqId)
            // this.form.id = Number(this.form.id)
            console.log('put Faq', this.form);
            optionsApi.updateFaq(this.form).then((res) => {
                console.log('put Faq', res);
                this.$message.success('修改成功')
                console.log('put Faq res', res);
                this.$router.push({ path: '/options/faq/list' })
            }).catch((err) => {
                console.log('put Faq', err);
                this.$message.error(err.data.msg)
            })
        },
    },
}
</script>

<style>
h1 {
    display: block;
    font-size: 2em;
    -webkit-margin-before: 0.67em;
    -webkit-margin-after: 0.67em;
    -webkit-margin-start: 0px;
    -webkit-margin-end: 0px;
    font-weight: bold;
}

h2 {
    display: block;
    font-size: 1.5em;
    -webkit-margin-before: 0.83em;
    -webkit-margin-after: 0.83em;
    -webkit-margin-start: 0px;
    -webkit-margin-end: 0px;
    font-weight: bold;
}

h3 {
    display: block;
    font-size: 1.17em;
    -webkit-margin-before: 1em;
    -webkit-margin-after: 1em;
    -webkit-margin-start: 0px;
    -webkit-margin-end: 0px;
    font-weight: bold;
}

h4 {
    display: block;
    -webkit-margin-before: 1.33em;
    -webkit-margin-after: 1.33em;
    -webkit-margin-start: 0px;
    -webkit-margin-end: 0px;
    font-weight: bold;
}

h5 {
    display: block;
    font-size: 0.83em;
    -webkit-margin-before: 1.67em;
    -webkit-margin-after: 1.67em;
    -webkit-margin-start: 0px;
    -webkit-margin-end: 0px;
    font-weight: bold;
}

h6 {
    display: block;
    font-size: 0.67em;
    -webkit-margin-before: 2.33em;
    -webkit-margin-after: 2.33em;
    -webkit-margin-start: 0px;
    -webkit-margin-end: 0px;
    font-weight: bold;
}


.md-editor {
    margin: 20px;
}

.faq-form {
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

.avatar-uploader-plus-icon {
    font-size: 28px;
    color: #8c939d;
    width: 231px;
    height: 178px;
    text-align: center;
}

.avatar {
    width: 231px;
    height: 178px;
    display: block;
}
</style>