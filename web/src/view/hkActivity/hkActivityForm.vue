<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="租户ID:" prop="tenantId">
          <el-input v-model="formData.tenantId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述:" prop="des">
          <el-input v-model="formData.des" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容:" prop="contentHtml">
          <el-input v-model="formData.contentHtml" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类_编号:" prop="classifyId">
          <el-input v-model.number="formData.classifyId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子_编号:" prop="circleId">
          <el-input v-model.number="formData.circleId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建者编号:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开始时间:" prop="startAt">
          <el-date-picker v-model="formData.startAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="介绍时间:" prop="endAt">
          <el-date-picker v-model="formData.endAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="活动地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地址所在区域:" prop="addressZone">
          <el-input v-model="formData.addressZone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地址经纬度:" prop="addressPos">
          <el-input v-model="formData.addressPos" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户人数:" prop="userNum">
          <el-input v-model.number="formData.userNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="参加人数:" prop="curUserNum">
          <el-input v-model.number="formData.curUserNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否公开：0 否 1 是:" prop="public">
          <el-input v-model.number="formData.public" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否付费：0 否 1 是:" prop="pay">
          <el-input v-model.number="formData.pay" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费货币：1人民、2积分、3代币:" prop="payCurrency">
          <el-input v-model.number="formData.payCurrency" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费金额:" prop="payNum">
          <el-input v-model.number="formData.payNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="动态数:" prop="dynamicNum">
          <el-input v-model.number="formData.dynamicNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收藏次数:" prop="collectNum">
          <el-input v-model.number="formData.collectNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否收藏：0否、1是:" prop="collect">
          <el-input v-model.number="formData.collect" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="聊天群编号:" prop="chatGroupId">
          <el-input v-model.number="formData.chatGroupId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核状态：0 未处理 1 通过，2驳回:" prop="checkStatus">
          <el-input v-model.number="formData.checkStatus" :clearable="true" placeholder="请输入" />
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
  name: 'HkActivity'
}
</script>

<script setup>
import {
  createHkActivity,
  updateHkActivity,
  findHkActivity
} from '@/api/hkActivity'

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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHkActivity({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehkActivity
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
