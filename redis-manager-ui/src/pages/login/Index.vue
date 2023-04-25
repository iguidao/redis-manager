<template>
    <div class="login">
        <!-- 引入卡片 -->
        <el-card class="login_card" shadow="always">
            <el-image class="logo_image" :src="`/src/assets/img/logo.png`" fit="cover"></el-image>
            <p class="login_title">登 录</p>
            <p class="login_desc">欢迎登录Redis-Manager平台</p>
            <el-form ref="ruleFormRef" :model="loginForm" :rules="rules">
                <el-form-item prop="username">
                    <el-input placeholder="请输入账号" v-model="loginForm.username" prefix-icon="User" />
                </el-form-item>
                <el-form-item prop="password">
                    <el-input type="password" placeholder="请输入密码" v-model="loginForm.password" prefix-icon="Lock"/>
                </el-form-item>
                <el-form-item>
                    <el-button :loading=false type="primary" @click="onLoginIn()">登录</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <div>
            <el-image class="github_logo" :src="`/src/assets/img/github-mark.png`" fut="civer" @click="ToGithub()"></el-image>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { reactive, ref, getCurrentInstance } from "vue";
import { ElMessage } from "element-plus";
import router from "../../router/index";
import {login} from '../../api/user';
const loginForm = reactive({
    username: '',
    password: ''
})
const ruleFormRef = ref();
const rules = reactive({
  username: [{ required: true, message: "账号不能为空", trigger: "blur" }],
  password: [{ required: true, message: "密码不能为空", trigger: "blur" }],
});
const onLoginIn = async () => {
  // router.push("/home");
    if (!ruleFormRef) return;
      ruleFormRef.value.validate(async (valid: any) => {
        if (valid) {
          const res = await login(loginForm);
          console.log("res: ",res.data.data)
          if (res.data.errorCode === 0 ) {
              sessionStorage.setItem('Authorization', res.data.data.token);
              sessionStorage.setItem('user_name',res.data.data.username);
              sessionStorage.setItem('user_type',res.data.data.usertype);
              // console.log("res token: ",res.data.data.token);
              router.push("/home");
          } else {
              ElMessage(res.data.msg);
          }
        } else {
          return false
        }
    });
}
const ToGithub = () => {
    window.open('https://github.com/iguidao/redis-manager', '_blank');
}
</script>

<style scoped>
.login {
    width: 100%;
    height: 100%;
    background-image: url('../../assets/img/sun.jpeg');
    position: absolute;
    left: 0;
    top: 0;
    background-size: 100% 100%;
}
.login_card {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    margin: auto;
    width: 20%;
    min-width: 300px;
    height: 450px;
    min-height: 450px;
    border-radius: 10px;
    text-align: center;
}
.logo_image {
  width: 50px;
  height: 50px;
  margin-top: 40px;
}
.login_title {
  font-size: 25px;
  font-weight: bold;
}
.login_desc {
  letter-spacing: 2px;
  color: #999a9a;
}

.el-button {
  width: 100%;
}

.github_logo {
  width: 40px;
  height: 40px;
  position: fixed;
  bottom: 60px;
  right: 20px;
}

</style>
