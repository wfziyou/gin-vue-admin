import service from '@/utils/request'

// @Tags HkRecordBrowsingUserHomepage
// @Summary 创建HkRecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkRecordBrowsingUserHomepage true "创建HkRecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkRecordBrowsingUserHomepage/createHkRecordBrowsingUserHomepage [post]
export const createHkRecordBrowsingUserHomepage = (data) => {
  return service({
    url: '/hkRecordBrowsingUserHomepage/createHkRecordBrowsingUserHomepage',
    method: 'post',
    data
  })
}

// @Tags HkRecordBrowsingUserHomepage
// @Summary 删除HkRecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkRecordBrowsingUserHomepage true "删除HkRecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkRecordBrowsingUserHomepage/deleteHkRecordBrowsingUserHomepage [delete]
export const deleteHkRecordBrowsingUserHomepage = (data) => {
  return service({
    url: '/hkRecordBrowsingUserHomepage/deleteHkRecordBrowsingUserHomepage',
    method: 'delete',
    data
  })
}

// @Tags HkRecordBrowsingUserHomepage
// @Summary 删除HkRecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkRecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkRecordBrowsingUserHomepage/deleteHkRecordBrowsingUserHomepage [delete]
export const deleteHkRecordBrowsingUserHomepageByIds = (data) => {
  return service({
    url: '/hkRecordBrowsingUserHomepage/deleteHkRecordBrowsingUserHomepageByIds',
    method: 'delete',
    data
  })
}

// @Tags HkRecordBrowsingUserHomepage
// @Summary 更新HkRecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkRecordBrowsingUserHomepage true "更新HkRecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkRecordBrowsingUserHomepage/updateHkRecordBrowsingUserHomepage [put]
export const updateHkRecordBrowsingUserHomepage = (data) => {
  return service({
    url: '/hkRecordBrowsingUserHomepage/updateHkRecordBrowsingUserHomepage',
    method: 'put',
    data
  })
}

// @Tags HkRecordBrowsingUserHomepage
// @Summary 用id查询HkRecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkRecordBrowsingUserHomepage true "用id查询HkRecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkRecordBrowsingUserHomepage/findHkRecordBrowsingUserHomepage [get]
export const findHkRecordBrowsingUserHomepage = (params) => {
  return service({
    url: '/hkRecordBrowsingUserHomepage/findHkRecordBrowsingUserHomepage',
    method: 'get',
    params
  })
}

// @Tags HkRecordBrowsingUserHomepage
// @Summary 分页获取HkRecordBrowsingUserHomepage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkRecordBrowsingUserHomepage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkRecordBrowsingUserHomepage/getHkRecordBrowsingUserHomepageList [get]
export const getHkRecordBrowsingUserHomepageList = (params) => {
  return service({
    url: '/hkRecordBrowsingUserHomepage/getHkRecordBrowsingUserHomepageList',
    method: 'get',
    params
  })
}
