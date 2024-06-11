<template>
  <div class="record">
    <!-- 搜索模块 -->
    <el-card class="topCard">
      <el-form ref="queryForm" :model="queryParams" :inline="true">
        <el-form-item
          prop="keyword"
          label="客户名/业务员/手机号: "
          @keyup.enter.native="handleQuery"
        >
          <el-input
            v-model="queryParams.keyword"
            clearable
            size="small"
            style="width: 200px"
            @keyup.enter.native="handleQuery"
          />
        </el-form-item>
        <el-form-item label-width="80px" label="投资单号: ">
          <el-input v-model="queryParams.investmentid" clearable size="small" />
        </el-form-item>
        <el-form-item label-width="80px" label="团队类型: ">
          <el-select
            v-model="queryParams.teamid"
            clearable
            size="small"
            @change="changeOptions"
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
            type="primary"
            icon="el-icon-search"
            size="mini"
            @click="handleQuery"
          >搜索</el-button>
          <el-button
            icon="el-icon-refresh"
            size="mini"
            @click="resetQuery"
          >重置</el-button>
        </el-col>
      </el-row>
    </el-card>

    <!-- 列表模块 -->
    <el-card>
      <el-table
        border
        :data="customerlist"
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column
          label="投资订单号"
          width="180"
          show-overflow-tooltip
          prop="investmentid"
        >
          <template v-slot="{ row }">
            {{ row.investmentid ? row.investmentid : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="客户姓名" prop="name" show-overflow-tooltip />
        <el-table-column label="手机号" prop="phone" show-overflow-tooltip />
        <el-table-column
          label="身份证号"
          prop="identity"
          show-overflow-tooltip
        />
        <el-table-column
          label="银行卡号"
          width="180"
          prop="bank"
          show-overflow-tooltip
        />
        <el-table-column label="开户行" prop="banknum" show-overflow-tooltip />
        <el-table-column
          label="性别"
          width="100"
          prop="sex"
          show-overflow-tooltip
        />
        <el-table-column label="金额" prop="amount" show-overflow-tooltip>
          <template v-slot="{ row }">
            {{ row.amount ? row.amount : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="备注" prop="remark" show-overflow-tooltip />
        <el-table-column label="审核类型" width="110" prop="edittype">
          <template slot-scope="scope">
            <span v-show="scope.row.edittype == 0">新增客户审核</span>
            <span v-show="scope.row.edittype == 1">修改客户审核</span>
            <span v-show="scope.row.edittype == 2">新增投资审核</span>
            <span v-show="scope.row.edittype == 3">修改投资审核</span>
            <span v-show="scope.row.edittype == 4">终止投资审核</span>
            <span v-show="scope.row.edittype == 5">修改客户审核</span>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="100" prop="status">
          <template slot-scope="scope">
            <span v-show="scope.row.status == 0" class="a">已通过</span>
            <span v-show="scope.row.status == 3" class="b">已拒绝</span>
          </template>
        </el-table-column>
        <el-table-column label="业务员姓名" width="100" prop="nick_name" />
        <el-table-column
          label="团队类型"
          show-overflow-tooltip
          prop="tick_name"
        />
        <el-table-column label="创建时间" width="160" prop="create_time">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.create_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="审核时间" width="160" prop="audit_time">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.audit_time) }}</span>
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
import { customerLogList } from '@/api/customer/customer'
import { teamList } from '@/api/profitconfig/profitconfig' // 团队类型列表
export default {
  name: 'Record',
  components: {},
  props: {},
  data() {
    return {
      // 团队选项
      teamOptions: [],
      total: 0,
      customerlist: [],
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        keyword: undefined, // 客户/业务员/手机号
        investmentid: '', // 投资单号
        teamid: '' // 团队
      }
    }
  },
  computed: {},
  watch: {},
  created() {
    this.auto()
    this.getData()
  },
  mounted() {},
  methods: {
    // 改变团队选项
    changeOptions(e) {
      this.queryParams.teamid = e
      this.auto()
    },
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
    // 获取列表
    auto() {
      customerLogList(this.queryParams).then(res => {
        this.customerlist = res.list
        this.total = res.total
      })
    },
    // 点击搜索
    handleQuery() {
      this.queryParams.pageIndex = 1
      // 获取数据
      this.auto()
    },
    // 点击重置
    resetQuery() {
      // 清空信息
      this.resetForm('queryForm')
      // 手动清空
      this.queryParams = {}
      // 获取列表
      this.handleQuery()
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

<style scoped>
.record {
  padding: 20px;
}
.topCard {
  margin-bottom: 20px;
  min-width: 1366px;
}
.el-form-item {
  margin-bottom: 0;
}
.a {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(170, 255, 170);
  color: rgb(0, 255, 0);
  background-color: rgb(235, 255, 235);
}
.b {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(255, 180, 180);
  color: rgb(255, 80, 80);
  background-color: rgb(255, 221, 221);
}
</style>
