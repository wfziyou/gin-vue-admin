import service from '@/utils/request'

// @Tags HkApply
// @Summary 创建HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkApply true "创建HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkApply/createHkApply [post]
export const createHkApply = (data) => {
  return service({
    url: '/hkApply/createHkApply',
    method: 'post',
    data
  })
}

// @Tags HkApply
// @Summary 删除HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkApply true "删除HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkApply/deleteHkApply [delete]
export const deleteHkApply = (data) => {
  return service({
    url: '/hkApply/deleteHkApply',
    method: 'delete',
    data
  })
}

// @Tags HkApply
// @Summary 删除HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkApply/deleteHkApply [delete]
export const deleteHkApplyByIds = (data) => {
  return service({
    url: '/hkApply/deleteHkApplyByIds',
    method: 'delete',
    data
  })
}

// @Tags HkApply
// @Summary 更新HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkApply true "更新HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkApply/updateHkApply [put]
export const updateHkApply = (data) => {
  return service({
    url: '/hkApply/updateHkApply',
    method: 'put',
    data
  })
}

// @Tags HkApply
// @Summary 用id查询HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkApply true "用id查询HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkApply/findHkApply [get]
export const findHkApply = (params) => {
  return service({
    url: '/hkApply/findHkApply',
    method: 'get',
    params
  })
}

// @Tags HkApply
// @Summary 分页获取HkApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkApply/getHkApplyList [get]
export const getHkApplyList = (params) => {
  return service({
    url: '/hkApply/getHkApplyList',
    method: 'get',
    params
  })
}
