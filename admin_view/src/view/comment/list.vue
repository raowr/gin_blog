<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="姓名">
          <el-input v-model="searchInfo.name" placeholder="name" />
        </el-form-item>
        <el-form-item label="内容">
          <el-input v-model="searchInfo.message" placeholder="message" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clear placeholder="请选择">
            <el-option key="all" label="全部" value="" />
            <el-option key="true" label="是" value="true" />
            <el-option key="false" label="否" value="false" />
          </el-select>
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
            @click="onShow"
        >批量显示</el-button>
        <el-button
            size="small"
            type="warning"
            icon="minus"
            @click="onHide"
        >批量隐藏</el-button>
      </div>
      <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="id" width="60" prop="ID"  />
        <el-table-column align="left" label="姓名" prop="name" width="160" />
        <el-table-column align="left" label="联系" prop="phone" width="160" />
        <el-table-column align="left" label="内容" prop="message" width="280" />
        <el-table-column align="center" label="状态" prop="status" width="80" >
          <template #default="scope">{{
              formatBoolean(scope.row.status)
            }}</template>
        </el-table-column>
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
              formatDate(scope.row.CreatedAt)
            }}</template>
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
import { listComment, showCommentIdsByIds, hideCommentIdsByIds} from '@/api/comment'
import { formatBoolean, formatDate } from '@/utils/format'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({created_at:[],status:''})
const created_at_start = ref(0)
const created_at_end = ref(0)
const commentIds = ref([])


// 查询
const getTableData = async() => {
  if (searchInfo.value.created_at.length > 0){
    created_at_start.value = searchInfo.value.created_at[0]
    created_at_end.value = searchInfo.value.created_at[1]
  }
  if (searchInfo.value.status === null || searchInfo.value.status === "") {
    searchInfo.value.status = -1
  }else if(searchInfo.value.status === 'true') {
    searchInfo.value.status = 1
  }else if(searchInfo.value.status === 'false'){
    searchInfo.value.status = 0
  }
  const table = await listComment({
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
  if (searchInfo.value.status === -1) {
    searchInfo.value.status = ''
  }else if(searchInfo.value.status === 1) {
    searchInfo.value.status = 'true'
  }else if(searchInfo.value.status === 0){
    searchInfo.value.status = 'false'
  }
}

getTableData()

// 条件搜索前端看此方法
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }

  getTableData()
}

const onReset = () => {
  searchInfo.value = {created_at:[],status:''}
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 批量操作
const handleSelectionChange = (val) => {
  commentIds.value = val
}

const onShow = async() => {
  const ids = commentIds.value.map(item => item.ID)
  const res = await showCommentIdsByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: res.msg
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const onHide = async() => {
  const ids = commentIds.value.map(item => item.ID)
  const res = await hideCommentIdsByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: res.msg
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}



</script>

<style scoped>

</style>
