import service from '@/utils/request'

// @Tags HkMiniProgram
// @Summary 创建HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkMiniProgram true "创建HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkMiniProgram/createHkMiniProgram [post]
export const createHkMiniProgram = (data) => {
  return service({
    url: '/hkMiniProgram/createHkMiniProgram',
    method: 'post',
    data
  })
}

// @Tags HkMiniProgram
// @Summary 删除HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkMiniProgram true "删除HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkMiniProgram/deleteHkMiniProgram [delete]
export const deleteHkMiniProgram = (data) => {
  return service({
    url: '/hkMiniProgram/deleteHkMiniProgram',
    method: 'delete',
    data
  })
}

// @Tags HkMiniProgram
// @Summary 删除HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkMiniProgram/deleteHkMiniProgram [delete]
export const deleteHkMiniProgramByIds = (data) => {
  return service({
    url: '/hkMiniProgram/deleteHkMiniProgramByIds',
    method: 'delete',
    data
  })
}

// @Tags HkMiniProgram
// @Summary 更新HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkMiniProgram true "更新HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkMiniProgram/updateHkMiniProgram [put]
export const updateHkMiniProgram = (data) => {
  return service({
    url: '/hkMiniProgram/updateHkMiniProgram',
    method: 'put',
    data
  })
}

// @Tags HkMiniProgram
// @Summary 用id查询HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkMiniProgram true "用id查询HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkMiniProgram/findHkMiniProgram [get]
export const findHkMiniProgram = (params) => {
  return service({
    url: '/hkMiniProgram/findHkMiniProgram',
    method: 'get',
    params
  })
}

// @Tags HkMiniProgram
// @Summary 分页获取HkMiniProgram列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkMiniProgram列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkMiniProgram/getHkMiniProgramList [get]
export const getHkMiniProgramList = (params) => {
  return service({
    url: '/hkMiniProgram/getHkMiniProgramList',
    method: 'get',
    params
  })
}
