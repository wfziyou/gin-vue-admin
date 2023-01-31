<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="租户ID:" prop="tenantId">
          <el-input v-model="formData.tenantId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="圈子_编号:" prop="circleId">
          <el-input v-model.number="formData.circleId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）:" prop="category">
          <el-input v-model.number="formData.category" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="帖子分类编号:" prop="groupId">
          <el-input v-model.number="formData.groupId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="SEO关键词:" prop="seoKey">
          <el-input v-model="formData.seoKey" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简介(SEO简介):" prop="seoIntroduce">
          <el-input v-model="formData.seoIntroduce" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="封面:" prop="coverImage">
          <el-input v-model="formData.coverImage" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="来源:" prop="source">
          <el-input v-model="formData.source" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发布时间:" prop="time">
          <el-date-picker v-model="formData.time" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容类型:0 markdown,1 html:" prop="contentType">
          <el-input v-model.number="formData.contentType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="markdown内容:" prop="contentMarkdown">
          <el-input v-model="formData.contentMarkdown" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="html内容:" prop="contentHtml">
          <el-input v-model="formData.contentHtml" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="视频地址:" prop="video">
          <el-input v-model="formData.video" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="附件:" prop="attachment">
          <el-input v-model="formData.attachment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标签:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="置顶(0否 1是):" prop="top">
          <el-input v-model.number="formData.top" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="精华(0否 1是):" prop="marrow">
          <el-input v-model.number="formData.marrow" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="问答最佳答案ID:" prop="commentId">
          <el-input v-model.number="formData.commentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="匿名发布(0否 1是):" prop="anonymity">
          <el-input v-model.number="formData.anonymity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发布者编号:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="阅读次数:" prop="readNum">
          <el-input v-model.number="formData.readNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="评论次数:" prop="commentNum">
          <el-input v-model.number="formData.commentNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分享次数:" prop="shareNum">
          <el-input v-model.number="formData.shareNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收藏次数:" prop="collectNum">
          <el-input v-model.number="formData.collectNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="点赞次数:" prop="likeNum">
          <el-input v-model.number="formData.likeNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核人:" prop="checkUser">
          <el-input v-model.number="formData.checkUser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核时间:" prop="checkTime">
          <el-date-picker v-model="formData.checkTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="审核状态(0草稿 1未审批 2通过 3拒绝):" prop="checkStatus">
          <el-input v-model.number="formData.checkStatus" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="评论权限(0关闭 1开启):" prop="powerComment">
          <el-input v-model.number="formData.powerComment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="匿名评论(0关闭 1开启):" prop="powerCommentAnonymity">
          <el-input v-model.number="formData.powerCommentAnonymity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费：0 否，1是:" prop="pay">
          <el-input v-model.number="formData.pay" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容付费：0 否，1是:" prop="payContent">
          <el-input v-model.number="formData.payContent" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容付费可查看百分比例:" prop="payContentLook">
          <el-input v-model.number="formData.payContentLook" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="附件付费：0 否，1是:" prop="payAttachment">
          <el-input v-model.number="formData.payAttachment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费货币：0 人民、1积分、2代币:" prop="payCurrency">
          <el-input v-model.number="formData.payCurrency" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付费金额:" prop="payNum">
          <el-input v-model.number="formData.payNum" :clearable="true" placeholder="请输入" />
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
  name: 'HkForumPosts'
}
</script>

<script setup>
import {
  createHkForumPosts,
  updateHkForumPosts,
  findHkForumPosts
} from '@/api/hkForumPosts'

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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHkForumPosts({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehkForumPosts
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
