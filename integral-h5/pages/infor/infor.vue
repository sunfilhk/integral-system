<template>
	<view class="infor">
		<view class="tab">
			<view>客户姓名：{{data.name}}</view>
			<view>性别：{{data.sex}}</view>
			<view>手机号：{{data.phone}}</view>
			<view>身份证号：{{data.identity}}</view>
			<view>银行卡号：{{data.bank}}</view>
			<view>开户行：{{data.banknum}}</view>
			<view>业务员姓名：{{data.nick_name}}</view>
			<view>创建时间：{{ parseTime(data.create_time)}}</view>
			<view class="status">状态：
				<text class="a" v-show="data.status==0">已通过</text>
				<text class="b" v-show="data.status==1">审核中</text>
				<text class="c" v-show="data.status==3">已拒绝</text>
			</view>
			<view class="btn">
				<view><button :disabled="data.status==1||flag==0" type="default" @click="customerid(data.id)">修改信息</button></view>
			</view>
		</view>
	</view>
</template>

<script>
	import { getCustomerByid } from '../../common/api.js';
	export default {
		data() {
			return {
				id:'',
				flag:'',
				data:{}
			}
		},
		onLoad(option) {
			this.id=option.id
			this.flag=option.flag
		},
		onShow() {
			this.auto();		
		},
		methods: {
			auto() {
				getCustomerByid({customerid:this.id}).then(res=> {
					this.data=res.data
				})
			},
			customerid(id) {
				uni.navigateTo({url:'/pages/incustomer/incustomer?id='+id})
			},
		}
	}
</script>

<style>
.infor {
	padding: 20px;
}
.tab {
		font-size: 15px;
		color: #444444;
	}
	.tab view {
		margin-top: 4px;
	}
	.tab .status .a {
		color: #00bc00;
	}
	.tab .status .b {
		color: #007AFF;
	}
	.tab .status .c {
		color: rgb(255, 80, 80);
	}
	.btn view {
		margin: 14px 0;
	}
	.btn button {
		height: 40px;
		font-size: 16px;
	}
</style>
