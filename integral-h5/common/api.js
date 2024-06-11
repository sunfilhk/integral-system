import request from './request.js'
import qs from 'qs'
export function getCaptcha(data) {
	return request({
		url: '/api/v1/getCaptcha',
		method: 'GET',
		data
	})
}
export function login(data) {
	return request({
		url: '/login',
		method: 'POST',
		data
	})
}
export function userPwd(data) {
	return request({
		url: '/api/v1/user/pwd',
		method: 'PUT',
		data
	})
}
export function customerList(data) {
	return request({
		url: '/api/v1/customerList',
		method: 'GET',
		data
	})
}
export function customerAdd(data) {
	data = qs.stringify(data)
	return request({
		url: '/api/v1/customerAdd?' + data,
		method: 'POST',
	})
}
export function customerEdit(data) {
	data = qs.stringify(data)
	return request({
		url: '/api/v1/customerEdit?' + data,
		method: 'POST',
	})
}
export function investmentList(data) {
	return request({
		url: '/api/v1/investmentList',
		method: 'GET',
		data
	})
}
export function investmentAdd(data) {
	data = qs.stringify(data)
	return request({
		url: '/api/v1/investmentAdd?' + data,
		method: 'POST',
	})
}
export function investmentEdit(data) {
	data = qs.stringify(data)
	return request({
		url: '/api/v1/investmentEdit?' + data,
		method: 'POST',
	})
}
export function investmentBreak(data) {
	data = qs.stringify(data)
	return request({
		url: '/api/v1/investmentBreak?' + data,
		method: 'POST',
	})
}
export function investmentRevoke(data) {
	data = qs.stringify(data)
	return request({
		url: '/api/v1/investmentRevoke?' + data,
		method: 'POST',
	})
}

// export function getConfigKey(data) {
//   return request({
//     url: '/api/v1/getConfigKey',
//     method: 'GET',
//     data
//   })
// }
// 更换接口
export function getConfigKey(data) {
	return request({
		url: '/api/v1/customerbili',
		method: 'POST',
		data
	})
}
export function getCustomerByid(data) {
	return request({
		url: '/api/v1/getCustomerByid',
		method: 'GET',
		data
	})
}
export function getInvestmentByid(data) {
	return request({
		url: '/api/v1/getInvestmentByid',
		method: 'GET',
		data
	})
}
export function customerAuditList(data) {
	return request({
		url: '/api/v1/customerAuditList',
		method: 'GET',
		data
	})
}
export function customerLogAudit(data) {
	data = qs.stringify(data)
	return request({
		url: '/api/v1/customerLogAudit?' + data,
		method: 'POST',
	})
}
export function customerLogList(data) {
	return request({
		url: '/api/v1/customerLogList',
		method: 'GET',
		data
	})
}
export function customerLog(data) {
	return request({
		url: '/api/v1/customerLog',
		method: 'GET',
		data
	})
}
export function getUserPerformance(data) {
	return request({
		url: '/api/v1/getUserPerformance',
		method: 'GET',
		data
	})
}
export function passingSysUserList(data) {
	return request({
		url: '/api/v1/passingSysUserList',
		method: 'get',
		data
	})
}
export function submitNewUser(data) {
	return request({
		url: '/api/v1/submitNewUser',
		method: 'post',
		data
	})
}
export function allowNewUserPass(data) {
	data = qs.stringify(data)
	return request({
		url: '/api/v1/allowNewUserPass?' + data,
		method: 'post',
	})
}
export function getUserInit(data) {
	return request({
		url: '/api/v1/sysUser/',
		method: 'get',
		data
	})
}
export function deptTree(data) {
	return request({
		url: '/api/v1/deptTree',
		method: 'get',
		data
	})
}
export function getUserReferrals(data) {
	return request({
		url: '/api/v1/getUserReferrals',
		method: 'get',
		data
	})
}
export function userAuditRecord(data) {
	return request({
		url: '/api/v1/userAuditRecord',
		method: 'get',
		data
	})
}
export function getUserProfits(data) {
	return request({
		url: '/api/v1/getUserProfits',
		method: 'get',
		data
	})
}
