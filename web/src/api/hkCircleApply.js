import service from '@/utils/request'

// @Tags HkCircleApply
// @Summary 创建HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleApply true "创建HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleApply/createHkCircleApply [post]
export const createHkCircleApply = (data) => {
  return service({
    url: '/hkCircleApply/createHkCircleApply',
    method: 'post',
    data
  })
}

// @Tags HkCircleApply
// @Summary 删除HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleApply true "删除HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleApply/deleteHkCircleApply [delete]
export const deleteHkCircleApply = (data) => {
  return service({
    url: '/hkCircleApply/deleteHkCircleApply',
    method: 'delete',
    data
  })
}

// @Tags HkCircleApply
// @Summary 删除HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleApply/deleteHkCircleApply [delete]
export const deleteHkCircleApplyByIds = (data) => {
  return service({
    url: '/hkCircleApply/deleteHkCircleApplyByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCircleApply
// @Summary 更新HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleApply true "更新HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleApply/updateHkCircleApply [put]
export const updateHkCircleApply = (data) => {
  return service({
    url: '/hkCircleApply/updateHkCircleApply',
    method: 'put',
    data
  })
}

// @Tags HkCircleApply
// @Summary 用id查询HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCircleApply true "用id查询HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleApply/findHkCircleApply [get]
export const findHkCircleApply = (params) => {
  return service({
    url: '/hkCircleApply/findHkCircleApply',
    method: 'get',
    params
  })
}

// @Tags HkCircleApply
// @Summary 分页获取HkCircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCircleApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleApply/getHkCircleApplyList [get]
export const getHkCircleApplyList = (params) => {
  return service({
    url: '/hkCircleApply/getHkCircleApplyList',
    method: 'get',
    params
  })
}
