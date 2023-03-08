import service from '@/utils/request'

// @Tags HkUserSign
// @Summary 创建HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserSign true "创建HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserSign/createHkUserSign [post]
export const createHkUserSign = (data) => {
  return service({
    url: '/hkUserSign/createHkUserSign',
    method: 'post',
    data
  })
}

// @Tags HkUserSign
// @Summary 删除HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserSign true "删除HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserSign/deleteHkUserSign [delete]
export const deleteHkUserSign = (data) => {
  return service({
    url: '/hkUserSign/deleteHkUserSign',
    method: 'delete',
    data
  })
}

// @Tags HkUserSign
// @Summary 删除HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserSign/deleteHkUserSign [delete]
export const deleteHkUserSignByIds = (data) => {
  return service({
    url: '/hkUserSign/deleteHkUserSignByIds',
    method: 'delete',
    data
  })
}

// @Tags HkUserSign
// @Summary 更新HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserSign true "更新HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserSign/updateHkUserSign [put]
export const updateHkUserSign = (data) => {
  return service({
    url: '/hkUserSign/updateHkUserSign',
    method: 'put',
    data
  })
}

// @Tags HkUserSign
// @Summary 用id查询HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkUserSign true "用id查询HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserSign/findHkUserSign [get]
export const findHkUserSign = (params) => {
  return service({
    url: '/hkUserSign/findHkUserSign',
    method: 'get',
    params
  })
}

// @Tags HkUserSign
// @Summary 分页获取HkUserSign列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkUserSign列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserSign/getHkUserSignList [get]
export const getHkUserSignList = (params) => {
  return service({
    url: '/hkUserSign/getHkUserSignList',
    method: 'get',
    params
  })
}
