<template>
	<view class="ininvestment">
		<uni-forms ref="form" :modelValue="formData" :rules="rules">
			<uni-forms-item label="金额" name="amount">
				<uni-easyinput type="text" v-model="formData.amount" placeholder="请输入金额" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="利润比例" name="profit">
				<uni-easyinput type="text" disabled v-model="formData.profit" placeholder="请输入利润比例" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="日期">
				<picker class="sexs" mode="date" :value="date" @change="bindDateChange">
					<view class="uni-input">{{date}}</view>
				</picker>
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="时间">
				<picker class="sexs" mode="time" :value="time" @change="bindTimeChange">
					<view class="uni-input">{{time}}</view>
				</picker>
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="备注" name="remark">
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
		getInvestmentByid,
		investmentAdd,
		investmentEdit
	} from '../../common/api.js';
	export default {
		data() {
			return {
				date: getDate({
					format: true
				}),
				time: '12:01',
				formData: {
					amount: undefined,
					profit: undefined,
					remark: undefined,
				},
				customerid: '',
				rules: {
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
			} else {
				this.customerid = option.customerid
				this.getConfigKey()
			}
		},
		methods: {
			auto(id) {
				getInvestmentByid({
					id: id
				}).then(res => {
					this.formData = res.data
					this.formData.profit = this.formData.profit * 100 + '%'
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
						investmentEdit(this.formData).then(res => {
							uni.navigateBack({
								delta: 1
							});
						})
					} else {
						this.formData.customerid = this.customerid
						investmentAdd(this.formData).then(res => {
							uni.navigateBack({
								delta: 1
							});
						})
					}
				})
			},
		}
	}
</script>

<style>
	.ininvestment {
		padding: 20px 20px 20px 26px;
	}

	.ininvestment button {
		height: 40px;
		font-size: 16px;
	}

	.xh {
		position: absolute;
		top: 8px;
		left: -10px;
		color: red;
		font-size: 18px;
	}

	.sexs {
		height: 36px;
		line-height: 36px;
		border: 1px solid #e5e5e5;
		border-radius: 6px;
		padding-left: 10px;
	}
</style>
