import service from '@/utils/request'

// @Tags HkActivityCollect
// @Summary 创建HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityCollect true "创建HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityCollect/createHkActivityCollect [post]
export const createHkActivityCollect = (data) => {
  return service({
    url: '/hkActivityCollect/createHkActivityCollect',
    method: 'post',
    data
  })
}

// @Tags HkActivityCollect
// @Summary 删除HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityCollect true "删除HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityCollect/deleteHkActivityCollect [delete]
export const deleteHkActivityCollect = (data) => {
  return service({
    url: '/hkActivityCollect/deleteHkActivityCollect',
    method: 'delete',
    data
  })
}

// @Tags HkActivityCollect
// @Summary 删除HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityCollect/deleteHkActivityCollect [delete]
export const deleteHkActivityCollectByIds = (data) => {
  return service({
    url: '/hkActivityCollect/deleteHkActivityCollectByIds',
    method: 'delete',
    data
  })
}

// @Tags HkActivityCollect
// @Summary 更新HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityCollect true "更新HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkActivityCollect/updateHkActivityCollect [put]
export const updateHkActivityCollect = (data) => {
  return service({
    url: '/hkActivityCollect/updateHkActivityCollect',
    method: 'put',
    data
  })
}

// @Tags HkActivityCollect
// @Summary 用id查询HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkActivityCollect true "用id查询HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkActivityCollect/findHkActivityCollect [get]
export const findHkActivityCollect = (params) => {
  return service({
    url: '/hkActivityCollect/findHkActivityCollect',
    method: 'get',
    params
  })
}

// @Tags HkActivityCollect
// @Summary 分页获取HkActivityCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkActivityCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityCollect/getHkActivityCollectList [get]
export const getHkActivityCollectList = (params) => {
  return service({
    url: '/hkActivityCollect/getHkActivityCollectList',
    method: 'get',
    params
  })
}
