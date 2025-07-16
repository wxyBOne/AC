<template>
    <div class="row-4">
        <!-- 第一行：标题和筛选 -->
        <div class="shop-row">
            <p class="shop-title">今日热卖</p>
            <div class="shop-options">
                <span class="shop-option" @click="sortBy('all')"
                    :class="{ 'active-option': selectedSort === 'all' }">综合</span>
                <span class="shop-option" @click="sortBy('price')"
                    :class="{ 'active-option': selectedSort === 'price' }">价格</span>
                <span class="shop-option" @click="sortBy('new')"
                    :class="{ 'active-option': selectedSort === 'new' }">最新</span>
            </div>
        </div>
        <!-- 第二行：商品列 -->
        <div class="product-row">
            <div class="product-card" v-for="product in paginatedProducts" :key="product.aid">
                <router-link :to="{ name: 'Product', params: { id: product.aid } }">
                    <!-- <img class="product-image" :src="require(`../upload-img/${product.art_image}`)" alt="Product Image"> -->
                    <LazyImage 
                        :src="require(`../upload-img/${product.art_image}`)"
                        :alt="product.title"
                        img-class="product-image"
                        @load="onProductImageLoad(product.aid)"
                        @error="onProductImageError(product.aid)"
                    />
                </router-link>
                <router-link :to="{ name: 'Product', params: { id: product.aid } }">
                    <p class="product-title">{{ product.title }}</p>
                </router-link>
                <p class="product-price">¥{{ product.art_price.toFixed(2) }}</p>
                <p class="product-creator">Created By @{{ product.creator_name }}</p>
            </div>
        </div>
        <!-- 分组件 -->
        <div class="pagination">
            <span class="page-link" @click="prevPage" :disabled="currentPage === 1">
                <img src="../img/page-l.png">
            </span>
            <span class="page-link" v-for="n in totalPages" :key="n" :class="{ 'active-page': currentPage === n }"
                @click="goToPage(n)">{{ n }}</span>
            <span class="page-link" @click="nextPage" :disabled="currentPage === totalPages">
                <img src="../img/page-r.png">
            </span>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import LazyImage from './LazyImage.vue'; // 导入懒加载组件

// 用于存储从后端获取的所有艺术品数据（商品数据），这里直接使用它来接收初始数据并作为后续操作的数据源
const artPieces = ref([]);
// 当前页显示的商品列表（经过分页处理后）
const paginatedProducts = ref([]);
// 选中的排序条件
const selectedSort = ref('');
// 每一页显示的商品数量
const itemsPerPage = ref(8);
// 当前的页码
const currentPage = ref(1);
// 总页数
const totalPages = ref(1);

// 获取艺术品数据（商品数据），从后端接口获取并进行相关初始化操作
const fetchArtPieces = async () => {
    try {
        const response = await axios.get('http://localhost:8001/art-pieces'); // 调用后端接口获取商品数据
        artPieces.value = response.data; // 存储后端返回的商品数据
        sortBy('all');
        // 排序后，重新计算总页数
        updateTotalPages();
    } catch (error) {
        console.error('获取艺术品失败:', error); // 处理错误
    }
};

// 计算总页数，接收商品数据列表作为参数
const updateTotalPages = () => {
    totalPages.value = Math.ceil(artPieces.value.length / itemsPerPage.value);
};

const updatePaginatedProducts = () => {
    const startIndex = (currentPage.value - 1) * itemsPerPage.value; // 计算当前页商品的起始索引
    const endIndex = startIndex + itemsPerPage.value; // 计算当前页商品的结束索引
    paginatedProducts.value = artPieces.value.slice(startIndex, endIndex); // 更新当前页的商品列表
};

// 商品排序，接收排序条件和商品数据列表作为参数
const sortBy = (criteria) => {
    let products = [...artPieces.value]; // 确保 products 是一个有效的数组

    if (criteria === 'all') {
        products.sort((a, b) => a.aid - b.aid); // 按 ID 升序排序
    } else if (criteria === 'price') {
        products.sort((a, b) => a.art_price - b.art_price); // 按价格升序排序
    } else if (criteria === 'new') {
        products.sort((a, b) => b.aid - a.aid); // 按 ID 降序排序
    }

    selectedSort.value = criteria; // 更新选中的排序条件
    artPieces.value = products; // 更新 artPieces 为排序后的数组
    currentPage.value = 1; // 重置当前页为第一页
    updatePaginatedProducts(); // 更新当前页显示的商品
};

const prevPage = () => {
    if (currentPage.value > 1) {
        currentPage.value--;
        updatePaginatedProducts(); // 使用当前的艺术品数据
    }
};

const nextPage = () => {
    if (currentPage.value < totalPages.value) {
        currentPage.value++;
        updatePaginatedProducts(); // 使用当前的艺术品数据
    }
};

const goToPage = (page) => {
    currentPage.value = page;
    updatePaginatedProducts(); // 使用当前的艺术品数据
};

// 图片加载事件处理
const onProductImageLoad = (productId) => {
    console.log(`商品 ${productId} 图片加载完成`);
};

const onProductImageError = (productId) => {
    console.error(`商品 ${productId} 图片加载失败`);
};

// 组件挂载时，调用获取数据的函数来初始化商品数据等相关操作
onMounted(() => {
    fetchArtPieces();
});
</script>

<style scoped>
@import url(../css/home.css);
</style>