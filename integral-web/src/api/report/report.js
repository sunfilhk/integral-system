import request from '@/utils/request'

export function statementSalesman(query) {
  return request({
    url: '/api/v1/statementSalesman',
    method: 'get',
    params: query
  })
}
export function statementCustomer(query) {
  return request({
    url: '/api/v1/statementCustomer',
    method: 'get',
    params: query
  })
}
export function statementConfigOnce(query) {
  return request({
    url: '/api/v1/statementConfigOnce',
    method: 'get',
    params: query
  })
}
export function statementSummary(query) {
  return request({
    url: '/api/v1/statementSummary',
    method: 'get',
    params: query
  })
}
// 导出业务报表
export function statementSalesmanExport(query) {
  return request({
    url: '/api/v1/statementSalesmanExport',
    method: 'get',
    params: query
  })
}
// 客户报表
export function statementCustomerExport(query) {
  return request({
    url: '/api/v1/statementCustomerExport',
    method: 'get',
    params: query
  })
}
export function statementConfigOnceExport(query) {
  return request({
    url: '/api/v1/statementConfigOnceExport',
    method: 'get',
    params: query
  })
}
export function statementSummaryExport(query) {
  return request({
    url: '/api/v1/statementSummaryExport',
    method: 'get',
    params: query
  })
}
export function statementCustomerSettle(query) {
  return request({
    url: '/api/v1/statementCustomerSettle',
    method: 'post',
    params: query
  })
}
export function statementSettleHistory(query) {
  return request({
    url: '/api/v1/statementSettleHistory',
    method: 'get',
    params: query
  })
}
export function statementNoSettlement(query) {
  return request({
    url: '/api/v1/statementNoSettlement',
    method: 'get',
    params: query
  })
}
export function statementIsSettlement(query) {
  return request({
    url: '/api/v1/statementIsSettlement',
    method: 'get',
    params: query
  })
}
export function statementSettlement(query) {
  return request({
    url: '/api/v1/statementSettlement',
    method: 'post',
    params: query
  })
}
// 查询明细
export function summrydetail(query) {
  return request({
    url: '/api/v1/summrydetail',
    method: 'get',
    params: query
  })
}
