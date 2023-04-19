<template>
  <div class="content">
    <el-tabs type="border-card" v-model="opkey">
      <el-tab-pane name="query" label="查询Key">
        <Cli :opkey="opkey"></Cli>
      </el-tab-pane>
      <el-tab-pane name="slow" label="查询慢查询">
        <Cli :opkey="opkey"></Cli>  
      </el-tab-pane>
      <el-tab-pane name="big" label="查询大Key">
        <Cli :opkey="opkey"></Cli>  
      </el-tab-pane>
      <el-tab-pane name="hot" label="查询热Key">
        <Cli :opkey="opkey"></Cli>  
      </el-tab-pane>
      <el-tab-pane name="all" label="查询1万Key">
        <Cli :opkey="opkey"></Cli>  
      </el-tab-pane>
      <el-tab-pane name="del" label="删除Key">
        <Cli :opkey="opkey"></Cli>  
      </el-tab-pane>
    </el-tabs>
   </div>
</template>

<script lang="ts" setup>
import Cli from '../../components/cli/Cli.vue'
import { ref } from 'vue';
const opkey = ref("")
</script>
<style lange="scss" scoped>
.content {
 margin: 20px 8px;
}
</style>
