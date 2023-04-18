<template>
   <div class="content">
    <el-card shadow="never">
      <el-row>
        <el-col :span="1">
          <span>系统配置</span>
        </el-col>
        <el-col :offset="18"  :span="3" style="min-width: 120px">
          <el-button size="small" type="primary" @click="dialogFormVisible = true">添加集群</el-button>
        </el-col>
      </el-row>
      <el-divider></el-divider>
      <el-table :data="tableData" stripe style="width: 100%">
        <el-table-column prop="Name" label="名称" width="180"/>
        <el-table-column prop="Key" label="key" width="180" />
        <el-table-column prop="Value" label="value" />
        <el-table-column prop="CreatedAt" :formatter="dateFormat" label="创建时间" />
        <el-table-column prop="UpdatedAt" :formatter="dateFormat" label="修改时间" />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
            <el-popconfirm
              title="确定要删除吗？" 
              confirm-button-text="确认" 
              cancel-button-text="取消" 
              confirm-button-type="danger" 
              cancel-button-type="primary" 
              @confirm="handleDelete(scope.row)">
              <template #reference>
                <el-button size="small" type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-dialog v-model="dialogFormVisible" title="配置变更" width="30%" align-center>
        <el-form :model="formcfg">
          <el-form-item label="配置项" :label-width="formLabelWidth">
              <el-select v-model="formcfg.key" placeholder="请选择需要变更的配置">
                <el-option 
                  v-for="item in options" 
                  :key="item.value" 
                  :label="item.label" 
                  :value="item.value" />
              </el-select>
          </el-form-item>
          <el-form-item label="配置值" :label-width="formLabelWidth">
              <el-input v-model="formcfg.value" autocomplete="off" />
          </el-form-item>
        </el-form>
        <template #footer>
        <span class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="handleChange()">
            Confirm
            </el-button>
        </span>
        </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, reactive} from 'vue';
import { listCfg, listDefaultCfg, delCfg, updateCfg } from '../../api/setting';
import moment from 'moment';
import { ElMessage } from 'element-plus';

// 配置列表
const tableData = ref<any[]>([])
const dateFormat = (row:any, column:any) => {
  const date = row[column.property]
  if (date===undefined) {
    return ''
  }
  return moment(date).utcOffset(8).format('YYYY-MM-DD HH:mm')
}
// 弹窗
const props = {
  expandTrigger: 'hover' as const,
}
const options = ref<any[]>([])
const dialogFormVisible = ref(false)
const formLabelWidth = '100px'
const formcfg = reactive({
  key: '',
  value: '',
})
// 添加删除

// 数据请求
const load = async () => {
  // console.log(await list())
  let data = (await listCfg()).data
  if (data.errorCode === 0 ) {
    tableData.value = data.data.lists
  } else {
    ElMessage.error(data.msg)
  }
  let dcfg = (await listDefaultCfg()).data
  if (dcfg.errorCode === 0 ) {
    options.value = dcfg.data
  } else {
    ElMessage.error(dcfg.msg)
  }
}
const handleEdit = (id:number, val:any) => {
  dialogFormVisible.value = true
  formcfg.key = val.Key
  formcfg.value = val.Value
}
const handleDelete = async (val:any) => {
  if (!val.Key) {
    ElMessage.error("未获取到的删除的内容")
  } else {
    formcfg.key = val.Key
    const res = await delCfg(formcfg)
    if (res.data.errorCode === 0) {
      ElMessage.success("删除成功")
      await load()
    } else {
      ElMessage.error(res.data.msg)
    }
  }
}
const handleChange = async () => {
  console.log(formcfg)
  if (!formcfg.key) {
    ElMessage.error("未获取到的变更配置项")
  } else if (!formcfg.value) {
    ElMessage.error("未获取到的变更配置值")
  } else {
    const res = await updateCfg(formcfg)
    if (res.data.errorCode === 0) {
      ElMessage.success("变更成功")
      await load()
    } else {
      ElMessage.error(res.data.msg)
    }
    dialogFormVisible.value = false
  }
}

// 启动执行
onMounted(async () => {
  await load()
})
</script>
<style lange="scss" scoped>
.content {
  margin: 20px 8px;
}

.el-button--text {
  margin-right: 15px;
}
.el-select {
  width: 300px;
}
.el-input {
  width: 300px;
}
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>