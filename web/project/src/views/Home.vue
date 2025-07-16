<!--
 * @Author: wxyBone 13638456960@163.com
 * @Date: 2024-06-16 18:48:36
 * @LastEditors: wxyBone 13638456960@163.com
 * @LastEditTime: 2024-12-25 17:01:26
 * @FilePath: \project\src\views\Home.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div class="home-container">
    <div class="home-room-out">
      <!-- 导航栏 -->
      <div class="navi">
        <router-link to="/">
          <img src="../img/logo.png" class="h-logo">
        </router-link>
        <p class="h-name">AC</p>
        <router-link to="/Login">
          <img src="../img/user.png" class="h-user" style="margin-right: 20px;">
        </router-link>
        <router-link to="/B">
          <img src="../img/bell1.png" class="h-mes">
        </router-link>
      </div>
      <div class="room">
        <router-view></router-view>
      </div>

      <!-- 悬浮框 -->
      <div class="floating-box">
        <div class="floating-item" @click="checkLogin">
          <img class="floating-icon" src="../img/smile.png" alt="My Icon" @click="showUserCenter = true">
          <p class="floating-text">个人</p>
        </div>

        <router-link to="/Cart" @click="checkLogin">
          <div class="floating-item">
            <img class="floating-icon" src="../img/cart.png" alt="Cart Icon">
            <p class="floating-text">购物车</p>
          </div>
        </router-link>

        <router-link to="/Shop">
          <div class="floating-item">
            <img class="floating-icon" src="../img/home.png" alt="home Icon" style="width: 25px;height:25px;">
            <p class="floating-text">首页</p>
          </div>
        </router-link>

        <router-link to="/Publish" @click="checkLogin">
          <div class="floating-item">
            <img class="floating-icon" src="../img/share.png" alt="Publish Icon">
            <p class="floating-text">发布</p>
          </div>
        </router-link>
      </div>

      <UserCenter></UserCenter>

    </div>
  </div>
</template>
  
<script setup>
import { ref, provide } from 'vue';
import { useStore } from 'vuex'; // 引入 Vuex
import { useRouter } from 'vue-router'; // 引入 Vue Router
import UserCenter from '@/components/UserCenter.vue';
const showUserCenter = ref(false);
// 子组件改变父组件传来的值，父组件的值会同步变化
provide('showUserCenter', showUserCenter);

const store = useStore(); // 获取 Vuex store
const router = useRouter(); // 获取路由实例

const checkLogin = () => {
  const token = store.getters.getToken; // 从 Vuex 获取 token
  if (!token) {
    window.alert('未登录，请重新登录。'); // 提示用户未登录
    router.push('/Login'); // 重定向到登录页面
  }
};
</script>
  
<style scoped>
@import url(../css/home.css);
</style>
  


