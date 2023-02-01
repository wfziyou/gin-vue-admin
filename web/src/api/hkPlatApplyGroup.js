import service from '@/utils/request'

// @Tags HkPlatApplyGroup
// @Summary 创建HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkPlatApplyGroup true "创建HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkPlatApplyGroup/createHkPlatApplyGroup [post]
export const createHkPlatApplyGroup = (data) => {
  return service({
    url: '/hkPlatApplyGroup/createHkPlatApplyGroup',
    method: 'post',
    data
  })
}

// @Tags HkPlatApplyGroup
// @Summary 删除HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkPlatApplyGroup true "删除HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkPlatApplyGroup/deleteHkPlatApplyGroup [delete]
export const deleteHkPlatApplyGroup = (data) => {
  return service({
    url: '/hkPlatApplyGroup/deleteHkPlatApplyGroup',
    method: 'delete',
    data
  })
}

// @Tags HkPlatApplyGroup
// @Summary 删除HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkPlatApplyGroup/deleteHkPlatApplyGroup [delete]
export const deleteHkPlatApplyGroupByIds = (data) => {
  return service({
    url: '/hkPlatApplyGroup/deleteHkPlatApplyGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags HkPlatApplyGroup
// @Summary 更新HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkPlatApplyGroup true "更新HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkPlatApplyGroup/updateHkPlatApplyGroup [put]
export const updateHkPlatApplyGroup = (data) => {
  return service({
    url: '/hkPlatApplyGroup/updateHkPlatApplyGroup',
    method: 'put',
    data
  })
}

// @Tags HkPlatApplyGroup
// @Summary 用id查询HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkPlatApplyGroup true "用id查询HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkPlatApplyGroup/findHkPlatApplyGroup [get]
export const findHkPlatApplyGroup = (params) => {
  return service({
    url: '/hkPlatApplyGroup/findHkPlatApplyGroup',
    method: 'get',
    params
  })
}

// @Tags HkPlatApplyGroup
// @Summary 分页获取HkPlatApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkPlatApplyGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkPlatApplyGroup/getHkPlatApplyGroupList [get]
export const getHkPlatApplyGroupList = (params) => {
  return service({
    url: '/hkPlatApplyGroup/getHkPlatApplyGroupList',
    method: 'get',
    params
  })
}
