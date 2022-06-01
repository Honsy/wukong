import request from '@/utils/request'

export function getDeviceList(params) {
  return request({
    url: '/api/v1/gb28181/devicelist',
    method: 'get',
    params
  })
}

