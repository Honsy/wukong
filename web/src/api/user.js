import request from '@/utils/request'

// 登录
export function login(data) {
  return request({
    url: '/api/v1/user/login',
    method: 'post',
    data
  })
}

// 注册
export function registerApi(data) {
  return request({
    url: '/api/v1/user/register',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/vue-admin-template/user/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/vue-admin-template/user/logout',
    method: 'post'
  })
}
