let server_url =
	process.env.NODE_ENV === 'development' ?
	'http://127.0.0.1:8001' :
	'http://127.0.0.1:8008' //环境配置更新
function service(options = {}) {
	options.url = `${server_url}${options.url}`
	//配置请求头
	let token = uni.getStorageSync('token')
	if (token != '') {
		options.header = {
			Authorization: `Bearer ${token}` //Bearer
		}
	}
	return new Promise((resolved, rejected) => {
		//成功
		options.success = res => {
			if (Number(res.data.code) == 200) {
				//请求成功
				resolved(res.data)
			} else if (Number(res.data.code) == 401) {
				uni.removeStorageSync('token')
				uni.redirectTo({
					url: '/pages/login/login'
				})
			} else {
				uni.showToast({
					icon: 'none',
					duration: 3000,
					title: `${res.data.msg}`
				})
				rejected(res.data.msg) //错误
			}
		}
		//错误
		options.fail = err => {
			rejected(err) //错误
		}
		uni.request(options)
	})
}
export default service