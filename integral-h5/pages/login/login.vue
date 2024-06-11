<template>
	<view class="login">
	    <view style="height: 140px;background: #51516c;"></view>
		<view style="height: 60px;background: #51516c;border-radius: 50% / 0 0 100% 100%;"></view>
		<uni-forms style="margin: 70px 50px 0;" ref="form" :modelValue="loginForm" :rules="rules">
			<uni-forms-item name="username">
				<uni-easyinput type="text" v-model="loginForm.username" placeholder="用户名" />
			</uni-forms-item>
			<uni-forms-item name="password">
				<uni-easyinput type="password" v-model="loginForm.password" maxlength="18" placeholder="密码" />
			</uni-forms-item>
			<uni-forms-item name="code">
				<uni-easyinput style=" width: 58%;float: left;" type="text" v-model="loginForm.code" placeholder="验证码" />
				<view style=" width: 38%;height: 36px;float: right;"><image style="height: 36px;width: 100%;" :src="codeUrl" mode="" @click="auto"></image></view>
			</uni-forms-item>
		</uni-forms>
		<button style="margin: 0 50px;background: #51516c;" type="primary" @click="submit">登录</button>
		<view class="logo"><image src="../../static/zy.png" mode=""></image></view>
	</view>
</template>

<script>
	import { getCaptcha,login } from '../../common/api.js';
	export default {
		data() {
			return {
				codeUrl:'',
				loginForm: {
				    username: '',
				    password: '',
				    rememberMe: false,
				    code: '',
				    uuid: ''
				},
				rules: {
					username: {
						rules: [
							{required: true,errorMessage: '请输入用户名',},
						]
					},
					password: {
						rules: [
							{required: true,errorMessage: '请输入密码',},
							{minLength: 6,errorMessage: '密码长度不能小于{minLength}个字符',}
						]
					},
					code: {
						rules: [
							{required: true,errorMessage: '请输入验证码',},
						]
					},
				}
			}
		},
		onLoad() {
			this.auto();
		},
		methods: {
			auto() {
				getCaptcha().then(res => {
                    this.codeUrl = res.data
                    this.loginForm.uuid = res.id
				})
			},
			submit() {
				this.$refs.form.validate().then(valid => {
					login(this.loginForm).then(res=> {
						uni.setStorageSync('token',res.token);
						uni.switchTab({url:'/pages/index/index'});
					})
				})
			}
		}
	}
</script>

<style>
	.logo {
		position: absolute;
		top: 140px;
		left: 50%;
		width: 100px;
		height: 100px;
		background: #d2b56f;
		border-radius: 16px;
		text-align: center;
		transform: translateX(-50px);
	}
	.logo image {
		width: 68px;
		height: 90px;
		margin-top: 5px;
	}
</style>
