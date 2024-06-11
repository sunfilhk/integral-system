<template>
  <div class="audrecord">
    <el-card>
      <el-table
        border
        :data="list"
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column label="序号" prop="index" type="index">
          <template v-slot="{ $index }">
            {{
              (queryParams.pageIndex - 1) * queryParams.pageSize + $index + 1
            }}
          </template>
        </el-table-column>
        <el-table-column label="姓名" prop="nick_name" />
        <el-table-column label="用户名" prop="username" />
        <el-table-column label="手机号" prop="phone" />
        <el-table-column label="推荐人" prop="referrer_name" />
        <el-table-column label="团队类型" prop="teamname" />
        <el-table-column label="状态" prop="is_pass">
          <template slot-scope="scope">
            <span v-show="scope.row.is_pass == 1" class="a">已通过</span>
            <span v-show="scope.row.is_pass == 0" class="b">已拒绝</span>
          </template>
        </el-table-column>
        <el-table-column label="操作时间" width="160" prop="pass_time">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.pass_time) }}</span>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页符 -->
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
import { userAuditRecord } from '@/api/system/audituser'
export default {
  name: 'Audrecord',
  components: {},
  props: {},
  data() {
    return {
      queryParams: {
        pageIndex: 1,
        pageSize: 10
      },
      total: 0,
      list: []
    }
  },
  computed: {},
  watch: {},
  created() {
    this.auto()
  },
  mounted() {},
  methods: {
    // 双击复制
    handleDblClick(row, cloumn, cell, event) {
      if (cloumn.property) {
        const txt = event.target.innerText
        this.$copyText(txt).then(res => {
          this.$message.success('复制成功')
        })
      }
    },
    auto() {
      userAuditRecord(this.queryParams).then(res => {
        this.list = res.data.list
        this.total = Number(res.data.totalCount)
      })
    }
  }
}
</script>

<style scoped>
.audrecord {
  padding: 20px;
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
