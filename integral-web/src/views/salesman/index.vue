<template>
  <div class="salesman">
    <!-- 搜索模块 -->
    <el-card class="topCard">
      <el-form
        ref="queryForm"
        class="indexForm"
        :model="queryParams"
        :inline="true"
        label-width="68px"
      >
        <el-form-item
          label-width="158px"
          prop="keyword"
          label="用户名/业务员/手机号: "
        >
          <el-input
            v-model="queryParams.keyword"
            clearable
            size="small"
            style="width: 200px"
            @keyup.enter.native="handleQuery"
          />
        </el-form-item>
        <el-form-item label="团队类型: " label-width="80px" prop="deptId">
          <el-select
            v-model="queryParams.team_id"
            clearable
            size="small"
            placeholder="请选择团队类型"
            style="height: 30px"
            @change="changeOptions"
          >
            <el-option
              v-for="item in deptnameOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
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
          <el-button
            v-permisaction="['system:syssalesman:default']"
            type="primary"
            size="mini"
            @click="defaults"
          >默认配置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 列表模块 -->
    <el-card>
      <el-table
        border
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        :data="list"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column label="序号" prop="index" align="center" type="index">
          <template v-slot="{ $index }">
            {{
              (queryParams.pageIndex - 1) * queryParams.pageSize + $index + 1
            }}
          </template>
        </el-table-column>
        <el-table-column label="姓名" prop="nick_name" align="center" />
        <el-table-column label="团队类型" prop="team_name" align="center" />
        <el-table-column label="用户等级" prop="user_level" align="center" />
        <el-table-column label="用户名" prop="username" align="center" />
        <el-table-column label="等级锁定" width="100" align="center">
          <template slot-scope="scope">
            <el-switch
              v-model="scope.row.level_lock"
              active-value="1"
              inactive-value="0"
              @change="handleStatusChange(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          width="220"
          class-name="small-padding fixed-width"
        >
          <template slot-scope="scope">
            <el-button
              v-permisaction="['system:syssalesman:edit']"
              size="mini"
              type="text"
              icon="el-icon-edit"
              @click="handleUpdate(scope.row)"
            >修改</el-button>
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

    <!-- 默认分配 -->
    <el-dialog :title="title" :visible.sync="open" width="1000px">
      <el-form>
        <el-form-item label="团队类型: " label-width="80px">
          <el-select
            v-model="team_id"
            size="small"
            placeholder="请选择团队类型"
            style="height: 30px"
            @change="changConfigOptions"
          >
            <el-option
              v-for="item in deptnameOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <el-table :data="defaultlist">
        <el-table-column label="等级" prop="grade" align="center">
          <template slot-scope="scope">
            <span style="color: red">V{{ scope.row.grade }}</span>
          </template>
        </el-table-column>
        <el-table-column label="金额（万）" prop="profitlevel" align="center">
          <template slot-scope="scope">
            <el-input
              v-if="isEdit == scope.$index"
              v-model.trim="scope.row.profitlevel"
              :disabled="scope.$index == 0"
              placeholder="请输入金额"
            />
            <span
              v-if="isEdit != scope.$index"
            >>
              {{
                scope.row.profitlevel > 0 ? scope.row.profitlevel : '0'
              }}</span>
          </template>
        </el-table-column>
        <el-table-column label="百分比" prop="percent" align="center">
          <template slot-scope="scope">
            <el-input
              v-if="isEdit == scope.$index"
              v-model.trim="scope.row.percent"
              oninput="value=value.match(/\d+\.?\d{0,2}/,'')"
              placeholder="请输入百分比"
            />
            <span
              v-if="isEdit != scope.$index"
            ><span v-show="scope.row.grade != '1'">+</span>
              {{ scope.row.percent > 0 ? scope.row.percent : '3' }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          class-name="small-padding fixed-width"
        >
          <template slot-scope="scope">
            <el-button
              v-if="isEdit == scope.$index"
              v-permisaction="['system:syssalesman:edit']"
              size="mini"
              type="text"
              icon="el-icon-edit"
              @click="handleEdit(scope.$index, scope.row, 1)"
            >确定</el-button>
            <el-button
              v-if="isEdit != scope.$index"
              v-permisaction="['system:syssalesman:edit']"
              size="mini"
              type="text"
              icon="el-icon-edit"
              @click="handleEdit(scope.$index, scope.row, 0)"
            >编辑</el-button>
            <el-button
              v-if="scope.$index != 0"
              v-permisaction="['system:syssalesman:remove']"
              size="mini"
              type="text"
              icon="el-icon-delete"
              @click="handleDelete(scope.$index, scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
        <el-button
          slot="append"
          style="
            width: 100%;
            border-radius: 0;
            border-top: 0;
            border-left: 0;
            border-right: 0;
          "
          @click="appendNew"
        >点击增加一行</el-button>
      </el-table>
      <div style="text-align: center; margin-top: 20px">
        <el-button size="mini" type="primary" @click="submits">提交</el-button>
        <el-button size="mini" @click="cancels">取消</el-button>
      </div>
    </el-dialog>

    <!-- 修改分配弹框 -->
    <el-dialog :title="titles" :visible.sync="show" width="500px">
      <el-form ref="form" :model="form" label-width="80px">
        <el-form-item label="当前等级" prop="levelid">
          <el-select
            v-model="form.levelid"
            style="width: 100%"
            placeholder="请选择"
          >
            <el-option
              v-for="dict in sexOptions"
              :key="dict.id"
              :label="dict.vipLevel + ' (>' + dict.levelValue + ')'"
              :value="dict.id"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <div style="text-align: center; margin-top: 20px">
        <el-button
          size="mini"
          type="primary"
          @click="insubmits"
        >提交</el-button>
        <el-button size="mini" @click="incancels">取消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  profitconfigList,
  profitconfigEdit,
  userList,
  teamList,
  getUserVipLevel,
  editUserVipLevel,
  editlevellock
} from '@/api/profitconfig/profitconfig'
export default {
  name: 'Salesman',
  components: {},
  props: {},
  data() {
    return {
      // 默认配置金额
      profitlevel: '',
      // 默认配置百分比
      percent: '',
      // 团队id
      team_id: '',
      // 团队选项
      deptnameOptions: [],
      total: 0,
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        keyword: undefined,
        team_id: null // 部门id
      },
      open: false,
      form: {
        levelid: '',
        vipLevel: '',
        nick_name: ''
      },
      list: [],
      // 默认配置列表
      defaultlist: [],
      isEdit: -99,
      ids: '',
      userid: '',
      title: '',
      titles: '',
      show: false,
      sexOptions: [],
      vips: []
    }
  },
  computed: {},
  watch: {
    // 侦听数组长度
    defaultlist: {
      handler: function(newVal, oldVal) {
        if (newVal.length) {
          for (const i in newVal) {
            // 实时获取grade的值
            newVal[i].grade = Number(i) + 1
          }
        }
      }
    },
    deep: true,
    immediate: true
  },
  created() {
    this.auto()
    this.getData()
  },
  mounted() {},
  methods: {
    // 修改配置团队选项
    async changConfigOptions(e) {
      const obj = {}
      obj.profittype = '2'
      if (obj) {
        obj.deptId = e
        obj.team_id = e
        const res = await profitconfigList(obj)

        this.defaultlist = res.data
        this.ids = this.defaultlist.map(item => item.id).join(',')
        for (var i = 0; i < this.defaultlist.length; i++) {
          this.defaultlist[i].grade = i + 1
          this.defaultlist[i].percent = this.defaultlist[i].percent * 100
          this.defaultlist[i].profitlevel =
            this.defaultlist[i].profitlevel / 10000
          this.defaultlist[i].team_id = this.team_id
        }
      }
    },
    // 修改主页面团队选项,
    changeOptions() {
      this.auto()
    },

    // 获取团队列表
    async getData() {
      const res = await teamList()
      this.deptnameOptions = res.list.map(v => {
        return {
          value: v.deptId, // 部门id
          label: v.deptname // 部门名称
        }
      })
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
    // 获取用户列表
    auto() {
      userList(this.queryParams).then(res => {
        console.log('userList:', res)
        this.list = res.list
        this.total = res.total
      })
    },
    // 修改等级
    handleStatusChange(row) {
      const text = row.level_lock === '1' ? '打开' : '关闭'
      this.$confirm(
        '确认要"' + text + '""' + row.username + '"的等级锁定吗?',
        '警告',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
        .then(function() {
          var params = {
            level_lock: row.level_lock,
            userid: row.userid,
            nick_name: row.nick_name
          }
          return editlevellock(params)
        })
        .then(() => {
          this.msgSuccess(text + '成功')
        })
        .catch(function() {
          row.level_lock = row.level_lock === '1' ? '0' : '1'
        })
    },
    // 点击修改
    handleUpdate(row) {
      this.show = true
      this.userid = row.userid
      this.titles = row.nick_name
      this.form.nick_name = row.nick_name
      getUserVipLevel({ userid: row.userid }).then(res => {
        this.sexOptions = res.data.vipList
        this.form.levelid = res.data.curLevelID
        this.vips = res.data.vipList
      })
    },
    // 获取配置列表
    async defaults() {
      // 默认团队
      this.team_id =
        this.team_id === '' ? this.deptnameOptions[0].value : this.team_id
      this.open = true
      this.title = '默认分配'
      this.isEdit = -99
      var params = {
        profittype: '2',
        deptId: this.team_id
      }
      await profitconfigList(params).then(res => {
        this.defaultlist = res.data
        this.ids = this.defaultlist.map(item => item.id).join(',')
        for (var i = 0; i < this.defaultlist.length; i++) {
          this.defaultlist[i].grade = i + 1
          this.defaultlist[i].percent = this.defaultlist[i].percent * 100
          this.defaultlist[i].profitlevel =
            this.defaultlist[i].profitlevel / 10000
          this.defaultlist[i].team_id = this.team_id
        }
      })
    },
    handleEdit(index, row, status) {
      switch (status) {
        case 0:
          this.isEdit = index
          break
        case 1:
          this.isEdit = -99
          break
      }
    },
    handleDelete(index, row) {
      this.isEdit = -99
      this.defaultlist.splice(index, 1)
    },
    appendNew() {
      this.defaultlist.push({
        grade: this.defaultlist.length + 1,
        percent: '',
        profitlevel: '',
        userid: '0',
        profittype: 2,
        team_id: this.team_id
      })
      this.isEdit = this.defaultlist.length - 1
    },
    submits() {
      this.isEdit = -99
      const list = JSON.parse(JSON.stringify(this.defaultlist))
      for (var i = 0; i < list.length; i++) {
        list[i].percent = list[i].percent * 0.01
        list[i].profitlevel = list[i].profitlevel * 10000
        list[i].team_id = this.team_id
        if (this.title != '默认分配') {
          list[i].userid = this.userid
          list[i].profittype = 1
        }
      }
      profitconfigEdit({ ids: this.ids, profit: list }).then(res => {
        this.open = false
        this.msgSuccess('操作成功')
      })
    },
    // 配置取消
    cancels() {
      this.open = false
    },
    // 修改提交
    insubmits() {
      for (var i = 0; i < this.vips.length; i++) {
        if (this.vips[i].id == this.form.levelid) {
          this.form.vipLevel = this.vips[i].vipLevel
        }
      }
      this.form.userid = this.userid
      editUserVipLevel(this.form).then(res => {
        this.msgSuccess('修改成功')
        this.show = false
      })
    },
    // 修改取消
    incancels() {
      this.show = false
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.auto()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.resetForm('queryForm')
      // 手动清空
      this.queryParams = {
        pageIndex: 1,
        pageSize: 10,
        keyword: undefined,
        team_id: null // 部门id
      }
      this.handleQuery()
    }
  }
}
</script>

<style scoped>
.salesman {
  padding: 20px;
}
.topCard {
  margin-bottom: 20px;
  min-width: 1366px;
}
.indexForm {
  min-width: 800px;
}
</style>
