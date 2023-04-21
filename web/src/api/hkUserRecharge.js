import service from '@/utils/request'

// @Tags HkUserRecharge
// @Summary 创建HkUserRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserRecharge true "创建HkUserRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserRecharge/createHkUserRecharge [post]
export const createHkUserRecharge = (data) => {
  return service({
    url: '/hkUserRecharge/createHkUserRecharge',
    method: 'post',
    data
  })
}

// @Tags HkUserRecharge
// @Summary 删除HkUserRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserRecharge true "删除HkUserRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserRecharge/deleteHkUserRecharge [delete]
export const deleteHkUserRecharge = (data) => {
  return service({
    url: '/hkUserRecharge/deleteHkUserRecharge',
    method: 'delete',
    data
  })
}

// @Tags HkUserRecharge
// @Summary 删除HkUserRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserRecharge/deleteHkUserRecharge [delete]
export const deleteHkUserRechargeByIds = (data) => {
  return service({
    url: '/hkUserRecharge/deleteHkUserRechargeByIds',
    method: 'delete',
    data
  })
}

// @Tags HkUserRecharge
// @Summary 更新HkUserRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserRecharge true "更新HkUserRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserRecharge/updateHkUserRecharge [put]
export const updateHkUserRecharge = (data) => {
  return service({
    url: '/hkUserRecharge/updateHkUserRecharge',
    method: 'put',
    data
  })
}

// @Tags HkUserRecharge
// @Summary 用id查询HkUserRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkUserRecharge true "用id查询HkUserRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserRecharge/findHkUserRecharge [get]
export const findHkUserRecharge = (params) => {
  return service({
    url: '/hkUserRecharge/findHkUserRecharge',
    method: 'get',
    params
  })
}

// @Tags HkUserRecharge
// @Summary 分页获取HkUserRecharge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkUserRecharge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserRecharge/getHkUserRechargeList [get]
export const getHkUserRechargeList = (params) => {
  return service({
    url: '/hkUserRecharge/getHkUserRechargeList',
    method: 'get',
    params
  })
}
