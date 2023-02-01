import service from '@/utils/request'

// @Tags HkUserExtend
// @Summary 创建HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserExtend true "创建HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserExtend/createHkUserExtend [post]
export const createHkUserExtend = (data) => {
  return service({
    url: '/hkUserExtend/createHkUserExtend',
    method: 'post',
    data
  })
}

// @Tags HkUserExtend
// @Summary 删除HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserExtend true "删除HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserExtend/deleteHkUserExtend [delete]
export const deleteHkUserExtend = (data) => {
  return service({
    url: '/hkUserExtend/deleteHkUserExtend',
    method: 'delete',
    data
  })
}

// @Tags HkUserExtend
// @Summary 删除HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserExtend/deleteHkUserExtend [delete]
export const deleteHkUserExtendByIds = (data) => {
  return service({
    url: '/hkUserExtend/deleteHkUserExtendByIds',
    method: 'delete',
    data
  })
}

// @Tags HkUserExtend
// @Summary 更新HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserExtend true "更新HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserExtend/updateHkUserExtend [put]
export const updateHkUserExtend = (data) => {
  return service({
    url: '/hkUserExtend/updateHkUserExtend',
    method: 'put',
    data
  })
}

// @Tags HkUserExtend
// @Summary 用id查询HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkUserExtend true "用id查询HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserExtend/findHkUserExtend [get]
export const findHkUserExtend = (params) => {
  return service({
    url: '/hkUserExtend/findHkUserExtend',
    method: 'get',
    params
  })
}

// @Tags HkUserExtend
// @Summary 分页获取HkUserExtend列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkUserExtend列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserExtend/getHkUserExtendList [get]
export const getHkUserExtendList = (params) => {
  return service({
    url: '/hkUserExtend/getHkUserExtendList',
    method: 'get',
    params
  })
}
