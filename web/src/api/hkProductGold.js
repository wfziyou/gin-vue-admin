import service from '@/utils/request'

// @Tags HkProductGold
// @Summary 创建HkProductGold
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkProductGold true "创建HkProductGold"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkProductGold/createHkProductGold [post]
export const createHkProductGold = (data) => {
  return service({
    url: '/hkProductGold/createHkProductGold',
    method: 'post',
    data
  })
}

// @Tags HkProductGold
// @Summary 删除HkProductGold
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkProductGold true "删除HkProductGold"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkProductGold/deleteHkProductGold [delete]
export const deleteHkProductGold = (data) => {
  return service({
    url: '/hkProductGold/deleteHkProductGold',
    method: 'delete',
    data
  })
}

// @Tags HkProductGold
// @Summary 删除HkProductGold
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkProductGold"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkProductGold/deleteHkProductGold [delete]
export const deleteHkProductGoldByIds = (data) => {
  return service({
    url: '/hkProductGold/deleteHkProductGoldByIds',
    method: 'delete',
    data
  })
}

// @Tags HkProductGold
// @Summary 更新HkProductGold
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkProductGold true "更新HkProductGold"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkProductGold/updateHkProductGold [put]
export const updateHkProductGold = (data) => {
  return service({
    url: '/hkProductGold/updateHkProductGold',
    method: 'put',
    data
  })
}

// @Tags HkProductGold
// @Summary 用id查询HkProductGold
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkProductGold true "用id查询HkProductGold"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkProductGold/findHkProductGold [get]
export const findHkProductGold = (params) => {
  return service({
    url: '/hkProductGold/findHkProductGold',
    method: 'get',
    params
  })
}

// @Tags HkProductGold
// @Summary 分页获取HkProductGold列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkProductGold列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkProductGold/getHkProductGoldList [get]
export const getHkProductGoldList = (params) => {
  return service({
    url: '/hkProductGold/getHkProductGoldList',
    method: 'get',
    params
  })
}
