<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户编号:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="关联id:" prop="linkId">
          <el-input v-model="formData.linkId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0 = 支出 1 = 获得:" prop="pm">
          <el-switch v-model="formData.pm" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="账单标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标题图标:" prop="titleIcon">
          <el-input v-model="formData.titleIcon" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="之前金额:" prop="beforeNumber">
          <el-input v-model.number="formData.beforeNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="改变金额:" prop="changeNumber">
          <el-input v-model.number="formData.changeNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="余额:" prop="balance">
          <el-input v-model.number="formData.balance" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:" prop="mark">
          <el-input v-model="formData.mark" :clearable="true" placeholder="请输入" />
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
        <el-form-item label="状态:0带确定 1有效 -1无效:" prop="status">
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
  name: 'HkGoldBill'
}
</script>

<script setup>
import {
  createHkGoldBill,
  updateHkGoldBill,
  findHkGoldBill
} from '@/api/hkGoldBill'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            userId: 0,
            linkId: '',
            pm: false,
            title: '',
            titleIcon: '',
            type: '',
            beforeNumber: 0,
            changeNumber: 0,
            balance: 0,
            mark: '',
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
      const res = await findHkGoldBill({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehkGoldBill
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
               res = await createHkGoldBill(formData.value)
               break
             case 'update':
               res = await updateHkGoldBill(formData.value)
               break
             default:
               res = await createHkGoldBill(formData.value)
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
