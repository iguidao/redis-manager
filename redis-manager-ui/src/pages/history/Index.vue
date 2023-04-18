<template>
  <div class="content">
   <el-card shadow="never">
     <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="CreatedAt" :formatter="dateFormat" label="时间" width="170" sortable />
       <el-table-column prop="UserId" label="用户ID" width="100"/>
       <el-table-column prop="OpInfo" label="地址" width="400" />
       <el-table-column prop="OpParams" label="事件" />
     </el-table>
   </el-card>
 </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import {list} from '../../api/history';
import moment from 'moment';
import { ElMessage } from "element-plus";

const tableData = ref<any[]>([])
const ctime = ref("")
const load = async () => {
 // console.log(await list())
 let data = (await list()).data
  if (data.errorCode === 0 ) {
    tableData.value = data.data
  } else {
    ElMessage.error(data.msg)
  }
}
const dateFormat = (row:any, column:any) => {
 const date = row[column.property]
 if (date===undefined) {
   return ''
 }
 return moment(date).utcOffset(8).format('YYYY-MM-DD HH:mm:ss')
}
onMounted(async () => {
 await load()
})
</script>
<style lange="scss" scoped>
.content {
 margin: 20px 8px;
}
</style>