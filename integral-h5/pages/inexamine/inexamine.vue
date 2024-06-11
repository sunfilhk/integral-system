<template>
	<view class="inexamine">
		<view class="tab">
			<view>客户姓名：{{data.name}}</view>
			<view>性别：{{data.sex}}</view>
			<view>手机号：{{data.phone}}</view>
			<view>身份证号：{{data.identity}}</view>
			<view>银行卡号：{{data.bank}}</view>
			<view>开户行：{{data.banknum}}</view>
			<view>业务员姓名：{{data.nick_name}}</view>
			<view>金额：{{data.amount}}</view>
			<view>备注：{{data.remark}}</view>
			<view>创建时间：{{ parseTime(data.create_time)}}</view>
			<view>审核类型：
				<text v-show="data.edittype==0">新增客户审核</text>
				<text v-show="data.edittype==1">修改客户审核</text>
				<text v-show="data.edittype==2">新增投资审核</text>
				<text v-show="data.edittype==3">修改投资审核</text>
				<text v-show="data.edittype==4">终止投资审核</text>
			</view>
			<view class="btn" v-if="data.flag==1">
				<view><button type="primary" @click="agree(data.id)">同意</button></view>
				<view><button type="default" @click="refuse(data.id)">驳回</button></view>
			</view>
		</view>
	</view>
</template>

<script>
	import { customerLog,customerLogAudit } from '../../common/api.js';
	export default {
		data() {
			return {
				data:{}
			}
		},
		onLoad(option) {
			this.auto(option.id);
		},
		methods: {
			auto(id) {
				customerLog({id}).then(res=> {
					this.data=res.data
				})
			},
			agree(id) {
				var params={
					id:id,
					status: '0'
				}
				customerLogAudit(params).then(res=> {
					uni.switchTab({url:'/pages/examines/examines'});
				})
			},
			refuse(id) {
				var params={
					id:id,
					status: '3'
				}
				customerLogAudit(params).then(res=> {
					uni.switchTab({url:'/pages/examines/examines'});
				})
			},
		}
	}
</script>

<style>
    .inexamine {
	    padding: 20px;
    }
    .tab {
		font-size: 15px;
		color: #444444;
	}
	.tab view {
		margin-top: 4px;
	}
	.btn view {
		margin: 14px 0;
	}
	.btn button {
		height: 40px;
		font-size: 16px;
	}
</style>
