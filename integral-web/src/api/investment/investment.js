import request from '@/utils/request'

export function investmentAdd(query) {
  return request({
    url: '/api/v1/investmentAdd',
    method: 'post',
    params: query
  })
}
// 更换新增接口
export function newinvestmentAdd(query) {
  return request({
    url: '/api/v1/newinvestmentAdd',
    method: 'post',
    params: query
  })
}
export function investmentList(query) {
  return request({
    url: '/api/v1/investmentList',
    method: 'get',
    params: query
  })
}
export function investmentEdit(query) {
  return request({
    url: '/api/v1/investmentEdit',
    method: 'post',
    params: query
  })
}

export function investmentBreak(query) {
  return request({
    url: '/api/v1/investmentBreak',
    method: 'post',
    params: query
  })
}

export function investmentRevoke(query) {
  return request({
    url: '/api/v1/investmentRevoke',
    method: 'post',
    params: query
  })
}
