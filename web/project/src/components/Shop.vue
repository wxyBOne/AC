<template>
    <div class="first-container">
        <div class="row-0">
            <p class="title0">Unlock art's mysterious door to endless possibilities</p>
            <p class="text0">Unleash art+blockchain. Creativity & innovation meet.</p>
        </div>

        <div class="light1"></div>
        <div class="light2"></div>

        <div class="row-1">
            <Search />
        </div>

        <!-- 活动播报 -->
        <!-- <div class="notification-container">
            <div class="notification-text-container">
                <img src="../img/megaphone.png" alt="喇叭图标" class="notification-icon" />
                <p class="notification-text">惊喜盲盒火热上线！开启神秘盒子，探索独一无二的数字藏品，邂逅艺术与惊喜的奇妙碰撞。</p>
            </div>
        </div> -->

        <div class="row-2">
            <div class="l-part2">
                <!-- 数据 -->
                <div class="data">
                    <h2>UP</h2>
                    <div class="data-chart"></div>
                </div>
                <!-- 幻灯片 -->
                <div class="slide-img">
                    <img :src="require(`../img/${activeImage}`)" alt="Slide Image" />
                </div>
                <!-- 底部点播 -->
                <div class="dots">
                    <span v-for="(_, index) in images" :key="index"
                        :class="['dot', { active: activeImageIndex === index }]"></span>
                </div>
            </div>
            <!-- 用户 -->
            <div class="r-part2">
                <div class="user-avatar"></div>
                <div class="username">
                    Hi, {{ userInfo.username ? ' ' + userInfo.username : ' 游客' }}
                </div>
                <div class="user-tag">Newbie Explorer</div>
                <router-link to="/Login">
                    <button class="login-register-btn" style="color: rgba(0, 0, 0, 0.459)">
                        {{ userInfo.username ? '切换账号' : '登录 / 注册' }}
                    </button>
                </router-link>
                <div class="vip-row">
                    <div class="vip-tit">会员权益</div>
                    <div class="vip-container">
                        <div class="vip">
                            <div class="vip-img">
                                <img src="../img/coin.png">
                            </div>
                            <p>积分</p>
                        </div>
                        <div class="vip">
                            <div class="vip-img">
                                <img src="../img/badge.png">
                            </div>
                            <p>徽章</p>
                        </div>
                        <div class="vip">
                            <div class="vip-img">
                                <img src="../img/check.png">
                            </div>
                            <p>额度</p>
                        </div>
                        <div class="vip">
                            <div class="vip-img">
                                <img src="../img/gift.png">
                            </div>
                            <p>优惠劵</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="row-3">
            <div class="l-part3">
                <img src="../img/car.gif">
            </div>

            <div class="r-part3">
                <div class="title3">交易记录</div>
                <div class="block-container" @click="checkLogin">
                    <!-- 已发布区块 -->
                    <router-link to="/Published">
                        <div class="block" @click="navigateToTransBlock('Published')">
                            <p>{{ p_sum }}</p>
                            <div class="block-title">
                                <img src="../img/cube.png" alt="Cube Icon">
                                <p>已发布</p>
                            </div>
                        </div>
                    </router-link>
                    <!-- 交易中区块 -->
                    <router-link to="/Purchased">
                        <div class="block" @click="navigateToTransBlock('Purchased')">
                            <p>{{ i_sum }}</p>
                            <div class="block-title">
                                <img src="../img/cube.png" alt="Cube Icon">
                                <p>已购入</p>
                            </div>
                        </div>
                    </router-link>
                    <!-- 交易成功区块 -->
                    <router-link to="/Sold">
                        <div class="block" @click="navigateToTransBlock('Sold')">
                            <p>{{ t_sum }}</p>
                            <div class="block-title">
                                <img src="../img/cube.png" alt="Cube Icon">
                                <p>已售出</p>
                            </div>
                        </div>
                    </router-link>
                </div>
            </div>
        </div>

        <!-- 商品列 -->
        <ProductItem></ProductItem>

        <!-- 底部 -->
        <div class="row-5">
            <p style="color: rgba(86, 109, 109, 0.774);">ArtChain©2024 - All rights reserved</p>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch, computed } from 'vue';
import Search from './Search.vue';
import ProductItem from './ProductItem.vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex'; // 引入 Vuex
import * as echarts from 'echarts';
import axios from 'axios';

const initChart = () => {
    const chartDom = document.querySelector('.data-chart');
    const myChart = echarts.init(chartDom);

    const option = {
        backgroundColor: 'transparent',
        tooltip: {
            trigger: 'axis',
            backgroundColor: 'rgba(26, 34, 35, 0.9)',
            borderColor: 'rgba(255, 255, 255, 0.1)',
            textStyle: {
                color: '#fff'
            },
            axisPointer: {
                type: 'line',
                lineStyle: {
                    color: 'rgba(255, 82, 162, 0.2)',
                    width: 2
                }
            }
        },
        grid: {
            top: '15%',
            left: '2%',
            right: '7%',
            bottom: '8%',
            containLabel: true,
            show: false // 隐藏网格线
        },
        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: ['00:00', '04:00', '08:00', '12:00', '16:00'],
            axisLine: {
                show: false
            },
            axisTick: {
                show: false
            },
            axisLabel: {
                color: 'rgba(168, 206, 210, 0.6)',
                fontSize: 12
            },
            splitLine: {
                show: false // 隐藏x轴分割线
            }
        },
        yAxis: {
            type: 'value',
            min: 0,
            max: 100,
            splitNumber: 2,
            axisLine: {
                show: false
            },
            axisTick: {
                show: false
            },
            axisLabel: {
                color: 'rgba(168, 206, 210, 0.6)',
                fontSize: 12
            },
            splitLine: {
                show: false // 隐藏y轴分割线
            }
        },
        series: [{
            name: 'Turnover',
            type: 'line',
            smooth: true,
            symbol: 'circle',
            symbolSize: 8,
            showSymbol: false,
            lineStyle: {
                width: 3,
                color: new echarts.graphic.LinearGradient(0, 0, 1, 0, [{
                    offset: 0,
                    color: 'rgba(49, 224, 255, 0.4)'
                }, {
                    offset: 1,
                    color: 'rgba(11, 72, 255, 0.5)'
                }])
            },
            itemStyle: {
                color: 'rgba(215, 216, 141, 0.3)',
                borderColor: 'rgba(126, 149, 149, 0.5)',
                borderWidth: 1
            },
            areaStyle: {
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                    offset: 0,
                    color: 'rgba(255, 0, 149, 0.5)'
                }, {
                    offset: 1,
                    color: 'rgba(215, 216, 141, 0.08)'
                }])
            },
            emphasis: {
                scale: true,
                focus: 'series'
            },
            data: [60, 80, 60, 70, 90]
        }]
    };

    myChart.setOption(option)

    // 响应式处理
    window.addEventListener('resize', () => {
        myChart.resize()
    })

    // 添加鼠标移入显示数据点的效果
    myChart.on('mouseover', () => {
        myChart.setOption({
            series: [{
                showSymbol: true
            }]
        })
    })

    myChart.on('mouseout', () => {
        myChart.setOption({
            series: [{
                showSymbol: false
            }]
        })
    })
}

onMounted(() => {
    initChart()
})

const router = useRouter();
const store = useStore(); // 获取 Vuex store
const userInfo = computed(() => store.getters.getUserInfo); // 获取用户信息

const checkLogin = () => {
    const token = store.getters.getToken; // 从 Vuex 获取 token
    if (!token) {
        window.alert('未登录，请重新登录。'); // 提示用户未登录
        router.push('/Login'); // 重定向到登录页面
    }
};

const p_sum = ref(0); // 已发布艺术品数量
const i_sum = ref(0); // 已购入艺术品数量
const t_sum = ref(0); // 已售出艺术品数量

// 传值
const navigateToTransBlock = (type) => {
    router.push({ path: '/TransBlock', query: { type } });
};

const fetchArtStats = async () => {
    try {
        const response = await axios.get('http://localhost:8001/user-art-stats', {
            headers: {
                'Authorization': localStorage.getItem('token') // 从本地存储获取 token
            }
        });
        p_sum.value = response.data.published;
        i_sum.value = response.data.purchased;
        t_sum.value = response.data.sold;
    } catch (error) {
        console.error('获取艺术品统计信息失败:', error);
    }
};

// 图片数组
const images = [
    'nft1.png',
    'nft6.png',
    'nft3.png',
];

// 当前活跃的图片索引
const activeImageIndex = ref(0);

// 获取当前活跃的图片
const activeImage = ref(images[activeImageIndex.value]);

// 切换图片的函数
const changeImage = () => {
    activeImageIndex.value = (activeImageIndex.value + 1) % images.length;
};

// 启动定时器，自动循环播放幻灯片
let interval;
onMounted(() => {
    fetchArtStats(); // 组件挂载时获取艺术品统计信息
    interval = setInterval(changeImage, 3000); // 每2秒切换一次图片
});

// 清除定时器，避免内存泄漏
onUnmounted(() => {
    clearInterval(interval);
});

// 监听当前活跃的索引，动态更新图片
watch(activeImageIndex, (newIndex) => {
    activeImage.value = images[newIndex];
});
</script>

<style scoped>
@import url(../css/home.css);
</style>