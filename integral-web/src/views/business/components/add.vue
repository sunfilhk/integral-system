<template>
  <div class="add">
    <el-dialog
      :visible.sync="isShow"
      :title="{ add: '新增投资', edit: '修改投资' }[mode]"
      width="400px"
      @close="closeEvent"
    >
      <el-form ref="form" label-width="90px" :model="form" :rules="rules">
        <el-form-item v-if="mode === 'add'" prop="name" label="选择客户: ">
          <el-input
            v-model="form.name"
            style="width: 250px"
            placeholder="请输入客户姓名查找用户绑定订单"
            @input="inputChange"
          />
        </el-form-item>
        <el-form-item v-if="mode === 'edit'" prop="name" label="投资客户: ">
          <el-input v-model="form.name" style="width: 250px" disabled />
        </el-form-item>
        <el-form-item prop="amount" label="投资金额: ">
          <el-input
            v-model="form.amount"
            style="width: 250px"
            type="text"
            placeholder="请输入投资金额"
          />
        </el-form-item>
        <el-form-item prop="date" label="投资时间: ">
          <el-date-picker
            v-model="form.date"
            type="datetime"
            placeholder="请选择投资日期和时间"
            style="width: 250px"
            value-format="yyyy-MM-dd HH:mm:ss"
            @change="changeTime"
          />
        </el-form-item>
        <el-form-item v-if="mode === 'add'" prop="profit" label="利润比例: ">
          <el-input
            v-model="form.profit"
            :disabled="true"
            style="width: 250px"
            oninput="value=value.match(/\d+\.?\d{0,2}/,'')"
          />
          %
        </el-form-item>
        <!-- 管理员可选利润比例 -->
        <!-- <el-form-item
          v-if="mode === 'add'"
          v-permission="['admin']"
          prop="profit"
          label="利润比例: "
        >
          <el-select
            v-model="adminProfit"
            class="m-2"
            placeholder="请选择客户对应团队利润比例"
            size="large"
            style="width: 250px"
            @change="changeProfitOptions"
          >
            <el-option
              v-for="item in profitOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item> -->
        <el-form-item v-if="mode === 'edit'" prop="profit" label="利润比例: ">
          <el-input
            v-model="form.profit"
            :disabled="true"
            style="width: 250px"
            oninput="value=value.match(/\d+\.?\d{0,2}/,'')"
          />
          %
        </el-form-item>
        <!-- 管理员可选利润比例 -->
        <!-- <el-form-item
          v-if="mode === 'edit'"
          v-permission="['admin']"
          prop="profit"
          label="利润比例: "
        >
          <el-select
            v-model="adminProfit"
            class="m-2"
            placeholder="请选择客户对应团队利润比例"
            size="large"
            @change="changeProfitOptions"
            style="width: 250px"
          >
            <el-option
              v-for="item in profitOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item> -->
        <el-form-item prop="remark" label="备注: ">
          <el-input
            v-model="form.remark"
            resize="none"
            type="textarea"
            :rows="2"
            style="width: 250px"
            placeholder="请输入备注"
          />
        </el-form-item>
      </el-form>
      <div v-show="show" class="small">
        <div v-for="(item, index) in data" :key="index" @click="choice(item)">
          {{ item.nick_name }} ({{ item.name }})
        </div>
      </div>
      <template #footer>
        <div style="text-align: center">
          <el-button
            type="primary"
            :disabled="loading"
            @click="submit"
          >确定</el-button>
          <el-button @click="isShow = false">取消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { investmentEdit, customerLists } from '@/api/customer/customer'
import { newinvestmentAdd } from '@/api/investment/investment'
export default {
  components: {},
  props: {},
  data() {
    return {
      // 管理员利润
      adminProfit: '',
      // 管理员利润选项
      profitOptions: [],
      // 选择客户列表
      data: [],
      // 控制选择客户列表框
      show: false,
      loading: false,
      mode: '',
      // 控制弹框显示
      isShow: false,
      // 投资时间
      time: '',
      // 限制投资时间为当前及之后
      // pickerOptions: {
      //   disabledDate(time) {
      //     return time.getTime() < Date.now() - 8.64e7
      //   }
      // },
      // 表单内容
      form: {
        name: '', // 新增客户名
        amount: '', // 投资金额
        date: '', // 投资时间
        profit: '', // 利润比例
        remark: '', // 备注
        customerid: '' // 客户id
      },
      // 验证规则
      rules: {
        name: [{ required: true, message: '用户名不能为空', trigger: 'blur' }],
        amount: [
          { required: true, message: '投资金额不能为空', trigger: 'blur' },
          {
            pattern: /^(([1-9]{1}\d*)|(0{1}))(\.\d{1,2})?$/,
            message: '投资金额不能小于0,且不能超过两位小数',
            trigger: 'blur'
          }
        ],
        date: [
          { required: true, message: '请选择投资日期和时间', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    // this.getProfit()
  },
  methods: {
    // 新增选择客户输入框
    async inputChange() {
      if (this.form.name.length != '') {
        const res = await customerLists({ keyword: this.form.name })
        this.data = res.list
        if (this.data.length > 0) {
          this.show = true
        } else {
          this.show = false
        }
      } else {
        this.data = []
        this.show = false
      }
    },
    async choice(item) {
      this.form.name = item.name
      this.form.profit = item.profit * 100
      this.form.customerid = item.customerid
      this.show = false
    },
    // 点击确定
    submit() {
      let forms = {}
      forms = { ...this.form }
      forms.profit = forms.profit * 0.01
      forms.date = this.parseTime(forms.date)
      this.loading = true
      this.$nextTick(() => {
        // 全局表单验证
        this.$refs['form'].validate(async result => {
          if (result) {
            if (this.mode === 'add') {
              await newinvestmentAdd(forms)
              // 提示
              this.$message.success('新增成功')
            } else if (this.mode === 'edit') {
              const res = await investmentEdit(forms)
              if (res.msg == '') {
                // 提示
                this.$message.success('修改成功')
              } else {
                this.$message.error(res.msg)
              }
            }
            // 关闭弹框
            this.isShow = false
            // 通知父刷新数据
            this.$emit('getList')
          }
          // 启用按钮
          this.loading = false
        })
      })
    },
    // 弹框关闭事件
    closeEvent() {
      // 清空校验规则
      this.$refs.form.resetFields()
      // 为防止清空不彻底,再手动清空一次
      this.form = {}
    },
    changeTime(e) {
      // 使用全局方法转化时间格式
      this.form.date = this.parseTime(e)
    }
  }
}
</script>
<style scoped lang="scss">
.small {
  position: absolute;
  top: 120px;
  left: 111px;
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
