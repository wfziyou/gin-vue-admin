import service from '@/utils/request'

// @Tags HkForumPosts
// @Summary 创建HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumPosts true "创建HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumPosts/createHkForumPosts [post]
export const createHkForumPosts = (data) => {
  return service({
    url: '/hkForumPosts/createHkForumPosts',
    method: 'post',
    data
  })
}

// @Tags HkForumPosts
// @Summary 删除HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumPosts true "删除HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumPosts/deleteHkForumPosts [delete]
export const deleteHkForumPosts = (data) => {
  return service({
    url: '/hkForumPosts/deleteHkForumPosts',
    method: 'delete',
    data
  })
}

// @Tags HkForumPosts
// @Summary 删除HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumPosts/deleteHkForumPosts [delete]
export const deleteHkForumPostsByIds = (data) => {
  return service({
    url: '/hkForumPosts/deleteHkForumPostsByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForumPosts
// @Summary 更新HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumPosts true "更新HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumPosts/updateHkForumPosts [put]
export const updateHkForumPosts = (data) => {
  return service({
    url: '/hkForumPosts/updateHkForumPosts',
    method: 'put',
    data
  })
}

// @Tags HkForumPosts
// @Summary 用id查询HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForumPosts true "用id查询HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumPosts/findHkForumPosts [get]
export const findHkForumPosts = (params) => {
  return service({
    url: '/hkForumPosts/findHkForumPosts',
    method: 'get',
    params
  })
}

// @Tags HkForumPosts
// @Summary 分页获取HkForumPosts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForumPosts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumPosts/getHkForumPostsList [get]
export const getHkForumPostsList = (params) => {
  return service({
    url: '/hkForumPosts/getHkForumPostsList',
    method: 'get',
    params
  })
}
