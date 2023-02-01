import service from '@/utils/request'

// @Tags HkForumTagGroup
// @Summary 创建HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTagGroup true "创建HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTagGroup/createHkForumTagGroup [post]
export const createHkForumTagGroup = (data) => {
  return service({
    url: '/hkForumTagGroup/createHkForumTagGroup',
    method: 'post',
    data
  })
}

// @Tags HkForumTagGroup
// @Summary 删除HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTagGroup true "删除HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTagGroup/deleteHkForumTagGroup [delete]
export const deleteHkForumTagGroup = (data) => {
  return service({
    url: '/hkForumTagGroup/deleteHkForumTagGroup',
    method: 'delete',
    data
  })
}

// @Tags HkForumTagGroup
// @Summary 删除HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTagGroup/deleteHkForumTagGroup [delete]
export const deleteHkForumTagGroupByIds = (data) => {
  return service({
    url: '/hkForumTagGroup/deleteHkForumTagGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForumTagGroup
// @Summary 更新HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForumTagGroup true "更新HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumTagGroup/updateHkForumTagGroup [put]
export const updateHkForumTagGroup = (data) => {
  return service({
    url: '/hkForumTagGroup/updateHkForumTagGroup',
    method: 'put',
    data
  })
}

// @Tags HkForumTagGroup
// @Summary 用id查询HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForumTagGroup true "用id查询HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumTagGroup/findHkForumTagGroup [get]
export const findHkForumTagGroup = (params) => {
  return service({
    url: '/hkForumTagGroup/findHkForumTagGroup',
    method: 'get',
    params
  })
}

// @Tags HkForumTagGroup
// @Summary 分页获取HkForumTagGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForumTagGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTagGroup/getHkForumTagGroupList [get]
export const getHkForumTagGroupList = (params) => {
  return service({
    url: '/hkForumTagGroup/getHkForumTagGroupList',
    method: 'get',
    params
  })
}
