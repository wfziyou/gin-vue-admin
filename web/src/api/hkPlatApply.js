import service from '@/utils/request'

// @Tags HkPlatApply
// @Summary 创建HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkPlatApply true "创建HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkPlatApply/createHkPlatApply [post]
export const createHkPlatApply = (data) => {
  return service({
    url: '/hkPlatApply/createHkPlatApply',
    method: 'post',
    data
  })
}

// @Tags HkPlatApply
// @Summary 删除HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkPlatApply true "删除HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkPlatApply/deleteHkPlatApply [delete]
export const deleteHkPlatApply = (data) => {
  return service({
    url: '/hkPlatApply/deleteHkPlatApply',
    method: 'delete',
    data
  })
}

// @Tags HkPlatApply
// @Summary 删除HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkPlatApply/deleteHkPlatApply [delete]
export const deleteHkPlatApplyByIds = (data) => {
  return service({
    url: '/hkPlatApply/deleteHkPlatApplyByIds',
    method: 'delete',
    data
  })
}

// @Tags HkPlatApply
// @Summary 更新HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkPlatApply true "更新HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkPlatApply/updateHkPlatApply [put]
export const updateHkPlatApply = (data) => {
  return service({
    url: '/hkPlatApply/updateHkPlatApply',
    method: 'put',
    data
  })
}

// @Tags HkPlatApply
// @Summary 用id查询HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkPlatApply true "用id查询HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkPlatApply/findHkPlatApply [get]
export const findHkPlatApply = (params) => {
  return service({
    url: '/hkPlatApply/findHkPlatApply',
    method: 'get',
    params
  })
}

// @Tags HkPlatApply
// @Summary 分页获取HkPlatApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkPlatApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkPlatApply/getHkPlatApplyList [get]
export const getHkPlatApplyList = (params) => {
  return service({
    url: '/hkPlatApply/getHkPlatApplyList',
    method: 'get',
    params
  })
}
