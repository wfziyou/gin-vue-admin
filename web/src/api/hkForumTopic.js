import service from '@/utils/request'

// @Tags HkForumTopic
// @Summary 创建HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTopic true "创建HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopic/createHkForumTopic [post]
export const createHkForumTopic = (data) => {
  return service({
    url: '/hkForumTopic/createHkForumTopic',
    method: 'post',
    data
  })
}

// @Tags HkForumTopic
// @Summary 删除HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTopic true "删除HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTopic/deleteHkForumTopic [delete]
export const deleteHkForumTopic = (data) => {
  return service({
    url: '/hkForumTopic/deleteHkForumTopic',
    method: 'delete',
    data
  })
}

// @Tags HkForumTopic
// @Summary 删除HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTopic/deleteHkForumTopic [delete]
export const deleteHkForumTopicByIds = (data) => {
  return service({
    url: '/hkForumTopic/deleteHkForumTopicByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForumTopic
// @Summary 更新HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTopic true "更新HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumTopic/updateHkForumTopic [put]
export const updateHkForumTopic = (data) => {
  return service({
    url: '/hkForumTopic/updateHkForumTopic',
    method: 'put',
    data
  })
}

// @Tags HkForumTopic
// @Summary 用id查询HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForumTopic true "用id查询HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumTopic/findHkForumTopic [get]
export const findHkForumTopic = (params) => {
  return service({
    url: '/hkForumTopic/findHkForumTopic',
    method: 'get',
    params
  })
}

// @Tags HkForumTopic
// @Summary 分页获取HkForumTopic列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForumTopic列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopic/getHkForumTopicList [get]
export const getHkForumTopicList = (params) => {
  return service({
    url: '/hkForumTopic/getHkForumTopicList',
    method: 'get',
    params
  })
}
