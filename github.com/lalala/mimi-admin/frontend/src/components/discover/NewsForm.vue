<template>
    <div class="news-form">
        <el-form :model="form" label-width="80px" :rules="formRules" ref="form">
            <el-row :span="24">
                <el-col :span="8">
                    <el-form-item label="标题" prop="title">
                        <el-input v-model="form.title" auto-complete="off"></el-input>
                    </el-form-item>
                </el-col>
                <el-col :span="4">
                    <el-form-item label="作者" prop="author">
                        <el-input v-model="form.author"></el-input>
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row :span="24">
                <el-col :span="12">
                    <el-form-item label="封面描述" prop="description">
                        <el-input v-model="form.description" auto-complete="off"></el-input>
                    </el-form-item>
                </el-col>
            </el-row>
            
            <el-row :span="24">
                <el-col :span="5">
                    
                    <el-tooltip class="item" effect="dark" content="点击可添加或修改" placement="bottom">
                        <el-form-item label="新闻封面">
                            <el-upload class="avatar-uploader" :action="uploadUrl" :show-file-list="false" :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
                                <img v-if="form.cover" :src="avatarUrl" class="avatar">
                                <i v-else class="el-icon-plus avatar-uploader-plus-icon" style="line-height: 178px;"></i>
                            </el-upload>
                        </el-form-item>
                    </el-tooltip>
                </el-col>
                
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
                <div>
                    <el-col :span="18">
                        <el-form-item label="正文内容" prop="content">
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
    import newsApi from '../../api/discover';
    import apiConst from '../../api/constant';
    import Handlebars from 'handlebars/dist/handlebars.min.js';
    import moment from 'moment';
    export default {
        components: {
            markdownEditor,
        },
        data() {
            return {
                newsId: 0,
                form: {
                    title: '',
                    author: '',
                    description: '',
                    content: '',
                    html: '',
                    cover: '',
                    createdTime: '',
                    innerHtml: '',
                    pvStr: '{{.PageViews}}',
                    isVisible: false
                },
                formRules: {
                    title: [{ required: true, message: '标题不能为空', trigger: 'blur' }],
                    author: [{ required: true, message: '作者不能为空', trigger: 'blur' }],
                    description: [{ required: true, message: '封面描述不能为空', trigger: 'blur' }],
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
            this.newsId = this.$route.params.id;
        if (this.newsId) {     //修改时加载的数据
            this.getNewsById()
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
            if (this.newsId) {
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
            let source = `<!DOCTYPE html>
            <html>
            <head>
             <title>{{title}}</title>
             <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
         </head>
         <body>
             <div class="article">
              <div class="article-header">
               <div>
                <h1 class="article-title">{{ title }}</h1>
            </div>
            <div>
                <p class="article-author">{{ createdTime }} &nbsp{{author}}</p>
            </div>
        </div>
        <div class="article-content">
          {{{ innerHtml }}}
      </div>
      <div class="article-footer">
       <p class="article-pageviews">{{{ pvStr }}}人浏览 &nbsp&nbsp&nbsp{{ createdTime }}</p>
   </div>
</div>
</body>
<style type="text/css">
 .article {
  margin: 12px 5px;
  font-family: "MicrosoftYaHei";
}	
.article-title {
  font-size: 1.6em;
  font-weight: normal;
  margin-bottom: 0;
}	
.article-author {
  color: #BBBBBB;
}
.article-pageviews {
  color: #BBBBBB;
  margin-top: 10%;
}
.article img {
  margin: 3% auto;
  width: 100%;
}
.article-content p {
    margin: 5% auto;
    line-height:160%;
}
.article-footer {
    margin-bottom: 5%
}
</style>
</html>`
let template = Handlebars.compile(source);
this.form.innerHtml = this.mdCompiler(this.form.content)
if (this.newsId) {
    console.log('put this.form', this.form);
    this.form.createdTime = moment.unix(this.form.created).format("YYYY-MM-DD")
    this.form.pvStr = '{{.PageViews}}'
    let result = template(this.form);
    console.log('result', result);
    this.form.html = result
    this.putNews()
} else {
    this.form.createdTime = moment().format("YYYY-MM-DD")
    console.log('this.form.Pvstt', this.form.pvStr);
    let result = template(this.form);
    console.log('result', result);
    this.form.html = result
    this.postNews();
}

},
setEmpty() {
    this.form.title = this.form.author = this.form.description = this.form.content = ""
},
imageuploaded(res) {
    if (res.errcode == 0) {
        this.src = res.data.src;
    }
},
mdCompiler(src) {
    var marked = require('marked');
    marked.setOptions({
        renderer: new marked.Renderer(),
        gfm: true,
        tables: true,
        breaks: false,
        pedantic: false,
        sanitize: false,
        smartLists: true,
        smartypants: false
    });
    console.log('mdCompiler: ', marked(this.form.content));
    return marked(this.form.content)
},
postNews() {
    console.log('post News', this.form);
    newsApi.addNews(this.form).then((res) => {
        this.$message.success('创建成功')
        console.log('post news res', res);
        this.$router.push({ path: '/news' })
    }).catch((error) => {
        console.log('postNews error', error);
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
getNewsById() {
    const vm = this
    if (this.newsId) {
        newsApi.getNewsById({ id: this.newsId }).then((res) => {
            console.log('getNewsById', res);
            if (res.data.msg !== undefined) {
                this.form = res.data.msg

                        // this.avatarUrl = apiConst.ASSETS_API + res.data.msg.url
                    }
                }).catch((err) => {

                })
            }
        },
        putNews() {
            this.form.id = this.newsId
            this.form.id = Number(this.form.id)
            console.log('put News', this.form);
            newsApi.updateNews(this.form).then((res) => {
                console.log('put News', res);
                this.$message.success('修改成功')
                console.log('put news res', res);
                this.$router.push({ path: '/news' })
            }).catch((err) => {
                console.log('put Banner', err);
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

    .news-form {
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