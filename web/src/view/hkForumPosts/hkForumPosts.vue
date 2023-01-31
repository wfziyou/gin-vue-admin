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
        <el-table-column align="left" label="圈子_编号" prop="circleId" width="120" />
        <el-table-column align="left" label="类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）" prop="category" width="120" />
        <el-table-column align="left" label="帖子分类编号" prop="groupId" width="120" />
        <el-table-column align="left" label="标题" prop="title" width="120" />
        <el-table-column align="left" label="SEO关键词" prop="seoKey" width="120" />
        <el-table-column align="left" label="简介(SEO简介)" prop="seoIntroduce" width="120" />
        <el-table-column align="left" label="封面" prop="coverImage" width="120" />
        <el-table-column align="left" label="来源" prop="source" width="120" />
         <el-table-column align="left" label="发布时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.time) }}</template>
         </el-table-column>
        <el-table-column align="left" label="内容" prop="content" width="120" />
        <el-table-column align="left" label="内容类型:0 markdown,1 html" prop="contentType" width="120" />
        <el-table-column align="left" label="markdown内容" prop="contentMarkdown" width="120" />
        <el-table-column align="left" label="html内容" prop="contentHtml" width="120" />
        <el-table-column align="left" label="视频地址" prop="video" width="120" />
        <el-table-column align="left" label="附件" prop="attachment" width="120" />
        <el-table-column align="left" label="标签" prop="tag" width="120" />
        <el-table-column align="left" label="置顶(0否 1是)" prop="top" width="120" />
        <el-table-column align="left" label="精华(0否 1是)" prop="marrow" width="120" />
        <el-table-column align="left" label="问答最佳答案ID" prop="commentId" width="120" />
        <el-table-column align="left" label="匿名发布(0否 1是)" prop="anonymity" width="120" />
        <el-table-column align="left" label="发布者编号" prop="userId" width="120" />
        <el-table-column align="left" label="阅读次数" prop="readNum" width="120" />
        <el-table-column align="left" label="评论次数" prop="commentNum" width="120" />
        <el-table-column align="left" label="分享次数" prop="shareNum" width="120" />
        <el-table-column align="left" label="收藏次数" prop="collectNum" width="120" />
        <el-table-column align="left" label="点赞次数" prop="likeNum" width="120" />
        <el-table-column align="left" label="审核人" prop="checkUser" width="120" />
         <el-table-column align="left" label="审核时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.checkTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="审核状态(0草稿 1未审批 2通过 3拒绝)" prop="checkStatus" width="120" />
        <el-table-column align="left" label="评论权限(0关闭 1开启)" prop="powerComment" width="120" />
        <el-table-column align="left" label="匿名评论(0关闭 1开启)" prop="powerCommentAnonymity" width="120" />
        <el-table-column align="left" label="付费：0 否，1是" prop="pay" width="120" />
        <el-table-column align="left" label="内容付费：0 否，1是" prop="payContent" width="120" />
        <el-table-column align="left" label="内容付费可查看百分比例" prop="payContentLook" width="120" />
        <el-table-column align="left" label="附件付费：0 否，1是" prop="payAttachment" width="120" />
        <el-table-column align="left" label="付费货币：0 人民、1积分、2代币" prop="payCurrency" width="120" />
        <el-table-column align="left" label="付费金额" prop="payNum" width="120" />
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
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateHkForumPostsFunc(scope.row)">变更</el-button>
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
        <el-form-item label="圈子_编号:"  prop="circleId" >
          <el-input v-model.number="formData.circleId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）:"  prop="category" >
          <el-input v-model.number="formData.category" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="帖子分类编号:"  prop="groupId" >
          <el-input v-model.number="formData.groupId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标题:"  prop="title" >
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="SEO关键词:"  prop="seoKey" >
          <el-input v-model="formData.seoKey" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简介(SEO简介):"  prop="seoIntroduce" >
          <el-input v-model="formData.seoIntroduce" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="封面:"  prop="coverImage" >
          <el-input v-model="formData.coverImage" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="来源:"  prop="source" >
          <el-input v-model="formData.source" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发布时间:"  prop="time" >
          <el-date-picker v-model="formData.time" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="内容:"  prop="content" >
          <el-input v-model="formData.content" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容类型:0 markdown,1 html:"  prop="contentType" >
          <el-input v-model.number="formData.contentType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="markdown内容:"  prop="contentMarkdown" >
          <el-input v-model="formData.contentMarkdown" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="html内容:"  prop="contentHtml" >
          <el-input v-model="formData.contentHtml" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="视频地址:"  prop="video" >
          <el-input v-model="formData.video" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="附件:"  prop="attachment" >
          <el-input v-model="formData.attachment" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标签:"  prop="tag" >
          <el-input v-model="formData.tag" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="置顶(0否 1是):"  prop="top" >
          <el-input v-model.number="formData.top" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="精华(0否 1是):"  prop="marrow" >
          <el-input v-model.number="formData.marrow" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问答最佳答案ID:"  prop="commentId" >
          <el-input v-model.number="formData.commentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="匿名发布(0否 1是):"  prop="anonymity" >
          <el-input v-model.number="formData.anonymity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发布者编号:"  prop="userId" >
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="阅读次数:"  prop="readNum" >
          <el-input v-model.number="formData.readNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="评论次数:"  prop="commentNum" >
          <el-input v-model.number="formData.commentNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分享次数:"  prop="shareNum" >
          <el-input v-model.number="formData.shareNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收藏次数:"  prop="collectNum" >
          <el-input v-model.number="formData.collectNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="点赞次数:"  prop="likeNum" >
          <el-input v-model.number="formData.likeNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核人:"  prop="checkUser" >
          <el-input v-model.number="formData.checkUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核时间:"  prop="checkTime" >
          <el-date-picker v-model="formData.checkTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="审核状态(0草稿 1未审批 2通过 3拒绝):"  prop="checkStatus" >
          <el-input v-model.number="formData.checkStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="评论权限(0关闭 1开启):"  prop="powerComment" >
          <el-input v-model.number="formData.powerComment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="匿名评论(0关闭 1开启):"  prop="powerCommentAnonymity" >
          <el-input v-model.number="formData.powerCommentAnonymity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费：0 否，1是:"  prop="pay" >
          <el-input v-model.number="formData.pay" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容付费：0 否，1是:"  prop="payContent" >
          <el-input v-model.number="formData.payContent" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容付费可查看百分比例:"  prop="payContentLook" >
          <el-input v-model.number="formData.payContentLook" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="附件付费：0 否，1是:"  prop="payAttachment" >
          <el-input v-model.number="formData.payAttachment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费货币：0 人民、1积分、2代币:"  prop="payCurrency" >
          <el-input v-model.number="formData.payCurrency" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费金额:"  prop="payNum" >
          <el-input v-model.number="formData.payNum" :clearable="true" placeholder="请输入" />
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
  name: 'HkForumPosts'
}
</script>

<script setup>
import {
  createHkForumPosts,
  deleteHkForumPosts,
  deleteHkForumPostsByIds,
  updateHkForumPosts,
  findHkForumPosts,
  getHkForumPostsList
} from '@/api/hkForumPosts'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        tenantId: '',
        circleId: 0,
        category: 0,
        groupId: 0,
        title: '',
        seoKey: '',
        seoIntroduce: '',
        coverImage: '',
        source: '',
        time: new Date(),
        content: '',
        contentType: 0,
        contentMarkdown: '',
        contentHtml: '',
        video: '',
        attachment: '',
        tag: '',
        top: 0,
        marrow: 0,
        commentId: 0,
        anonymity: 0,
        userId: 0,
        readNum: 0,
        commentNum: 0,
        shareNum: 0,
        collectNum: 0,
        likeNum: 0,
        checkUser: 0,
        checkTime: new Date(),
        checkStatus: 0,
        powerComment: 0,
        powerCommentAnonymity: 0,
        pay: 0,
        payContent: 0,
        payContentLook: 0,
        payAttachment: 0,
        payCurrency: 0,
        payNum: 0,
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
  const table = await getHkForumPostsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteHkForumPostsFunc(row)
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
      const res = await deleteHkForumPostsByIds({ ids })
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
const updateHkForumPostsFunc = async(row) => {
    const res = await findHkForumPosts({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rehkForumPosts
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteHkForumPostsFunc = async (row) => {
    const res = await deleteHkForumPosts({ ID: row.ID })
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
        circleId: 0,
        category: 0,
        groupId: 0,
        title: '',
        seoKey: '',
        seoIntroduce: '',
        coverImage: '',
        source: '',
        time: new Date(),
        content: '',
        contentType: 0,
        contentMarkdown: '',
        contentHtml: '',
        video: '',
        attachment: '',
        tag: '',
        top: 0,
        marrow: 0,
        commentId: 0,
        anonymity: 0,
        userId: 0,
        readNum: 0,
        commentNum: 0,
        shareNum: 0,
        collectNum: 0,
        likeNum: 0,
        checkUser: 0,
        checkTime: new Date(),
        checkStatus: 0,
        powerComment: 0,
        powerCommentAnonymity: 0,
        pay: 0,
        payContent: 0,
        payContentLook: 0,
        payAttachment: 0,
        payCurrency: 0,
        payNum: 0,
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
                  res = await createHkForumPosts(formData.value)
                  break
                case 'update':
                  res = await updateHkForumPosts(formData.value)
                  break
                default:
                  res = await createHkForumPosts(formData.value)
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
