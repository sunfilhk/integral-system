<template>
	<view class="mine">
		<view class="img">
			<view><image src="https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80" mode=""></image></view>
			<view style="margin-bottom: 10px;font-size: 18px;">{{data.nick_name}}</view>
		    <button style="margin-bottom: 10px;" size="mini" type="warn" @click="logut">退出登录</button>
			<view class="title">
				<view>
					<view>等级</view>
				    <view>
						<image v-if="star1" src="../../static/star.png" mode=""></image>
						<image v-if="star2" src="../../static/star.png" mode=""></image>
						<image v-if="star3" src="../../static/star.png" mode=""></image>
						<image v-if="star4" src="../../static/star.png" mode=""></image>
						<image v-if="star5" src="../../static/star.png" mode=""></image>
						<image v-if="star6" src="../../static/star.png" mode=""></image>
					</view>
				</view>
				<view>
					<view>称号</view>
				    <view style="color: red;">{{data.vip_title}}</view>
				</view>
			</view>
			<view class="title">
				<view>
					<view>当日业绩</view>
					<view>{{moneyFormat(Number(data.today_performance))}}</view>
				</view>
				<view>
					<view>总业绩</view>
					<view>{{moneyFormat(Number(data.total_performance))}}</view>
				</view>
			</view>
			<view class="title">
				<view>
					<view>当日业务奖励</view>
					<view>{{moneyFormat(Number(datas.today))}}</view>
				</view>
				<view>
					<view>昨日业务奖励</view>
					<view>{{moneyFormat(Number(datas.yesterday))}}</view>
				</view>
			</view>
			<view class="title">
				<view style="border-bottom: none;">
					<view>累计业务奖励</view>
					<view>{{moneyFormat(Number(datas.total))}}</view>
				</view>
				<view style="border-bottom: none;"></view>
			</view>
		</view>
		<view class="menu" @click="salesman">
			下级业务员审核
			<view><image src="../../static/arrow-right.png" mode=""></image></view>
		</view>
		<view class="menu" @click="subordinate">
			下级业务员
			<view><image src="../../static/arrow-right.png" mode=""></image></view>
		</view>
		<view class="menu" @click="audituser">
			业务员审核记录
			<view><image src="../../static/arrow-right.png" mode=""></image></view>
		</view>
		<view class="menu" style="border-bottom: 1px solid #cccccc;margin-bottom: 20px;" @click="pass">
			修改密码
			<view><image src="../../static/arrow-right.png" mode=""></image></view>
		</view>
	</view>
</template>

<script>
	import { getUserPerformance,getUserProfits } from '../../common/api.js';
	export default {
		data() {
			return {
				data:{},
				datas:{},
				star1:false,
				star2:false,
				star3:false,
				star4:false,
				star5:false,
				star6:false,
			}
		},
		onLoad() {
			this.auto()
		},
		onPullDownRefresh() {
			this.auto()
		    uni.stopPullDownRefresh();
		},
		methods: {
			auto() {
				getUserPerformance().then(res=> {
					this.data=res.data
					if (res.data.vip_level=='V1') {
						this.star1=true
					}else if (res.data.vip_level=='V2') {
						this.star1=true
						this.star2=true
					}else if (res.data.vip_level=='V3') {
						this.star1=true
						this.star2=true
						this.star3=true
					}else if (res.data.vip_level=='V4') {
						this.star1=true
						this.star2=true
						this.star3=true
						this.star4=true
					}else if (res.data.vip_level=='V5') {
						this.star1=true
						this.star2=true
						this.star3=true
						this.star4=true
						this.star5=true
					}else if (res.data.vip_level=='V6') {
						this.star1=true
						this.star2=true
						this.star3=true
						this.star4=true
						this.star5=true
						this.star6=true
					}
				})
				getUserProfits().then(res=> {
					this.datas=res.data
				})
			},
			logut() {
				uni.removeStorageSync('token');
				uni.reLaunch({url:'/pages/login/login'});
			},
			pass() {
				uni.navigateTo({url:'/pages/pass/pass'})
			},
			salesman() {
				uni.navigateTo({url:'/pages/salesman/salesman'})
			},
			audituser() {
				uni.navigateTo({url:'/pages/audituser/audituser'})
			},
			subordinate() {
				uni.navigateTo({url:'/pages/subordinate/subordinate'})
			}
		}
	}
</script>

<style>
.img {
	text-align: center;
	padding: 20px;
}
.img image{
	width: 100px;
	height: 100px;
	border-radius: 50px;
}
.title {
	display: flex;
}
.title image {
	width: 24px;
	height: 24px;
}
.title>view {
	flex: 1;
	padding: 10px 0;
}
.title>view:nth-child(1) {
	border-right: 1px solid #cccccc;
	border-bottom: 1px solid #cccccc;
}
.title>view:nth-child(2) {
	border-bottom: 1px solid #cccccc;
}
.title>view view:nth-child(1) {
	font-size: 14px;
	color: #666666;
}
.title>view view:nth-child(2) {
	font-size: 17px;
	margin-top: 4px;
}
.menu {
	padding: 20px 20px;
	border-top: 1px solid #cccccc;
	position: relative;
	font-size: 15px;
}
.menu view {
	position: absolute;
	top: 20px;
	right: 20px;
}
.menu view image {
	width: 22px;
	height: 22px;
}
</style>
