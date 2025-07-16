<template>
    <transition name="showCenter">
        <div class="center-container" v-if="showUserCenter">
            <div class="center-content">
                <!-- 左侧 -->
                <div class="center-img">
                    <img :src="require(`../img/${userImage}`)" alt="User center">
                </div>

                <div class="center-tit">
                    <p class="centerName">{{ userInfo.username ? userInfo.username : 'Guest' }}</p>
                    <p class="centerTag">@Newbie Explorer</p>
                </div>

                <div class="centerCancel" @click="cancelEdit">
                    <img src="../img/cancel.png">
                </div>

                <!-- 展示页 -->
                <div class="center-info" v-if="!isEditing">
                    <button id="to-setting-user" @click="isEditing = true">
                        修改信息
                    </button>
                </div>

                <!-- 修改页 -->
                <div class="center-info" v-if="isEditing">
                    <div class="input-group" style="margin-top: 50px;">
                        <label for="username">用户名*</label>
                        <input type="text" id="username" placeholder="输入用户名">
                    </div>
                    <div class="input-group">
                        <label for="password">密码*</label>
                        <input type="password" id="password" placeholder="输入密码">
                    </div>
                    <div class="setting-but">
                        <button id="setting-user" @click="saveEdit" style="margin-right: 40px;">确定</button>
                        <button id="setting-user" @click="isEditing = false">取消</button>
                    </div>
                </div>
            </div>
        </div>
    </transition>
</template>

<script setup>
import { ref, inject, computed } from 'vue';
import { useStore } from 'vuex'; // 引入 Vuex
import router from '@/router';
import axios from 'axios';

const store = useStore(); // 获取 Vuex store
const userInfo = computed(() => store.getters.getUserInfo); // 获取用户信息
const showUserCenter = inject('showUserCenter');
const userImage = ref('back-bottom.jpg');
const isEditing = ref(false);

const cancelEdit = () => {
    showUserCenter.value = false;
    isEditing.value = false;
};

const saveEdit = async () => {
    // 输入验证
    if (username.value.trim() === "") {
        window.alert('用户信息不能为空！'); // 用户名不能为空
        return;
    }

    if (password.value.trim() === "") {
        window.alert('用户信息不能为空！'); // 密码不能为空
        return;
    }

    // 这里可以添加对新密码的验证
    if (password.value.length < 6) {
        window.alert('密码长度必须至少为6位！');
        return;
    }

    // 这里可以调用后端接口更新用户信息
    try {
        const token = localStorage.getItem('token'); // 从 Local Storage 获取 token
        const response = await axios.post('http://localhost:8001/update-user', {
            username: username.value,
            password: password.value // 发送新密码
        }, {
            headers: {
                Authorization: token // 将 token 添加到请求头
            }
        });

        if (response.data) {
            window.alert('修改成功！');
            localStorage.setItem('username', username.value);
            showUserCenter.value = false;
            isEditing.value = false;
            // 更新用户信息到 Vuex
            store.dispatch('updateUserInfo', {
                username: username.value, // 更新用户名
                privateKey: userInfo.value.privateKey // 保持私钥不变
            });
        }
    } catch (error) {
        if (error.response.status === 401) {
            window.alert('未登录，请重新登录。');
            router.push('/Login');
        } else if (error.response.status === 409) {
            window.alert('用户名已存在！'); // 用户名重复检查
        } else {
            window.alert('修改失败，请重试。');
        }
    }
};
</script>

<style scoped>
@import url(../css/usercenter.css);
</style>
