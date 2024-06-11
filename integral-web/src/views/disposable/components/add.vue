<template>
  <div class="add">
    <el-dialog
      :title="{ add: '添加用户', edit: '修改用户' }[mode]"
      :visible.sync="show"
      width="400px"
      @close="closeEvent"
    >
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="团队类型: " prop="deptId">
          <el-select
            v-model="form.teamname"
            clearable
            size="small"
            placeholder="请选择团队类型"
            style="height: 30px"
          >
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="form.username"
            placeholder="请输入姓名/用户名/手机号查询"
            style="width: 250px"
            @input="chang"
          />
        </el-form-item>
        <el-form-item label="百分比" prop="percent">
          <el-input
            v-model="form.percent"
            clearable
            placeholder="请输入百分比(如：5%)"
            oninput="value=value.match(/\d+\.?\d{0,2}/,'')"
            style="width: 250px"
          />
          %
        </el-form-item>
      </el-form>
      <div v-show="isShow" class="small">
        <div v-for="(item, index) in data" :key="index" @click="choice(item)">
          {{ item.nick_name }} ({{ item.username }})
        </div>
      </div>
      <template #footer>
        <div>
          <el-button type="primary" @click="submit">确定</el-button>
          <el-button @click="show = false">取消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  profitconfigOnce,
  updateProfitConfigOnce,
  getonceinto,
  userList
} from '@/api/profitconfig/profitconfig'

export default {
  components: {},
  props: {},
  data() {
    return {
      mode: '',
      // 控制弹框显示
      show: false,
      // 显示用户搜索列表
      isShow: false,
      // 列表
      data: [],
      // 团队选项
      options: [],
      form: {
        teamname: '',
        username: '',
        percent: ''
      },
      rules: {
        username: [
          { required: true, message: '用户名不能为空', trigger: 'blur' }
        ],
        percent: [
          { required: true, message: '百分比不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getData()
  },
  mounted() {},
  methods: {
    // 获取团队
    async getData() {
      const res = await getonceinto()
      this.options = res.list.map(v => {
        return {
          value: v.teamname, // 团队id
          label: v.teamname // 团队名称
        }
      })
    },
    // 点击确定
    submit() {
      let forms = {}
      forms = { ...this.form }
      forms.percent = forms.percent * 0.01
      this.$refs['form'].validate(async result => {
        if (result) {
          if (this.mode === 'add') {
            await profitconfigOnce(forms)
            // 提示
            this.$message.success('新增成功')
          } else if (this.mode === 'edit') {
            await updateProfitConfigOnce(forms)
            // 提示
            this.$message.success('修改成功')
          }
          // 关闭弹框
          this.show = false
          // 刷新数据
          this.$emit('auto')
        }
      })
    },
    // 清空数据与验证
    closeEvent() {
      this.$refs.form.resetFields()
      this.form = {
        deptId: '',
        username: '',
        percent: ''
      }
    },
    // 获取用户列表
    chang() {
      if (this.form.username.length != '') {
        userList({ keyword: this.form.username }).then(res => {
          this.data = res.list
          if (this.data.length > 0) {
            this.isShow = true
          } else {
            this.isShow = false
          }
        })
      } else {
        this.data = []
        this.isShow = false
      }
    },
    choice(item) {
      this.form.username = item.username
      this.form.nick_name = item.nick_name
      this.form.userid = item.userid
      this.isShow = false
    }
  }
}
</script>
<style scoped lang="scss">
.small {
  position: absolute;
  top: 177px;
  left: 100px;
  width: 250px;
  height: 258px;
  border: 1px solid rgb(0, 119, 255);
  background-color: #fff;
  border-radius: 4px;
  overflow: hidden;
  overflow-y: auto;
}
.small div {
  padding: 8px 20px;
}
.small div:hover {
  background-color: #eeeeee;
}
</style>
