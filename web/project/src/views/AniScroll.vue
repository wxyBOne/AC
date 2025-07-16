<template>
    <div class="scroll-container">
        <!-- 介绍 -->
        <div id="part0" ref="part0Ref">
            <transition name="slide0">
                <div class="intro-room" v-show="isShow['0']">
                    <p class="intro-title">ABOUT US</p>
                    <p class="introduce">
                        The ArtChain platform is committed to constructing a reliable art trading system. It effectively
                        guarantees the authenticity and traceability of artworks and infuses robust vitality and unwavering
                        confidence into the market. Selecting AC signifies choosing a brand-new, safe, efficient, and
                        sustainable art trading approach, thereby making your art investment journey even more seamless and
                        reassuring.
                    </p>
                    <!-- <img src="../img/tl(6) 3.png" class="popo"> -->
                </div>
            </transition>
        </div>

        <!-- 特色 -->
        <div id="part1" ref="part1Ref">
            <transition name="slide1">
                <div class="slogan" v-show="isShow['1']">
                    <i class="slo1">True art trade.</i>
                    <i class="slo2">Authentic art link.</i>
                    <!-- 雕塑 -->
                    <img src="../img/VV(1).png" class="sculpture">
                    <div class="left-text">
                        <p class="text-title">No Fakes</p>
                        <p class="text">Professional team authenticates artworks for a reliable trading haven.</p>
                    </div>
                    <div class="right-text">
                        <p class="text-title">No Middlemen</p>
                        <p class="text">Platform offers one-on-one art transactions without middlemen for transparency.</p>
                    </div>
                </div>
            </transition>
        </div>

        <!-- 幻灯片 -->
        <div id="part2" ref="part2Ref">
            <transition name="slide2">
                <div class="meet" v-show="isShow['2']">
                    <div class="meet-title">
                        <p class="symbol">“</p>
                        <p style="margin-top: -30px;">Discover Your Artistic Delight. Meet the Artwork You Love.</p>
                    </div>
                    <!-- 图片 -->
                    <div class="slide-container">
                        <div class="slide-room">
                            <img src="../img/01.jpg" class="slide-image">
                            <p>Vincent van Gogh (荷兰, 1853-1890)</p>
                        </div>
                        <div class="slide-room">
                            <img src="../img/02.jpg" class="slide-image">
                            <p>Vincent van Gogh (荷兰, 1853-1890)</p>
                        </div>
                        <div class="slide-room">
                            <img src="../img/03.jpg" class="slide-image">
                            <p>Vincent van Gogh (荷兰, 1853-1890)</p>
                        </div>
                    </div>
                </div>
            </transition>
        </div>

        <!-- 评论区 -->
        <div id="part3" ref="part3Ref">
            <transition name="slide3">
                <div class="comments-section" v-show="isShow['3']">
                    <div class="back-text">
                        <p style="margin-left: -90vh;">Their Voice</p>
                        <p style="margin:-12vh 0 0 80vh;">Your Choice</p>
                    </div>

                    <div class="comments">
                        <div class="l-card">
                            <img src="../img/back-bottom.jpg" alt="User Avatar" class="avatar">
                            <div class="user-info">
                                <p class="user-name">John Doe</p>
                                <p class="user-email">johndoe@example.com</p>
                                <div class="comment-content">
                                    <p class="comment-text">"AC makes art trading simple and reliable. I like it very much."
                                    </p>
                                    <p class="comment-time">Nov 17, 2024</p>
                                </div>
                            </div>
                        </div>
                        <div class="r-card">
                            <img src="../img/back-bottom.jpg" alt="User Avatar" class="avatar">
                            <div class="user-info">
                                <p class="user-name">Alice Smith</p>
                                <p class="user-email">alicesmith@example.com</p>
                                <div class="comment-content">
                                    <p class="comment-text">"The platform is intuitive and user-friendly. Love it!"</p>
                                    <p class="comment-time">Nov 20, 2024</p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="comments">
                        <div class="l-card">
                            <img src="../img/back-bottom.jpg" alt="User Avatar" class="avatar">
                            <div class="user-info">
                                <p class="user-name">Bob Johnson</p>
                                <p class="user-email">bobjohnson@example.com</p>
                                <div class="comment-content">
                                    <p class="comment-text">"Unique artworks are super cool. Great experience."</p>
                                    <p class="comment-time">Nov 19, 2024</p>
                                </div>
                            </div>
                        </div>
                        <div class="r-card">
                            <img src="../img/back-bottom.jpg" alt="User Avatar" class="avatar">
                            <div class="user-info">
                                <p class="user-name">Eve Williams</p>
                                <p class="user-email">evewilliams@example.com</p>
                                <div class="comment-content">
                                    <p class="comment-text">"Superb! It makes art trading purer."
                                    </p>
                                    <p class="comment-time">Nov 15, 2024</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </transition>
        </div>

        <!-- 底部 -->
        <div id="part4" ref="part4Ref">
            <div class="bottom" v-show="isShow['4']">
                <div class="bottom-title">ArtChain</div>
                <div class="rights">
                    <p style="color: rgba(172, 35, 124);">©2024·AC</p>
                    <p style="color:rgba(203, 255, 265, 0.755)">All Rights Reserved</p>
                </div>
                <img src="../img/logo.png" class="bottom-logo">
                <div class="contact">
                    <p>Enjoy the beauty of art together.</p>
                </div>
                <!-- 进入首页 -->
                <router-link to="/Home" class="into">
                    <img src="../img/arrow-to-right.png" class="arrow-right">
                </router-link>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, inject, defineEmits } from 'vue';

// 从父组件注入 containerRef，用来引用滚动容器
const containerRef = inject('containerRef');
// 缓存部分的 DOM 引用
const part0Ref = ref(null);
const part1Ref = ref(null);
const part2Ref = ref(null);
const part3Ref = ref(null);
const part4Ref = ref(null);

// 用一个对象来保存每个部分的显示状态
const isShow = ref({
    '0': false,
    '1': false,
    '2': false,
    '3': false,
    '4': false
});

// 定义可以发出的自定义事件
const emits = defineEmits(['setStayFalse', 'setStayTrue']);

const onScroll = () => {
    // 获取滚动容器的矩形信息
    const containerRect = containerRef.value.getBoundingClientRect();

    // 存储每个部分的矩形信息
    const sections = [
        { id: "0", el: part0Ref.value },
        { id: "1", el: part1Ref.value },
        { id: "2", el: part2Ref.value },
        { id: "3", el: part3Ref.value },
        { id: "4", el: part4Ref.value },
    ];

    // 判断部分是否已经完全离开视口以及是否进入视口
    sections.forEach((section) => {
        // 获取当前部分的矩形信息
        const sectionRect = section.el.getBoundingClientRect();

        // 判断部分是否离开视口的条件
        const isOutOfView = sectionRect.top - containerRect.top > containerRect.height * 2 / 3 ||
            sectionRect.bottom - containerRect.top < containerRect.height / 3;

        // 判断部分是否进入视口的条件
        const isInView = (sectionRect.top - containerRect.top < containerRect.height / 2 && !isShow.value[section.id]) ||
            (sectionRect.top - containerRect.top < containerRect.height && sectionRect.bottom - containerRect.top > containerRect.height / 2 && !isShow.value[section.id]);

        if (isOutOfView) {
            // 当前部分如果离开视口，则隐藏
            if (isShow.value[section.id]) {
                isShow.value[section.id] = false;
            }

            // 离开子组件部分，触发自定义事件通知父组件
            if (section.id == 0) emits('setStayTrue');

        } else if (isInView) {
            // 判断部分是否进入视口并设置显示状态
            if (!isShow.value[section.id]) {
                isShow.value[section.id] = true;
            }

            // 进入子组件部分，触发自定义事件通知父组件
            if (section.id == 0) emits('setStayFalse');
        }
    });
};

// 在组件挂载时添加滚动事件监听器
onMounted(() => {
    containerRef.value.addEventListener("scroll", onScroll);
});

// 在组件卸载时移除滚动事件监听器
onUnmounted(() => {
    if (containerRef.value) {
        containerRef.value.removeEventListener("scroll", onScroll);
    }
});
</script>


<style scoped>
@import url(../css/aniscroll.css);
</style>