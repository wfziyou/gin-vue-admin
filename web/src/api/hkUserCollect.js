import service from '@/utils/request'

// @Tags HkUserCollect
// @Summary 创建HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserCollect true "创建HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserCollect/createHkUserCollect [post]
export const createHkUserCollect = (data) => {
  return service({
    url: '/hkUserCollect/createHkUserCollect',
    method: 'post',
    data
  })
}

// @Tags HkUserCollect
// @Summary 删除HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserCollect true "删除HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserCollect/deleteHkUserCollect [delete]
export const deleteHkUserCollect = (data) => {
  return service({
    url: '/hkUserCollect/deleteHkUserCollect',
    method: 'delete',
    data
  })
}

// @Tags HkUserCollect
// @Summary 删除HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserCollect/deleteHkUserCollect [delete]
export const deleteHkUserCollectByIds = (data) => {
  return service({
    url: '/hkUserCollect/deleteHkUserCollectByIds',
    method: 'delete',
    data
  })
}

// @Tags HkUserCollect
// @Summary 更新HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserCollect true "更新HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserCollect/updateHkUserCollect [put]
export const updateHkUserCollect = (data) => {
  return service({
    url: '/hkUserCollect/updateHkUserCollect',
    method: 'put',
    data
  })
}

// @Tags HkUserCollect
// @Summary 用id查询HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkUserCollect true "用id查询HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserCollect/findHkUserCollect [get]
export const findHkUserCollect = (params) => {
  return service({
    url: '/hkUserCollect/findHkUserCollect',
    method: 'get',
    params
  })
}

// @Tags HkUserCollect
// @Summary 分页获取HkUserCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkUserCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserCollect/getHkUserCollectList [get]
export const getHkUserCollectList = (params) => {
  return service({
    url: '/hkUserCollect/getHkUserCollectList',
    method: 'get',
    params
  })
}
