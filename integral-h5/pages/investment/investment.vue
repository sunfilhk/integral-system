<template>
	<view class="investment">
			<button :disabled='flag==0' class="newinvestment" size="mini" type="primary" @click="handleAdd">新增投资</button>
			<view class="tab" v-for="(item,index) in list" :key="index">
				<view>金额：{{item.amount}}</view>
				<view>利润比例：{{item.profit*100+'%'}}</view>
				<view>修改时间：{{parseTime(item.update_time)}}</view>
				<view class="status">状态：
					<text class="a" v-show="item.status==0">已通过</text>
					<text class="b" v-show="item.status==1">审核中</text>
					<text class="c" v-show="item.status==3">已拒绝</text>
					<text class="b" v-show="item.status==4">终止审核</text>
					<text class="c" v-show="item.status==5">已终止</text>
					<text class="c" v-show="item.status==6">已终止</text>
				</view>
				<view class="btm" v-if="item.flag1==1">
					<view><button :class="{active:flag==0||item.status==1||item.status==3||item.status==4||item.status==5||item.status==6}" :disabled='flag==0||item.status==1||item.status==3||item.status==4||item.status==5||item.status==6' size="mini" type="default" @click="handleRevoke(item.id)">撤销投资</button></view>
				</view>
				<view class="btn" v-if="item.flag1==1">
					<view><button :class="{active:flag==0||item.alterbool==1||item.status==1||item.status==4||item.status==5||item.status==6}" :disabled='flag==0||item.alterbool==1||item.status==1||item.status==4||item.status==5||item.status==6' size="mini" type="default" @click="handleAddid(item.id)">修改投资</button></view>
					<view><button :class="{active:flag==0||item.status==1||item.status==3||item.status==4||item.status==5||item.status==6}" :disabled='flag==0||item.status==1||item.status==3||item.status==4||item.status==5||item.status==6' size="mini" type="default" @click="handleStop(item.id)">终止投资</button></view>
				</view>
			</view>
			<view v-show="isLoadMore">
			    <uni-load-more :status="loadStatus" ></uni-load-more>
			</view>
	</view>
</template>

<script>
	import { investmentList,investmentBreak,investmentRevoke } from '../../common/api.js';
	export default {
		data() {
			return {
				flag:'',
				list:[],
				queryParams: {
				    pageIndex: 1,
				    pageSize: 10,
				    customerid:''
				},
				loadStatus:'loading',  //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				isLoadMore:false,  //是否加载中
			}
		},
		onLoad(option) {
			this.queryParams.customerid=option.id
			this.flag=option.flag
		},
		onShow() {
			this.list=[]
			this.auto();		
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
				investmentList(this.queryParams).then(res=> {
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
			},
			handleAdd() {
				uni.navigateTo({url:'/pages/ininvestment/ininvestment?customerid='+this.queryParams.customerid})
			},
			handleAddid(id) {
				uni.navigateTo({url:'/pages/ininvestment/ininvestment?id='+id})
			},
			handleStop(id) {
				var params={
					id:id,
					customerid:this.queryParams.customerid
				}
				investmentBreak(params).then(res=> {
					uni.showToast({icon: 'success',title: '操作成功'});
					this.list=[]
					this.auto();
				})
			},
			handleRevoke(id) {
				var that=this
				var params={
					id:id,
					customerid:this.queryParams.customerid
				}
				uni.showModal({
				    title: '提示',
					content: '确定撤销这笔投资吗？',
					success: function (res) {
						if (res.confirm) {
							investmentRevoke(params).then(res=> {
								uni.showToast({icon: 'success',title: '操作成功'});
								that.list=[]
								that.auto();
							})
					    } else if (res.cancel) {
						    // console.log('用户点击取消');
					    }
				    }
				});
			}
		}
	}
</script>

<style>
    .newinvestment {
		margin: 20px 20px 10px;
	}
	.tab {
		font-size: 14px;
		color: #444444;
		border-top: 1px solid #dddddd;
		padding: 11px 20px 15px;
		position: relative;
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
	.btn {
		position: absolute;
		right: 10px;
		top: 15px;
	}
	.btn button::after{ border: none;}
	.btn button {
		background: #ffffff;
		color: #007AFF;
	}
	.btm {
		position: absolute;
		right: 140px;
		bottom: 6px;
	}
	.btm button::after{ border: none;}
	.btm button {
		background-color:transparent;
		color: rgb(255, 80, 80);
	}
	.active {
		color: #aaaaaa !important;
	}
</style>
