import service from '@/utils/request'

// @Tags HkUserCircleApply
// @Summary 创建HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserCircleApply true "创建HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserCircleApply/createHkUserCircleApply [post]
export const createHkUserCircleApply = (data) => {
  return service({
    url: '/hkUserCircleApply/createHkUserCircleApply',
    method: 'post',
    data
  })
}

// @Tags HkUserCircleApply
// @Summary 删除HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserCircleApply true "删除HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserCircleApply/deleteHkUserCircleApply [delete]
export const deleteHkUserCircleApply = (data) => {
  return service({
    url: '/hkUserCircleApply/deleteHkUserCircleApply',
    method: 'delete',
    data
  })
}

// @Tags HkUserCircleApply
// @Summary 删除HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserCircleApply/deleteHkUserCircleApply [delete]
export const deleteHkUserCircleApplyByIds = (data) => {
  return service({
    url: '/hkUserCircleApply/deleteHkUserCircleApplyByIds',
    method: 'delete',
    data
  })
}

// @Tags HkUserCircleApply
// @Summary 更新HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserCircleApply true "更新HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserCircleApply/updateHkUserCircleApply [put]
export const updateHkUserCircleApply = (data) => {
  return service({
    url: '/hkUserCircleApply/updateHkUserCircleApply',
    method: 'put',
    data
  })
}

// @Tags HkUserCircleApply
// @Summary 用id查询HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkUserCircleApply true "用id查询HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserCircleApply/findHkUserCircleApply [get]
export const findHkUserCircleApply = (params) => {
  return service({
    url: '/hkUserCircleApply/findHkUserCircleApply',
    method: 'get',
    params
  })
}

// @Tags HkUserCircleApply
// @Summary 分页获取HkUserCircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkUserCircleApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserCircleApply/getHkUserCircleApplyList [get]
export const getHkUserCircleApplyList = (params) => {
  return service({
    url: '/hkUserCircleApply/getHkUserCircleApplyList',
    method: 'get',
    params
  })
}
