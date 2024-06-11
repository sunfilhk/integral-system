<template>
  <div style="padding: 20px">
    <!-- <component :is="currentRole" /> -->
    <el-card class="topCard">
      <div class="data">
        <el-row>
          <el-col :span="6">
            <div class="a">用户姓名：{{ data.nick_name }}</div>
          </el-col>
          <el-col :span="6">
            <div class="a">历史总业绩: {{ data.total_performance }}</div>
          </el-col>
          <el-col :span="6">
            <div class="a">当月业绩: {{ data.month }}</div>
          </el-col>
          <el-col :span="6">
            <div class="a">当日业绩: {{ data.today_performance }}</div>
          </el-col>
          <el-col :span="24">
            <div class="a">vip等级: {{ data.vip_level }}</div>
          </el-col>
          <el-col :span="6">
            <div class="a">vip称号: {{ data.vip_title }}</div>
          </el-col>
          <el-col :span="6">
            <div class="a">历史总利润: {{ profit.total }}</div>
          </el-col>
          <el-col :span="6">
            <div class="a">当月利润: {{ profit.month }}</div>
          </el-col>
          <el-col :span="6">
            <div class="a">当日利润: {{ profit.today }}</div>
          </el-col>
        </el-row>
      </div>
    </el-card>
    <el-card>
      <div style="margin: 20px 0 20px; font-weight: 700">下级业务员</div>
      <el-table
        border
        :data="list"
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        style="min-width: 800px"
      >
        <el-table-column sortable label="等级" prop="vip" width="100" />
        <el-table-column label="姓名" prop="nick_name" />
        <el-table-column label="用户名" prop="username" />
        <el-table-column label="手机号" prop="phone" />
        <el-table-column sortable label="团队业绩" prop="performance" />
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
import { mapGetters } from 'vuex'
// import adminDashboard from './admin'
// import editorDashboard from './editor'
import {
  getUserPerformance,
  getUserReferrals,
  getUserProfits
} from '@/api/user'
import { profitconfigList } from '@/api/profitconfig/profitconfig'

export default {
  name: 'Dashboard',
  // components: { adminDashboard, editorDashboard },
  data() {
    return {
      // currentRole: 'adminDashboard'
      data: {},
      deptId: '',
      // 利润
      profit: {},
      list: [],
      queryParams: {
        pageIndex: 1,
        pageSize: 10
      },
      total: 0,
      arr: []
    }
  },
  computed: {
    ...mapGetters(['roles'])
  },
  created() {
    // if (!this.roles.includes('admin')) {
    //   this.currentRole = 'editorDashboard'
    // }
    this.auto()
    this.getData()
  },
  methods: {
    // 月日历史利润
    async getData() {
      const res = await getUserProfits()
      this.profit = res.data
    },
    async auto() {
      await getUserPerformance().then(res => {
        this.data = res.data
        this.deptId = res.data.deptId
      })
      profitconfigList({ profittype: '2', deptId: this.deptId }).then(res => {
        this.arr = res.data.map(item => item.profitlevel)
        // 获取下级
        getUserReferrals(this.queryParams).then(res => {
          this.total = res.total
          // for (var i = 0; i < res.list.length; i++) {
          //   if (res.list[i].lifts == 0) {
          //     if (res.list[i].accumulative < res.list[i].set_level) {
          //       for (var x = 0; x < this.arr.length; x++) {
          //         if (res.list[i].set_level >= this.arr[x]) {
          //           res.list[i].vip = 'V' + (x + 1)
          //         }
          //       }
          //     } else if (res.list[i].accumulative >= res.list[i].set_level) {
          //       for (var y = 0; y < this.arr.length; y++) {
          //         if (res.list[i].accumulative >= this.arr[y]) {
          //           res.list[i].vip = 'V' + (y + 1)
          //         }
          //       }
          //     }
          //   } else if (res.list[i].lifts == 1) {
          //     for (var z = 0; z < this.arr.length; z++) {
          //       if (res.list[i].set_level >= this.arr[z]) {
          //         res.list[i].vip = 'V' + (z + 1)
          //       }
          //     }
          //   }
          // }
          this.list = res.list
        })
      })
    }
  }
}
</script>
<style scoped>
.topCard {
  margin-bottom: 20px;
}
.data {
  padding: 20px;
  /* border: 1px solid #dddddd; */
  background-color: #f9f9f9;
  color: #333;
  font-weight: 700;
  min-width: 800px;
}
.data .a {
  margin-bottom: 10px;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
