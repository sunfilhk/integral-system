import request from '@/utils/request'

export function profitconfigOnce(query) {
  return request({
    url: '/api/v1/profitconfigOnce',
    method: 'post',
    params: query
  })
}
// 获取新增/修改弹框团队
export function getonceinto(query) {
  return request({
    url: '/api/v1/getonceinto',
    method: 'get',
    params: query
  })
}

// 一次性分配列表
export function profitconfigList(query) {
  return request({
    url: '/api/v1/profitconfigList',
    method: 'get',
    params: query
  })
}

// 客户分配列表
export function profitList(query) {
  return request({
    url: '/api/v1/getProfitcustomers',
    method: 'get',
    params: query
  })
}
export function profitconfigEdit(data) {
  return request({
    url: '/api/v1/profitconfigEdit',
    method: 'post',
    data
  })
}

export function userList(query) {
  return request({
    url: '/api/v1/userList',
    method: 'get',
    params: query
  })
}
// 获取团队类型列表
export function teamList(query) {
  return request({
    url: '/api/v1/teamList',
    method: 'get',
    params: query
  })
}

export function updateProfitConfigOnce(query) {
  return request({
    url: '/api/v1/updateProfitConfigOnce',
    method: 'post',
    params: query
  })
}
export function delProfitconfigOnce(query) {
  return request({
    url: '/api/v1/delProfitconfigOnce',
    method: 'post',
    params: query
  })
}
// export function sysUserList(query) {
//   return request({
//     url: '/api/v1/sysUserList',
//     method: 'get',
//     params: query
//   })
// }
export function getUserVipLevel(query) {
  return request({
    url: '/api/v1/getUserVipLevel',
    method: 'get',
    params: query
  })
}
export function editUserVipLevel(query) {
  return request({
    url: '/api/v1/editUserVipLevel',
    method: 'post',
    params: query
  })
}
export function editlevellock(query) {
  return request({
    url: '/api/v1/editlevellock',
    method: 'post',
    params: query
  })
}
