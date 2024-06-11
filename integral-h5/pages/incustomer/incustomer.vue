<template>
	<view class="incustomer">
		<uni-forms ref="form" :modelValue="formData" :rules="rules">
			<uni-forms-item label="客户姓名" name="name">
				<uni-easyinput type="text" v-model="formData.name" placeholder="请输入客户姓名" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="客户性别" name="sex">
				<picker class="sexs" @change="bindSex" :value="index" :range="array">
					<view>{{array[index]}}</view>
				</picker>
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="手机号" name="phone">
				<uni-easyinput type="text" v-model="formData.phone" maxlength="11" placeholder="请输入手机号" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="身份证" name="identity">
				<uni-easyinput type="text" v-model="formData.identity" maxlength="18" placeholder="请输入身份证" />
			</uni-forms-item>
			<uni-forms-item label="银行卡" name="bank">
				<uni-easyinput type="text" v-model="formData.bank" minlength="10" placeholder="请输入银行卡" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="开户行" name="banknum">
				<uni-easyinput type="text" v-model="formData.banknum" placeholder="请输入开户行" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item v-if="show1" label="金额" name="amount">
				<uni-easyinput type="text" v-model="formData.amount" placeholder="请输入金额" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item v-if="show1" label="利润比例" name="profit">
				<uni-easyinput type="text" disabled v-model="formData.profit" placeholder="请输入利润比例" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item v-if="show1" label="日期">
				<picker class="sexs" mode="date" :value="date" @change="bindDateChange">
					<view class="uni-input">{{date}}</view>
				</picker>
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item v-if="show1" label="时间">
				<picker class="sexs" mode="time" :value="time" @change="bindTimeChange">
					<view class="uni-input">{{time}}</view>
				</picker>
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item v-if="show1" label="备注" name="remark">
				<uni-easyinput type="text" v-model="formData.remark" placeholder="请输入备注" />
			</uni-forms-item>
		</uni-forms>
		<button type="primary" @click="submit">提交</button>
	</view>
</template>

<script>
	function getDate(type) {
		const date = new Date();

		let year = date.getFullYear();
		let month = date.getMonth() + 1;
		let day = date.getDate();

		if (type === 'start') {
			year = year - 10;
		} else if (type === 'end') {
			year = year + 10;
		}
		month = month > 9 ? month : '0' + month;;
		day = day > 9 ? day : '0' + day;

		return `${year}-${month}-${day}`;
	}
	import {
		getConfigKey,
		customerAdd,
		getCustomerByid,
		customerEdit
	} from '../../common/api.js';
	export default {
		data() {
			return {
				show1: false,
				date: getDate({
					format: true
				}),
				time: '12:01',
				formData: {
					name: undefined,
					sex: undefined,
					phone: undefined,
					identity: undefined,
					bank: undefined,
					banknum: undefined,
					amount: undefined,
					profit: undefined,
					remark: undefined,
				},
				array: ['男', '女'],
				index: 0,
				rules: {
					name: {
						rules: [{
							required: true,
							trigger: 'bind',
							errorMessage: '请输入客户姓名'
						}, ]
					},
					phone: {
						rules: [{
								required: true,
								errorMessage: '请输入手机号',
							},
							{
								pattern: /^1[3|4|5|7|8|9][0-9]\d{8}$/,
								errorMessage: '请输入正确的手机号码',
							}
						]
					},
					identity: {
						rules: [
							// {required: true,errorMessage: '请输入身份证',},
							// {
							//     pattern: /^[1-9]\d{5}(18|19|20|(3\d))\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$/,
							//     errorMessage: '请输入正确的身份证号',
							// }
						]
					},
					bank: {
						rules: [{
								required: true,
								errorMessage: '请输入银行卡',
							},
							{
								minLength: 10,
								errorMessage: '请输入正确的银行卡号',
							},
							// {
							// 	pattern: /^([1-9]{1})(\d{15}|\d{18})$/,
							// 	errorMessage: '请输入正确的银行卡号',
							// }
						]
					},
					banknum: {
						rules: [{
							required: true,
							errorMessage: '请输入开户行',
						}, ]
					},
					amount: {
						rules: [{
								required: true,
								errorMessage: '请输入金额',
							},
							{
								pattern: /^[0-9]*[1-9][0-9]*$/,
								errorMessage: '金额为数字且不能为小数',
							}
						]
					},
				}
			}
		},
		onLoad(option) {
			if (option.id != undefined) {
				this.auto(option.id);
				this.show1 = false
			} else {
				this.formData.sex = this.array[this.index]
				this.getConfigKey()
				this.show1 = true
			}
		},
		methods: {
			auto(id) {
				getCustomerByid({
					customerid: id
				}).then(res => {
					this.formData = res.data
					this.index = this.array.indexOf(res.data.sex)
				})
			},
			getConfigKey() {
				getConfigKey({
					configKey: 'customerratio'
				}).then(res => {
					this.formData.profit = parseFloat(res.data.profit).toFixed(5) * 100 + '%'
				})
			},
			bindDateChange: function(e) {
				this.date = e.target.value
			},
			bindTimeChange: function(e) {
				this.time = e.target.value
			},
			submit() {
				this.$refs.form.validate().then(valid => {
					if (this.formData.amount >= 1000000000) {
						uni.showToast({
							icon: 'none',
							title: '金额不能超过十亿'
						});
						return
					}
					this.formData.profit = parseFloat(this.formData.profit.replace("%", "")) * 0.01
					this.formData.date = this.date + ' ' + this.time + ':00'
					if (this.formData.id !== undefined) {
						customerEdit(this.formData).then(res => {
							uni.switchTab({
								url: '/pages/customer/customer'
							});
						})
					} else {
						customerAdd(this.formData).then(res => {
							uni.switchTab({
								url: '/pages/customer/customer'
							});
						})
					}
				})
			},
			bindSex(e) {
				this.index = e.target.value
				this.formData.sex = this.array[this.index]
			}
		}
	}
</script>

<style>
	.incustomer {
		padding: 20px 20px 20px 26px;
	}

	.incustomer button {
		height: 40px;
		font-size: 16px;
	}

	.sexs {
		height: 36px;
		line-height: 36px;
		border: 1px solid #e5e5e5;
		border-radius: 6px;
		padding-left: 10px;
	}

	.xh {
		position: absolute;
		top: 8px;
		left: -10px;
		color: red;
		font-size: 18px;
	}
</style>
