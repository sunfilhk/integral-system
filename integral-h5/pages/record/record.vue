<template>
	<view class="record">
		<view class="tab" v-for="(item,index) in list" :key="index">
			<view>客户姓名：{{item.name}}</view>
			<view>性别：{{item.sex}}</view>
			<view>手机号：{{item.phone}}</view>
			<view>身份证号：{{item.identity}}</view>
			<view>银行卡号：{{item.bank}}</view>
			<view>开户行：{{item.banknum}}</view>
			<view>金额：{{item.amount}}</view>
			<view>备注：{{item.remark}}</view>
			<view>创建时间：{{ parseTime(item.create_time)}}</view>
			<view>审核类型：
				<text v-show="item.edittype==0">新增客户审核</text>
				<text v-show="item.edittype==1">修改客户审核</text>
				<text v-show="item.edittype==2">新增投资审核</text>
				<text v-show="item.edittype==3">修改投资审核</text>
				<text v-show="item.edittype==4">终止投资审核</text>
			</view>
			<view class="status">状态：
				<text class="a" v-show="item.status==0">已通过</text>
				<text class="c" v-show="item.status==3">已拒绝</text>
			</view>
		</view>
		<view v-show="isLoadMore">
		    <uni-load-more :status="loadStatus" ></uni-load-more>
		</view>
	</view>
</template>

<script>
	import { customerLogList } from '../../common/api.js';
	export default {
		data() {
			return {
				list:[],
				queryParams: {
				    pageIndex: 1,
				    pageSize: 10,
				},
				loadStatus:'loading',  //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				isLoadMore:false,  //是否加载中
			}
		},
		onLoad() {
			this.auto()
		},
		onReachBottom(){  //上拉触底函数
		    if(!this.isLoadMore){  //此处判断，上锁，防止重复请求
		        this.isLoadMore=true
		        this.queryParams.pageIndex+=1
		        this.auto()
		    }
		},
		onPullDownRefresh() {
			this.list=[]
			this.queryParams.pageIndex=1
			this.auto()
		    uni.stopPullDownRefresh();
		},
		methods: {
			auto() {
				customerLogList(this.queryParams).then(res=> {
					if(res.list){
					    this.list=this.list.concat(res.list)
					    if(res.list.length<this.queryParams.pageSize){  //判断接口返回数据量小于请求数据量，则表示此为最后一页
					        this.isLoadMore=true                                             
					        this.loadStatus='nomore'
					    }else{
					        this.isLoadMore=false
					    }
					}else{
					    this.isLoadMore=true
					    this.loadStatus='nomore'
					}
				})
			}
		}
	}
</script>

<style>
    .tab {
		font-size: 14px;
		color: #444444;
		border-top: 1px solid #dddddd;
		padding: 6px 20px 10px;
		position: relative;
	}
	.tab view {
		margin-top: 4px;
	}
	.tab .status .a {
		color: #00bc00;
	}
	.tab .status .c {
		color: rgb(255, 80, 80);
	}
</style>
