import service from '@/utils/request'

// @Tags HkForumTopicGroup
// @Summary 创建HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTopicGroup true "创建HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopicGroup/createHkForumTopicGroup [post]
export const createHkForumTopicGroup = (data) => {
  return service({
    url: '/hkForumTopicGroup/createHkForumTopicGroup',
    method: 'post',
    data
  })
}

// @Tags HkForumTopicGroup
// @Summary 删除HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTopicGroup true "删除HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTopicGroup/deleteHkForumTopicGroup [delete]
export const deleteHkForumTopicGroup = (data) => {
  return service({
    url: '/hkForumTopicGroup/deleteHkForumTopicGroup',
    method: 'delete',
    data
  })
}

// @Tags HkForumTopicGroup
// @Summary 删除HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTopicGroup/deleteHkForumTopicGroup [delete]
export const deleteHkForumTopicGroupByIds = (data) => {
  return service({
    url: '/hkForumTopicGroup/deleteHkForumTopicGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForumTopicGroup
// @Summary 更新HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTopicGroup true "更新HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumTopicGroup/updateHkForumTopicGroup [put]
export const updateHkForumTopicGroup = (data) => {
  return service({
    url: '/hkForumTopicGroup/updateHkForumTopicGroup',
    method: 'put',
    data
  })
}

// @Tags HkForumTopicGroup
// @Summary 用id查询HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForumTopicGroup true "用id查询HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumTopicGroup/findHkForumTopicGroup [get]
export const findHkForumTopicGroup = (params) => {
  return service({
    url: '/hkForumTopicGroup/findHkForumTopicGroup',
    method: 'get',
    params
  })
}

// @Tags HkForumTopicGroup
// @Summary 分页获取HkForumTopicGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForumTopicGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopicGroup/getHkForumTopicGroupList [get]
export const getHkForumTopicGroupList = (params) => {
  return service({
    url: '/hkForumTopicGroup/getHkForumTopicGroupList',
    method: 'get',
    params
  })
}
