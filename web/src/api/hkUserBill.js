import service from '@/utils/request'

// @Tags HkUserBill
// @Summary 创建HkUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserBill true "创建HkUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserBill/createHkUserBill [post]
export const createHkUserBill = (data) => {
  return service({
    url: '/hkUserBill/createHkUserBill',
    method: 'post',
    data
  })
}

// @Tags HkUserBill
// @Summary 删除HkUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserBill true "删除HkUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserBill/deleteHkUserBill [delete]
export const deleteHkUserBill = (data) => {
  return service({
    url: '/hkUserBill/deleteHkUserBill',
    method: 'delete',
    data
  })
}

// @Tags HkUserBill
// @Summary 删除HkUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserBill/deleteHkUserBill [delete]
export const deleteHkUserBillByIds = (data) => {
  return service({
    url: '/hkUserBill/deleteHkUserBillByIds',
    method: 'delete',
    data
  })
}

// @Tags HkUserBill
// @Summary 更新HkUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserBill true "更新HkUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserBill/updateHkUserBill [put]
export const updateHkUserBill = (data) => {
  return service({
    url: '/hkUserBill/updateHkUserBill',
    method: 'put',
    data
  })
}

// @Tags HkUserBill
// @Summary 用id查询HkUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkUserBill true "用id查询HkUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserBill/findHkUserBill [get]
export const findHkUserBill = (params) => {
  return service({
    url: '/hkUserBill/findHkUserBill',
    method: 'get',
    params
  })
}

// @Tags HkUserBill
// @Summary 分页获取HkUserBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkUserBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserBill/getHkUserBillList [get]
export const getHkUserBillList = (params) => {
  return service({
    url: '/hkUserBill/getHkUserBillList',
    method: 'get',
    params
  })
}
