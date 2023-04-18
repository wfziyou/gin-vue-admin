import service from '@/utils/request'

// @Tags HkRecordBrowsingCircleHomepage
// @Summary 创建HkRecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkRecordBrowsingCircleHomepage true "创建HkRecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkRecordBrowsingCircleHomepage/createHkRecordBrowsingCircleHomepage [post]
export const createHkRecordBrowsingCircleHomepage = (data) => {
  return service({
    url: '/hkRecordBrowsingCircleHomepage/createHkRecordBrowsingCircleHomepage',
    method: 'post',
    data
  })
}

// @Tags HkRecordBrowsingCircleHomepage
// @Summary 删除HkRecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkRecordBrowsingCircleHomepage true "删除HkRecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkRecordBrowsingCircleHomepage/deleteHkRecordBrowsingCircleHomepage [delete]
export const deleteHkRecordBrowsingCircleHomepage = (data) => {
  return service({
    url: '/hkRecordBrowsingCircleHomepage/deleteHkRecordBrowsingCircleHomepage',
    method: 'delete',
    data
  })
}

// @Tags HkRecordBrowsingCircleHomepage
// @Summary 删除HkRecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkRecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkRecordBrowsingCircleHomepage/deleteHkRecordBrowsingCircleHomepage [delete]
export const deleteHkRecordBrowsingCircleHomepageByIds = (data) => {
  return service({
    url: '/hkRecordBrowsingCircleHomepage/deleteHkRecordBrowsingCircleHomepageByIds',
    method: 'delete',
    data
  })
}

// @Tags HkRecordBrowsingCircleHomepage
// @Summary 更新HkRecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkRecordBrowsingCircleHomepage true "更新HkRecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkRecordBrowsingCircleHomepage/updateHkRecordBrowsingCircleHomepage [put]
export const updateHkRecordBrowsingCircleHomepage = (data) => {
  return service({
    url: '/hkRecordBrowsingCircleHomepage/updateHkRecordBrowsingCircleHomepage',
    method: 'put',
    data
  })
}

// @Tags HkRecordBrowsingCircleHomepage
// @Summary 用id查询HkRecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkRecordBrowsingCircleHomepage true "用id查询HkRecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkRecordBrowsingCircleHomepage/findHkRecordBrowsingCircleHomepage [get]
export const findHkRecordBrowsingCircleHomepage = (params) => {
  return service({
    url: '/hkRecordBrowsingCircleHomepage/findHkRecordBrowsingCircleHomepage',
    method: 'get',
    params
  })
}

// @Tags HkRecordBrowsingCircleHomepage
// @Summary 分页获取HkRecordBrowsingCircleHomepage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkRecordBrowsingCircleHomepage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkRecordBrowsingCircleHomepage/getHkRecordBrowsingCircleHomepageList [get]
export const getHkRecordBrowsingCircleHomepageList = (params) => {
  return service({
    url: '/hkRecordBrowsingCircleHomepage/getHkRecordBrowsingCircleHomepageList',
    method: 'get',
    params
  })
}
