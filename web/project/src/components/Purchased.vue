<template>
    <div class="published-room">
        <div v-for="order in orders" :key="order.oid" class="transaction-card">
            <div class="pub-avatar"></div>
            <div class="pub-user-name">{{ order.seller_username }}</div>
            <div class="transaction-status">购买成功</div>
            <div class="pub-info">
                <img :src="require(`../upload-img/${order.art_image}`)" alt="Product Image" />
                <div class="pub-name">{{ order.art_title }}</div>
            </div>
            <div class="price">¥{{ order.price.toFixed(2) }}</div>
            <div class="more-link">交易哈希：{{ getShortenedContractHash(order.transaction_hash) }}</div>
            <div class="transaction-date">{{ formatDateWithDayjs(order.create_time) }}</div>
        </div>
    </div>
</template>
  
<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import dayjs from 'dayjs';

const orders = ref([]);

// 定义一个方法来获取缩略的合约标识
const getShortenedContractHash = (hash) => {
    if (hash) {
        const start = hash.slice(0, 6);
        const end = hash.slice(-4);
        return start + '...' + end;
    }
    return hash;
};

// 在模板中使用 dayjs 格式化日期
const formatDateWithDayjs = (dateStr) => {
    return dayjs(dateStr).format('YYYY-MM-DD HH:mm:ss');
};

// 获取购买订单
const fetchPurchasedOrders = async () => {
    try {
        const response = await axios.get('http://localhost:8001/purchased-orders', {
            headers: {
                'Authorization': localStorage.getItem('token')
            }
        });
        orders.value = response.data;
    } catch (error) {
        console.error('获取购买订单失败:', error);
    }
};

onMounted(() => {
    fetchPurchasedOrders();
});
</script>
  
<style scoped>
.published-room {
    display: flex;
    flex-wrap: wrap;
    gap: 25px;
}

.transaction-card {
    display: block;
    background-color: rgba(0, 0, 0, 0.219);
    border-radius: 20px;
    padding: 20px 30px;
    margin-bottom: 0px;
    width: 43.5%;
    height: auto;
    transition: all 0.2s ease-in-out;
}

.transaction-card:hover {
    background-color: #131819;
}

.pub-avatar {
    width: 45px;
    height: 45px;
    /* margin-top: 10px; */
    border: 1px solid rgba(165, 217, 221, 0.745);
    border-radius: 50%;
    background-image: url('../img/back-bottom.jpg');
    background-size: cover;
    background-position: center;
    opacity: 0.8;
    float: left;
}

.pub-user-name {
    font-weight: bold;
    color: rgba(224, 96, 158, 0.829);
    font-size: 17px;
    margin-top: 10px;
    margin-left: 60px;
}

.transaction-status {
    margin-top: -22px;
    color: #bdf5f89c;
    font-size: 16px;
    float: right;
}

.pub-info {
    display: flex;
    align-items: center;
    clear: both;
    margin-top: 50px;
}

.pub-info img {
    width: 100px;
    float: left;
    margin-right: 10px;
    border: none;
    opacity: 0.8;
    border-radius: 15px;
}

.pub-name {
    color: rgb(181, 206, 207);
    font-size: 18px;
    margin-left: 10px;
    font-weight: bold;
}

.price {
    color: rgba(240, 92, 164, 0.829);
    font-weight: bold;
    font-size: 20px;
    margin-top: -64px;
    float: right;
}

.more-link {
    color: #dffdff83;
    float: left;
    font-size: 15px;
    margin: 22px 0 10px;
}

.transaction-date {
    float: right;
    color: #dffdff83;
    font-size: 15px;
    margin: 20px 0 10px;
}
</style>