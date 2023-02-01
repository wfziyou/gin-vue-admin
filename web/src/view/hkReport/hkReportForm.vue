<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="租户ID:" prop="tenantId">
          <el-input v-model="formData.tenantId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="被举报用户编号:" prop="reportUserId">
          <el-input v-model.number="formData.reportUserId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="举报用户编号:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子_编号:" prop="circleId">
          <el-input v-model.number="formData.circleId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="举报原因_编号:" prop="reasonId">
          <el-input v-model.number="formData.reasonId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题:" prop="reportType">
          <el-input v-model.number="formData.reportType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="举报内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容附件:" prop="contentAttachment">
          <el-input v-model="formData.contentAttachment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="处理状态：0 未处理、1已处理:" prop="curStatus">
          <el-input v-model.number="formData.curStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="操作用户_编号:" prop="handleUserId">
          <el-input v-model.number="formData.handleUserId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="处理方式:0无效处理（不予处理）、账号禁言:" prop="handleType">
          <el-input v-model.number="formData.handleType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="禁言时长_编号:" prop="durationId">
          <el-input v-model.number="formData.durationId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="积分处理:0不扣分、1扣分:" prop="handleScore">
          <el-input v-model.number="formData.handleScore" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="经验:" prop="scoreExperience">
          <el-input v-model.number="formData.scoreExperience" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="社区积分:" prop="scoreCommunity">
          <el-input v-model.number="formData.scoreCommunity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="购物积分:" prop="scoreBuy">
          <el-input v-model.number="formData.scoreBuy" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="下载值:" prop="scoreDownload">
          <el-input v-model.number="formData.scoreDownload" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否解除：0 否、是:" prop="unlock">
          <el-input v-model.number="formData.unlock" :clearable="true" placeholder="请输入" />
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
  name: 'HkReport'
}
</script>

<script setup>
import {
  createHkReport,
  updateHkReport,
  findHkReport
} from '@/api/hkReport'

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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHkReport({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehkReport
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
