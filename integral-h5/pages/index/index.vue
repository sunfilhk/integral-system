<template>
	<view class="index">
		<image src="../../static/zyyc.png" mode=""></image>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				
			}
		},
		onLoad() {
			
		},
		onPullDownRefresh() {
		    uni.stopPullDownRefresh();
		},
		methods: {
			
		}
	}
</script>

<style>
	.index{
	        position: absolute;
	        top: 0;
	        left: 0;
	        right: 0;
	        bottom: 0;
	        display: flex;
	        justify-content: center;
	        align-items: center;
	    }
		image {
			width: 200px;
			height: 270px;
		}
</style>
