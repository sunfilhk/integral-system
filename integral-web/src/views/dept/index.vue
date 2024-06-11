<template>
  <div class="app-container">
    <!-- 头部搜索 -->
    <el-card class="topCard">
      <el-form :inline="true">
        <el-form-item label="部门名称">
          <el-input
            v-model="queryParams.deptName"
            placeholder="请输入部门名称"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
          />
        </el-form-item>
        <el-form-item label="状态">
          <el-select
            v-model="queryParams.status"
            placeholder="部门状态"
            clearable
            size="small"
          >
            <el-option
              v-for="dict in statusOptions"
              :key="dict.dictValue"
              :label="dict.dictLabel"
              :value="dict.dictValue"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button
            class="filter-item"
            type="primary"
            icon="el-icon-search"
            size="mini"
            @click="handleQuery"
          >搜索</el-button>
          <el-button
            v-permisaction="['system:sysdept:add']"
            class="filter-item"
            type="primary"
            icon="el-icon-plus"
            size="mini"
            @click="handleAdd"
          >新增</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <!-- 表格 -->
    <el-card>
      <el-table
        v-loading="loading"
        border
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        :data="deptList"
        row-key="deptId"
        default-expand-all
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column
          prop="deptname"
          label="部门名称"
          show-overflow-tooltip
        />
        <el-table-column prop="sort" label="排序" align="center" />
        <el-table-column
          prop="status"
          label="状态"
          align="center"
          :formatter="statusFormat"
        >
          <template slot-scope="scope">
            <el-tag
              :type="scope.row.status === '1' ? 'danger' : 'success'"
              disable-transitions
            >{{ statusFormat(scope.row) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          align="center"
          prop="createTime"
          width="220"
        >
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createTime) }}</span>
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
              v-permisaction="['system:sysdept:edit']"
              size="mini"
              type="text"
              icon="el-icon-edit"
              @click="handleUpdate(scope.row)"
            >修改</el-button>
            <el-button
              v-permisaction="['system:sysdept:add']"
              size="mini"
              type="text"
              icon="el-icon-plus"
              @click="handleAdd(scope.row)"
            >新增</el-button>
            <el-button
              v-if="scope.row.p_id != 0"
              v-permisaction="['system:sysdept:remove']"
              size="mini"
              type="text"
              icon="el-icon-delete"
              @click="handleDelete(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加或修改部门对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="600px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-row>
          <el-col v-if="form.parentId !== 0" :span="24">
            <el-form-item label="上级部门" prop="parentId">
              <treeselect
                v-model="form.parent_id"
                :options="deptOptions"
                :normalizer="normalizer"
                :show-count="true"
                placeholder="选择上级部门"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="部门名称" prop="deptName">
              <el-input v-model="form.deptname" placeholder="请输入部门名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="显示排序" prop="orderNum">
              <el-input-number
                v-model="form.sort"
                controls-position="right"
                :min="0"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="负责人" prop="leader">
              <el-input
                v-model="form.leader"
                placeholder="请输入姓名、用户名、手机号查找"
                maxlength="20"
                @input="chang"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="phone">
              <el-input
                v-model="form.phone"
                placeholder="请输入联系电话"
                maxlength="11"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="邮箱" prop="email">
              <el-input
                v-model="form.email"
                placeholder="请输入邮箱"
                maxlength="50"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="部门状态">
              <el-radio-group v-model="form.status">
                <el-radio
                  v-for="dict in statusOptions"
                  :key="dict.dictValue"
                  :label="dict.dictValue"
                >{{ dict.dictLabel }}</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div v-show="show" class="small">
        <div v-for="(item, index) in arr" :key="index" @click="choice(item)">
          {{ item.nick_name }} ({{ item.username }})
        </div>
      </div>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  getDeptList,
  getDept,
  delDept,
  addDept,
  updateDept
} from '@/api/system/dept'
import { userList } from '@/api/profitconfig/profitconfig'
import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'

export default {
  name: 'Dept',
  components: { Treeselect },
  data() {
    return {
      show: false,
      arr: [],
      // 遮罩层
      loading: true,
      // 表格树数据
      deptList: [],
      // 部门树选项
      deptOptions: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 状态数据字典
      statusOptions: [],
      // 查询参数
      queryParams: {
        deptname: undefined,
        status: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        parent_id: [
          { required: true, message: '上级部门不能为空', trigger: 'blur' }
        ],
        deptname: [
          { required: true, message: '部门名称不能为空', trigger: 'blur' }
        ],
        sort: [
          { required: true, message: '菜单顺序不能为空', trigger: 'blur' }
        ],
        email: [
          {
            type: 'email',
            message: "'请输入正确的邮箱地址",
            trigger: ['blur', 'change']
          }
        ],
        phone: [
          {
            pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/,
            message: '请输入正确的手机号码',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('sys_normal_disable').then(response => {
      this.statusOptions = response.data
    })
  },
  methods: {
    /** 查询部门列表 */
    getList() {
      this.loading = true
      getDeptList(this.queryParams).then(response => {
        this.deptList = response.data
        this.loading = false
      })
    },
    /** 转换部门数据结构 */
    normalizer(node) {
      // 剔除子节点为空的项
      if (node.children && !node.children.length) {
        delete node.children
      }
      return {
        id: node.deptId, // 部门id
        label: node.deptname, // 部门名称
        children: node.children // 子节点
      }
    },
    /** 查询部门下拉树结构 */
    getTreeselect() {
      getDeptList().then(response => {
        this.deptOptions = []
        const dept = { deptId: 0, deptname: '主类目', children: [] }
        // console.log('dept:', dept)
        dept.children = response.data
        this.deptOptions.push(dept)
      })
    },
    // 字典状态字典翻译
    statusFormat(row) {
      return this.selectDictLabel(this.statusOptions, row.status)
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        deptId: undefined,
        parent_id: undefined,
        deptname: undefined,
        sorc: undefined,
        leader: undefined,
        phone: undefined,
        email: undefined,
        status: '0'
      }
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.getList()
    },
    /** 新增按钮操作 */
    handleAdd(row) {
      this.reset()
      this.getTreeselect()
      if (row !== undefined) {
        this.form.parent_id = row.deptId
      }
      this.open = true
      this.title = '添加部门'
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      this.getTreeselect()

      getDept(row.deptId).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改部门'
      })
    },
    chang() {
      if (this.form.leader.length != '') {
        userList({ keyword: this.form.leader }).then(res => {
          this.arr = res.list
          if (this.arr.length > 0) {
            this.show = true
          } else {
            this.show = false
          }
        })
      } else {
        this.arr = []
        this.show = false
      }
    },
    choice(item) {
      this.form.leader = item.nick_name
      this.form.leaderID = parseInt(item.userid)
      this.show = false
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.deptId !== undefined) {
            updateDept(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            if (this.form.leaderID == null) {
              this.msgError('没有选中负责人')
              return
            }
            addDept(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('新增成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      this.$confirm(
        '是否确认删除名称为"' + row.deptname + '"的数据项?',
        '警告',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
        .then(function() {
          return delDept(row.deptId)
        })
        .then(() => {
          this.getList()
          this.msgSuccess('删除成功')
        })
        .catch(function() {})
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
.small {
  position: absolute;
  top: 242px;
  left: 100px;
  width: 200px;
  height: 258px;
  border: 1px solid rgb(0, 119, 255);
  background-color: #fff;
  border-radius: 4px;
  overflow: hidden;
  overflow-y: auto;
}
.small div {
  padding: 8px 20px;
}
.small div:hover {
  background-color: #eeeeee;
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
