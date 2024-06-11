<template>
  <div class="disposable">
    <!-- 搜索模块 -->
    <el-card class="topCard">
      <el-form ref="from" class="indexForm" :model="queryParams">
        <el-form-item label="团队类型: " label-width="80px" prop="deptid">
          <el-select
            v-model="queryParams.deptid"
            clearable
            size="small"
            placeholder="请选择团队类型"
            style="height: 30px"
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
            v-permisaction="['system:sysdisposable:add']"
            type="primary"
            icon="el-icon-plus"
            size="mini"
            @click="handleAdd"
          >新增</el-button>
        </el-form-item>
      </el-form>
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
        <el-table-column label="序号" prop="index" type="index" />
        <el-table-column label="姓名" prop="nick_name" />
        <el-table-column label="用户名" prop="username" />
        <el-table-column label="团队类型" prop="teamname" />
        <el-table-column sortable label="分润比例" prop="percent">
          <template v-slot="{ row }">
            <span>{{ parseFloat(row.percent * 100).toFixed(2) }}%</span>
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
            <el-button
              v-permisaction="['system:sysdisposable:remove']"
              size="mini"
              type="text"
              icon="el-icon-delete"
              @click="handleDelete(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加或修改弹框 -->
    <Add ref="add" @auto="auto" />
  </div>
</template>

<script>
import {
  profitconfigList,
  delProfitconfigOnce,
  teamList
} from '@/api/profitconfig/profitconfig'
import Add from './components/add.vue'
export default {
  name: 'Disposable',
  components: { Add },
  props: {},
  data() {
    return {
      // 用于数据筛选
      listData: [],
      // 团队选项
      teamOptions: [],
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        deptid: null // 部门id
      },
      total: 0,
      open: false,
      title: '',
      percent: '',
      form: {},
      list: [],
      data: [],
      show: false,
      optionsList: [],
      forms: {}
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
    // 获取列表数据方法
    async auto() {
      const res = await profitconfigList({ profittype: '0' })
      this.list = res.data
      this.listData = res.data
      this.total = res.data.length
    },
    // 修改团队选项
    async changeOptions(e) {
      let obj = {}
      obj = this.listData.find(v => v.deptId === e)
      if (obj) {
        obj.profittype = '0'
        const res = await profitconfigList(obj)
        this.list = res.data
      } else if (!obj) {
        if (!e) {
          const res = await profitconfigList({ profittype: '0' })
          this.list = res.data
          this.listData = res.data
        } else if (e) {
          this.list = []
        }
      }
    },
    // 双击复制
    handleDblClick(row, cloumn, cell, event) {
      if (cloumn.property) {
        const txt = event.target.innerText
        this.$copyText(txt).then(res => {
          this.$message.success('复制成功')
        })
      }
    },
    // 点击新增
    handleAdd() {
      // 显示弹框
      this.$refs.add.show = true
      // 标题
      this.$refs.add.mode = 'add'
    },
    // 点击修改
    handleUpdate(row) {
      // 显示弹框
      this.$refs.add.show = true
      // 标题
      this.$refs.add.mode = 'edit'
      this.forms = JSON.parse(JSON.stringify(row))
      this.forms.percent = (this.forms.percent * 100).toFixed(2)
      // 数据回显
      this.$refs.add.form = JSON.parse(JSON.stringify(this.forms))
    },
    // 点击删除
    handleDelete(row) {
      delProfitconfigOnce({ id: row.id }).then(res => {
        this.msgSuccess('删除成功')
        this.auto()
      })
    }
  }
}
</script>

<style scoped>
.disposable {
  padding: 20px;
}
.topCard {
  margin-bottom: 20px;
  min-width: 800px;
}
.indexForm {
  display: flex;
}
.indexForm .el-form-item {
  margin-bottom: 0;
}
</style>
