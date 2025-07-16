<!--
 * @Author: wxyBone 13638456960@163.com
 * @Date: 2024-11-05 11:06:41
 * @LastEditors: wxyBone 13638456960@163.com
 * @LastEditTime: 2024-12-27 10:02:20
 * @FilePath: \Vue\dapp_vue\project\src\views\AnimationPage.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div class="container" ref="containerRef">
        <!-- 主页 -->
        <div class="room-out">
            <!-- 加载层 -->
            <loading></loading>

            <!-- 导航栏 -->
            <div class="head">
                <img src="../img/logo.png" class="logo">
                <p class="name">AC</p>
                <router-link to="/Login">
                    <p class="losu">Login/Create Account</p>
                </router-link>
                <div class="Enter">
                    <router-link to="/Home" class="enter_">
                        Enter AC
                    </router-link>
                </div>
            </div>

            <!-- 进入首页 -->
            <div class="Enter_">
                <router-link to="/Home" class="enter_">
                    Explore AC Now!
                </router-link>
            </div>
            <!-- 视觉差容器 -->
            <div id="room" ref="scene">
                <!-- 星星容器 -->
                <div id="stars" ref="starsContainer"></div>
                <!-- 球 -->
                <div class="ball" data-depth="0.4"></div>
                <div class="ball-inner-shadow-1" data-depth="0.2"></div>
                <div class="ball-inner-shadow-2" data-depth="0.2"></div>
                <!-- 标题 -->
                <p class="title">ArtChain</p>

                <Transition name="d">
                    <div class="down" v-show="isStay">
                        <p class="tip">Scorll Down</p>
                        <img src="../img/arrowDown.png" class="arrow-down" alt="arrow">
                    </div>
                </Transition>
            </div>
        </div>

        <!-- 下滑内容部分 -->
        <ScroRoom @setStayFalse="StayFalse" @setStayTrue="StayTrue" />
    </div>
</template>

<script setup>
import { ref, onMounted, provide, onBeforeUnmount } from 'vue';
import Parallax from 'parallax-js';
import loading from "@/components/loading.vue"
import ScroRoom from "./AniScroll.vue"

// 引用 DOM 元素
const containerRef = ref(null);
// 提供 containerRef 给子组件
provide('containerRef', containerRef);

const scene = ref(null);
const starsContainer = ref(null);
const isStay = ref(true);

onMounted(() => {
    // 视觉差容器
    new Parallax(scene.value);

    document.addEventListener('mousemove', moveParticles);

    // 创建星星并初始化星星动画
    createStars();
});

onBeforeUnmount(() => {
    document.removeEventListener('mousemove', moveParticles);
});


const StayFalse = () => {
    // 子组件离开时设置 isStay 为 false
    isStay.value = false;
};
const StayTrue = () => {
    // 子组件进入时设置 isStay 为 true
    isStay.value = true;
};


// 鼠标移动时触发
const moveParticles = (e) => {
    const x = e.pageX;
    const y = e.pageY;

    createParticle(x, y);
};
// 跟随鼠标粒子
const createParticle = (x, y) => {
    const particle = document.createElement('div');
    particle.className = 'particle';

    // 设置初始位置
    particle.style.left = `${x - 5}px`;  // 粒子中心定位
    particle.style.top = `${y - 5}px`;

    // 设置随机大小、透明度和运动方向
    const size = Math.random() * 1.5 + 1; // 随机大小
    const opacity = Math.random() * 0.5 + 0.3; // 随机透明度
    const duration = Math.random() * 1.5 + 1; // 随机动画时长

    particle.style.width = `${size}px`;
    particle.style.height = `${size}px`;
    particle.style.pointerEvents = 'none';
    particle.style.opacity = opacity;

    // 动画效果：粒子会在一段时间内逐渐消失并移动
    particle.style.animation = `particleAnimation ${duration}s linear infinite, particleMovement ${duration}s linear infinite`;

    // 设置粒子运动的方向（加入一些随机位移）
    const randomX = (Math.random() - 0.5) * 100; // 随机水平位移
    const randomY = (Math.random() - 0.5) * 100; // 随机垂直位移

    particle.style.transform = `translate(${randomX}px, ${randomY}px)`;

    // 加入 DOM 中
    document.body.appendChild(particle);

    // 清理粒子
    setTimeout(() => {
        particle.remove();
    }, duration * 1000); // 粒子在动画结束后移除
};


// 创建星尘
const createStars = () => {
    const numberOfStars = 900;
    const styleSheet = document.styleSheets[0]; // 获取页面的第一个样式表

    // 预定义动画（所有星星共享一个动画）
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
        star.className = 'star';
        star.style.width = Math.random() * 1 + 1 + 'px';  // 星星大小
        star.style.height = star.style.width;
        star.style.position = 'absolute';

        // 生成不同角度的路径
        const pathPositionX = Math.random() * 100; // 横向位置（0-100%）
        const pathPositionY1 = 43 + Math.sin(pathPositionX / 100 * Math.PI) * 20; // 第一条路径的纵向位置
        const pathPositionY2 = 55 + Math.cos(pathPositionX / 200 * Math.PI) * 5; // 第二条路径的纵向位置
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

        // 将星星添加到容器
        starsContainer.value.appendChild(star);
    }
};

</script>
  
<style scoped>
@import url(../css/animation.css);
</style>