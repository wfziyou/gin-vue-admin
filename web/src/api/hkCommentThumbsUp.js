import service from '@/utils/request'

// @Tags HkCommentThumbsUp
// @Summary 创建HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCommentThumbsUp true "创建HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCommentThumbsUp/createHkCommentThumbsUp [post]
export const createHkCommentThumbsUp = (data) => {
  return service({
    url: '/hkCommentThumbsUp/createHkCommentThumbsUp',
    method: 'post',
    data
  })
}

// @Tags HkCommentThumbsUp
// @Summary 删除HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCommentThumbsUp true "删除HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCommentThumbsUp/deleteHkCommentThumbsUp [delete]
export const deleteHkCommentThumbsUp = (data) => {
  return service({
    url: '/hkCommentThumbsUp/deleteHkCommentThumbsUp',
    method: 'delete',
    data
  })
}

// @Tags HkCommentThumbsUp
// @Summary 删除HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCommentThumbsUp/deleteHkCommentThumbsUp [delete]
export const deleteHkCommentThumbsUpByIds = (data) => {
  return service({
    url: '/hkCommentThumbsUp/deleteHkCommentThumbsUpByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCommentThumbsUp
// @Summary 更新HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCommentThumbsUp true "更新HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCommentThumbsUp/updateHkCommentThumbsUp [put]
export const updateHkCommentThumbsUp = (data) => {
  return service({
    url: '/hkCommentThumbsUp/updateHkCommentThumbsUp',
    method: 'put',
    data
  })
}

// @Tags HkCommentThumbsUp
// @Summary 用id查询HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCommentThumbsUp true "用id查询HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCommentThumbsUp/findHkCommentThumbsUp [get]
export const findHkCommentThumbsUp = (params) => {
  return service({
    url: '/hkCommentThumbsUp/findHkCommentThumbsUp',
    method: 'get',
    params
  })
}

// @Tags HkCommentThumbsUp
// @Summary 分页获取HkCommentThumbsUp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCommentThumbsUp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCommentThumbsUp/getHkCommentThumbsUpList [get]
export const getHkCommentThumbsUpList = (params) => {
  return service({
    url: '/hkCommentThumbsUp/getHkCommentThumbsUpList',
    method: 'get',
    params
  })
}
