<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="租户ID:" prop="tenantId">
          <el-input v-model="formData.tenantId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型：0官方圈子、1用户圈子、2小区:" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子Logo:" prop="logo">
          <el-input v-model="formData.logo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子分类_编号:" prop="circleClassifyId">
          <el-input v-model.number="formData.circleClassifyId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子标语:" prop="slogan">
          <el-input v-model="formData.slogan" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子简介:" prop="des">
          <el-input v-model="formData.des" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子规约:" prop="protocol">
          <el-input v-model="formData.protocol" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子背景图:" prop="backImage">
          <el-input v-model="formData.backImage" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支持内容类别(json数组)：0视频、1动态、2资讯、3公告、4文章、5问答、6建议:" prop="supportCategory">
          <el-input v-model="formData.supportCategory" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费：0 否，1是:" prop="pay">
          <el-input v-model.number="formData.pay" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否开启版块内容人工审核：0 否，1是:" prop="process">
          <el-input v-model.number="formData.process" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）:" prop="property">
          <el-input v-model.number="formData.property" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到:" prop="view">
          <el-input v-model.number="formData.view" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户:" prop="powerAdd">
          <el-input v-model.number="formData.powerAdd" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组:" prop="powerView">
          <el-input v-model.number="formData.powerView" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组:" prop="powerPublish">
          <el-input v-model.number="formData.powerPublish" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组:" prop="powerComment">
          <el-input v-model.number="formData.powerComment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子加入权限_指定部门和成员(json数组):" prop="powerAddUser">
          <el-input v-model="formData.powerAddUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内浏览权限_指定部门和用户(json数组):" prop="powerViewUser">
          <el-input v-model="formData.powerViewUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内发布权限_指定部门和用户(json数组):" prop="powerPublishUser">
          <el-input v-model="formData.powerPublishUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子内评论权限_指定部门和用户(json数组):" prop="powerCommentUser">
          <el-input v-model="formData.powerCommentUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="不受限用户组(json数组):" prop="noLimitUserGroup">
          <el-input v-model="formData.noLimitUserGroup" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="新注册用户默认关注：0 否，1是:" prop="newUserFocus">
          <el-input v-model.number="formData.newUserFocus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入" />
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
  name: 'HkCircle'
}
</script>

<script setup>
import {
  createHkCircle,
  updateHkCircle,
  findHkCircle
} from '@/api/hkCircle'

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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHkCircle({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehkCircle
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
