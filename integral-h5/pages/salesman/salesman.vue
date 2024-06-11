<template>
	<view class="salesman">
		<button class="newsalesman" size="mini" type="primary" @click="handleAdd">新增下级</button>
		<view class="tab" v-for="(item,index) in list" :key="index">
			<view>姓名：{{item.nick_name}}</view>
			<view>推荐人：{{item.referrer_name}}</view>
			<view>创建时间：{{parseTime(item.create_time)}}</view>
			<view class="btn" v-if="item.flag==1">
				<view><button size="mini" type="default" @click="adopt(item.user_id)">通过</button></view>
				<view><button size="mini" type="default" @click="refuse(item.user_id)">拒绝</button></view>
			</view>
		</view>
		<view v-show="isLoadMore">
		    <uni-load-more :status="loadStatus" ></uni-load-more>
		</view>
	</view>
</template>

<script>
	import { passingSysUserList,submitNewUser,allowNewUserPass } from '../../common/api.js';
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
		onShow() {
			this.list=[]
			this.auto()
		},
		methods: {
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
			auto() {
				passingSysUserList(this.queryParams).then(res=> {
					if(res.data.list){
					    this.list=this.list.concat(res.data.list)
					    if(res.data.list.length<this.queryParams.pageSize){  //判断接口返回数据量小于请求数据量，则表示此为最后一页
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
			adopt(id) {
				var params={
				        user_id:id,
				        is_allow:1
				    }
				allowNewUserPass(params).then(res=> {
					uni.showToast({icon: 'success',title: '操作成功'});
					this.queryParams.pageIndex=1
				    this.list=[]
				    this.auto();
				})
			},
			refuse(id) {
				var params={
				        user_id:id,
				        is_allow:1
				    }
				allowNewUserPass(params).then(res=> {
					uni.showToast({icon: 'success',title: '操作成功'});
					this.queryParams.pageIndex=1
				    this.list=[]
				    this.auto();
				})
			},
			handleAdd() {
				uni.navigateTo({url:'/pages/level/level'})
			}
		}
	}
</script>

<style>
.newsalesman {
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
	.btn {
		position: absolute;
		right: 10px;
		top: 10px;
	}
	.btn button::after{ border: none;}
	.btn button {
		background: #ffffff;
		color: #007AFF;
	}
</style>
