<template>
    <div class="published-room">
        <div v-for="art in artPieces" :key="art.aid" class="transaction-card">
            <div class="pub-info">
                <img :src="require(`../upload-img/${art.art_image}`)" alt="Product Image" />
                <div class="pub-name">{{ art.title }}</div>
            </div>
            <div class="price">¥{{ art.art_price.toFixed(2) }}</div>
            <div class="more-link">认证标识：{{ getShortenedContractHash(art.contract_hash) }}</div>
            <div class="publish-status">发布成功</div>
        </div>
    </div>
</template>
  
<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

// 定义一个方法来获取缩略的合约标识
const getShortenedContractHash = (hash) => {
    if (hash) {
        const start = hash.slice(0, 6);
        const end = hash.slice(-4);
        return start + '...' + end;
    }
    return hash;
};

const artPieces = ref([]);

// 获取已发布艺术品
const fetchPublishedArt = async () => {
    try {
        const response = await axios.get('http://localhost:8001/published-art', {
            headers: {
                'Authorization': localStorage.getItem('token')
            }
        });
        artPieces.value = response.data;
    } catch (error) {
        console.error('获取已发布艺术品失败:', error);
    }
};

onMounted(() => {
    fetchPublishedArt();
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

.pub-info {
    display: flex;
    align-items: center;
    clear: both;
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
    font-size: 15px;
    float: left;
    margin: 22px 0 0;
}

.transaction-date {
    float: right;
    color: #dffdff83;
    margin: 20px 0 0;
}

.publish-status {
    float: right;
    color: #bdf5f89c;
    font-size: 16px;
    margin: 20px 0 10px;
}
</style>