<template>
    <div class="scrub" @click="startDisappear" :class="{ 'fade-out': isDisappearing }">
        <div class="load">
            <div class="loader">
                <span v-for="i in 15" :key="i"></span>
            </div>
            <!-- 百分比显示 -->
            <div class="percentage">
                <span v-if="percentage <= 100">{{ percentage }}%</span>
                <span v-else>Click anywhere to continue</span>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

// 控制 scrup 消失的状态
const isDisappearing = ref(false);
const isEnd = ref(false);
// 用来显示加载百分比
const percentage = ref(0);

// 更新百分比的函数
const updatePercentage = () => {
    if (percentage.value < 105) {
        percentage.value += 5; // 每次增加5%
    }
};

// 设置定时器每100ms更新一次百分比
onMounted(() => {
    const interval = setInterval(() => {
        updatePercentage();
        if (percentage.value >= 105) {
            clearInterval(interval); // 当达到105%时，停止更新
            isEnd.value = true;
        }
    }, 200); // 每200ms更新一次
});

// 点击 scrub 时触发消失动画
const startDisappear = () => {
    if (isEnd.value) {
        isDisappearing.value = true;
    }
};
</script>

<style scoped>
.load {
    margin: 0px 0 0 100px;
}

.scrub {
    position: sticky;
    top: 0;
    background-color: rgb(0, 0, 0);
    width: 100%;
    height: 100vh;
    backdrop-filter: blur(50px);
    z-index: 10;
    transition: all 2s ease-in-out;
}

/* 进入页 */
.fade-out {
    opacity: 0;
    display: none;
    filter: blur(20px);
    pointer-events: none;
}

.percentage {
    font-size: 65px;
    font-weight: 700;
    color: rgba(52, 95, 98, 0.5);
    text-shadow: 3px 3px 2px rgba(126, 59, 103, 0.4);
    letter-spacing: 2px;
    font-style: italic;
    white-space: nowrap;
    text-align: right;
    bottom: 50px;
    margin-right: 50px;
    display: block;
}

/* 动画过渡效果 */
.percentage span {
    display: inline-block;
    animation: fade 1.5s infinite ease-in-out;
}

@keyframes fade {
    0% {
        opacity: 0.5;
    }

    50% {
        opacity: 1;
    }

    100% {
        opacity: 0.5;
    }
}

.loader {
    width: 300px;
    height: 380px;
    opacity: 0.5;
}

.loader span {
    display: block;
    background: rgba(190, 113, 150, 0.202);
    width: 3px;
    height: 115%;
    border-radius: 50%;
    margin-right: 38px;
    float: left;
}

.loader span:last-child {
    margin-right: 0px;
}

.loader span:nth-child(1) {
    animation: load 2.5s 1.4s infinite linear;
}

.loader span:nth-child(2) {
    animation: load 2.5s 1.2s infinite linear;
}

.loader span:nth-child(3) {
    animation: load 2.5s 1s infinite linear;
}

.loader span:nth-child(4) {
    animation: load 2.5s 0.8s infinite linear;
}

.loader span:nth-child(5) {
    animation: load 2.5s 0.6s infinite linear;
}

.loader span:nth-child(6) {
    animation: load 2.5s 0.4s infinite linear;
}

.loader span:nth-child(7) {
    animation: load 2.5s 0.2s infinite linear;
}

.loader span:nth-child(8) {
    animation: load 2.5s 0s infinite linear;
}

.loader span:nth-child(9) {
    animation: load 2.5s 0.2s infinite linear;
}

.loader span:nth-child(10) {
    animation: load 2.5s 0.4s infinite linear;
}

.loader span:nth-child(11) {
    animation: load 2.5s 0.6s infinite linear;
}

.loader span:nth-child(12) {
    animation: load 2.5s 0.8s infinite linear;
}

.loader span:nth-child(13) {
    animation: load 2.5s 1s infinite linear;
}

.loader span:nth-child(14) {
    animation: load 2.5s 1.2s infinite linear;
}

.loader span:nth-child(15) {
    animation: load 2.5s 1.4s infinite linear;
}

@keyframes load {
    0% {
        background: #ff00a611;
        transform: scaleY(0.08);
    }

    50% {
        background: rgba(112, 237, 233, 0.25);
        transform: scaleY(1);
    }

    100% {
        background: #ff00a611;
        transform: scaleY(0.08);
    }
}

/* 定义fade过渡动画 */
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.5s ease-in-out;
}

.fade-enter,
.fade-leave-to {
    opacity: 0;
}
</style>
