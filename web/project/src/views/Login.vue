<template>
    <div class="login-container">
        <!-- 星星容器 -->
        <div id="l-stars" ref="l_starsContainer"></div>
        <div class="login-room-out">
            <div class="logo-row">
                <img src="../img/logo.png" alt="Logo" class="login-logo" />
            </div>
            <div class="login-content">
                <div class="toggle-border">
                    <div class="toggle-row">
                        <button :class="{ active: isLogin }" @click="isLogin = true">登录</button>
                        <button :class="{ active: !isLogin }" @click="isLogin = false">注册</button>
                    </div>
                </div>

                <div class="title-row">
                    欢迎来到ArtChain平台
                </div>
                <div class="subtitle-row">
                    {{ isLogin ? '登录，开启您的艺术品价值保障之旅' : '注册，开启您的艺术品价值保障之旅' }}
                </div>
                <div class="name-input-row">
                    <div class="input-container">
                        <input type="text" placeholder="输入用户名" v-model="username" class="input-field" />
                    </div>
                </div>
                <div class="password-input-row">
                    <div class="input-container">
                        <input :type="showPassword ? 'text' : 'password'" placeholder="输入密码" v-model="password"
                            class="input-field" />
                        <el-icon @click="showPassword = !showPassword" class="eye-icon">
                            <i v-if="showPassword" class="el-icon-view">
                                <View />
                            </i>
                            <i v-else class="el-icon-view-off">
                                <Hide />
                            </i>
                        </el-icon>
                    </div>
                </div>
                <div v-if="!isLogin" class="confirm-password-input-row" v-show="!isLogin">
                    <div class="input-container">
                        <input :type="showConfirmPassword ? 'text' : 'password'" placeholder="再次输入密码"
                            v-model="confirmPassword" class="input-field" />
                        <el-icon @click="showConfirmPassword = !showConfirmPassword" class="eye-icon">
                            <i v-if="showConfirmPassword" class="el-icon-view">
                                <View />
                            </i>
                            <i v-else class="el-icon-view-off">
                                <Hide />
                            </i>
                        </el-icon>
                    </div>
                </div>
                <div class="forget-password-row">
                    <p class="forget-password" @click="showResetModal = true">忘记密码?</p>
                </div>
                <div class="login-button-row">
                    <button class="login-button" @click="handleAction">
                        {{ isLogin ? '登录' : '注册' }}
                    </button>
                </div>
            </div>
        </div>

        <!-- 重置密码弹窗 -->
        <div v-if="showResetModal" class="modal">
            <div class="modal-content">
                <span class="close" @click="showResetModal = false">&times;</span>
                <h2>重置用户密码</h2>
                <div class="input-container">
                    <input type="text" placeholder="输入私钥" v-model="resetPrivateKey" class="input-field" />
                </div>
                <div class="input-container">
                    <input :type="showNewPassword ? 'text' : 'password'" placeholder="输入新密码" v-model="newPassword"
                        class="input-field" />
                    <el-icon @click="showNewPassword = !showNewPassword" class="eye-icon">
                        <i v-if="showNewPassword" class="el-icon-view">
                            <View />
                        </i>
                        <i v-else class="el-icon-view-off">
                            <Hide />
                        </i>
                    </el-icon>
                </div>
                <button @click="resetPassword">确定</button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex'; // 引入 Vuex
import axios from 'axios';

const router = useRouter();  // 获取路由实例
const store = useStore(); // 获取 Vuex store

const isLogin = ref(true);
const username = ref('');
const password = ref('');
const confirmPassword = ref('');
const showPassword = ref(false); // 控制密码可见性
const showConfirmPassword = ref(false);
const showNewPassword = ref(false);
const showResetModal = ref(false);
const resetPrivateKey = ref('');
const newPassword = ref('');

const handleAction = async () => {
    if (isLogin.value) {
        // 登录逻辑
        if (!username.value || !password.value) {
            window.alert('用户信息不能为空！'); // 输入框不可为空
            return;
        }

        try {
            const response = await axios.post('http://localhost:8001/login', {
                username: username.value,
                password: password.value
            });

            if (response.data) {
                window.alert('登录成功！');
                // 保存 token 和用户信息到本地存储
                localStorage.setItem('token', response.data.token);
                localStorage.setItem('username', response.data.username);
                localStorage.setItem('privateKey', response.data.privateKey);
                // 更新用户信息到 Vuex
                store.dispatch('updateUserInfo', {
                    username: response.data.username,
                    privateKey: response.data.privateKey
                });
                setTimeout(() => {
                    router.push('/Shop');
                }, 1000); // 延时 1 秒后跳转
            }
        } catch (error) {
            if (error.response.status === 401) {
                window.alert('用户未注册或密码错误！'); // 用户未注册或密码错误
            } else {
                window.alert('登录失败！请检查用户名和密码。');
            }
        }
    } else {
        // 注册逻辑
        if (!username.value || !password.value || !confirmPassword.value) {
            window.alert('用户信息不能为空！'); // 输入框不可为空
            return;
        }

        if (password.value.length < 6) {
            window.alert('密码不能少于6位！'); // 密码长度检查
            return;
        }

        if (password.value !== confirmPassword.value) {
            window.alert('两次输入的密码不一致！'); // 密码一致性检查
            return;
        }

        try {
            const response = await axios.post('http://localhost:8001/register', {
                username: username.value,
                password: password.value
            });

            if (response.data) {
                window.alert('注册成功！');
                window.alert(`您的私钥是：${response.data.privateKey}\n请务必牢记，私钥如同您账户的钥匙，一旦丢失可能无法找回。`); // 显示用户私钥
                isLogin.value = true; // 切换回登录状态
            }
        } catch (error) {
            if (error.response.status === 409) {
                window.alert('用户名已存在！'); // 用户名重复检查
            } else {
                window.alert('注册失败！请重试。');
            }
        }
    }
};

// 重置密码函数
const resetPassword = async () => {
    // 表单验证
    if (!resetPrivateKey.value || !newPassword.value) {
        window.alert('私钥和新密码不能为空')
        return
    }

    // 验证密码长度
    if (newPassword.value.length < 6) {
        window.alert('密码长度不能少于6位')
        return
    }

    try {
        const response = await axios.post('http://localhost:8001/reset-password', {
            privateKey: resetPrivateKey.value,
            newPassword: newPassword.value
        })

        if (response.status === 200) {
            window.alert('密码重置成功')
            // 重置表单
            resetPrivateKey.value = ''
            newPassword.value = ''
            // 关闭弹窗
            showResetModal.value = false
        }
    } catch (error) {
        if (error.response) {
            // 处理后端返回的错误信息
            window.alert(error.response.data.error || '密码重置失败')
        } else {
            // 处理网络错误
            window.alert('网络错误，请稍后重试')
        }
    }
}

const l_starsContainer = ref(null);

onMounted(() => {
    // 创建星星并初始化星星动画
    createStars();
});

// 创建星尘
const createStars = () => {
    const numberOfStars = 900;
    const styleSheet = document.styleSheets[0]; // 获取页面的第一个样式表

    // 预定义动画（所有星星共享一个动画
    const keyframes = `
        @keyframes float {
            0% {
                transform: translateY(0px);
            }
            50% {
                transform: translateY(-20px);
            }
            100% {
                transform: translateY(0px);
            }
        }
    `;
    styleSheet.insertRule(keyframes, styleSheet.cssRules.length);  // 插入一个公共的 @keyframes 动画

    for (let i = 0; i < numberOfStars; i++) {
        const star = document.createElement('div');
        star.className = 'l-star';
        star.style.width = Math.random() * 1 + 1 + 'px';  // 星星大小
        star.style.height = star.style.width;
        star.style.position = 'absolute';

        // 生成不同角度路径
        const pathPositionX = Math.random() * 100; // 横向位置（0-100%）
        const pathPositionY1 = 43 + Math.sin(pathPositionX / 100 * Math.PI) * 20; // 第一条路径的纵向位置
        const pathPositionY2 = 55 + Math.cos(pathPositionX / 200 * Math.PI) * 5; // 二条路径的向位置
        const pathPositionY3 = Math.random() * 10 + 55;
        const pathPositionY4 = Math.random() * 20 + 55;

        // 随机选择路径1、路径2或路径3
        const pathPositionY = Math.random() > 0.8 ? pathPositionY1 : (Math.random() > 0.8 ? pathPositionY2 : (Math.random() > 0.5 ? pathPositionY4 : pathPositionY3));

        // 设置星星位置
        star.style.left = pathPositionX + 'vw';  // 横向位置，使用 vw（视口宽度）单位
        star.style.top = pathPositionY + 'vh';   // 纵向位置，模拟弯曲轨迹

        // 动画持续时间和延迟
        star.style.animationDuration = Math.random() * 5 + 5 + 's';  // 动画持续时间
        star.style.animationDelay = Math.random() * 3 + 3 + 's';  // 动画延迟时间

        // 随机透明度
        star.style.opacity = Math.random();

        // 为星星添加浮动动画（所有星星共享同一个动画）
        star.style.animation = `float ${Math.random() * 5 + 5}s ease-in-out infinite`;

        // 将星星添加到器
        l_starsContainer.value.appendChild(star);
    }
};
</script>

<style scoped>
@import url(../css/login.css);
</style>