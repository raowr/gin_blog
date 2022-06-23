<template>
<div>
  <div class="gva-search-box">
    <el-form :inline="true" :model="searchInfo">
      <el-form-item label="标题">
        <el-input v-model="searchInfo.title" placeholder="标题" />
      </el-form-item>
      <el-form-item label="创建时间">
      <div class="block">
        <el-date-picker
            v-model="searchInfo.created_at"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期">
        </el-date-picker>
      </div>
      </el-form-item>
      <el-form-item>
        <el-button
            size="small"
            type="primary"
            icon="search"
            @click="onSubmit"
        >查询</el-button>
        <el-button
            size="small"
            icon="refresh"
            @click="onReset"
        >重置</el-button>
      </el-form-item>
    </el-form>
  </div>

  <div class="gva-table-box">
    <div class="gva-btn-list">
      <el-button
          size="small"
          type="primary"
          icon="plus"
          @click="openCreat"
      >新增</el-button>
    </div>
    <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column
          align="left"
          label="标题"
          prop="title"
          width="160"
      />
      <el-table-column align="left" label="简述" prop="content" width="280" >
        <template #default="scope">{{
            formatText(scope.row.content)
          }}</template>
      </el-table-column>
      <el-table-column align="center" label="浏览" prop="browses" width="80" />
      <el-table-column align="center" label="评论" prop="comments" width="80" />
      <el-table-column align="left" label="日期" width="180">
        <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
      </el-table-column>

      <el-table-column align="left" label="操作">
        <template #default="scope">
          <el-button
              size="small"
              icon="document"
              type="primary" link
              @click="toEdit(scope.row)"
          >编辑</el-button>
          <el-popover
              v-model="scope.row.visible"
              placement="top"
              width="160"
          >
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px">
              <el-button
                  size="small"
                  type="primary" link
                  @click="scope.row.visible = false"
              >取消</el-button>
              <el-button
                  type="primary"
                  size="small"
                  @click="deleteSysArticleFunc(scope.row)"
              >确定</el-button>
            </div>
            <template #reference>
              <el-button
                  type="primary" link
                  icon="delete"
                  size="small"
                  style="margin-left: 10px"
                  @click="scope.row.visible = true"
              >删除</el-button>
            </template>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <div class="gva-pagination">
      <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
      />
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: 'List',
}
</script>

<script setup>
import { listArticle,deleteArticle } from '@/api/article'
import { formatBoolean, formatDate } from '@/utils/format'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
const router = useRouter()

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({created_at:[]})
const created_at_start = ref(0)
const created_at_end = ref(0)


// 条件搜索前端看此方法
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 查询
const getTableData = async() => {
  if (searchInfo.value.created_at.length > 0){
    created_at_start.value = searchInfo.value.created_at[0]
    created_at_end.value = searchInfo.value.created_at[1]
  }
  const table = await listArticle({
    page: page.value,
    pageSize: pageSize.value,
    created_at_start:searchInfo.value.created_at[0],
    created_at_end:searchInfo.value.created_at[1],
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const onReset = () => {
  searchInfo.value = {created_at:[]}
}

const openCreat = () =>{
  router.push({
    name: 'creatArticle',
  })
}

const toEdit = (row) =>{
  router.push({
    name: 'editArticle',
    params: {
      id: row.ID,
    },
  })
}

const deleteSysArticleFunc = async(row) => {
  row.visible = false
  const res = await deleteArticle({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const formatText =(text) =>{
  return text.slice(0,150)
}

</script>

<style scoped>

</style>
