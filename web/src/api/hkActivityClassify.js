import service from '@/utils/request'

// @Tags HkActivityClassify
// @Summary 创建HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityClassify true "创建HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityClassify/createHkActivityClassify [post]
export const createHkActivityClassify = (data) => {
  return service({
    url: '/hkActivityClassify/createHkActivityClassify',
    method: 'post',
    data
  })
}

// @Tags HkActivityClassify
// @Summary 删除HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityClassify true "删除HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityClassify/deleteHkActivityClassify [delete]
export const deleteHkActivityClassify = (data) => {
  return service({
    url: '/hkActivityClassify/deleteHkActivityClassify',
    method: 'delete',
    data
  })
}

// @Tags HkActivityClassify
// @Summary 删除HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityClassify/deleteHkActivityClassify [delete]
export const deleteHkActivityClassifyByIds = (data) => {
  return service({
    url: '/hkActivityClassify/deleteHkActivityClassifyByIds',
    method: 'delete',
    data
  })
}

// @Tags HkActivityClassify
// @Summary 更新HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityClassify true "更新HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkActivityClassify/updateHkActivityClassify [put]
export const updateHkActivityClassify = (data) => {
  return service({
    url: '/hkActivityClassify/updateHkActivityClassify',
    method: 'put',
    data
  })
}

// @Tags HkActivityClassify
// @Summary 用id查询HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkActivityClassify true "用id查询HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkActivityClassify/findHkActivityClassify [get]
export const findHkActivityClassify = (params) => {
  return service({
    url: '/hkActivityClassify/findHkActivityClassify',
    method: 'get',
    params
  })
}

// @Tags HkActivityClassify
// @Summary 分页获取HkActivityClassify列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkActivityClassify列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityClassify/getHkActivityClassifyList [get]
export const getHkActivityClassifyList = (params) => {
  return service({
    url: '/hkActivityClassify/getHkActivityClassifyList',
    method: 'get',
    params
  })
}
