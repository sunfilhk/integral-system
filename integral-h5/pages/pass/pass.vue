<template>
	<view class="pass">
		<uni-forms ref="form" :modelValue="form" :rules="rules">
			<uni-forms-item label="旧密码" name="oldPassword">
				<uni-easyinput type="password" v-model="form.oldPassword" maxlength="18" placeholder="请输入旧密码" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="新密码" name="newPassword">
				<uni-easyinput type="password" v-model="form.newPassword" maxlength="18" placeholder="请输入新密码" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="确认密码" name="password">
				<uni-easyinput type="password" v-model="form.password" maxlength="18" placeholder="请再次输入新密码" />
				<text class="xh">*</text>
			</uni-forms-item>
		</uni-forms>
		<button type="primary" @click="submit">确认</button>
	</view>
</template>

<script>
	import { userPwd } from '../../common/api.js';
	export default {
		data() {
			return {
				form:{
					oldPassword:undefined,
					newPassword:undefined,
					password:undefined
				},
				rules: {
					oldPassword: {
						rules: [
							{required: true,errorMessage: '请输入旧密码',},
							{minLength: 6,maxLength: 18,errorMessage: '密码长度在 {minLength} 到 {maxLength} 个字符',}
						]
					},
					newPassword: {
						rules: [
							{required: true,errorMessage: '请输入新密码',},
							{minLength: 6,maxLength: 18,errorMessage: '密码长度在 {minLength} 到 {maxLength} 个字符',}
						]
					},
					password: {
						rules: [
							{required: true,errorMessage: '请再次输入新密码',},
							{minLength: 6,maxLength: 18,errorMessage: '密码长度在 {minLength} 到 {maxLength} 个字符',}
						]
					},
				}
			}
		},
		methods: {
			submit() { 
				if (this.form.newPassword!=this.form.password) {
					uni.showToast({icon: 'error',title: '密码不一致'});
					return
				}
				userPwd(this.form).then(res=> {
					uni.showToast({icon: 'success',title: '密码修改成功'});
					this.form={
						oldPassword:undefined,
						newPassword:undefined,
						password:undefined
					}
				})
			}
		}
	}
</script>

<style>
.pass {
	padding: 20px 20px 20px 26px;
}
.xh {
		position: absolute;
		top: 8px;
		left: -10px;
		color: red;
		font-size: 18px;
	}
</style>
