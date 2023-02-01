import service from '@/utils/request'

// @Tags HkForumComment
// @Summary 创建HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumComment true "创建HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumComment/createHkForumComment [post]
export const createHkForumComment = (data) => {
  return service({
    url: '/hkForumComment/createHkForumComment',
    method: 'post',
    data
  })
}

// @Tags HkForumComment
// @Summary 删除HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumComment true "删除HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumComment/deleteHkForumComment [delete]
export const deleteHkForumComment = (data) => {
  return service({
    url: '/hkForumComment/deleteHkForumComment',
    method: 'delete',
    data
  })
}

// @Tags HkForumComment
// @Summary 删除HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumComment/deleteHkForumComment [delete]
export const deleteHkForumCommentByIds = (data) => {
  return service({
    url: '/hkForumComment/deleteHkForumCommentByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForumComment
// @Summary 更新HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumComment true "更新HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumComment/updateHkForumComment [put]
export const updateHkForumComment = (data) => {
  return service({
    url: '/hkForumComment/updateHkForumComment',
    method: 'put',
    data
  })
}

// @Tags HkForumComment
// @Summary 用id查询HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForumComment true "用id查询HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumComment/findHkForumComment [get]
export const findHkForumComment = (params) => {
  return service({
    url: '/hkForumComment/findHkForumComment',
    method: 'get',
    params
  })
}

// @Tags HkForumComment
// @Summary 分页获取HkForumComment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForumComment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumComment/getHkForumCommentList [get]
export const getHkForumCommentList = (params) => {
  return service({
    url: '/hkForumComment/getHkForumCommentList',
    method: 'get',
    params
  })
}
