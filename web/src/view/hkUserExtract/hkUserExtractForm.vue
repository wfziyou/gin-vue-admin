<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户编号:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="名称:" prop="realName">
          <el-input v-model="formData.realName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="bank = 银行卡 alipay = 支付宝wx=微信:" prop="extractType">
          <el-input v-model="formData.extractType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行卡:" prop="bankCode">
          <el-input v-model="formData.bankCode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户地址:" prop="bankAddress">
          <el-input v-model="formData.bankAddress" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付宝账号:" prop="alipayCode">
          <el-input v-model="formData.alipayCode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="微信号:" prop="wechat">
          <el-input v-model="formData.wechat" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="提现金额:" prop="extractPrice">
          <el-input-number v-model="formData.extractPrice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="mark字段:" prop="mark">
          <el-input v-model="formData.mark" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="balance字段:" prop="balance">
          <el-input-number v-model="formData.balance" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="无效原因:" prop="failMsg">
          <el-input v-model="formData.failMsg" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="failTime字段:" prop="failTime">
          <el-date-picker v-model="formData.failTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
        <el-form-item label="状态:0审核中 1已提现 2未通过:" prop="status">
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
  name: 'HkUserExtract'
}
</script>

<script setup>
import {
  createHkUserExtract,
  updateHkUserExtract,
  findHkUserExtract
} from '@/api/hkUserExtract'

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
            realName: '',
            extractType: '',
            bankCode: '',
            bankAddress: '',
            alipayCode: '',
            wechat: '',
            extractPrice: 0,
            mark: '',
            balance: 0,
            failMsg: '',
            failTime: new Date(),
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
      const res = await findHkUserExtract({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehkUserExtract
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
               res = await createHkUserExtract(formData.value)
               break
             case 'update':
               res = await updateHkUserExtract(formData.value)
               break
             default:
               res = await createHkUserExtract(formData.value)
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
