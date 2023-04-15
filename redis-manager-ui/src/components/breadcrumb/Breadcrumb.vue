<template>
  <div>
    <!-- 面包屑 -->
    <el-breadcrumb separator="/">
      <el-breadcrumb-item v-for="item in routers" :key="item.path" :to="{ path: item?.path }">
            {{item?.meta?.title}}
      </el-breadcrumb-item>
    </el-breadcrumb>
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import router from '../../router';
// console.log(router.currentRoute.value.matched)
const routers = computed(()=>{
    // 过滤掉没有meta的 
    return router.currentRoute.value.matched.filter(item=>item.meta.title)
})
</script>

<style lange="scss" scoped>
.content {
  margin: 20px 8px;
}
</style>