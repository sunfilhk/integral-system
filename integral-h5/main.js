import Vue from 'vue'
import App from './App'
import {
	parseTime,
	moneyFormat
} from './common/costum.js'
Vue.prototype.parseTime = parseTime
Vue.prototype.moneyFormat = moneyFormat

// // #ifdef H5
// // 提交前需要注释  本地调试使用
// import Vconsole from 'vconsole'
// let vConsole = new Vconsole();
// export default vConsole
// // #endif

Vue.config.productionTip = false
App.mpType = 'app'

const app = new Vue({
	...App
})
app.$mount()
