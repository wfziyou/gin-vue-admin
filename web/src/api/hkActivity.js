import service from '@/utils/request'

// @Tags HkActivity
// @Summary 创建HkActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivity true "创建HkActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivity/createHkActivity [post]
export const createHkActivity = (data) => {
  return service({
    url: '/hkActivity/createHkActivity',
    method: 'post',
    data
  })
}

// @Tags HkActivity
// @Summary 删除HkActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivity true "删除HkActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivity/deleteHkActivity [delete]
export const deleteHkActivity = (data) => {
  return service({
    url: '/hkActivity/deleteHkActivity',
    method: 'delete',
    data
  })
}

// @Tags HkActivity
// @Summary 删除HkActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivity/deleteHkActivity [delete]
export const deleteHkActivityByIds = (data) => {
  return service({
    url: '/hkActivity/deleteHkActivityByIds',
    method: 'delete',
    data
  })
}

// @Tags HkActivity
// @Summary 更新HkActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivity true "更新HkActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkActivity/updateHkActivity [put]
export const updateHkActivity = (data) => {
  return service({
    url: '/hkActivity/updateHkActivity',
    method: 'put',
    data
  })
}

// @Tags HkActivity
// @Summary 用id查询HkActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkActivity true "用id查询HkActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkActivity/findHkActivity [get]
export const findHkActivity = (params) => {
  return service({
    url: '/hkActivity/findHkActivity',
    method: 'get',
    params
  })
}

// @Tags HkActivity
// @Summary 分页获取HkActivity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkActivity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivity/getHkActivityList [get]
export const getHkActivityList = (params) => {
  return service({
    url: '/hkActivity/getHkActivityList',
    method: 'get',
    params
  })
}
