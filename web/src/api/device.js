import request from '@/utils/request'

// 获取GB设备
export function getDeviceList(params) {
  return request({
    url: '/api/v1/gb28181/deviceList',
    method: 'get',
    params
  })
}

// 获取GB设备子设备
export function getSubDeviceList(params) {
  return request({
    url: '/api/v1/gb28181/subDeviceList',
    method: 'get',
    params
  })
}

// 播放GB设备子设备
export function playCameraWithDeviceId(data) {
  return request({
    url: '/api/v1/gb28181/playVideo',
    method: 'POST',
    data
  })
}