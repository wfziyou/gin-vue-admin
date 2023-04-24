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
        <el-form-item label="类型编号:" prop="typeId">
          <el-input v-model.number="formData.typeId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述:" prop="des">
          <el-input v-model="formData.des" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="附件:" prop="attachment">
          <el-input v-model="formData.attachment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="处理状态：0 未处理、1 处理中、2 拒绝、3 完成:" prop="checkStatus">
          <el-input v-model.number="formData.checkStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="处理描述:" prop="process">
          <el-input v-model="formData.process" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:" prop="createUser">
          <el-input v-model.number="formData.createUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建部门:" prop="createDept">
          <el-input v-model.number="formData.createDept" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="修改人:" prop="updateUser">
          <el-input v-model.number="formData.updateUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否已删除:" prop="isDel">
          <el-input v-model.number="formData.isDel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HkFeedback'
}
</script>

<script setup>
import {
  createHkFeedback,
  updateHkFeedback,
  findHkFeedback
} from '@/api/hkFeedback'

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
            typeId: 0,
            type: '',
            des: '',
            attachment: '',
            checkStatus: 0,
            process: '',
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHkFeedback({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehkFeedback
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
               res = await createHkFeedback(formData.value)
               break
             case 'update':
               res = await updateHkFeedback(formData.value)
               break
             default:
               res = await createHkFeedback(formData.value)
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
