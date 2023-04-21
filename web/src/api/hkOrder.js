import service from '@/utils/request'

// @Tags HkOrder
// @Summary 创建HkOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkOrder true "创建HkOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkOrder/createHkOrder [post]
export const createHkOrder = (data) => {
  return service({
    url: '/hkOrder/createHkOrder',
    method: 'post',
    data
  })
}

// @Tags HkOrder
// @Summary 删除HkOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkOrder true "删除HkOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkOrder/deleteHkOrder [delete]
export const deleteHkOrder = (data) => {
  return service({
    url: '/hkOrder/deleteHkOrder',
    method: 'delete',
    data
  })
}

// @Tags HkOrder
// @Summary 删除HkOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkOrder/deleteHkOrder [delete]
export const deleteHkOrderByIds = (data) => {
  return service({
    url: '/hkOrder/deleteHkOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags HkOrder
// @Summary 更新HkOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkOrder true "更新HkOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkOrder/updateHkOrder [put]
export const updateHkOrder = (data) => {
  return service({
    url: '/hkOrder/updateHkOrder',
    method: 'put',
    data
  })
}

// @Tags HkOrder
// @Summary 用id查询HkOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkOrder true "用id查询HkOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkOrder/findHkOrder [get]
export const findHkOrder = (params) => {
  return service({
    url: '/hkOrder/findHkOrder',
    method: 'get',
    params
  })
}

// @Tags HkOrder
// @Summary 分页获取HkOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkOrder/getHkOrderList [get]
export const getHkOrderList = (params) => {
  return service({
    url: '/hkOrder/getHkOrderList',
    method: 'get',
    params
  })
}
