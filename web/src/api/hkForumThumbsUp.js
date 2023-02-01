import service from '@/utils/request'

// @Tags HkForumThumbsUp
// @Summary 创建HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumThumbsUp true "创建HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumThumbsUp/createHkForumThumbsUp [post]
export const createHkForumThumbsUp = (data) => {
  return service({
    url: '/hkForumThumbsUp/createHkForumThumbsUp',
    method: 'post',
    data
  })
}

// @Tags HkForumThumbsUp
// @Summary 删除HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumThumbsUp true "删除HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumThumbsUp/deleteHkForumThumbsUp [delete]
export const deleteHkForumThumbsUp = (data) => {
  return service({
    url: '/hkForumThumbsUp/deleteHkForumThumbsUp',
    method: 'delete',
    data
  })
}

// @Tags HkForumThumbsUp
// @Summary 删除HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumThumbsUp/deleteHkForumThumbsUp [delete]
export const deleteHkForumThumbsUpByIds = (data) => {
  return service({
    url: '/hkForumThumbsUp/deleteHkForumThumbsUpByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForumThumbsUp
// @Summary 更新HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumThumbsUp true "更新HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumThumbsUp/updateHkForumThumbsUp [put]
export const updateHkForumThumbsUp = (data) => {
  return service({
    url: '/hkForumThumbsUp/updateHkForumThumbsUp',
    method: 'put',
    data
  })
}

// @Tags HkForumThumbsUp
// @Summary 用id查询HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForumThumbsUp true "用id查询HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumThumbsUp/findHkForumThumbsUp [get]
export const findHkForumThumbsUp = (params) => {
  return service({
    url: '/hkForumThumbsUp/findHkForumThumbsUp',
    method: 'get',
    params
  })
}

// @Tags HkForumThumbsUp
// @Summary 分页获取HkForumThumbsUp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForumThumbsUp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumThumbsUp/getHkForumThumbsUpList [get]
export const getHkForumThumbsUpList = (params) => {
  return service({
    url: '/hkForumThumbsUp/getHkForumThumbsUpList',
    method: 'get',
    params
  })
}
