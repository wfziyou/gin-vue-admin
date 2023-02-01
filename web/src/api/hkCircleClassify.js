import service from '@/utils/request'

// @Tags HkCircleClassify
// @Summary 创建HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleClassify true "创建HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleClassify/createHkCircleClassify [post]
export const createHkCircleClassify = (data) => {
  return service({
    url: '/hkCircleClassify/createHkCircleClassify',
    method: 'post',
    data
  })
}

// @Tags HkCircleClassify
// @Summary 删除HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleClassify true "删除HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleClassify/deleteHkCircleClassify [delete]
export const deleteHkCircleClassify = (data) => {
  return service({
    url: '/hkCircleClassify/deleteHkCircleClassify',
    method: 'delete',
    data
  })
}

// @Tags HkCircleClassify
// @Summary 删除HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleClassify/deleteHkCircleClassify [delete]
export const deleteHkCircleClassifyByIds = (data) => {
  return service({
    url: '/hkCircleClassify/deleteHkCircleClassifyByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCircleClassify
// @Summary 更新HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleClassify true "更新HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleClassify/updateHkCircleClassify [put]
export const updateHkCircleClassify = (data) => {
  return service({
    url: '/hkCircleClassify/updateHkCircleClassify',
    method: 'put',
    data
  })
}

// @Tags HkCircleClassify
// @Summary 用id查询HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCircleClassify true "用id查询HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleClassify/findHkCircleClassify [get]
export const findHkCircleClassify = (params) => {
  return service({
    url: '/hkCircleClassify/findHkCircleClassify',
    method: 'get',
    params
  })
}

// @Tags HkCircleClassify
// @Summary 分页获取HkCircleClassify列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCircleClassify列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleClassify/getHkCircleClassifyList [get]
export const getHkCircleClassifyList = (params) => {
  return service({
    url: '/hkCircleClassify/getHkCircleClassifyList',
    method: 'get',
    params
  })
}
