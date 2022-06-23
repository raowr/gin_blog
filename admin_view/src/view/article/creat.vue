<template>
  <div>
    <el-form ref="article" :model="articleData" :rules="rules" size="medium" label-width="100px">
      <el-form-item label="标题" prop="title">
        <el-input v-model="articleData.title" placeholder="请输入标题" clearable :style="{width: '40%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="封面" prop="cover">
        <div style="display:inline-block" @click="openHeaderChange">
          <img v-if="articleData.cover" class="header-img-box" :src="(articleData.cover && articleData.cover.slice(0, 4) !== 'http')?path+articleData.cover:articleData.cover">
          <div v-else class="header-img-box">从媒体库选择</div>
        </div>
      </el-form-item>
      <el-form-item label="内容" prop="content">
        <editor :id="tinymceId" :init="init" v-model="articleData.content" placeholder="请输入内容" :height="300" :style="{width: '40%'}"></editor>
      </el-form-item>
      <el-form-item size="large">
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>
    <ChooseImg ref="chooseImg" :target="articleData" :target-key="`cover`" />
  </div>
</template>

<script setup>
import { infoArticle } from '@/api/article'
import ChooseImg from '@/components/chooseImg/index.vue'
import { ref, watch } from 'vue'
const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}
const path = ref(import.meta.env.VITE_BASE_API + '/')

</script>

<script>
import tinymce from 'tinymce/tinymce'
import Editor from '@tinymce/tinymce-vue'
import 'tinymce/themes/silver'
import 'tinymce/icons/default/icons'
import 'tinymce/models/dom'
import { addArticle } from '@/api/article.js'
import { ElMessage } from 'element-plus'

export default {
  name: 'CreatArticle',
  components: {
    Editor,
  },
  props: [],
  data() {
    return {
      articleData: {
        title: '',
        cover: '',
        content: '',
      },
      rules: {
        title: [{
          required: true,
          message: '请输入标题',
          trigger: 'blur'
        }],
        content: [{
          required: true,
          message: '请输入内容',
          trigger: 'blur'
        }],
        cover: [{
          required: true,
          message: '请选择封面',
          trigger: 'blur'
        }],
      },
      tinymceId: this.id,
      init: {
        language_url: '/tinymce/langs/zh-Hans.js',
        language: 'zh_CN',
        skin_url: '/tinymce/skins/ui/oxide',
        content_css: '/tinymce/skins/content/default/content.css',
      }
    }
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {
    tinymce.init({})
  },
  methods: {
    submitForm() {
      this.$refs['article'].validate(async valid => {
        if (!valid) return
        // TODO 提交表单
        const res = await addArticle(this.articleData)
        if (res.code === 0 ){
          ElMessage({ type: 'success', message: '添加成功' })
        }
      })
    },
    resetForm() {
      this.$refs['article'].resetFields()
    },
    coverBeforeUpload(file) {
      let isRightSize = file.size / 1024 / 1024 < 2
      if (!isRightSize) {
        this.$message.error('文件大小超过 2MB')
      }
      return isRightSize
    }
  }
}

</script>
<style>
.el-upload__tip {
  line-height: 1.2;
}
.header-img-box {
  width: 250px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}

</style>
