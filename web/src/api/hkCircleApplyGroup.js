import service from '@/utils/request'

// @Tags HkCircleApplyGroup
// @Summary 创建HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleApplyGroup true "创建HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleApplyGroup/createHkCircleApplyGroup [post]
export const createHkCircleApplyGroup = (data) => {
  return service({
    url: '/hkCircleApplyGroup/createHkCircleApplyGroup',
    method: 'post',
    data
  })
}

// @Tags HkCircleApplyGroup
// @Summary 删除HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleApplyGroup true "删除HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleApplyGroup/deleteHkCircleApplyGroup [delete]
export const deleteHkCircleApplyGroup = (data) => {
  return service({
    url: '/hkCircleApplyGroup/deleteHkCircleApplyGroup',
    method: 'delete',
    data
  })
}

// @Tags HkCircleApplyGroup
// @Summary 删除HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleApplyGroup/deleteHkCircleApplyGroup [delete]
export const deleteHkCircleApplyGroupByIds = (data) => {
  return service({
    url: '/hkCircleApplyGroup/deleteHkCircleApplyGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCircleApplyGroup
// @Summary 更新HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleApplyGroup true "更新HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleApplyGroup/updateHkCircleApplyGroup [put]
export const updateHkCircleApplyGroup = (data) => {
  return service({
    url: '/hkCircleApplyGroup/updateHkCircleApplyGroup',
    method: 'put',
    data
  })
}

// @Tags HkCircleApplyGroup
// @Summary 用id查询HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCircleApplyGroup true "用id查询HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleApplyGroup/findHkCircleApplyGroup [get]
export const findHkCircleApplyGroup = (params) => {
  return service({
    url: '/hkCircleApplyGroup/findHkCircleApplyGroup',
    method: 'get',
    params
  })
}

// @Tags HkCircleApplyGroup
// @Summary 分页获取HkCircleApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCircleApplyGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleApplyGroup/getHkCircleApplyGroupList [get]
export const getHkCircleApplyGroupList = (params) => {
  return service({
    url: '/hkCircleApplyGroup/getHkCircleApplyGroupList',
    method: 'get',
    params
  })
}
