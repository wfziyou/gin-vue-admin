<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单号:" prop="orderId">
          <el-input v-model="formData.orderId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="额外订单号:" prop="extendOrderId">
          <el-input v-model="formData.extendOrderId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户编号:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品ID:" prop="productId">
          <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品数量:" prop="cartNum">
          <el-input v-model.number="formData.cartNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单总价:" prop="totalPrice">
          <el-input-number v-model="formData.totalPrice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="实际支付金额:" prop="payPrice">
          <el-input-number v-model="formData.payPrice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="支付状态:" prop="paid">
          <el-switch v-model="formData.paid" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="支付时间:" prop="payTime">
          <el-date-picker v-model="formData.payTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="支付方式:" prop="payType">
          <el-input v-model="formData.payType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0 未退款 1 申请中 2 已退款:" prop="refundStatus">
          <el-switch v-model="formData.refundStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="退款图片:" prop="refundReasonWapImg">
          <el-input v-model="formData.refundReasonWapImg" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退款用户说明:" prop="refundReasonWapExplain">
          <el-input v-model="formData.refundReasonWapExplain" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退款时间:" prop="refundReasonTime">
          <el-date-picker v-model="formData.refundReasonTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="前台退款原因:" prop="refundReasonWap">
          <el-input v-model="formData.refundReasonWap" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="不退款的理由:" prop="refundReason">
          <el-input v-model="formData.refundReason" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退款金额:" prop="refundPrice">
          <el-input-number v-model="formData.refundPrice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="消费赚取积分:" prop="gainIntegral">
          <el-input-number v-model="formData.gainIntegral" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="给用户退了多少积分:" prop="backIntegral">
          <el-input-number v-model="formData.backIntegral" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="备注:" prop="mark">
          <el-input v-model="formData.mark" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="唯一id(md5加密)类似id:" prop="unique">
          <el-input v-model="formData.unique" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="管理员备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="成本价:" prop="cost">
          <el-input-number v-model="formData.cost" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="核销码:" prop="verifyCode">
          <el-input v-model="formData.verifyCode" :clearable="true" placeholder="请输入" />
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
        <el-form-item label="状态:0待发货 1已完成 2申请退款 3已退货 4已退款 5拒绝退款:" prop="status">
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
  name: 'HkOrder'
}
</script>

<script setup>
import {
  createHkOrder,
  updateHkOrder,
  findHkOrder
} from '@/api/hkOrder'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            orderId: '',
            extendOrderId: '',
            userId: 0,
            productId: 0,
            cartNum: 0,
            totalPrice: 0,
            payPrice: 0,
            paid: false,
            payTime: new Date(),
            payType: '',
            refundStatus: false,
            refundReasonWapImg: '',
            refundReasonWapExplain: '',
            refundReasonTime: new Date(),
            refundReasonWap: '',
            refundReason: '',
            refundPrice: 0,
            gainIntegral: 0,
            backIntegral: 0,
            mark: '',
            unique: '',
            remark: '',
            cost: 0,
            verifyCode: '',
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
      const res = await findHkOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehkOrder
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
               res = await createHkOrder(formData.value)
               break
             case 'update':
               res = await updateHkOrder(formData.value)
               break
             default:
               res = await createHkOrder(formData.value)
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
