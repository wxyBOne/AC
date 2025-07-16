<template>
    <div class="b-out-room">
        <header class="header-b">
            <div class="logo-section-b">
                <img src="../img/logo.png" class="logo-icon-b">
                <span class="logo-text-b">AC</span>
            </div>

            <nav class="nav-section-b">
                <router-link v-for="item in navItems" :key="item.path" :to="item.path"
                    :class="['nav-link-b', { 'active-b': currentRoute === item.path }]">
                    {{ item.name }}
                </router-link>
            </nav>

            <div class="user-section-b">
                <div class="icon-wrapper-b">
                    <el-icon>
                        <Setting />
                    </el-icon>
                </div>
                <div class="icon-wrapper-b">
                    <router-link to="/AdminLog">
                        <el-icon style="margin-top:2px">
                            <SwitchButton />
                        </el-icon>
                    </router-link>
                </div>
                <div class="avatar-wrapper-b" @click="togglePopup">
                    <img src="../img/back-bottom.jpg" alt="avatar">
                    <!-- 悬浮框 -->
                    <div class="admin-info-popup" v-show="showPopup">
                        <div class="logout-btn" @click="logout">
                            切换
                        </div>
                        <div class="divider"></div>
                        <div class="admin-content">
                            <div class="admin-avatar">
                                <img src="../img/back-bottom.jpg" alt="avatar">
                            </div>
                            <div class="admin-info">
                                <div class="admin-name">{{ adminName }}</div>
                                <div class="admin-role">欢迎来到管理空间</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </header>
        <!-- 左侧导航栏 -->
        <div id="main-b">
            <div class="side-nav-b">
                <div class="nav-items-b">
                    <div class="nav-item-b active-b">
                        <el-icon>
                            <Grid />
                        </el-icon>
                    </div>
                    <div class="nav-item-b">
                        <router-link to="/Shop">
                            <el-icon>
                                <HomeFilled />
                            </el-icon>
                        </router-link>
                    </div>
                </div>
            </div>

            <router-view id="content-b" />
        </div>
    </div>
</template>
  
<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { Bell, Present } from '@element-plus/icons-vue'

const route = useRoute();
const router = useRouter()
const currentRoute = computed(() => route.path);

const showPopup = ref(false)
// 从localStorage读取真实的管理员名字
const adminName = ref(localStorage.getItem('admin_name') || '未登录')

const togglePopup = () => {
    showPopup.value = !showPopup.value
}

const logout = () => {
    router.push('/AdminLog')
}

const navItems = [
    { name: '用户管理', path: '/UserAdmin' },
    { name: '数据大屏', path: '/Overview' },
    { name: '艺术品审核', path: '/ArtCheck' },
];

onMounted(() => {
    document.addEventListener('click', (e) => {
        const popup = document.querySelector('.admin-info-popup')
        const trigger = document.querySelector('.avatar-wrapper-b') // 触发按钮
        if (popup && trigger && !popup.contains(e.target) && !trigger.contains(e.target)) {
            showPopup.value = false
        }
    })
})
</script>
  
<style>
@import url(../css/b.css);
</style>