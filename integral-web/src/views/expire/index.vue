<template>
  <div class="expire">
    <!-- 搜索区域 -->
    <el-card class="topCard">
      <el-form
        ref="form"
        :inline="true"
        class="form"
        label-width="120px"
        :model="queryParams"
      >
        <!-- 合同到期时间 -->
        <el-form-item prop="days" label="合同到期时间: ">
          <el-select
            v-model="queryParams.days"
            clearable
            size="small"
            @change="changeExpiration_date"
          >
            <el-option
              v-for="item in contractOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <!-- 按钮 -->
        <el-form-item label-width="20px">
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-search"
            @click="searchClick"
          >搜索</el-button>
          <el-button
            size="mini"
            icon="el-icon-refresh"
            @click="resetClick"
          >重置</el-button>
        </el-form-item>
      </el-form>

      <!-- 表格区域 -->
      <el-card>
        <el-table
          width="100%"
          :data="businessList"
          border
          :header-cell-style="{ textAlign: 'center' }"
          :cell-style="{ textAlign: 'center' }"
          @cell-dblclick="handleDblClick"
        >
          <el-table-column
            label="序号"
            width="50"
            type="index"
            :show-overflow-tooltip="true"
          >
            <template v-slot="{ $index }">
              {{
                (queryParams.pageIndex - 1) * queryParams.pageSize + $index + 1
              }}
            </template>
          </el-table-column>
          <el-table-column
            prop="id"
            label="投资订单号"
            width="170"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            :show-overflow-tooltip="true"
            prop="name"
            label="客户姓名"
          />
          <el-table-column
            :show-overflow-tooltip="true"
            prop="phone"
            label="手机号"
            width="115"
          />
          <el-table-column
            :show-overflow-tooltip="true"
            prop="id_car"
            label="身份证号"
          />
          <el-table-column
            :show-overflow-tooltip="true"
            prop="bank_num"
            label="银行卡号"
          />
          <el-table-column
            :show-overflow-tooltip="true"
            prop="amount"
            label="投资金额"
            width="100"
          >
            <template v-slot="{ row }">
              {{ row.amount > 0 ? row.amount : '-' }}
            </template>
          </el-table-column>
          <el-table-column
            :show-overflow-tooltip="true"
            prop="profit"
            label="分润比例"
          >
            <template v-slot="{ row }">
              <span>{{ parseFloat(row.profit * 100).toFixed(2) }}%</span>
            </template>
          </el-table-column>
          <el-table-column
            :show-overflow-tooltip="true"
            prop="profits"
            label="分润金额"
            width="88"
          />
          <el-table-column
            :show-overflow-tooltip="true"
            prop="teamname"
            label="团队类型"
          />
          <el-table-column
            :show-overflow-tooltip="true"
            prop="invest_time"
            label="投资时间"
          >
            <template v-slot="{ row }">
              {{ parseTime(row.invest_time) }}
            </template>
          </el-table-column>
          <el-table-column
            :show-overflow-tooltip="true"
            prop="expiration_date"
            label="合同到期时间"
            width="170"
          >
            <template v-slot="{ row }">
              {{ parseTime(row.expiration_date) }}
            </template>
          </el-table-column>
          <el-table-column
            :show-overflow-tooltip="true"
            prop="form"
            label="类型"
          >
            <template v-slot="{ row }">
              <div v-if="row.form === 0">客户分润</div>
            </template>
          </el-table-column>
          <el-table-column
            :show-overflow-tooltip="true"
            prop="remark"
            label="备注"
          />
          <el-table-column
            :show-overflow-tooltip="true"
            prop="nick_name"
            label="业务员"
            width="90"
          />
          <el-table-column
            :show-overflow-tooltip="true"
            prop="status"
            label="状态"
            width="120"
          >
            <template v-slot="{ row }">
              <span v-show="row.status == 0" class="a">已通过</span>
              <span v-show="row.status == 3" class="b">已拒绝</span>
              <span v-show="row.status == 1" class="c">审核中</span>
              <span v-show="row.status == 4" class="c">已终止</span>
              <span v-show="row.status == 5" class="b">已终止</span>
              <span v-show="row.status == 6" class="b">已终止</span>
            </template>
          </el-table-column>
          <el-table-column
            :show-overflow-tooltip="true"
            prop="order_time"
            label="订单创建时间"
            width="100"
          >
            <template v-slot="{ row }">
              {{ parseTime(row.order_time) }}
            </template>
          </el-table-column>
          <!-- 操作按钮 -->
          <!-- <el-table-column
            :show-overflow-tooltip="true"
            prop="investment_id"
            label="操作"
            width="150"
          >
            <template v-slot="{ row }">
              <div v-if="row.status !== '0'">
                <el-button size="mini" type="text" disabled style="color: gray"
                  >无需操作</el-button
                >
              </div>
              <div v-else>
                <el-button
                  size="mini"
                  icon="el-icon-edit"
                  type="text"
                  @click="edit(row)"
                  >修改</el-button
                >
                <el-button
                  v-permisaction="['system:syscustomer:stop']"
                  size="mini"
                  type="text"
                  @click="terminationClick(row)"
                  >终止</el-button
                >
                <el-button
                  v-permisaction="['system:syscustomer:revoke']"
                  size="mini"
                  type="text"
                  @click="revocation(row)"
                  >撤销</el-button
                >
              </div>
            </template>
          </el-table-column> -->
        </el-table>
        <!-- 分页符 -->
        <div>
          <div class="page">
            <pagination
              v-show="total > 0"
              :total="total"
              :page.sync="queryParams.pageIndex"
              :limit.sync="queryParams.pageSize"
              @pagination="clickPage"
            />
          </div>
        </div>
      </el-card>
    </el-card>
  </div>
</template>

<script>
import { searchsexpire } from '@/api/customer/customer'
export default {
  name: 'Expire',
  components: {},
  props: {},
  data() {
    return {
      queryParams: {
        pageIndex: 1, // 页码
        pageSize: 10, // 页容量
        days: '1' // 合同到期时间
      },
      // 总数
      total: 0,
      // 业务管理列表
      businessList: [],
      // 合同到期
      contractTime: '1',
      // 合同到期选项
      contractOptions: [
        {
          value: '1',
          label: '当天到期'
        },
        {
          value: '7',
          label: '7天内到期'
        },
        {
          value: '14',
          label: '14天内到期'
        },
        {
          value: '30',
          label: '30天内到期'
        }
      ]
    }
  },
  created() {
    this.getList()
  },
  methods: {
    // 改变合同到期时间
    async changeExpiration_date(e) {
      this.queryParams.days = e
    },
    // 获取列表
    async getList() {
      const res = await searchsexpire({
        status: '0',
        days: this.queryParams.days,
        pageIndex: this.queryParams.pageIndex,
        pageSize: this.queryParams.pageSize
      })
      this.businessList = res.list
      this.total = res.total
    },
    // 点击搜索
    async searchClick() {
      this.getList()
    },
    // 重置事件
    resetClick() {
      // 设置日期选择器初始化显示一个月时间
      // 清空表单
      this.$refs.form.resetFields()
      // 手动再重置一次
      this.form = {}
      // 刷新
      this.getList()
    },
    // 点击分页
    clickPage() {
      this.getList()
    },
    // 双击复制
    handleDblClick(row, cloumn, cell, event) {
      if (cloumn.property) {
        const txt = event.target.innerText
        this.$copyText(txt).then(res => {
          this.$message.success('复制成功')
        })
      }
    }
  }
}
</script>
<style scoped lang="scss">
.a {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(170, 255, 170);
  color: rgb(0, 255, 98);
  background-color: rgb(235, 255, 235);
}
.b {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(255, 180, 180);
  color: rgb(255, 80, 80);
  background-color: rgb(255, 221, 221);
}
.c {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(149, 195, 255);
  color: rgb(0, 110, 255);
  background-color: rgb(231, 241, 255);
}
</style>
