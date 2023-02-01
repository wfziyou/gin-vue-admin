import service from '@/utils/request'

// @Tags HkForumPostsGroup
// @Summary 创建HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumPostsGroup true "创建HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumPostsGroup/createHkForumPostsGroup [post]
export const createHkForumPostsGroup = (data) => {
  return service({
    url: '/hkForumPostsGroup/createHkForumPostsGroup',
    method: 'post',
    data
  })
}

// @Tags HkForumPostsGroup
// @Summary 删除HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumPostsGroup true "删除HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumPostsGroup/deleteHkForumPostsGroup [delete]
export const deleteHkForumPostsGroup = (data) => {
  return service({
    url: '/hkForumPostsGroup/deleteHkForumPostsGroup',
    method: 'delete',
    data
  })
}

// @Tags HkForumPostsGroup
// @Summary 删除HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumPostsGroup/deleteHkForumPostsGroup [delete]
export const deleteHkForumPostsGroupByIds = (data) => {
  return service({
    url: '/hkForumPostsGroup/deleteHkForumPostsGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForumPostsGroup
// @Summary 更新HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumPostsGroup true "更新HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumPostsGroup/updateHkForumPostsGroup [put]
export const updateHkForumPostsGroup = (data) => {
  return service({
    url: '/hkForumPostsGroup/updateHkForumPostsGroup',
    method: 'put',
    data
  })
}

// @Tags HkForumPostsGroup
// @Summary 用id查询HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForumPostsGroup true "用id查询HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumPostsGroup/findHkForumPostsGroup [get]
export const findHkForumPostsGroup = (params) => {
  return service({
    url: '/hkForumPostsGroup/findHkForumPostsGroup',
    method: 'get',
    params
  })
}

// @Tags HkForumPostsGroup
// @Summary 分页获取HkForumPostsGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForumPostsGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumPostsGroup/getHkForumPostsGroupList [get]
export const getHkForumPostsGroupList = (params) => {
  return service({
    url: '/hkForumPostsGroup/getHkForumPostsGroupList',
    method: 'get',
    params
  })
}
