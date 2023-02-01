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
        <el-table-column align="left" label="类型：0官方圈子、1用户圈子、2小区" prop="type" width="120" />
        <el-table-column align="left" label="圈子名称" prop="name" width="120" />
        <el-table-column align="left" label="圈子Logo" prop="logo" width="120" />
        <el-table-column align="left" label="圈子分类_编号" prop="circleClassifyId" width="120" />
        <el-table-column align="left" label="圈子标语" prop="slogan" width="120" />
        <el-table-column align="left" label="圈子简介" prop="des" width="120" />
        <el-table-column align="left" label="圈子规约" prop="protocol" width="120" />
        <el-table-column align="left" label="圈子背景图" prop="backImage" width="120" />
        <el-table-column align="left" label="支持内容类别(json数组)：0视频、1动态、2资讯、3公告、4文章、5问答、6建议" prop="supportCategory" width="120" />
        <el-table-column align="left" label="付费：0 否，1是" prop="pay" width="120" />
        <el-table-column align="left" label="是否开启版块内容人工审核：0 否，1是" prop="process" width="120" />
        <el-table-column align="left" label="圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）" prop="property" width="120" />
        <el-table-column align="left" label="板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到" prop="view" width="120" />
        <el-table-column align="left" label="圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户" prop="powerAdd" width="120" />
        <el-table-column align="left" label="圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组" prop="powerView" width="120" />
        <el-table-column align="left" label="圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组" prop="powerPublish" width="120" />
        <el-table-column align="left" label="圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组" prop="powerComment" width="120" />
        <el-table-column align="left" label="圈子加入权限_指定部门和成员(json数组)" prop="powerAddUser" width="120" />
        <el-table-column align="left" label="圈子内浏览权限_指定部门和用户(json数组)" prop="powerViewUser" width="120" />
        <el-table-column align="left" label="圈子内发布权限_指定部门和用户(json数组)" prop="powerPublishUser" width="120" />
        <el-table-column align="left" label="圈子内评论权限_指定部门和用户(json数组)" prop="powerCommentUser" width="120" />
        <el-table-column align="left" label="不受限用户组(json数组)" prop="noLimitUserGroup" width="120" />
        <el-table-column align="left" label="新注册用户默认关注：0 否，1是" prop="newUserFocus" width="120" />
        <el-table-column align="left" label="排序" prop="sort" width="120" />
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
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateHkCircleFunc(scope.row)">变更</el-button>
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
        <el-form-item label="类型：0官方圈子、1用户圈子、2小区:"  prop="type" >
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子名称:"  prop="name" >
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子Logo:"  prop="logo" >
          <el-input v-model="formData.logo" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子分类_编号:"  prop="circleClassifyId" >
          <el-input v-model.number="formData.circleClassifyId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子标语:"  prop="slogan" >
          <el-input v-model="formData.slogan" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子简介:"  prop="des" >
          <el-input v-model="formData.des" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子规约:"  prop="protocol" >
          <el-input v-model="formData.protocol" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子背景图:"  prop="backImage" >
          <el-input v-model="formData.backImage" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支持内容类别(json数组)：0视频、1动态、2资讯、3公告、4文章、5问答、6建议:"  prop="supportCategory" >
          <el-input v-model="formData.supportCategory" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费：0 否，1是:"  prop="pay" >
          <el-input v-model.number="formData.pay" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否开启版块内容人工审核：0 否，1是:"  prop="process" >
          <el-input v-model.number="formData.process" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）:"  prop="property" >
          <el-input v-model.number="formData.property" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到:"  prop="view" >
          <el-input v-model.number="formData.view" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户:"  prop="powerAdd" >
          <el-input v-model.number="formData.powerAdd" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组:"  prop="powerView" >
          <el-input v-model.number="formData.powerView" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组:"  prop="powerPublish" >
          <el-input v-model.number="formData.powerPublish" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组:"  prop="powerComment" >
          <el-input v-model.number="formData.powerComment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子加入权限_指定部门和成员(json数组):"  prop="powerAddUser" >
          <el-input v-model="formData.powerAddUser" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内浏览权限_指定部门和用户(json数组):"  prop="powerViewUser" >
          <el-input v-model="formData.powerViewUser" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内发布权限_指定部门和用户(json数组):"  prop="powerPublishUser" >
          <el-input v-model="formData.powerPublishUser" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内评论权限_指定部门和用户(json数组):"  prop="powerCommentUser" >
          <el-input v-model="formData.powerCommentUser" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="不受限用户组(json数组):"  prop="noLimitUserGroup" >
          <el-input v-model="formData.noLimitUserGroup" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="新注册用户默认关注：0 否，1是:"  prop="newUserFocus" >
          <el-input v-model.number="formData.newUserFocus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="排序:"  prop="sort" >
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入" />
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
  name: 'HkCircle'
}
</script>

<script setup>
import {
  createHkCircle,
  deleteHkCircle,
  deleteHkCircleByIds,
  updateHkCircle,
  findHkCircle,
  getHkCircleList
} from '@/api/hkCircle'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        tenantId: '',
        type: 0,
        name: '',
        logo: '',
        circleClassifyId: 0,
        slogan: '',
        des: '',
        protocol: '',
        backImage: '',
        supportCategory: '',
        pay: 0,
        process: 0,
        property: 0,
        view: 0,
        powerAdd: 0,
        powerView: 0,
        powerPublish: 0,
        powerComment: 0,
        powerAddUser: '',
        powerViewUser: '',
        powerPublishUser: '',
        powerCommentUser: '',
        noLimitUserGroup: '',
        newUserFocus: 0,
        sort: 0,
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
  const table = await getHkCircleList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteHkCircleFunc(row)
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
      const res = await deleteHkCircleByIds({ ids })
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
const updateHkCircleFunc = async(row) => {
    const res = await findHkCircle({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rehkCircle
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteHkCircleFunc = async (row) => {
    const res = await deleteHkCircle({ ID: row.ID })
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
        type: 0,
        name: '',
        logo: '',
        circleClassifyId: 0,
        slogan: '',
        des: '',
        protocol: '',
        backImage: '',
        supportCategory: '',
        pay: 0,
        process: 0,
        property: 0,
        view: 0,
        powerAdd: 0,
        powerView: 0,
        powerPublish: 0,
        powerComment: 0,
        powerAddUser: '',
        powerViewUser: '',
        powerPublishUser: '',
        powerCommentUser: '',
        noLimitUserGroup: '',
        newUserFocus: 0,
        sort: 0,
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
                  res = await createHkCircle(formData.value)
                  break
                case 'update':
                  res = await updateHkCircle(formData.value)
                  break
                default:
                  res = await createHkCircle(formData.value)
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
