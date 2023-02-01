import service from '@/utils/request'

// @Tags HkForumTag
// @Summary 创建HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTag true "创建HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTag/createHkForumTag [post]
export const createHkForumTag = (data) => {
  return service({
    url: '/hkForumTag/createHkForumTag',
    method: 'post',
    data
  })
}

// @Tags HkForumTag
// @Summary 删除HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTag true "删除HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTag/deleteHkForumTag [delete]
export const deleteHkForumTag = (data) => {
  return service({
    url: '/hkForumTag/deleteHkForumTag',
    method: 'delete',
    data
  })
}

// @Tags HkForumTag
// @Summary 删除HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTag/deleteHkForumTag [delete]
export const deleteHkForumTagByIds = (data) => {
  return service({
    url: '/hkForumTag/deleteHkForumTagByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForumTag
// @Summary 更新HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTag true "更新HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumTag/updateHkForumTag [put]
export const updateHkForumTag = (data) => {
  return service({
    url: '/hkForumTag/updateHkForumTag',
    method: 'put',
    data
  })
}

// @Tags HkForumTag
// @Summary 用id查询HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForumTag true "用id查询HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumTag/findHkForumTag [get]
export const findHkForumTag = (params) => {
  return service({
    url: '/hkForumTag/findHkForumTag',
    method: 'get',
    params
  })
}

// @Tags HkForumTag
// @Summary 分页获取HkForumTag列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForumTag列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTag/getHkForumTagList [get]
export const getHkForumTagList = (params) => {
  return service({
    url: '/hkForumTag/getHkForumTagList',
    method: 'get',
    params
  })
}
