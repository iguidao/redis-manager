<template>
  <div class="content">
   <el-card shadow="never">
     <el-row>
       <el-col :span="2">
         <span>权限配置</span>
       </el-col>
       <el-col :offset="18"  :span="3" style="min-width: 120px">
         <el-button size="small" type="primary" @click="dialogFormVisible = true">添加权限</el-button>
       </el-col>
     </el-row>
     <el-divider></el-divider>
     <el-table :data="tableData" v-loading="loading" stripe style="width: 100%">
       <el-table-column prop="identity" label="身份" width="180" />
       <el-table-column prop="path" label="接口" />
       <el-table-column prop="method" label="meth" width="180"/>
       <el-table-column prop="note" label="备注" />
       <el-table-column label="操作">
         <template #default="scope">
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
   <el-dialog v-model="dialogFormVisible" title="权限添加" width="30%" align-center>
       <el-form :model="formcfg">
         <el-form-item label="身份" :label-width="formLabelWidth">
             <el-select v-model="formcfg.identity" placeholder="选择身份">
               <el-option 
                 v-for="item in identityoptions" 
                 :key="item.value" 
                 :label="item.label" 
                 :value="item.value" />
             </el-select>
         </el-form-item>
         <el-form-item label="接口" :label-width="formLabelWidth">
             <el-select v-model="formcfg.path" placeholder="选择接口">
               <el-option 
                 v-for="item in pathoptions" 
                 :key="item.value" 
                 :label="item.label" 
                 :value="item.value" />
             </el-select>
         </el-form-item>
         <el-form-item label="Method" :label-width="formLabelWidth">
             <el-select v-model="formcfg.method" placeholder="选择Method">
               <el-option 
                 v-for="item in methodoptions" 
                 :key="item.value" 
                 :label="item.label" 
                 :value="item.value" />
             </el-select>
         </el-form-item>
       </el-form>
       <template #footer>
       <span class="dialog-footer">
           <el-button @click="handleCancel()">取消</el-button>
           <el-button type="primary" @click="handleChange()">确定</el-button>
       </span>
       </template>
   </el-dialog>
 </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, reactive} from 'vue';
import { ruleadd, ruledel, rulelist, rulecfg } from '../../api/rule';
import { usertypelist } from '../../api/user'
import moment from 'moment';
import { ElMessage } from 'element-plus';

// 配置列表
const loading = ref(false)
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
const identityoptions = ref<any[]>([])
const pathoptions = ref<any[]>([])
const methodoptions = ref<any[]>([])
const dialogFormVisible = ref(false)
const formLabelWidth = '100px'
const formcfg = reactive({
  identity:'',
  path: '',
  method: '',
})
// 添加删除

// 数据请求
const load = async () => {
 // console.log(await list())
  let data = (await rulelist()).data
  if (data.errorCode === 0 ) {
    loading.value = false
    tableData.value = data.data
  } else {
    loading.value = false
    ElMessage.error(data.msg)
  }
  let tcfg = (await usertypelist()).data
  if (tcfg.errorCode === 0 ) {
      identityoptions.value = tcfg.data
  } else {
    ElMessage.error(tcfg.msg)
  }
  let rcfg = (await rulecfg()).data
  if (rcfg.errorCode === 0 ) {
      pathoptions.value = rcfg.data.url
      methodoptions.value = rcfg.data.method
  } else {
    ElMessage.error(rcfg.msg)
  }

}
const handleChange = async () => {
  console.log(formcfg)
  if (!formcfg.identity) {
    ElMessage.error("未获取到的添加的配置")
  } else if (!formcfg.path) {
    ElMessage.error("未获取到的添加的配置")
  } else if (!formcfg.method) {
    ElMessage.error("未获取到的添加的配置")
  } else {
    const res = await ruleadd(formcfg)
    if (res.data.errorCode === 0) {
      ElMessage.success("添加成功")
      formcfg.identity = ""
      formcfg.path = ""
      formcfg.method = ""
      await load()
    } else {
      ElMessage.error(res.data.msg)
    }
    dialogFormVisible.value = false
  }
}
const handleDelete = async (val:any) => {
 if (!val.identity) {
   ElMessage.error("未获取到的删除的配置")
 } else if (!val.path) {
  ElMessage.error("未获取到的删除的配置")
 } else if (!val.method) {
  ElMessage.error("未获取到的删除的配置")
 } else {
   formcfg.identity = val.identity
   formcfg.path = val.path
   formcfg.method = val.method
   const res = await ruledel(formcfg)
   if (res.data.errorCode === 0) {
     ElMessage.success("删除成功")
     await load()
   } else {
     ElMessage.error(res.data.msg)
   }
 }
}

const handleCancel = () => {
  dialogFormVisible.value = false
  formcfg.identity = ""
  formcfg.path = ""
  formcfg.method = ""
}
// 启动执行
onMounted(async () => {
 loading.value = true
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