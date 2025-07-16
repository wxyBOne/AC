<!--
 * @Author: wxyBone 13638456960@163.com
 * @Date: 2024-12-04 11:12:53
 * @LastEditors: wxyBone 13638456960@163.com
 * @LastEditTime: 2024-12-26 17:14:16
 * @FilePath: \project\src\components\Product.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div class="product-container">
        <div class="product-room-1">
            <div class="product-image-out">
                <img :src="proImg" alt="Product Image" class="product-image-section">
            </div>
            <div class="product-info-section">
                <p class="product-title-section">{{ product.title }}</p>
                <div class="product-style">{{ product.tag }}</div>
                <div class="product-row">
                    <div class="product-price-section">
                        <span>¥{{ proPrice }}</span>
                    </div>
                </div>
                <button class="buy-button" @click="addToCart">加入购物车</button>

                <!-- 小提示：圆圈复选框 -->
                <div class="agreement-row">
                    <input type="checkbox" id="auto-agreement" v-model="autoAgreement" class="a-circle-checkbox" />
                    <label for="auto-agreement" class="agreement">我已阅读并同意《购买须知》。此商品是虚拟商品，购买后不支持退换。</label>
                </div>

                <div class="certificate-info">
                    <p><strong>存证信息</strong></p>
                    <div class="certificate-row">
                        <p>藏品名称:</p>
                        <p id="cer-info">{{ product.title }}</p>
                    </div>
                    <div class="certificate-row">
                        <p>类型:</p>
                        <p id="cer-info">数字版权</p>
                    </div>
                    <div class="certificate-row">
                        <p>创作方:</p>
                        <p id="cer-info">{{ product.creator_name }}</p>
                    </div>
                    <div class="certificate-row">
                        <p>授权方:</p>
                        <p id="cer-info">{{ product.creator_name }}</p>
                    </div>
                    <div class="certificate-row">
                        <p>发行方:</p>
                        <p id="cer-info">{{ product.creator_name }}</p>
                    </div>
                    <div class="certificate-row">
                        <p>认证标识:</p>
                        <div id="cer-info" style="display:flex;align-items: center;">
                            <p>{{ getShortenedContractHash(product.contract_hash) }}</p>
                            <el-icon class="copy-icon" @click="copyContractHash" title="复制完整哈希值">
                                <CopyDocument />
                            </el-icon>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="product-room-2">
            <h2>购买须知</h2>
            <ol>
                <li>
                    <strong>合法合规交易</strong><br />
                    依据我国法律要求，数字版权交易必须合法合规进行。坚决反对任何形式的非法、侵权行为，包括但不限于版权炒作、场外交易、欺诈等。请用户仔细甄别网络欺诈行为及相关风险。
                </li>
                <li>
                    <strong>用户身份限制</strong><br />
                    数字版权交易需购买前完成实名认证，未满 18 周岁人群禁止购买。
                </li>
                <li>
                    <strong>退换货政策</strong><br />
                    数字版权商品一经出售，不支持退货。
                </li>
                <li>
                    <strong>交易限制</strong><br />
                    本平台所售数字版权商品一经交易，不允许二次交易。
                </li>
                <li>
                    <strong>信息授权</strong><br />
                    如您为参与任何活动或获得任何赋能而对数字版权进行转让（出售、赠与等）、受让、合成、兑换、持有、使用、处置、销毁等操作，您授权同意本平台向活动举办方、赋能提供方、数字版权发售方等提供您的联系方式、版权信息、相关数量、收货地址（如需）等必要信息。
                </li>
            </ol>
        </div>
        <div class="light3"></div>
        <div class="light4"></div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { CopyDocument } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import router from '@/router';
import axios from 'axios';

const autoAgreement = ref(true);
const route = useRoute();
const product = ref({}); // 用于存储商品数据
let proImg;
let proPrice;

const fetchProduct = async () => {
    const productId = route.params.id; // 获取传递的商品 ID
    try {
        const response = await axios.get(`http://localhost:8001/art-pieces/${productId}`); // 调用后端接口获取商品数据
        product.value = response.data; // 存储商品数据
        proImg = require(`../upload-img/${product.value.art_image}`);
        proPrice = product.value.art_price.toFixed(2);
    } catch (error) {
        console.error('获取商品失败:', error); // 处理错误
    }
};

const copyContractHash = () => {
    console.log('完整认证哈希:', product.value.contract_hash)  // 修改这里

    try {
        navigator.clipboard.writeText(product.value.contract_hash)  // 修改这里
        console.log('复制成功')
        ElMessage({
            message: '认证哈希已复制',
            type: 'success'
        })
    } catch (err) {
        console.log('复制失败:', err)
        ElMessage.error('复制失败')
    }
}

const getShortenedContractHash = (hash) => {
    if (hash) {
        const start = hash.slice(0, 6);
        const end = hash.slice(-4);
        return start + '...' + end;
    }
    return hash;
};

const addToCart = async () => {
    if (!autoAgreement.value) { // 检查复选框是否选中
        window.alert('请先阅读并同意《购买须知》。'); // 弹窗提示
        return;
    }
    const token = localStorage.getItem('token'); // 从 Local Storage 获取 token
    try {
        const response = await axios.post('http://localhost:8001/add-to-cart', {
            art_id: product.value.aid // 传递艺术品 ID
        }, {
            headers: {
                Authorization: token,
                'Content-Type': 'multipart/form-data' // 设置请求头
            }
        });
        window.alert(response.data.message); // 弹窗提示
    } catch (error) {
        if (error.response.status === 401) {
            window.alert('未登录，请重新登录。'); // 提示用户未登录
            router.push('/Login'); // 重定向到登录页面
        } else if (error.response.status === 403) {
            window.alert('您不能购买自己发布的艺术商品'); // 弹窗提示不能购买自己发布的
        } else {
            window.alert('该商品已在购物车中，请勿重复添加。');
        }
    }
};

onMounted(() => {
    fetchProduct(); // 组件挂载时获取商品数据
});
</script>

<style scoped>
@import url(../css/product.css);
</style>