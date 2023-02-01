<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="租户ID" prop="tenantId" width="120" />
        <el-table-column align="left" label="被举报用户编号" prop="reportUserId" width="120" />
        <el-table-column align="left" label="举报用户编号" prop="userId" width="120" />
        <el-table-column align="left" label="圈子_编号" prop="circleId" width="120" />
        <el-table-column align="left" label="举报原因_编号" prop="reasonId" width="120" />
        <el-table-column align="left" label="举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题" prop="reportType" width="120" />
        <el-table-column align="left" label="举报内容" prop="content" width="120" />
        <el-table-column align="left" label="内容附件" prop="contentAttachment" width="120" />
        <el-table-column align="left" label="处理状态：0 未处理、1已处理" prop="curStatus" width="120" />
        <el-table-column align="left" label="操作用户_编号" prop="handleUserId" width="120" />
        <el-table-column align="left" label="处理方式:0无效处理（不予处理）、账号禁言" prop="handleType" width="120" />
        <el-table-column align="left" label="禁言时长_编号" prop="durationId" width="120" />
        <el-table-column align="left" label="积分处理:0不扣分、1扣分" prop="handleScore" width="120" />
        <el-table-column align="left" label="经验" prop="scoreExperience" width="120" />
        <el-table-column align="left" label="社区积分" prop="scoreCommunity" width="120" />
        <el-table-column align="left" label="购物积分" prop="scoreBuy" width="120" />
        <el-table-column align="left" label="下载值" prop="scoreDownload" width="120" />
        <el-table-column align="left" label="是否解除：0 否、是" prop="unlock" width="120" />
        <el-table-column align="left" label="创建人" prop="createUser" width="120" />
        <el-table-column align="left" label="创建部门" prop="createDept" width="120" />
         <el-table-column align="left" label="创建时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.createTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="修改人" prop="updateUser" width="120" />
         <el-table-column align="left" label="修改时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.updateTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120" />
        <el-table-column align="left" label="是否已删除" prop="isDeleted" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateHkReportFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="租户ID:"  prop="tenantId" >
          <el-input v-model="formData.tenantId" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="被举报用户编号:"  prop="reportUserId" >
          <el-input v-model.number="formData.reportUserId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="举报用户编号:"  prop="userId" >
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子_编号:"  prop="circleId" >
          <el-input v-model.number="formData.circleId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="举报原因_编号:"  prop="reasonId" >
          <el-input v-model.number="formData.reasonId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题:"  prop="reportType" >
          <el-input v-model.number="formData.reportType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="举报内容:"  prop="content" >
          <el-input v-model="formData.content" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容附件:"  prop="contentAttachment" >
          <el-input v-model="formData.contentAttachment" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="处理状态：0 未处理、1已处理:"  prop="curStatus" >
          <el-input v-model.number="formData.curStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="操作用户_编号:"  prop="handleUserId" >
          <el-input v-model.number="formData.handleUserId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="处理方式:0无效处理（不予处理）、账号禁言:"  prop="handleType" >
          <el-input v-model.number="formData.handleType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="禁言时长_编号:"  prop="durationId" >
          <el-input v-model.number="formData.durationId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="积分处理:0不扣分、1扣分:"  prop="handleScore" >
          <el-input v-model.number="formData.handleScore" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="经验:"  prop="scoreExperience" >
          <el-input v-model.number="formData.scoreExperience" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="社区积分:"  prop="scoreCommunity" >
          <el-input v-model.number="formData.scoreCommunity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="购物积分:"  prop="scoreBuy" >
          <el-input v-model.number="formData.scoreBuy" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="下载值:"  prop="scoreDownload" >
          <el-input v-model.number="formData.scoreDownload" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否解除：0 否、是:"  prop="unlock" >
          <el-input v-model.number="formData.unlock" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:"  prop="createUser" >
          <el-input v-model.number="formData.createUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建部门:"  prop="createDept" >
          <el-input v-model.number="formData.createDept" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建时间:"  prop="createTime" >
          <el-date-picker v-model="formData.createTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="修改人:"  prop="updateUser" >
          <el-input v-model.number="formData.updateUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="修改时间:"  prop="updateTime" >
          <el-date-picker v-model="formData.updateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="状态:"  prop="status" >
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否已删除:"  prop="isDeleted" >
          <el-input v-model.number="formData.isDeleted" :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'HkReport'
}
</script>

<script setup>
import {
  createHkReport,
  deleteHkReport,
  deleteHkReportByIds,
  updateHkReport,
  findHkReport,
  getHkReportList
} from '@/api/hkReport'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        tenantId: '',
        reportUserId: 0,
        userId: 0,
        circleId: 0,
        reasonId: 0,
        reportType: 0,
        content: '',
        contentAttachment: '',
        curStatus: 0,
        handleUserId: 0,
        handleType: 0,
        durationId: 0,
        handleScore: 0,
        scoreExperience: 0,
        scoreCommunity: 0,
        scoreBuy: 0,
        scoreDownload: 0,
        unlock: 0,
        createUser: 0,
        createDept: 0,
        createTime: new Date(),
        updateUser: 0,
        updateTime: new Date(),
        status: 0,
        isDeleted: 0,
        })

// 验证规则
const rule = reactive({
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getHkReportList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteHkReportFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteHkReportByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateHkReportFunc = async(row) => {
    const res = await findHkReport({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rehkReport
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteHkReportFunc = async (row) => {
    const res = await deleteHkReport({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        tenantId: '',
        reportUserId: 0,
        userId: 0,
        circleId: 0,
        reasonId: 0,
        reportType: 0,
        content: '',
        contentAttachment: '',
        curStatus: 0,
        handleUserId: 0,
        handleType: 0,
        durationId: 0,
        handleScore: 0,
        scoreExperience: 0,
        scoreCommunity: 0,
        scoreBuy: 0,
        scoreDownload: 0,
        unlock: 0,
        createUser: 0,
        createDept: 0,
        createTime: new Date(),
        updateUser: 0,
        updateTime: new Date(),
        status: 0,
        isDeleted: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createHkReport(formData.value)
                  break
                case 'update':
                  res = await updateHkReport(formData.value)
                  break
                default:
                  res = await createHkReport(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
