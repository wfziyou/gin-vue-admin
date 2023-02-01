import service from '@/utils/request'

// @Tags HkProtocol
// @Summary 创建HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkProtocol true "创建HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkProtocol/createHkProtocol [post]
export const createHkProtocol = (data) => {
  return service({
    url: '/hkProtocol/createHkProtocol',
    method: 'post',
    data
  })
}

// @Tags HkProtocol
// @Summary 删除HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkProtocol true "删除HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkProtocol/deleteHkProtocol [delete]
export const deleteHkProtocol = (data) => {
  return service({
    url: '/hkProtocol/deleteHkProtocol',
    method: 'delete',
    data
  })
}

// @Tags HkProtocol
// @Summary 删除HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkProtocol/deleteHkProtocol [delete]
export const deleteHkProtocolByIds = (data) => {
  return service({
    url: '/hkProtocol/deleteHkProtocolByIds',
    method: 'delete',
    data
  })
}

// @Tags HkProtocol
// @Summary 更新HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkProtocol true "更新HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkProtocol/updateHkProtocol [put]
export const updateHkProtocol = (data) => {
  return service({
    url: '/hkProtocol/updateHkProtocol',
    method: 'put',
    data
  })
}

// @Tags HkProtocol
// @Summary 用id查询HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkProtocol true "用id查询HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkProtocol/findHkProtocol [get]
export const findHkProtocol = (params) => {
  return service({
    url: '/hkProtocol/findHkProtocol',
    method: 'get',
    params
  })
}

// @Tags HkProtocol
// @Summary 分页获取HkProtocol列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkProtocol列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkProtocol/getHkProtocolList [get]
export const getHkProtocolList = (params) => {
  return service({
    url: '/hkProtocol/getHkProtocolList',
    method: 'get',
    params
  })
}
