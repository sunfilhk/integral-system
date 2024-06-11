<template>
	<view class="level">
		<uni-forms ref="form" :modelValue="formData" :rules="rules">
			<uni-forms-item label="姓名" name="nickName">
				<uni-easyinput type="text" v-model="formData.nickName" placeholder="请输入姓名" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="手机号" name="phone">
				<uni-easyinput type="text" v-model="formData.phone" maxlength="11" placeholder="请输入手机号" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="邮箱" name="email">
				<uni-easyinput type="text" v-model="formData.email" placeholder="请输入邮箱" />
			</uni-forms-item>
			<uni-forms-item label="用户名" name="username">
				<uni-easyinput type="text" v-model="formData.username" placeholder="请输入用户名" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="用户密码" name="password">
				<uni-easyinput type="password" v-model="formData.password" maxlength="18" placeholder="请输入用户密码" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="银行户名" name="bank_user_name">
				<uni-easyinput type="text" v-model="formData.bank_user_name" placeholder="请输入银行户名" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="银行卡号" name="bank_num">
				<uni-easyinput type="text" v-model="formData.bank_num" minlength="10" placeholder="请输入银行卡号" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="开户行" name="open_bank">
				<uni-easyinput type="text" v-model="formData.open_bank" placeholder="请输入开户行" />
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="客户性别" name="sex">
				<picker class="sexs" @change="bindSex" :value="formData.sex" :range="array">
					<view>{{array[formData.sex]}}</view>
				</picker>
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="角色">
				<picker class="sexs" @change="bindRole" :value="roleindex" :range="rolearr">
					<view>{{rolearr[roleindex]}}</view>
				</picker>
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="岗位">
				<picker class="sexs" @change="bindPost" :value="postindex" :range="postarr">
					<view>{{postarr[postindex]}}</view>
				</picker>
				<text class="xh">*</text>
			</uni-forms-item>
			<uni-forms-item label="归属部门">
				<picker class="sexs" @change="bindDept" :value="deptindex" :range="deptarr">
					<view>{{deptarr[deptindex]}}</view>
				</picker>
				<text class="xh">*</text>
			</uni-forms-item>
		</uni-forms>
		<button type="primary" @click="submit">提交</button>
	</view>
</template>

<script>
	import {
		getUserInit,
		deptTree,
		submitNewUser
	} from '../../common/api.js';
	export default {
		data() {
			return {
				formData: {
					nickName: undefined,
					phone: undefined,
					email: undefined,
					username: undefined,
					password: undefined,
					bank_user_name: undefined,
					bank_num: undefined,
					open_bank: undefined,
					sex: '0',
					status: '0',
				},
				array: ['男', '女'],
				posts: [],
				roles: [],
				postarr: [],
				postindex: '1',
				rolearr: [],
				roleindex: '0',
				depts: [],
				deptarr: [],
				deptindex: '0',
				rules: {
					nickName: {
						rules: [{
							required: true,
							errorMessage: '请输入姓名',
						}, ]
					},
					phone: {
						rules: [{
								required: true,
								errorMessage: '请输入手机号',
							},
							{
								pattern: /^1[3|4|5|7|8][0-9]\d{8}$/,
								errorMessage: '请输入正确的手机号码',
							}
						]
					},
					username: {
						rules: [{
							required: true,
							errorMessage: '请输入用户名',
						}, ]
					},
					password: {
						rules: [{
								required: true,
								errorMessage: '请输入用户密码',
							},
							{
								minLength: 6,
								errorMessage: '密码长度不能小于{minLength}个字符',
							}
						]
					},
					bank_user_name: {
						rules: [{
							required: true,
							errorMessage: '请输入用户密码',
						}, ]
					},
					bank_num: {
						rules: [{
								required: true,
								errorMessage: '请输入银行卡',
							},
							// 限制银行卡号最低输入10个数
							{
								minLength: 10,
								errorMessage: '请输入正确的银行卡号',
							}
							// {
							// 	pattern: /^([1-9]{1})(\d{15}|\d{18})$/,
							// 	errorMessage: '请输入正确的银行卡号',
							// }
						]
					},
					open_bank: {
						rules: [{
							required: true,
							errorMessage: '请输入开户行',
						}, ]
					}
				}
			}
		},
		onLoad() {
			this.auto()
		},
		methods: {
			auto() {
				getUserInit().then(res => {
					this.posts = res.data.posts
					this.roles = res.data.roles
					res.data.posts.forEach(item => {
						this.postarr.push(item.postName)
					})
					res.data.roles.forEach(item => {
						this.rolearr.push(item.roleName)
					})
					this.formData.roleId = this.roles[this.roleindex].roleId
					this.formData.postId = this.posts[this.postindex].postId
				})
				deptTree().then(res => {
					// this.depts = res.data[0].children
					this.depts = res.data
					this.depts.forEach(item => {
						this.deptarr.push(item.deptname)
					})
					this.formData.deptId = this.depts[this.deptindex].deptId
				})
			},
			bindSex(e) {
				this.formData.sex = e.target.value
			},
			bindRole(e) {
				this.roleindex = e.target.value
				this.formData.roleId = this.roles[e.target.value].roleId
			},
			bindPost(e) {
				this.postindex = e.target.value
				this.formData.postId = this.posts[e.target.value].postId
			},
			bindDept(e) {
				this.deptindex = e.target.value
				this.formData.deptId = this.depts[e.target.value].deptId
			},
			submit() {
				this.formData.sex = this.formData.sex + ''
				this.$refs.form.validate().then(valid => {
					submitNewUser(this.formData).then(response => {
						uni.navigateBack({
							delta: 1
						});
					})
				})
			}
		}
	}
</script>

<style>
	.level {
		padding: 20px 20px 20px 26px;
	}

	.sexs {
		height: 36px;
		line-height: 36px;
		border: 1px solid #e5e5e5;
		border-radius: 6px;
		padding-left: 10px;
	}

	.xh {
		position: absolute;
		top: 8px;
		left: -10px;
		color: red;
		font-size: 18px;
	}
</style>
