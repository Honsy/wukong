import request from '@/utils/request'

// 获取GB设备
export function getDeviceList(params) {
  return request({
    url: '/api/v1/gb28181/devicelist',
    method: 'get',
    params
  })
}

// 获取GB设备子设备
export function getSubDeviceList(params) {
  return request({
    url: '/api/v1/gb28181/subdevicelist',
    method: 'get',
    params
  })
}