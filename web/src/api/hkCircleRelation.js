import service from '@/utils/request'

// @Tags HkCircleRelation
// @Summary 创建HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleRelation true "创建HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRelation/createHkCircleRelation [post]
export const createHkCircleRelation = (data) => {
  return service({
    url: '/hkCircleRelation/createHkCircleRelation',
    method: 'post',
    data
  })
}

// @Tags HkCircleRelation
// @Summary 删除HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleRelation true "删除HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleRelation/deleteHkCircleRelation [delete]
export const deleteHkCircleRelation = (data) => {
  return service({
    url: '/hkCircleRelation/deleteHkCircleRelation',
    method: 'delete',
    data
  })
}

// @Tags HkCircleRelation
// @Summary 删除HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleRelation/deleteHkCircleRelation [delete]
export const deleteHkCircleRelationByIds = (data) => {
  return service({
    url: '/hkCircleRelation/deleteHkCircleRelationByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCircleRelation
// @Summary 更新HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleRelation true "更新HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleRelation/updateHkCircleRelation [put]
export const updateHkCircleRelation = (data) => {
  return service({
    url: '/hkCircleRelation/updateHkCircleRelation',
    method: 'put',
    data
  })
}

// @Tags HkCircleRelation
// @Summary 用id查询HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCircleRelation true "用id查询HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleRelation/findHkCircleRelation [get]
export const findHkCircleRelation = (params) => {
  return service({
    url: '/hkCircleRelation/findHkCircleRelation',
    method: 'get',
    params
  })
}

// @Tags HkCircleRelation
// @Summary 分页获取HkCircleRelation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCircleRelation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRelation/getHkCircleRelationList [get]
export const getHkCircleRelationList = (params) => {
  return service({
    url: '/hkCircleRelation/getHkCircleRelationList',
    method: 'get',
    params
  })
}
