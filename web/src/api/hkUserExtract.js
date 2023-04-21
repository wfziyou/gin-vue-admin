import service from '@/utils/request'

// @Tags HkUserExtract
// @Summary 创建HkUserExtract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserExtract true "创建HkUserExtract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserExtract/createHkUserExtract [post]
export const createHkUserExtract = (data) => {
  return service({
    url: '/hkUserExtract/createHkUserExtract',
    method: 'post',
    data
  })
}

// @Tags HkUserExtract
// @Summary 删除HkUserExtract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserExtract true "删除HkUserExtract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserExtract/deleteHkUserExtract [delete]
export const deleteHkUserExtract = (data) => {
  return service({
    url: '/hkUserExtract/deleteHkUserExtract',
    method: 'delete',
    data
  })
}

// @Tags HkUserExtract
// @Summary 删除HkUserExtract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserExtract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserExtract/deleteHkUserExtract [delete]
export const deleteHkUserExtractByIds = (data) => {
  return service({
    url: '/hkUserExtract/deleteHkUserExtractByIds',
    method: 'delete',
    data
  })
}

// @Tags HkUserExtract
// @Summary 更新HkUserExtract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserExtract true "更新HkUserExtract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserExtract/updateHkUserExtract [put]
export const updateHkUserExtract = (data) => {
  return service({
    url: '/hkUserExtract/updateHkUserExtract',
    method: 'put',
    data
  })
}

// @Tags HkUserExtract
// @Summary 用id查询HkUserExtract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkUserExtract true "用id查询HkUserExtract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserExtract/findHkUserExtract [get]
export const findHkUserExtract = (params) => {
  return service({
    url: '/hkUserExtract/findHkUserExtract',
    method: 'get',
    params
  })
}

// @Tags HkUserExtract
// @Summary 分页获取HkUserExtract列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkUserExtract列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserExtract/getHkUserExtractList [get]
export const getHkUserExtractList = (params) => {
  return service({
    url: '/hkUserExtract/getHkUserExtractList',
    method: 'get',
    params
  })
}
