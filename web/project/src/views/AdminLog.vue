<template>
    <div class="login-container">
        <div class="login-box">
            <!-- 左侧 Logo 和标题区域 -->
            <div class="left-section">
                <div class="logo-section">
                    <img src="../img/logo.png" class="logo-icon">
                    <span class="logo-text">ArtChain</span>
                </div>
                <h2 class="welcome-text">欢迎回来</h2>
                <p class="sub-text">Please enter your details to sign in</p>
            </div>

            <!-- 右侧登录表单区域 -->
            <div class="form-section">
                <div class="input-group">
                    <label>用户名</label>
                    <div class="input-wrapper">
                        <el-input v-model="name" placeholder="输入用户名" :prefix-icon="User" />
                    </div>
                </div>

                <div class="input-group">
                    <label>密码</label>
                    <div class="input-wrapper">
                        <el-input v-model="password" type="password" placeholder="输入用户密码" :prefix-icon="Lock" />
                    </div>
                </div>

                <el-button type="primary" class="login-btn" @click="handleLogin">
                    登录
                </el-button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { User, Lock, View } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const name = ref('')
const password = ref('')

const handleLogin = async () => {
    if (!name.value || !password.value) {
        window.alert('输入不能为空')
        return
    }

    try {
        const response = await fetch('http://localhost:8001/admin/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                adminname: name.value,
                password: password.value,
            }),
        })

        const data = await response.json()

        if (response.ok) {
            localStorage.setItem('admin_token', data.token)
            localStorage.setItem('admin_id', data.adminID.toString())
            localStorage.setItem('admin_name', data.adminName)
            // 跳转到管理员首页
            router.push('/Overview')
        } else {
            window.alert(data.error || '登录失败')
        }
    } catch (error) {
        window.alert('网络错误，请稍后重试')
    }
}
</script>

<style scoped>
.login-container {
    min-height: 100vh;
    background-color: #0f1314;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 20px;
}

.login-box {
    background-color: rgba(40, 48, 49, 0.223);
    backdrop-filter: blur(20px);
    border-radius: 16px;
    padding: 60px;
    width: 100%;
    max-width: 900px;
    display: flex;
    gap: 60px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.left-section {
    flex: 1;
    padding-right: 40px;
    border-right: 1px solid rgba(168, 206, 210, 0.1);
}

.logo-section {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 40px;
}

.logo-icon {
    width: 28px;
    height: 32px;
}

.logo-text {
    color: rgb(168, 206, 210);
    font-size: 24px;
    font-weight: 600;
}

.welcome-text {
    color: rgb(189, 208, 211);
    font-size: 32px;
    font-weight: 600;
    margin-bottom: 16px;
}

.sub-text {
    color: rgba(168, 206, 210, 0.6);
    font-size: 15px;
}

.form-section {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 15px;
}

.input-group {
    display: flex;
    flex-direction: column;
    gap: 9px;
}

.input-group label {
    color: rgba(168, 206, 210, 0.8);
    font-size: 16px;
}

.input-wrapper :deep(.el-input__wrapper) {
    background-color: rgba(168, 206, 210, 0.05);
    box-shadow: none;
    border: 1px solid rgba(168, 206, 210, 0.1);
}

.input-wrapper :deep(.el-input__wrapper:hover) {
    border-color: rgba(168, 206, 210, 0.2);
}

.input-wrapper :deep(.el-input__wrapper.is-focus) {
    border-color: #a8ced286;
}

.input-wrapper :deep(.el-input__inner) {
    color: rgb(189, 208, 211);
    height: 42px;
}

.input-wrapper :deep(.el-input__prefix-inner),
.input-wrapper :deep(.el-input__suffix-inner) {
    color: rgba(168, 206, 210, 0.5);
}

.options-row {
    margin-top: -35px;
}

:deep(.el-checkbox__label) {
    color: rgba(168, 206, 210, 0.8);
}

:deep(.el-checkbox__input.is-checked .el-checkbox__inner) {
    background-color: #a8ced2;
    border-color: #a8ced2;
}

.forgot-link {
    color: #a8ced2d4;
    float: right;
    text-decoration: none;
    font-size: 14px;
}

.login-btn {
    width: 100%;
    border-radius: 5px;
    height: 42px;
    font-size: 16px;
    color: rgba(168, 206, 210, 0.781);
    background-color: rgba(107, 148, 155, 0.877);
    border: none;
    font-weight: 600;
    transition: all 0.2s ease-in-out;
}

.login-btn:hover {
    background-color: #8fb8bda7;
}

.signup-section {
    text-align: center;
    color: rgba(168, 206, 210, 0.6);
    font-size: 14px;
}

.signup-link {
    color: #a8ced2;
    text-decoration: none;
    margin-left: 8px;
}

/* 响应式设计 */
@media (max-width: 768px) {
    .login-box {
        flex-direction: column;
        gap: 40px;
        padding: 30px;
    }

    .left-section {
        padding-right: 0;
        border-right: none;
        border-bottom: 1px solid rgba(168, 206, 210, 0.1);
        padding-bottom: 30px;
    }
}
</style>