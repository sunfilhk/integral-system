<template>
    <div class="ment">
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['system:sysment:export']"
              type="warning"
              icon="el-icon-download"
              size="mini"
              @click="handleExport"
            >导出投资列表</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['system:sysment:delete']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              @click="handleDelete"
            >删除所有投资</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['system:sysment:import']"
              type="warning"
              icon="el-icon-upload"
              size="mini"
              @click="handleImport"
            >导入投资列表</el-button>
          </el-col>
        </el-row>
        <el-dialog title="上传" :visible.sync="port" width="500px">
          <el-upload
            :action="url"
            :headers="obj">
            <el-button size="small" type="primary">点击上传</el-button>
          </el-upload>
        </el-dialog>
    </div>
</template>

<script>
import { investmentExport,emptyInvestment } from '@/api/customer/customer'
import { getToken } from '@/utils/auth'
export default {
    name:'Ment',
    props: {},
    data() {
        return {
          port: false,
          url:'',
          obj:{},
        };
    },
    computed: {},
    created() {
        this.auto()
    },
    mounted() {},
    watch: {},
    methods: {
        auto() {
          this.url=process.env.VUE_APP_BASE_API+'/api/v1/investmentImport'
          this.obj.Authorization='Bearer '+getToken()
        },
        handleExport() {
            this.$confirm('确认导出投资列表所有数据?', '警告', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(function() {
              investmentExport().then(res=> {
                location.href=process.env.VUE_APP_BASE_API+'/'+res.data
              })
            })
        },
        handleDelete() {
            this.$confirm('确认删除所有投资列表数据?', '警告', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(function() {
              return emptyInvestment()
            }).then(() => {
              this.msgSuccess('删除成功')
            })
        },
        handleImport() {
            this.port=true
        },
    },
    components: {},
};
</script>

<style scoped>
.ment {
    padding: 20px;
}
</style>
