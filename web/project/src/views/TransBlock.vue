<template>
    <div class="trans-block-container">
        <div class="trans-nav">
            <button v-for="(item, index) in navItems" :key="index"
                :class="['nav-button', { active: activeNav === item.name }]" @click="setActiveNav(item.name)">
                <router-link :to="`/${item.name}`" class="nav-block">
                    <img :src="item.icon" alt="Icon" class="nav-icon" />
                    <p>{{ item.label }}</p>
                </router-link>
            </button>
        </div>

        <div class="content-area">
            <router-view />
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';

const navItems = [
    { name: 'Published', label: '已发布', icon: require('../img/cube.png') },
    { name: 'Purchased', label: '已购入', icon: require('../img/cube.png') },
    { name: 'Sold', label: '已售出', icon: require('../img/cube.png') },
];

const activeNav = ref(navItems[0].name);
const route = useRoute();

const setActiveNav = (navName) => {
    activeNav.value = navName;
};

// 根据路由参数设置初始导航
onMounted(() => {
    const type = route.query.type;
    if (type) {
        setActiveNav(type);
    }
});
</script>

<style scoped>
@import url(../css/transBlock.css);
</style>