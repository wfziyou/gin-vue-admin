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
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
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
        <el-table-column align="left" label="标题" prop="title" width="120" />
        <el-table-column align="left" label="描述" prop="des" width="120" />
        <el-table-column align="left" label="内容" prop="contentHtml" width="120" />
        <el-table-column align="left" label="分类_编号" prop="classifyId" width="120" />
        <el-table-column align="left" label="圈子_编号" prop="circleId" width="120" />
        <el-table-column align="left" label="创建者编号" prop="userId" width="120" />
         <el-table-column align="left" label="开始时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.startAt) }}</template>
         </el-table-column>
         <el-table-column align="left" label="介绍时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.endAt) }}</template>
         </el-table-column>
        <el-table-column align="left" label="活动地址" prop="address" width="120" />
        <el-table-column align="left" label="地址所在区域" prop="addressZone" width="120" />
        <el-table-column align="left" label="地址经纬度" prop="addressPos" width="120" />
        <el-table-column align="left" label="用户人数" prop="userNum" width="120" />
        <el-table-column align="left" label="参加人数" prop="curUserNum" width="120" />
        <el-table-column align="left" label="是否公开：0 否 1 是" prop="public" width="120" />
        <el-table-column align="left" label="是否付费：0 否 1 是" prop="pay" width="120" />
        <el-table-column align="left" label="付费货币：1人民、2积分、3代币" prop="payCurrency" width="120" />
        <el-table-column align="left" label="付费金额" prop="payNum" width="120" />
        <el-table-column align="left" label="动态数" prop="dynamicNum" width="120" />
        <el-table-column align="left" label="收藏次数" prop="collectNum" width="120" />
        <el-table-column align="left" label="是否收藏：0否、1是" prop="collect" width="120" />
        <el-table-column align="left" label="聊天群编号" prop="chatGroupId" width="120" />
        <el-table-column align="left" label="审核状态：0 未处理 1 通过，2驳回" prop="checkStatus" width="120" />
        <el-table-column align="left" label="创建人" prop="createUser" width="120" />
        <el-table-column align="left" label="创建部门" prop="createDept" width="120" />
        <el-table-column align="left" label="修改人" prop="updateUser" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120" />
        <el-table-column align="left" label="是否已删除" prop="isDel" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateHkActivityFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
        <el-form-item label="标题:"  prop="title" >
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述:"  prop="des" >
          <el-input v-model="formData.des" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容:"  prop="contentHtml" >
          <el-input v-model="formData.contentHtml" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类_编号:"  prop="classifyId" >
          <el-input v-model.number="formData.classifyId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子_编号:"  prop="circleId" >
          <el-input v-model.number="formData.circleId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建者编号:"  prop="userId" >
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开始时间:"  prop="startAt" >
          <el-date-picker v-model="formData.startAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="介绍时间:"  prop="endAt" >
          <el-date-picker v-model="formData.endAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="活动地址:"  prop="address" >
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地址所在区域:"  prop="addressZone" >
          <el-input v-model="formData.addressZone" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地址经纬度:"  prop="addressPos" >
          <el-input v-model="formData.addressPos" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户人数:"  prop="userNum" >
          <el-input v-model.number="formData.userNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="参加人数:"  prop="curUserNum" >
          <el-input v-model.number="formData.curUserNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否公开：0 否 1 是:"  prop="public" >
          <el-input v-model.number="formData.public" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否付费：0 否 1 是:"  prop="pay" >
          <el-input v-model.number="formData.pay" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费货币：1人民、2积分、3代币:"  prop="payCurrency" >
          <el-input v-model.number="formData.payCurrency" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费金额:"  prop="payNum" >
          <el-input v-model.number="formData.payNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="动态数:"  prop="dynamicNum" >
          <el-input v-model.number="formData.dynamicNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收藏次数:"  prop="collectNum" >
          <el-input v-model.number="formData.collectNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否收藏：0否、1是:"  prop="collect" >
          <el-input v-model.number="formData.collect" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="聊天群编号:"  prop="chatGroupId" >
          <el-input v-model.number="formData.chatGroupId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核状态：0 未处理 1 通过，2驳回:"  prop="checkStatus" >
          <el-input v-model.number="formData.checkStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:"  prop="createUser" >
          <el-input v-model.number="formData.createUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建部门:"  prop="createDept" >
          <el-input v-model.number="formData.createDept" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="修改人:"  prop="updateUser" >
          <el-input v-model.number="formData.updateUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:"  prop="status" >
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否已删除:"  prop="isDel" >
          <el-input v-model.number="formData.isDel" :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'HkActivity'
}
</script>

<script setup>
import {
  createHkActivity,
  deleteHkActivity,
  deleteHkActivityByIds,
  updateHkActivity,
  findHkActivity,
  getHkActivityList
} from '@/api/hkActivity'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        tenantId: '',
        title: '',
        des: '',
        contentHtml: '',
        classifyId: 0,
        circleId: 0,
        userId: 0,
        startAt: new Date(),
        endAt: new Date(),
        address: '',
        addressZone: '',
        addressPos: '',
        userNum: 0,
        curUserNum: 0,
        public: 0,
        pay: 0,
        payCurrency: 0,
        payNum: 0,
        dynamicNum: 0,
        collectNum: 0,
        collect: 0,
        chatGroupId: 0,
        checkStatus: 0,
        createUser: 0,
        createDept: 0,
        updateUser: 0,
        status: 0,
        isDel: 0,
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
  const table = await getHkActivityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteHkActivityFunc(row)
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
      const res = await deleteHkActivityByIds({ ids })
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
const updateHkActivityFunc = async(row) => {
    const res = await findHkActivity({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rehkActivity
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteHkActivityFunc = async (row) => {
    const res = await deleteHkActivity({ ID: row.ID })
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
        title: '',
        des: '',
        contentHtml: '',
        classifyId: 0,
        circleId: 0,
        userId: 0,
        startAt: new Date(),
        endAt: new Date(),
        address: '',
        addressZone: '',
        addressPos: '',
        userNum: 0,
        curUserNum: 0,
        public: 0,
        pay: 0,
        payCurrency: 0,
        payNum: 0,
        dynamicNum: 0,
        collectNum: 0,
        collect: 0,
        chatGroupId: 0,
        checkStatus: 0,
        createUser: 0,
        createDept: 0,
        updateUser: 0,
        status: 0,
        isDel: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createHkActivity(formData.value)
                  break
                case 'update':
                  res = await updateHkActivity(formData.value)
                  break
                default:
                  res = await createHkActivity(formData.value)
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
