<template>
  <div class="examine">
    <!-- 搜索模块 -->
    <el-card class="topCard">
      <el-form ref="queryForm" :model="queryParams" :inline="true">
        <el-form-item prop="keyword" label="客户名/业务员/手机号: ">
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
            v-model="queryParams.deptId"
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
        <el-table-column
          label="手机号"
          width="110"
          prop="phone"
          show-overflow-tooltip
        />
        <el-table-column
          label="身份证号"
          prop="identity"
          show-overflow-tooltip
        />
        <el-table-column label="银行卡号" prop="bank" show-overflow-tooltip />
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
        <el-table-column
          label="审核类型"
          width="140"
          show-overflow-tooltip
          prop="edittype"
        >
          <template slot-scope="scope">
            <span v-show="scope.row.edittype == 0">新增客户投资审核</span>
            <span v-show="scope.row.edittype == 1">修改客户审核</span>
            <span v-show="scope.row.edittype == 2">新增投资审核</span>
            <span v-show="scope.row.edittype == 3">修改投资审核</span>
            <span v-show="scope.row.edittype == 4">终止投资审核</span>
          </template>
        </el-table-column>
        <el-table-column label="业务员姓名" width="100" prop="nick_name" />
        <el-table-column
          label="团队类型"
          show-overflow-tooltip
          prop="tick_name"
        />
        <el-table-column label="投资时间" width="160" prop="invest_time">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.invest_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="160" prop="create_time">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.create_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="130"
          class-name="small-padding fixed-width"
        >
          <template slot-scope="scope">
            <el-button
              v-permisaction="['system:sysexamine:agree']"
              size="mini"
              type="text"
              icon="el-icon-check"
              @click="agree(scope.row)"
            >同意</el-button>
            <el-button
              v-permisaction="['system:sysexamine:refuse']"
              size="mini"
              type="text"
              icon="el-icon-close"
              @click="refuse(scope.row)"
            >拒绝</el-button>
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
import { customerAuditList, customerLogAudit } from '@/api/customer/customer'
import { teamList } from '@/api/profitconfig/profitconfig' // 团队类型列表
export default {
  name: 'Examine',
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
        keyword: '', // 客户/业务员/手机号
        investmentid: '', // 客户审核-投资单号
        deptId: '' // 团队
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
    changeOptions() {
      // 刷新数据
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
    // 获取列表数据
    auto() {
      customerAuditList(this.queryParams).then(res => {
        this.customerlist = res.list
        this.total = res.total
      })
    },
    // 点击同意
    agree(row) {
      var params = {
        id: row.id,
        status: '0'
      }
      customerLogAudit(params).then(res => {
        this.msgSuccess('审核成功')
        this.auto()
      })
    },
    // 点击拒绝
    refuse(row) {
      var params = {
        id: row.id,
        status: '3'
      }
      customerLogAudit(params).then(res => {
        // 提示
        this.msgSuccess('审核成功')
        // 刷新
        this.auto()
      })
    },
    handleQuery() {
      this.queryParams.pageIndex = 1
      // 刷新
      this.auto()
    },

    // 重置
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
.examine {
  padding: 20px;
}
.topCard {
  margin-bottom: 20px;
  min-width: 1366px;
}
.el-form-item {
  margin: 0 10px 0 0;
}
</style>
