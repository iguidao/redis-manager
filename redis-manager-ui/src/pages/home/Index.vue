<template>
  <div>
    <el-container class="home-container">
      <!-- header -->
      <el-header>
        <el-row>
          <el-col :span="2">
            <div class="grid-content logo-wrapper">
              <el-icon  v-model="isCollapse" @click="isCollapse = !isCollapse">
                <Fold v-if="isCollapse"/>
                <Expand v-if="!isCollapse"/>
              </el-icon>
              <p class="system_name">Redis Manager</p>
            </div>
          </el-col>
          <el-col :offset="14" :span="8" style="min-width: 150px">
              <el-dropdown style="float: right; margin: 20px 10px">
                  <span class="system_name">iguidao</span>
                  <template #dropdown>
                      <el-dropdown-menu>
                          <el-dropdown-item @click.native="logout">退出系统</el-dropdown-item>
                      </el-dropdown-menu>
                  </template>
              </el-dropdown>
            </el-col>
        </el-row>
      </el-header>
      <el-container style="overflow: auto">
        <!-- 菜单 -->
        <el-aside>
          <el-menu
            active-text-color="#ffd04b"
            background-color="#545c64"
            class="el-menu-vertical-demo"
            text-color="#fff"
            :collapse="isCollapse"
            router :default-active="activePath"
          >
            <el-menu-item  index="/dashboard" @click="saveActiveNav('/dashboard')">
                <el-icon><HomeFilled /></el-icon>
                <span>概览</span>
            </el-menu-item >
            <el-sub-menu index="1">
                <template #title>
                  <el-icon><TakeawayBox /></el-icon>
                  <span>Redis集群</span>
                </template>
                <el-menu-item index="/redis/index" @click="saveActiveNav('/redis/index')">
                  <el-icon><CaretRight /></el-icon>
                  <span>自建Redis集群</span>
                </el-menu-item>
                <el-menu-item index="/txredis/index" @click="saveActiveNav('/txredis/index')">
                  <el-icon><CaretRight /></el-icon>
                  <span>腾讯Redis集群</span>
                </el-menu-item>
                <el-menu-item index="/aliredis/index" @click="saveActiveNav('/aliredis/index')">
                  <el-icon><CaretRight /></el-icon>
                  <span>阿里Redis集群</span>
                </el-menu-item>
                <el-menu-item index="/codis/index" @click="saveActiveNav('/codis/index')">
                  <el-icon><CaretRight /></el-icon>
                  <span>自建Codis集群</span>
                </el-menu-item>
            </el-sub-menu>
            <el-menu-item index="/command/query" @click="saveActiveNav('/command/query')">
              <el-icon><MagicStick /></el-icon>
              <span>数据查询</span>
            </el-menu-item>
            <el-sub-menu index="2">
              <template #title>
                <el-icon><User /></el-icon>
                <span>用户管理</span>
              </template>
              <el-menu-item index="/user/index" @click="saveActiveNav('/user/index')">
                <el-icon><CaretRight /></el-icon>
                <span>用户列表</span>
              </el-menu-item>
              <el-menu-item index="/user/permission" @click="saveActiveNav('/user/permission')">
                <el-icon><CaretRight /></el-icon>
                <span>权限管理</span>
              </el-menu-item>
            </el-sub-menu>
            <el-sub-menu index="3">
              <template #title>
                <el-icon><Setting /></el-icon>
                <span>系统设置</span>
              </template>
              <el-menu-item index="/setting/index" @click="saveActiveNav('/setting/index')">
                <el-icon><CaretRight /></el-icon>
                <span>全局配置</span>
              </el-menu-item>
            </el-sub-menu>
            <el-menu-item index="/history/index" @click="saveActiveNav('/history/index')">
              <el-icon><Tickets /></el-icon>
              <span>历史记录</span>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <el-container>
          <el-main>
            <!-- 面包屑 -->
            <Breadcrumb></Breadcrumb>
            <!-- 主要内容 -->
            <router-view></router-view>
          </el-main>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import Breadcrumb from '../../components/breadcrumb/Breadcrumb.vue'
import { onBeforeMount, ref } from 'vue'
import router from "../../router/index"

// 挂载 DOM 之前
onBeforeMount(() => {
    activePath.value || null == sessionStorage.getItem("activePath")
        ? sessionStorage.getItem("activePath")
        : "/home"
})
let isCollapse = ref(false);
let activePath = ref("");
// 保存链接的激活状态
const saveActiveNav = (path: string) => {
    sessionStorage.setItem("activePath", path);
    activePath.value = path;
}
const logout = () => {
    // 清除缓存
    sessionStorage.clear();
    router.push("/login");
}
</script>

<style lange="scss" scoped>
/* header */
.home-container {
  position: absolute;
  height: 100%;
  top: 0px;
  left: 0px;
  width: 100%;
  background: #f2f3f5;
}

.el-header {
  background: #ffffff;
  padding: 0 10px;
  overflow: hidden;
}

.grid-content {
  min-height: 40px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.logo-wrapper {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 150px;
  cursor: pointer;
  padding-left: 22px;
}
.system_name {
  color: #000000;
  font-size: 18px;
}
/* other */
.el-aside {
  background: #545c64;
  width: auto !important;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
.el-menu-item.is-active {
  color: #fff !important;
  font-size: 15px;
  font-weight: bold;
  background-color: #a7a7a7 !important;
  border-radius: 2px;
  height: 50px;
  line-height: 50px;
  box-sizing: border-box;
  margin: 2px 5px 0px 2px;
}
</style>