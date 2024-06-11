<template>
  <div class="client">
    <!-- 头部搜索区域 -->
    <el-card class="topCard">
      <el-form ref="form" class="form">
        <el-form-item label="团队类型: " label-width="80px">
          <el-select
            v-model="searchTeamid"
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
        <el-form-item label-width="10px">
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-search"
          >搜索</el-button>
          <el-button
            size="mini"
            icon="el-icon-refresh"
            @click="resetClick"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 表格区域 -->
    <el-card>
      <el-table
        border
        :data="list"
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
      >
        <el-table-column label="团队" prop="deptname" />

        <el-table-column sortable label="客户利润比例" prop="profit">
          <template v-slot="{ row }">
            <span>{{ parseFloat(row.profit * 100).toFixed(2) }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="220"
          class-name="small-padding fixed-width"
        >
          <template slot-scope="scope">
            <el-button
              v-permisaction="['system:sysdisposable:edit']"
              size="mini"
              type="text"
              icon="el-icon-edit"
              @click="handleUpdate(scope.row)"
            >修改</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!--修改弹框区域 -->
    <el-dialog title="修改客户利润比例" :visible.sync="show" width="400px">
      <div style="text-align: center">
        <h3>团队: {{ deptname }}团队</h3>
      </div>

      <div style="text-align: center; font-weight: 700">
        客户利润比例:
        <el-input
          v-model="profit"
          clearable
          placeholder="请输入百分比(如：5%)"
          oninput="value=value.match(/\d+\.?\d{0,2}/,'')"
          style="width: 200px"
        />
        %
      </div>
      <div slot="footer" class="dialog-footer" style="text-align: center">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="show = false">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { updateProfitcustomer } from '@/api/customer/customer'
import { profitList, teamList } from '@/api/profitconfig/profitconfig'
export default {
  name: 'Client',
  components: {},
  props: {},
  data() {
    return {
      // 列表数据
      list: [],
      // 用于数据筛选
      listData: [],
      // 团队选项
      teamOptions: [
        {
          value: 1, // 默认值
          label: '中悦云创' // 默认值
        }
      ],
      // 搜索团队id
      searchTeamid: '',
      // 团队名称
      deptname: '',
      // 利润比例
      profit: '',
      // 修改团队比例
      updateProfit: { profit: '', deptId: '' },
      // 控制弹框显示
      show: false
    }
  },
  created() {
    // 获取数据
    this.getData()
    this.getTeam()
  },
  // mounted() {
  //   //获取列表数据
  //   this.getData()
  // },
  methods: {
    // 获取团队列表
    async getTeam() {
      const res = await teamList()
      this.teamOptions = res.list.map(v => {
        return {
          value: v.deptId, // 部门id
          label: v.deptname // 部门名称
        }
      })
    },

    // 获取列表数据方法
    async getData() {
      const res = await profitList({ status: '0' })
      this.list = res.data
      this.listData = res.data
      // 团队选项
      this.teamOptions = res.data.map(item => {
        return {
          value: item.deptId,
          label: item.deptname
        }
      })
    },

    // 修改团队选项
    async changeOptions(e) {
      let obj = {}
      obj = this.listData.find(v => v.deptId === e)
      const res = await profitList(obj)
      this.list = res.data
    },
    // 点击修改
    handleUpdate(row) {
      // 显示弹框
      this.show = true
      // 数据回显
      this.deptname = row.deptname
      this.updateProfit.deptId = row.deptId
      this.profit = (row.profit * 100).toFixed(2)
    },
    // 点击提交
    async submitForm() {
      // 修改团队利润百分比
      this.updateProfit.profit = this.profit * 0.01
      await updateProfitcustomer(this.updateProfit)
      // 刷新
      this.getData()
      this.searchTeamid = ''
      // 提示
      this.$message.success('修改成功')
      // 关闭弹框
      this.show = false
    },

    // 点击重置
    resetClick() {
      // 清空信息
      this.searchTeamid = ''
      // 刷新
      this.getData()
    }
  }
}
</script>
<style scoped lang="scss">
.client {
  padding: 20px;
}
.form {
  display: flex;
}
.topCard {
  margin-bottom: 20px;
  .el-form-item {
    margin: 0;
  }
}
</style>
