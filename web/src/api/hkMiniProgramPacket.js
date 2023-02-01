import service from '@/utils/request'

// @Tags HkMiniProgramPacket
// @Summary 创建HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkMiniProgramPacket true "创建HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkMiniProgramPacket/createHkMiniProgramPacket [post]
export const createHkMiniProgramPacket = (data) => {
  return service({
    url: '/hkMiniProgramPacket/createHkMiniProgramPacket',
    method: 'post',
    data
  })
}

// @Tags HkMiniProgramPacket
// @Summary 删除HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkMiniProgramPacket true "删除HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkMiniProgramPacket/deleteHkMiniProgramPacket [delete]
export const deleteHkMiniProgramPacket = (data) => {
  return service({
    url: '/hkMiniProgramPacket/deleteHkMiniProgramPacket',
    method: 'delete',
    data
  })
}

// @Tags HkMiniProgramPacket
// @Summary 删除HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkMiniProgramPacket/deleteHkMiniProgramPacket [delete]
export const deleteHkMiniProgramPacketByIds = (data) => {
  return service({
    url: '/hkMiniProgramPacket/deleteHkMiniProgramPacketByIds',
    method: 'delete',
    data
  })
}

// @Tags HkMiniProgramPacket
// @Summary 更新HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkMiniProgramPacket true "更新HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkMiniProgramPacket/updateHkMiniProgramPacket [put]
export const updateHkMiniProgramPacket = (data) => {
  return service({
    url: '/hkMiniProgramPacket/updateHkMiniProgramPacket',
    method: 'put',
    data
  })
}

// @Tags HkMiniProgramPacket
// @Summary 用id查询HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkMiniProgramPacket true "用id查询HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkMiniProgramPacket/findHkMiniProgramPacket [get]
export const findHkMiniProgramPacket = (params) => {
  return service({
    url: '/hkMiniProgramPacket/findHkMiniProgramPacket',
    method: 'get',
    params
  })
}

// @Tags HkMiniProgramPacket
// @Summary 分页获取HkMiniProgramPacket列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkMiniProgramPacket列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkMiniProgramPacket/getHkMiniProgramPacketList [get]
export const getHkMiniProgramPacketList = (params) => {
  return service({
    url: '/hkMiniProgramPacket/getHkMiniProgramPacketList',
    method: 'get',
    params
  })
}
