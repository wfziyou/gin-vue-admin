import service from '@/utils/request'

// @Tags HkCircleRelationRequest
// @Summary 创建HkCircleRelationRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleRelationRequest true "创建HkCircleRelationRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRelationRequest/createHkCircleRelationRequest [post]
export const createHkCircleRelationRequest = (data) => {
  return service({
    url: '/hkCircleRelationRequest/createHkCircleRelationRequest',
    method: 'post',
    data
  })
}

// @Tags HkCircleRelationRequest
// @Summary 删除HkCircleRelationRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleRelationRequest true "删除HkCircleRelationRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleRelationRequest/deleteHkCircleRelationRequest [delete]
export const deleteHkCircleRelationRequest = (data) => {
  return service({
    url: '/hkCircleRelationRequest/deleteHkCircleRelationRequest',
    method: 'delete',
    data
  })
}

// @Tags HkCircleRelationRequest
// @Summary 删除HkCircleRelationRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleRelationRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleRelationRequest/deleteHkCircleRelationRequest [delete]
export const deleteHkCircleRelationRequestByIds = (data) => {
  return service({
    url: '/hkCircleRelationRequest/deleteHkCircleRelationRequestByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCircleRelationRequest
// @Summary 更新HkCircleRelationRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleRelationRequest true "更新HkCircleRelationRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleRelationRequest/updateHkCircleRelationRequest [put]
export const updateHkCircleRelationRequest = (data) => {
  return service({
    url: '/hkCircleRelationRequest/updateHkCircleRelationRequest',
    method: 'put',
    data
  })
}

// @Tags HkCircleRelationRequest
// @Summary 用id查询HkCircleRelationRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCircleRelationRequest true "用id查询HkCircleRelationRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleRelationRequest/findHkCircleRelationRequest [get]
export const findHkCircleRelationRequest = (params) => {
  return service({
    url: '/hkCircleRelationRequest/findHkCircleRelationRequest',
    method: 'get',
    params
  })
}

// @Tags HkCircleRelationRequest
// @Summary 分页获取HkCircleRelationRequest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCircleRelationRequest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRelationRequest/getHkCircleRelationRequestList [get]
export const getHkCircleRelationRequestList = (params) => {
  return service({
    url: '/hkCircleRelationRequest/getHkCircleRelationRequestList',
    method: 'get',
    params
  })
}
