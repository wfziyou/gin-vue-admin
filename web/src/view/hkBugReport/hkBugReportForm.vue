<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="租户ID:" prop="tenantId">
          <el-input v-model="formData.tenantId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户编号:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="操作步骤:" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="操作步骤附件:" prop="contentAttachment">
          <el-input v-model="formData.contentAttachment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="预期结果:" prop="expectedResult">
          <el-input v-model="formData.expectedResult" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="实际结果:" prop="actualResult">
          <el-input v-model="formData.actualResult" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="实际结果附件:" prop="actualResultAttachment">
          <el-input v-model="formData.actualResultAttachment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他信息:" prop="otherInfo">
          <el-input v-model="formData.otherInfo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:" prop="createUser">
          <el-input v-model.number="formData.createUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建部门:" prop="createDept">
          <el-input v-model.number="formData.createDept" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建时间:" prop="createTime">
          <el-date-picker v-model="formData.createTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="修改人:" prop="updateUser">
          <el-input v-model.number="formData.updateUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="修改时间:" prop="updateTime">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否已删除:" prop="isDeleted">
          <el-input v-model.number="formData.isDeleted" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HkBugReport'
}
</script>

<script setup>
import {
  createHkBugReport,
  updateHkBugReport,
  findHkBugReport
} from '@/api/hkBugReport'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            tenantId: '',
            userId: 0,
            title: '',
            content: '',
            contentAttachment: '',
            expectedResult: '',
            actualResult: '',
            actualResultAttachment: '',
            otherInfo: '',
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHkBugReport({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehkBugReport
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createHkBugReport(formData.value)
               break
             case 'update':
               res = await updateHkBugReport(formData.value)
               break
             default:
               res = await createHkBugReport(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
