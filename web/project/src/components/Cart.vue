<template>
    <div class="cart-container">
        <p class="cart-title">购物车</p>
        <div class="cart-list">
            <div class="cart-item" v-for="(cart, index) in carts" :key="index" @click="toggleCartSelected(cart)">
                <input type="checkbox" :id="'cart-checkbox-' + index" v-model="cart.selected" @click.stop>
                <img :src="require(`../upload-img/${cart.art_image}`)" alt="cart" class="cart-image">
                <div class="cart-info">
                    <p class="cart-name">{{ cart.title }}</p>
                    <p class="cart-tag">{{ cart.tag }}</p>
                    <p class="cart-price">¥{{ cart.art_price.toFixed(2) }}</p>
                </div>
                <img src="../img/trashbin.png" class="cart-delete" @click="removeFromCart(cart.aid)">
            </div>
            <div class="cart-empty" v-if="carts.length === 0">
                <img src="../img/empty-cart.png">
                <p>您的购物车暂无商品</p>
            </div>
        </div>
        <div class="cart-summary">
            <div class="select-all">
                <input type="checkbox" v-model="isSelectAll" @change="handleSelectAll">
                <label for="select-all">全选</label>
            </div>
            <div class="total-price">
                总价：<span>¥{{ totalPrice }}</span>
            </div>
            <button class="checkout-button" @click="checkout">结算 ({{ selectedItemsCount }})</button>
        </div>
    </div>
</template>

<script setup>
import { ref, watch, watchEffect, onMounted } from 'vue';
import axios from 'axios';
import { ethers } from 'ethers';
import ArtCopyright from '@/contracts/ArtCopyright.json'

const carts = ref([]); // 用于存储购物车商品数据
const isSelectAll = ref(false); // 全选状态
const totalPrice = ref(0); // 计算总价格
const selectedItemsCount = ref(0); // 统计选中商品的数量

// 修改区块链相关的状态管理
const contract = ref(null);
const provider = ref(null);
const signer = ref(null);

// 修改初始化函数
const initBlockchain = async () => {
    try {
        if (typeof window.ethereum === 'undefined') {
            throw new Error('请安装 MetaMask 钱包');
        }

        // 请求账户访问
        const accounts = await window.ethereum.request({
            method: 'eth_requestAccounts'
        });
        console.log('当前账户:', accounts[0]);

        // 创建 provider
        const web3Provider = new ethers.providers.Web3Provider(window.ethereum);
        const signer = web3Provider.getSigner();

        // 创建合约实例
        contract.value = new ethers.Contract(
            process.env.VUE_APP_CONTRACT_ADDRESS,
            ArtCopyright.abi,
            signer
        );

        return { provider: web3Provider, signer };
    } catch (error) {
        console.error('区块链初始化失败:', error);
        throw error;
    }
};

// 原有的获取购物车商品函数
const fetchCartItems = async () => {
    try {
        const response = await axios.get('http://localhost:8001/cart-items', {
            headers: {
                'Authorization': localStorage.getItem('token')
            }
        });
        carts.value = response.data;
    } catch (error) {
        console.error('获取购物车失败:', error);
    }
};

// 计算总价格、全选操作、监听函数
watchEffect(() => {
    let sum = 0;
    carts.value.forEach(cart => {
        if (cart.selected) {
            sum += cart.art_price;
        }
    });
    totalPrice.value = sum.toFixed(2);
});

const toggleCartSelected = (cart) => {
  cart.selected = !cart.selected;
};

const handleSelectAll = () => {
    carts.value.forEach(cart => {
        cart.selected = isSelectAll.value;
    });
};

watch(() => carts.value.map(cart => cart.selected), (newSelectedStates) => {
    if (carts.value.length === 0) {
        isSelectAll.value = false;
    } else {
        isSelectAll.value = newSelectedStates.every(state => state);
    }
});

watch(() => carts.value.map(cart => cart.selected), (newSelectedStates) => {
    selectedItemsCount.value = newSelectedStates.filter(state => state).length;
});

// 原有的删除购物车商品函数
const removeFromCart = async (aid) => {
    try {
        await axios.delete(`http://localhost:8001/cart-items/${aid}`, {
            headers: {
                'Authorization': localStorage.getItem('token')
            }
        });
        window.alert('商品已从购物车中删除。');
        fetchCartItems();
    } catch (error) {
        console.error('删除购物车商品失败:', error);
    }
};


// 添加检查函数
const checkArtworkOnChain = async (artId) => {
    try {
        const provider = new ethers.providers.Web3Provider(window.ethereum);
        const signer = provider.getSigner();
        const currentAccount = await signer.getAddress();

        console.log('当前账户:', currentAccount);

        // 创建合约实例
        contract.value = new ethers.Contract(
            process.env.VUE_APP_CONTRACT_ADDRESS,
            ArtCopyright.abi,
            signer
        );

        // 调用合约的 ownerOf 方法
        const owner = await contract.value.ownerOf(artId);
        console.log('艺术品所有者:', owner);
        return true;
    } catch (error) {
        console.error('检查艺术品状态失败:', error);
        return false;
    }
};

// 修改结算函数
const checkout = async () => {
    const selectedItems = carts.value.filter(cart => cart.selected);
    if (selectedItems.length === 0) {
        alert('请至少选择一个商品进行结算。');
        return;
    }

    try {
        const provider = new ethers.providers.Web3Provider(window.ethereum);
        const signer = provider.getSigner();
        const currentAccount = await signer.getAddress();

        console.log('当前账户:', currentAccount);

        contract.value = new ethers.Contract(
            process.env.VUE_APP_CONTRACT_ADDRESS,
            ArtCopyright.abi,
            signer
        );

        for (const item of selectedItems) {
            try {
                console.log('正在处理商品:', item);
                console.log('合约地址:', process.env.VUE_APP_CONTRACT_ADDRESS);
                console.log('商品ID:', item.aid);
                console.log('商品价格:', item.art_price);
                console.log('合约哈希:', item.contract_hash);

                // 检查艺术品状态
                const exists = await checkArtworkOnChain(item.aid);
                if (!exists) {
                    throw new Error('该艺术品在链上不存在，请联系卖家确认');
                }

                // 转换价格为 wei
                const priceInWei = ethers.utils.parseEther(item.art_price.toString());
                console.log('支付金额(wei):', priceInWei.toString());

                // 估算 gas
                try {
                    const gasEstimate = await contract.value.estimateGas.purchaseArtwork(
                        item.aid,
                        { value: priceInWei }
                    );
                    console.log('预估 gas:', gasEstimate.toString());
                } catch (error) {
                    console.error('Gas 估算失败:', error);
                }

                // 购买艺术品
                const tx = await contract.value.purchaseArtwork(
                    item.aid,
                    {
                        value: priceInWei,
                        gasLimit: 500000  // 增加 gas 限制
                    }
                );

                console.log('购买交易已发送:', tx.hash);
                const receipt = await tx.wait();
                console.log('购买完成:', receipt);

                // 更新后端
                await axios.post('http://localhost:8001/checkout',
                    {
                        artworkId: item.aid,
                        transactionHash: receipt.transactionHash
                    },
                    {
                        headers: {
                            'Authorization': localStorage.getItem('token')
                        }
                    }
                );

                window.alert(`购买成功！交易哈希: ${receipt.transactionHash}`);
            } catch (error) {
                console.error('处理失败，详细错误:', error);

                if (error.code === 4001) {
                    window.alert('交易已被取消');
                } else if (error.message.includes('execution reverted')) {
                    const reason = error.data?.message || error.message;
                    if (reason.includes('Artwork does not exist')) {
                        window.alert('该艺术品在链上不存在，请联系卖家确认');
                    } else if (reason.includes('Artwork is not for sale')) {
                        window.alert('该艺术品已不在售');
                    } else if (reason.includes('Insufficient payment')) {
                        window.alert('支付金额不足');
                    } else if (reason.includes('Cannot buy your own artwork')) {
                        window.alert('不能购买自己的艺术品');
                    } else {
                        window.alert(`操作失败: ${reason}`);
                    }
                } else {
                    window.alert(`操作失败: ${error.message}`);
                }
                return;
            }
        }

        await fetchCartItems();
    } catch (error) {
        console.error('结算过程发生错误:', error);
        window.alert(error.message);
    }
};

// 组件加载时初始化
onMounted(async () => {
    await initBlockchain();
    await fetchCartItems();
});
</script>
  
<style scoped>
@import url(../css/cart.css);
</style>