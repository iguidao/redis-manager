<template>
   <div class="content">
    <el-card shadow="never">
      <el-table :data="tableData" stripe style="width: 100%">
        <el-table-column prop="Name" label="参数" width="180" />
        <el-table-column prop="Value" label="值" width="400" />
        <el-table-column prop="Note" label="说名" />
        <el-table-column prop="CreatedAt" label="创建时间" width="120" />
        <el-table-column prop="UpdatedAt" label="修改时间" width="120" />
      </el-table>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import {list} from '../../api/rconfig';
// 

const tableData = ref<any[]>([])
const ctime = ref("")
const load = async () => {
  // console.log(await list())
  let data = (await list()).data
  tableData.value = data.data.lists
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