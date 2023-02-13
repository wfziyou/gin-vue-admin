import service from '@/utils/request'

// @Tags HkForumTopicPostsMapping
// @Summary 创建HkForumTopicPostsMapping
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTopicPostsMapping true "创建HkForumTopicPostsMapping"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopicPostsMapping/createHkForumTopicPostsMapping [post]
export const createHkForumTopicPostsMapping = (data) => {
  return service({
    url: '/hkForumTopicPostsMapping/createHkForumTopicPostsMapping',
    method: 'post',
    data
  })
}

// @Tags HkForumTopicPostsMapping
// @Summary 删除HkForumTopicPostsMapping
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTopicPostsMapping true "删除HkForumTopicPostsMapping"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTopicPostsMapping/deleteHkForumTopicPostsMapping [delete]
export const deleteHkForumTopicPostsMapping = (data) => {
  return service({
    url: '/hkForumTopicPostsMapping/deleteHkForumTopicPostsMapping',
    method: 'delete',
    data
  })
}

// @Tags HkForumTopicPostsMapping
// @Summary 删除HkForumTopicPostsMapping
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumTopicPostsMapping"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTopicPostsMapping/deleteHkForumTopicPostsMapping [delete]
export const deleteHkForumTopicPostsMappingByIds = (data) => {
  return service({
    url: '/hkForumTopicPostsMapping/deleteHkForumTopicPostsMappingByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForumTopicPostsMapping
// @Summary 更新HkForumTopicPostsMapping
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTopicPostsMapping true "更新HkForumTopicPostsMapping"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumTopicPostsMapping/updateHkForumTopicPostsMapping [put]
export const updateHkForumTopicPostsMapping = (data) => {
  return service({
    url: '/hkForumTopicPostsMapping/updateHkForumTopicPostsMapping',
    method: 'put',
    data
  })
}

// @Tags HkForumTopicPostsMapping
// @Summary 用id查询HkForumTopicPostsMapping
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForumTopicPostsMapping true "用id查询HkForumTopicPostsMapping"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumTopicPostsMapping/findHkForumTopicPostsMapping [get]
export const findHkForumTopicPostsMapping = (params) => {
  return service({
    url: '/hkForumTopicPostsMapping/findHkForumTopicPostsMapping',
    method: 'get',
    params
  })
}

// @Tags HkForumTopicPostsMapping
// @Summary 分页获取HkForumTopicPostsMapping列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForumTopicPostsMapping列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopicPostsMapping/getHkForumTopicPostsMappingList [get]
export const getHkForumTopicPostsMappingList = (params) => {
  return service({
    url: '/hkForumTopicPostsMapping/getHkForumTopicPostsMappingList',
    method: 'get',
    params
  })
}
