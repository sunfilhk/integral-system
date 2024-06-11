import request from '@/utils/request'

export function customerAdd(query) {
  return request({
    url: '/api/v1/customerAdd',
    method: 'post',
    params: query
  })
}

export function customerList(query) {
  return request({
    url: '/api/v1/customerList',
    method: 'get',
    params: query
  })
}
export function customerEdit(query) {
  return request({
    url: '/api/v1/customerEdit',
    method: 'post',
    params: query
  })
}

export function customerLogList(query) {
  return request({
    url: '/api/v1/customerLogList',
    method: 'get',
    params: query
  })
}
export function customerAuditList(query) {
  return request({
    url: '/api/v1/customerAuditList',
    method: 'get',
    params: query
  })
}
export function customerLogAudit(query) {
  return request({
    url: '/api/v1/customerLogAudit',
    method: 'post',
    params: query
  })
}
export function individualPerformance(query) {
  return request({
    url: '/api/v1/individualPerformance',
    method: 'get',
    params: query
  })
}
export function customerProfitEdit(query) {
  return request({
    url: '/api/v1/customerProfitEdit',
    method: 'post',
    params: query
  })
}

export function getConfigKey(query) {
  return request({
    url: '/api/v1/getConfigKey',
    method: 'get',
    params: query
  })
}

// 新增接口利润
export function getProfitcustomer(query) {
  return request({
    url: '/api/v1/customerbili',
    method: 'post',
    params: query
  })
}
export function updateConfig(query) {
  return request({
    url: '/api/v1/updateConfig',
    method: 'post',
    params: query
  })
}
export function investmentExport(query) {
  return request({
    url: '/api/v1/investmentExport',
    method: 'get',
    params: query
  })
}
export function emptyInvestment(query) {
  return request({
    url: '/api/v1/emptyInvestment',
    method: 'post',
    params: query
  })
}

// 修改客户利润百分比
export function updateProfitcustomer(data) {
  return request({
    url: '/api/v1/updateProfitcustomer',
    method: 'post',
    params: data
  })
}
// 业务管理
export function investListss(data) {
  return request({
    url: '/api/v1/investListss',
    method: 'get',
    params: data
  })
}

// 修改投资
export function investmentEdit(data) {
  return request({
    url: '/api/v1/investmentEdit',
    method: 'post',
    params: data
  })
}
// 终止
export function investmentBreak(data) {
  return request({
    url: '/api/v1/investmentBreak',
    method: 'post',
    params: data
  })
}
// 撤销
export function investmentRevoke(data) {
  return request({
    url: '/api/v1/investmentRevoke',
    method: 'post',
    params: data
  })
}

// 合同到期时间
export function searchsexpire(data) {
  return request({
    url: '/api/v1/searchsexpire',
    method: 'get',
    params: data
  })
}

// 获取新增投资选择客户
export function customerLists(query) {
  return request({
    url: '/api/v1/customerLists',
    method: 'get',
    params: query
  })
}
