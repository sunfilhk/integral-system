<template>
  <div class="unsettled">
    <!-- 搜索模块 -->
    <el-card class="topCard">
      <el-form ref="form" :model="queryParams" :inline="true" class="form">
        <!--日期 -->
        <el-form-item prop="update_time" label="日期: ">
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
        <el-form-item prop="id" label="投资单号: ">
          <el-input
            v-model="queryParams.id"
            clearable
            size="small"
            style="width: 240px"
            @keyup.enter.native="handleQuery"
          />
        </el-form-item>
        <el-form-item prop="keyword" label="客户名/业务员/手机号: ">
          <el-input
            v-model="queryParams.keyword"
            clearable
            size="small"
            style="width: 240px"
            @keyup.enter.native="handleQuery"
          />
        </el-form-item>
        <el-form-item label="团队类型: ">
          <el-select
            v-model="queryParams.deptid"
            clearable
            size="small"
            placeholder="请选择"
          >
            <el-option
              v-for="item in teamOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <!-- 按钮 -->
      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5" style="margin: 10px 0 0 0">
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
        :data="profit"
        border
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        show-summary
        :summary-method="getSummaries"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column
          :show-overflow-tooltip="true"
          label="投资订单号"
          prop="id"
          align="center"
          width="170"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          label="订单创建时间"
          prop="create_time"
          align="center"
          width="170"
        >
          <template v-slot="{ row }">
            {{ parseTime(row.create_time) }}
          </template>
        </el-table-column>
        <el-table-column
          :show-overflow-tooltip="true"
          label="客户姓名"
          prop="customername"
          align="center"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          label="业务员姓名"
          prop="nick_name"
          align="center"
          width="120"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          label="团队类型"
          prop="teamname"
          align="center"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          label="业绩"
          prop="amount"
          align="center"
          width="130"
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.amount) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          v-for="(item, index) in keyslist"
          :key="index"
          :show-overflow-tooltip="true"
          :label="item == '' ? '暂无数据' : item"
          :prop="wordslist[index]"
          align="center"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          label="操作"
          align="center"
          width="80"
          class-name="small-padding fixed-width"
        >
          <template slot-scope="scope">
            <el-button
              v-permisaction="['system:sysunsettled:set']"
              size="mini"
              type="text"
              icon="el-icon-document"
              @click="settlement(scope.row)"
            >结算</el-button>
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
import { statementNoSettlement, statementSettlement } from '@/api/report/report'
import { teamList } from '@/api/profitconfig/profitconfig' // 团队类型列表
export default {
  name: 'Unsettled',
  components: {},
  props: {},
  data() {
    return {
      profit: [],
      total: 0,
      summation: {},
      keyslist: '',
      wordslist: '',
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        start: '', // 结算开始
        end: '', // 结算截止
        id: '', // 投资订单号
        keyword: '', // 客户名/业务员/手机号
        deptid: '' // 团队类型
      },
      // 团队选项
      teamOptions: [],
      // 日期
      date: [],
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now() - 8.64e6
        }
      }
    }
  },
  computed: {},
  watch: {},
  created() {
    const end = new Date()
    const start = new Date()
    start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
    this.queryParams.start = this.parseTime(start).slice(0, 10)
    this.queryParams.end = this.parseTime(end).slice(0, 10)
    this.date = [this.queryParams.start, this.queryParams.end]
    this.auto()
    this.getData()
  },
  mounted() {},
  methods: {
    // 获取团队列表
    async getData() {
      const res = await teamList()
      this.teamOptions = res.list.map(v => {
        return {
          value: v.deptId, // 部门id
          label: v.deptname // 部门名称
        }
      })
    },
    auto() {
      statementNoSettlement(this.queryParams).then(res => {
        if (JSON.stringify(res.list) != '{}') {
          this.profit = res.list.profit
          this.total = res.total
          this.keyslist = res.list.keys.split(',')
          this.wordslist = res.list.words.split(',')
          this.summation = res.summation
          for (var i = 0; i < this.profit.length; i++) {
            for (var x = 0; x < this.wordslist.length; x++) {
              this.profit[i][this.wordslist[x]] = this.moneyFormat(
                this.profit[i][this.wordslist[x]]
              )
            }
          }
        } else {
          this.profit = []
          this.total = 0
          this.keyslist = ''
          this.wordslist = ''
        }
      })
    },
    // 修改日期
    changeDate(e) {
      this.queryParams.start = e[0]
      this.queryParams.end = e[1]
      this.auto()
    },
    settlement(row) {
      statementSettlement({ id: row.id }).then(res => {
        this.msgSuccess('结算成功')
        this.auto()
      })
    },
    getSummaries(param) {
      const { columns } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '合计'
          return
        }
        // const values = data.map(item => Number(item[column.property]));
        // if (!values.every(value => isNaN(value))) {
        sums[index] = this.moneyFormat(this.summation[column.property])
        // } else {
        //   sums[index] = '';
        // }
      })
      return sums
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
        end: this.date[1], // 结束时间
        id: '', // 投资订单号
        keyword: '', // 客户名/业务员/手机号
        deptid: '' // 团队类型
      }
      // 刷新
      this.auto()
    },
    // 搜索
    handleQuery() {
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
.unsettled {
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
