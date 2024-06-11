import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function refreshtoken(data) {
  return request({
    url: '/refreshtoken',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/api/v1/getinfo',
    method: 'get'
  })
}

export function logout(data) {
  return request({
    url: '/api/v1/logout',
    method: 'post',
    data
  })
}

export function getUserPerformance() {
  return request({
    url: '/api/v1/getUserPerformance',
    method: 'get'
  })
}
// 获取利润
export function getUserProfits() {
  return request({
    url: '/api/v1/getUserProfits',
    method: 'get'
  })
}
export function getUserReferrals(query) {
  return request({
    url: '/api/v1/getUserReferrals',
    method: 'get',
    params: query
  })
}
