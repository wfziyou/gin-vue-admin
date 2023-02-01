import service from '@/utils/request'

// @Tags HkCircle
// @Summary 创建HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircle true "创建HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircle/createHkCircle [post]
export const createHkCircle = (data) => {
  return service({
    url: '/hkCircle/createHkCircle',
    method: 'post',
    data
  })
}

// @Tags HkCircle
// @Summary 删除HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircle true "删除HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircle/deleteHkCircle [delete]
export const deleteHkCircle = (data) => {
  return service({
    url: '/hkCircle/deleteHkCircle',
    method: 'delete',
    data
  })
}

// @Tags HkCircle
// @Summary 删除HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircle/deleteHkCircle [delete]
export const deleteHkCircleByIds = (data) => {
  return service({
    url: '/hkCircle/deleteHkCircleByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCircle
// @Summary 更新HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircle true "更新HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircle/updateHkCircle [put]
export const updateHkCircle = (data) => {
  return service({
    url: '/hkCircle/updateHkCircle',
    method: 'put',
    data
  })
}

// @Tags HkCircle
// @Summary 用id查询HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCircle true "用id查询HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircle/findHkCircle [get]
export const findHkCircle = (params) => {
  return service({
    url: '/hkCircle/findHkCircle',
    method: 'get',
    params
  })
}

// @Tags HkCircle
// @Summary 分页获取HkCircle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCircle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircle/getHkCircleList [get]
export const getHkCircleList = (params) => {
  return service({
    url: '/hkCircle/getHkCircleList',
    method: 'get',
    params
  })
}
