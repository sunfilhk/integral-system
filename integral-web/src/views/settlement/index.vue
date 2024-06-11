<template>
  <div class="settlement">
    <!--搜索模块 -->
    <el-card class="topCard">
      <el-form ref="form" class="form" label-width="100px" :model="queryParams">
        <!--投资日期 -->
        <el-form-item prop="update_time" label="结算日期: ">
          <el-date-picker
            v-model="date"
            size="small"
            type="daterange"
            align="right"
            unlink-panels
            :clearable="false"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            :picker-options="pickerOptions"
            :default-value="Date.now() - 30 * 24 * 3600 * 1000"
            value-format="yyyy-MM-dd"
            style="width: 240px"
            @change="changeDate"
          />
        </el-form-item>
        <!-- 订单号 -->
        <el-form-item label-width="100px" label="投资订单号: ">
          <el-input v-model="queryParams.invest_id" size="small" />
        </el-form-item>
        <!-- 客户名/手机号 -->
        <el-form-item label-width="180px" label="客户名/业务员/手机号: ">
          <el-input v-model="queryParams.keyword" size="small" />
        </el-form-item>
        <!-- 业务员 -->
      </el-form>
      <!-- 按钮 -->
      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5" style="margin: 10px 0 0 30px">
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
        </el-col>
      </el-row>
    </el-card>
    <!-- 列表模块 -->
    <el-card>
      <el-table
        border
        :data="list"
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column
          label="投资订单号"
          prop="invest_id"
          align="center"
          show-overflow-tooltip
          width="120"
        />
        <el-table-column
          label="订单创建时间"
          prop="create_time"
          align="center"
          show-overflow-tooltip
          width="120"
        >
          <template v-slot="{ row }">
            {{ parseTime(row.create_time) }}
          </template>
        </el-table-column>
        <el-table-column
          label="姓名"
          prop="name"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="银行卡号"
          width="190"
          prop="bank"
          align="center"
        />
        <el-table-column
          label="开户行"
          prop="banknum"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="业绩"
          prop="amount"
          align="center"
          width="150px"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(+scope.row.amount) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="结算利润"
          prop="profits"
          align="center"
          width="150px"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(+scope.row.profits) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="业务员"
          prop="nickname"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="结算时间"
          width="155"
          align="center"
          prop="settle_time"
        >
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.settle_time) }}</span>
          </template>
        </el-table-column>
      </el-table>
      <pagination
        v-show="total > 0"
        :total="total"
        :page.sync="queryParams.pageIndex"
        :limit.sync="queryParams.pageSize"
        @pagination="auto"
      />
    </el-card>
  </div>
</template>

<script>
import { statementSettleHistory } from '@/api/report/report'
export default {
  name: 'Settlement',
  components: {},
  props: {},
  data() {
    return {
      // 限制时间选择截止当天
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now() - 8.64e6
        }
      },
      // 结算日期
      date: [],
      list: [],
      total: 0,
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        start: '', // 结算开始
        end: '', // 结算截止
        invest_id: '', // 投资订单号
        keyword: '' // 客户名/业务员/手机号
      }
    }
  },
  computed: {},
  watch: {},
  created() {
    const start = new Date()
    const end = new Date()
    start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
    this.queryParams.start = this.parseTime(start).slice(0, 10)
    this.queryParams.end = this.parseTime(end).slice(0, 10)
    // 初始投资日期
    this.date = [this.queryParams.start, this.queryParams.end]
    this.auto()
  },
  mounted() {},
  methods: {
    // 选择日期
    changeDate(e) {
      this.queryParams.start = e[0]
      this.queryParams.end = e[1]
      // 刷新数据
      this.auto()
    },

    auto() {
      statementSettleHistory(this.queryParams).then(res => {
        this.list = res.list
        this.total = res.total
      })
    },
    // 点击搜索
    searchClick() {
      this.auto()
    },
    // 重置
    resetClick() {
      // 日期复原
      const start = new Date()
      const end = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      this.queryParams.start = this.parseTime(start).slice(0, 10)
      this.queryParams.end = this.parseTime(end).slice(0, 10)
      this.date = [this.queryParams.start, this.queryParams.end]
      // 清空表单
      this.$refs.form.resetFields()
      // 手动清空
      this.queryParams = {
        pageIndex: 1,
        pageSize: 10,
        start: this.date[0], // 开始时间
        end: this.date[1], // 截止时间
        invest_id: '', // 投资订单号
        keyword: '' // 客户名/业务员/手机号
      }
      // 刷新
      this.auto()
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
.settlement {
  padding: 20px;
}
.topCard {
  margin-bottom: 20px;
  .form {
    min-width: 1366px;
    display: flex;
    flex-wrap: wrap;
    .el-form-item {
      margin: 0 10px 5px 0;
      ::v-deep .el-input--small .el-input__inner {
        width: 240px;
      }
    }
  }
}
</style>
