import service from '@/utils/request'

// @Tags HkThirdPlatform
// @Summary 创建HkThirdPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkThirdPlatform true "创建HkThirdPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkThirdPlatform/createHkThirdPlatform [post]
export const createHkThirdPlatform = (data) => {
  return service({
    url: '/hkThirdPlatform/createHkThirdPlatform',
    method: 'post',
    data
  })
}

// @Tags HkThirdPlatform
// @Summary 删除HkThirdPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkThirdPlatform true "删除HkThirdPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkThirdPlatform/deleteHkThirdPlatform [delete]
export const deleteHkThirdPlatform = (data) => {
  return service({
    url: '/hkThirdPlatform/deleteHkThirdPlatform',
    method: 'delete',
    data
  })
}

// @Tags HkThirdPlatform
// @Summary 删除HkThirdPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkThirdPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkThirdPlatform/deleteHkThirdPlatform [delete]
export const deleteHkThirdPlatformByIds = (data) => {
  return service({
    url: '/hkThirdPlatform/deleteHkThirdPlatformByIds',
    method: 'delete',
    data
  })
}

// @Tags HkThirdPlatform
// @Summary 更新HkThirdPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkThirdPlatform true "更新HkThirdPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkThirdPlatform/updateHkThirdPlatform [put]
export const updateHkThirdPlatform = (data) => {
  return service({
    url: '/hkThirdPlatform/updateHkThirdPlatform',
    method: 'put',
    data
  })
}

// @Tags HkThirdPlatform
// @Summary 用id查询HkThirdPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkThirdPlatform true "用id查询HkThirdPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkThirdPlatform/findHkThirdPlatform [get]
export const findHkThirdPlatform = (params) => {
  return service({
    url: '/hkThirdPlatform/findHkThirdPlatform',
    method: 'get',
    params
  })
}

// @Tags HkThirdPlatform
// @Summary 分页获取HkThirdPlatform列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkThirdPlatform列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkThirdPlatform/getHkThirdPlatformList [get]
export const getHkThirdPlatformList = (params) => {
  return service({
    url: '/hkThirdPlatform/getHkThirdPlatformList',
    method: 'get',
    params
  })
}
