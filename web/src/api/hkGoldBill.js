import service from '@/utils/request'

// @Tags HkGoldBill
// @Summary 创建HkGoldBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkGoldBill true "创建HkGoldBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkGoldBill/createHkGoldBill [post]
export const createHkGoldBill = (data) => {
  return service({
    url: '/hkGoldBill/createHkGoldBill',
    method: 'post',
    data
  })
}

// @Tags HkGoldBill
// @Summary 删除HkGoldBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkGoldBill true "删除HkGoldBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkGoldBill/deleteHkGoldBill [delete]
export const deleteHkGoldBill = (data) => {
  return service({
    url: '/hkGoldBill/deleteHkGoldBill',
    method: 'delete',
    data
  })
}

// @Tags HkGoldBill
// @Summary 删除HkGoldBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkGoldBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkGoldBill/deleteHkGoldBill [delete]
export const deleteHkGoldBillByIds = (data) => {
  return service({
    url: '/hkGoldBill/deleteHkGoldBillByIds',
    method: 'delete',
    data
  })
}

// @Tags HkGoldBill
// @Summary 更新HkGoldBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkGoldBill true "更新HkGoldBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkGoldBill/updateHkGoldBill [put]
export const updateHkGoldBill = (data) => {
  return service({
    url: '/hkGoldBill/updateHkGoldBill',
    method: 'put',
    data
  })
}

// @Tags HkGoldBill
// @Summary 用id查询HkGoldBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkGoldBill true "用id查询HkGoldBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkGoldBill/findHkGoldBill [get]
export const findHkGoldBill = (params) => {
  return service({
    url: '/hkGoldBill/findHkGoldBill',
    method: 'get',
    params
  })
}

// @Tags HkGoldBill
// @Summary 分页获取HkGoldBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkGoldBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkGoldBill/getHkGoldBillList [get]
export const getHkGoldBillList = (params) => {
  return service({
    url: '/hkGoldBill/getHkGoldBillList',
    method: 'get',
    params
  })
}
