<!--
 * @Author: wxyBone 13638456960@163.com
 * @Date: 2024-12-17 10:48:02
 * @LastEditors: wxyBone 13638456960@163.com
 * @LastEditTime: 2024-12-27 09:59:47
 * @FilePath: \project\src\components\Overview.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div class="portfolio-container-b">
        <div class="main-content-b">
            <h1 class="title-b">AC链证交易监测系统</h1>

            <div class="nav-tabs-b">
                <button v-for="tab in tabs" :key="tab" :class="['tab-btn-b', { 'active-b': activeTab === tab }]"
                    @click="activeTab = tab">
                    <el-icon>
                        <Position />
                    </el-icon>
                    {{ tab }}
                </button>
            </div>

            <div class="content-grid-b">
                <div class="panel-b wallet-panel-b">
                    <div class="wallet-dropdown-b">
                        <button class="dropdown-btn-b">
                            更新<el-icon>
                                <ArrowDown />
                            </el-icon>
                        </button>
                    </div>
                    <div class="growth-info-b">
                        <h2>与上一区块相比，交易量变化率为 {{ blockchainData.volumeChange }}%</h2>
                        <button class="details-btn-b">
                            查看细节
                            <el-icon>
                                <ArrowRight />
                            </el-icon>
                        </button>
                    </div>
                </div>

                <div class="panel-b chart-panel-b">
                    <TradingViewChart />
                </div>

                <div class="panel-b exchange-panel-b">
                    <h3>区块信息</h3>
                    <div class="exchange-inputs-b">
                        <div class="currency-input-b">
                            <div class="currency-info-b">
                                <el-icon>
                                    <Menu />
                                </el-icon>
                                <span>市场交易量</span>
                            </div>
                            <input type="text" :value="totalTxCount" class="amount-input-b" readonly>
                            <span class="tether-b">Txs</span>
                        </div>
                        <div class="currency-input-b">
                            <div class="currency-info-b">
                                <el-icon>
                                    <Menu />
                                </el-icon>
                                <span>区块高度</span>
                            </div>
                            <input type="text" :value="blockchainData.latestBlock" class="amount-input-b" readonly>
                            <span class="tether-b">Block</span>
                        </div>
                        <div class="currency-input-b">
                            <div class="currency-info-b">
                                <el-icon>
                                    <Menu />
                                </el-icon>
                                <span>Gas 价格</span>
                            </div>
                            <input type="text" :value="blockchainData.gasPrice" class="amount-input-b" readonly>
                            <span class="tether-b">Gwei</span>
                        </div>
                    </div>
                </div>
            </div>

            <div class="bottom-grid-b">
                <div class="panel-b wallet-stats-b">
                    <div class="wallet-header-b">
                        <div class="header-left-b">
                            <div class="wallet-title-b">
                                <el-icon>
                                    <Menu />
                                </el-icon>
                                ETH
                            </div>
                        </div>
                        <span class="time-tag-b">当前区块编号 #{{ blockchainData.latestBlock }}</span>
                        <el-icon class="Refresh-icon-b">
                            <Refresh />
                        </el-icon>
                    </div>
                    <div class="wallet-balance-b">
                        <div class="balance-info-b">
                            <h2>{{ blockchainData.totalVolume }} <p>ETH</p>
                            </h2>
                            <span class="change-tag-b">{{ blockchainData.volumeChange > 0 ? '+' : '' }}{{
                                blockchainData.volumeChange }}%</span>
                        </div>
                        <p class="revenue-text-b">记录当前区块的交易花费</p>
                    </div>
                    <WalletChart style="margin:7px 0 0 -10px;" />
                </div>

                <BlockData></BlockData>

                <div class="panel-b favorites-panel-b">
                    <div class="favorites-header-b">
                        <h3>新品上市</h3>
                        <div class="sort-buttons-b">
                            <button class="sort-btn-b" @click="sortBy('high')">
                                最高 <el-icon>
                                    <ArrowUp />
                                </el-icon>
                            </button>
                            <button class="sort-btn-b" @click="sortBy('low')">
                                最低 <el-icon>
                                    <ArrowDown />
                                </el-icon>
                            </button>
                        </div>
                    </div>
                    <div class="favorites-list-b">
                        <div v-for="(piece, index) in displayedArtPieces" :key="index" class="favorite-item-b">
                            <el-icon class="Star-icon-b">
                                <Star />
                            </el-icon>
                            <div class="favorite-info-b">
                                <span class="token-name-b">{{ piece.title }}</span>
                                <span class="market-cap-b">¥{{ piece.art_price.toFixed(2) }}</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, computed, onUnmounted } from 'vue';
import { ethers } from 'ethers';
import {
    Grid, Wallet, Position, SwitchButton, ChatDotRound, Bell,
    ArrowDown, ArrowRight, ArrowUp, Lock, Menu, Operation,
    View, CaretTop, CaretBottom, Trophy, Star, Refresh
} from '@element-plus/icons-vue';
import BlockData from './BlockData.vue'
import TradingViewChart from './TradingViewChart-b.vue';
import WalletChart from './WalletChart-b.vue';
import MiniChart from './MiniChart-b.vue';
import axios from 'axios';

const totalTxCount = ref(0);

const fetchTransactionCount = async () => {
    try {
        const response = await axios.get('http://localhost:8001/admin/transaction/count', {
            headers: {
                'Authorization': localStorage.getItem('admin_token')
            }
        })
        totalTxCount.value = response.data.count;
    } catch (error) {
        console.error('获取交易总量失败:', error);
    }
};

onMounted(() => {
    fetchTransactionCount();
});

// 初始化 provider
const provider = new ethers.providers.JsonRpcProvider('http://localhost:7545');

// 区块链数据状态
const blockchainData = ref({
    latestBlock: 0,
    gasPrice: 0,
    txCount: 0,
    totalVolume: 0,
    volumeChange: 0,
    lastBlockVolume: 0,
});

// 更新获取区块链数据的函数
const fetchBlockchainData = async () => {
    try {
        // 获取最新区块号和Gas价格
        const latestBlock = await provider.getBlockNumber();
        const gasPrice = await provider.getGasPrice();

        // 获取当前区块和前一个区块
        const block = await provider.getBlock(latestBlock);
        const previousBlock = await provider.getBlock(latestBlock - 1);

        // 计算包含ETH转账的区块数量和交易量
        let ethTxBlockCount = 0;
        let totalVolume = 0;
        let lastBlockVolume = 0;

        // 获取当前区块的交易
        if (block && block.transactions.length > 0) {
            let hasEthTransfer = false;
            for (const txHash of block.transactions) {
                try {
                    const tx = await provider.getTransaction(txHash);
                    if (tx && tx.value && !tx.value.isZero()) {
                        totalVolume += Number(ethers.utils.formatEther(tx.value));
                        hasEthTransfer = true;
                    }
                } catch (e) {
                    console.error('Error processing transaction:', e);
                }
            }
            if (hasEthTransfer) ethTxBlockCount++;
        }

        // 获取前一个区块的交易
        if (previousBlock && previousBlock.transactions.length > 0) {
            for (const txHash of previousBlock.transactions) {
                try {
                    const tx = await provider.getTransaction(txHash);
                    if (tx && tx.value && !tx.value.isZero()) {
                        lastBlockVolume += Number(ethers.utils.formatEther(tx.value));
                    }
                } catch (e) {
                    console.error('Error processing previous block transaction:', e);
                }
            }
        }

        let volumeChange = lastBlockVolume === 0
            ? (totalVolume === 0 ? 0 : 100)
            : ((totalVolume - lastBlockVolume) / lastBlockVolume * 100).toFixed(2);

        blockchainData.value = {
            latestBlock,
            gasPrice: ethers.utils.formatUnits(gasPrice, 'gwei'),
            txCount: block.transactions.length,
            totalVolume: totalVolume.toFixed(2),
            volumeChange,
            lastBlockVolume: lastBlockVolume.toFixed(2),
        };

    } catch (error) {
        console.error('获取区块链数据失败:', error);
        console.error('错误详情:', error.message);
    }
};

// 保持原有的艺术品相关代码
const artPieces = ref([]);
const fetchArtPieces = async () => {
    try {
        const response = await axios.get('http://localhost:8001/art-pieces');
        artPieces.value = response.data;
        sortBy('all');
    } catch (error) {
        console.error('获取艺术品失败:', error);
    }
};

const sortBy = (criteria) => {
    let products = [...artPieces.value];
    if (criteria === 'all') {
        products.sort((a, b) => a.aid - b.aid);
    } else if (criteria === 'high') {
        products.sort((a, b) => b.art_price - a.art_price);
    } else if (criteria === 'low') {
        products.sort((a, b) => a.art_price - b.art_price);
    }
    artPieces.value = products;
};

const displayedArtPieces = computed(() => {
    return artPieces.value.slice(0, 5);
});

const tabs = ['Overview', 'Blocks', 'Transactions'];
const activeTab = ref('Overview');

// 定时更新
let updateInterval;
onMounted(() => {
    fetchBlockchainData();
    fetchArtPieces();
    updateInterval = setInterval(fetchBlockchainData, 5000);
});

onUnmounted(() => {
    if (updateInterval) {
        clearInterval(updateInterval);
    }
});
</script>

<style scoped>
.portfolio-container-b {
    display: flex;
    min-height: calc(100vh - 70px);
    background-color: #0f1314;
}

/* 主内容区样式 */
.main-content-b {
    padding: 30px;
}

.title-b {
    font-size: 30px;
    font-weight: 600;
    color: rgba(213, 252, 255, 0.75);
    margin-bottom: 15px;
}

/* 导航标签样式 */
.nav-tabs-b {
    display: flex;
    gap: 10px;
    margin-bottom: 30px;
}

.tab-btn-b {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px 16px;
    border: 1px solid rgba(168, 206, 210, 0.325);
    background-color: rgba(168, 206, 210, 0);
    border-radius: 20px;
    color: rgba(168, 206, 210, 0.7);
    cursor: pointer;
    transition: all 0.3s;
    font-size: 14px;
}

.tab-btn-b .el-icon {
    font-size: 16px;
}

.tab-btn-b.active-b {
    background-color: #2a2f3f7b;
    color: rgb(189, 208, 211);
}

/* 内容网格样式 */
.content-grid-b {
    display: grid;
    grid-template-columns: 1fr 2fr 1fr;
    gap: 24px;
    margin-bottom: 24px;
}

.panel-b {
    background-color: rgba(40, 48, 49, 0.223);
    border-radius: 15px;
    padding: 20px;
}

/* 钱包面板样式 */
.wallet-panel-b {
    background: linear-gradient(45deg, rgba(215, 216, 141, 0.1), rgba(255, 0, 149, 0.15), rgba(11, 72, 255, 0.1), rgba(49, 224, 255, 0.15));
    position: relative;
    overflow: hidden;
}

.wallet-panel-b::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: linear-gradient(45deg, rgba(255, 105, 180, 0.1), transparent);
    pointer-events: none;
}

.wallet-dropdown-b {
    margin-bottom: 20px;
}

.dropdown-btn-b {
    display: flex;
    align-items: center;
    gap: 7px;
    padding: 8px 16px;
    background-color: rgba(168, 206, 210, 0.1);
    border: none;
    font-size: 13px;
    border-radius: 8px;
    color: rgb(168, 206, 209);
    cursor: pointer;
}

.growth-info-b {
    color: rgba(189, 208, 211, 0.9);
    position: relative;
    z-index: 1;
}

.growth-info-b h2 {
    font-size: 25px;
    margin-bottom: 15px;
    line-height: 1.4;
}

.details-btn-b {
    display: flex;
    align-items: center;
    gap: 5px;
    padding: 8px 16px;
    background-color: rgba(168, 206, 210, 0.1);
    border: none;
    font-size: 13px;
    border-radius: 8px;
    color: rgba(168, 206, 210, 0.7);
    cursor: pointer;
}

/* 图表面板样式 */
.chart-panel-b {
    background-color: rgba(40, 48, 49, 0.223);
}

.chart-header-b {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.pair-selector-b {
    display: flex;
    align-items: center;
    gap: 8px;
    color: rgb(189, 208, 211);
    font-size: 14px;
}

.price-info-b {
    text-align: right;
}

.price-b {
    font-size: 24px;
    color: rgb(189, 208, 211);
}

.currency-b {
    color: rgba(168, 206, 210, 0.6);
    font-size: 16px;
}

.change-badge-b {
    display: inline-block;
    padding: 4px 8px;
    background-color: rgba(76, 175, 80, 0.1);
    color: rgba(67, 242, 255, 0.6);
    border-radius: 6px;
    margin-top: 5px;
}

.time-filters-b {
    display: flex;
    justify-content: center;
    gap: 10px;
    margin-top: 20px;
}

.filter-btn-b {
    padding: 6px 12px;
    background: none;
    border: none;
    color: rgba(168, 206, 210, 0.6);
    cursor: pointer;
    border-radius: 6px;
    transition: all 0.3s;
    font-size: 13px;
}

.filter-btn-b.active-b {
    background-color: rgba(168, 206, 210, 0.1);
    color: rgb(189, 208, 211);
}

/* 兑换面板样式 */
.exchange-panel-b {
    background-color: rgba(40, 48, 49, 0.223);
}

.exchange-panel-b h3 {
    color: rgb(189, 208, 211);
    margin: 5px 0 20px;
    font-size: 17px;
}

.exchange-inputs-b {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.currency-input-b {
    background-color: rgba(168, 206, 210, 0.1);
    padding: 15px 15px;
    border-radius: 12px;
}

.currency-info-b {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 10px;
}

.currency-info-b .el-icon {
    font-size: 16px;
    color: rgba(196, 61, 149, 0.818);
}

.currency-info-b span {
    color: rgb(160, 194, 198);
    font-size: 15px;
}

.amount-input-b {
    width: 100%;
    background: none;
    border: none;
    color: rgb(189, 208, 211);
    font-size: 20px;
    font-weight: 600;
    outline: none;
}

.tether-b {
    color: rgba(168, 206, 210, 0.6);
    font-size: 12px;
    margin-top: 4px;
    display: block;
}

.swap-btn-b {
    width: 100%;
    padding: 12px;
    background-color: rgba(168, 206, 210, 0.1);
    border: none;
    border-radius: 12px;
    color: rgb(189, 208, 211);
    font-size: 16px;
    cursor: pointer;
    margin-top: 20px;
    transition: all 0.3s;
}

.swap-btn-b:hover {
    background-color: rgba(168, 206, 210, 0.2);
}

/* 底部网格样式 */
.bottom-grid-b {
    display: grid;
    grid-template-columns: 1fr 2fr 1fr;
    gap: 24px;
}

/* 钱包统计样式 */
.wallet-stats-b {
    background-color: rgba(40, 48, 49, 0.223);
    padding: 2px 0 0 18px;
}

.wallet-header-b {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px;
    border-bottom: 1px solid rgba(168, 206, 210, 0.05);
    margin-bottom: 10px;
}

.header-left-b {
    display: flex;
    align-items: center;
    gap: 15px;
}

.wallet-title-b {
    margin-left: -15px;
    display: flex;
    align-items: center;
    gap: 8px;
    color: rgb(189, 208, 211);
    font-weight: 500;
}

.time-tag-b {
    color: rgba(168, 206, 210, 0.6);
    font-size: 13px;
    margin-left: 30px;
    padding: 5px 13px;
    background-color: rgba(168, 206, 210, 0.1);
    border-radius: 20px;
}

.Refresh-icon-b {
    color: rgba(168, 206, 210, 0.6);
    font-size: 20px;
}

.wallet-balance-b {
    margin-bottom: 0px;
}

.balance-info-b {
    display: flex;
    align-items: center;
    gap: 10px;
}

.balance-info-b h2 {
    display: flex;
    align-items: center;
    color: rgba(213, 252, 255, 0.7);
    font-size: 28px;
}

.balance-info-b h2 p {
    margin-left: 5px;
    color: rgba(213, 252, 255, 0.65);
    font-size: 23px;
}

.change-tag-b {
    padding: 4px 8px;
    background-color: rgba(213, 252, 255, 0.1);
    color: rgba(213, 252, 255, 0.6);
    border-radius: 6px;
    font-size: 13px;
}

.revenue-text-b {
    color: rgba(168, 206, 210, 0.7);
    margin-top: 5px;
    font-size: 14px;
}

/* 资产表格样式 */
.assets-panel-b {
    width: 105%;
    background-color: rgba(40, 48, 49, 0.223);
    padding: 0;
    overflow: hidden;
}

.assets-header-b {
    padding: 20px;
    border-bottom: 1px solid rgba(168, 206, 210, 0.05);
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.view-toggles-b {
    display: flex;
    gap: 10px;
}

.toggle-btn-b {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px 16px;
    background: none;
    border: none;
    color: rgba(168, 206, 210, 0.7);
    cursor: pointer;
    border-radius: 8px;
    font-size: 14px;
}

.toggle-btn-b.active-b {
    background-color: rgba(168, 206, 210, 0.1);
    color: rgb(189, 208, 211);
}

.view-modes-b {
    display: flex;
    gap: 8px;
}

.mode-btn-b {
    width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    color: rgba(168, 206, 210, 0.6);
    cursor: pointer;
    border-radius: 6px;
    transition: all 0.3s;
}

.mode-btn-b:hover,
.mode-btn-b.active-b {
    background-color: rgba(168, 206, 210, 0.1);
    color: rgb(189, 208, 211);
}

.assets-table-b {
    width: 100%;
}

.table-header-b {
    display: grid;
    grid-template-columns: 1.1fr 1fr 0.8fr 0.9fr 1.1fr 0.6fr;
    padding: 15px 20px;
    color: rgba(168, 206, 210, 0.6);
    font-size: 13px;
}

.th-b {
    display: flex;
    align-items: center;
    gap: 5px;
    cursor: pointer;
}

.th-b .el-icon {
    font-size: 12px;
    opacity: 0.6;
}

.table-body-b {
    font-size: 15px;
    padding: 0 20px;
}

.table-row-b {
    display: grid;
    grid-template-columns: 1fr 1fr 0.8fr 0.8fr 1.1fr 0.6fr;
    padding: 10px 0;
    border-top: 1px solid rgba(168, 206, 210, 0.05);
    align-items: center;
}

.token-col-b {
    display: flex;
    align-items: center;
    gap: 12px;
}

.token-col-b .Trophy {
    width: 24px;
    height: 24px;
    color: rgba(196, 61, 149, 0.818);
    border-radius: 50%;
}

.token-col-b span {
    font-size: 15px;
    margin-left: -10px;
    color: rgb(189, 208, 211);
    font-weight: 500;
}

.change-col-b {
    display: flex;
    align-items: center;
    gap: 4px;
    font-weight: 500;
}

.change-col-b.positive-b {
    color: rgba(168, 206, 210, 0.9);

}

.change-col-b.negative-b {
    color: rgba(168, 206, 210, 0.9);

}

.price-col-b,
.cap-col-b,
.change-col-b {
    color: rgba(168, 206, 210, 0.9);
}

.trade-btn-b {
    padding: 6px 16px;
    background-color: rgba(168, 206, 210, 0.1);
    border: none;
    border-radius: 8px;
    color: rgb(189, 208, 211);
    cursor: pointer;
    transition: all 0.3s;
    font-size: 13px;
}

.trade-btn-b:hover {
    background-color: rgba(168, 206, 210, 0.15);
}

/* 收藏夹样式 */
.favorites-panel-b {
    padding: 2px 20px 5px;
    margin-left: 9%;
    width: 92%;
    background-color: rgba(40, 48, 49, 0.223);
}

.favorites-header-b {
    display: flex;
    flex-direction: column;
    padding: 20px 0 10px;
    border-bottom: 1px solid rgba(168, 206, 210, 0.05);
}

.favorites-panel-b h3 {
    display: block;
    color: rgb(189, 208, 211);
    font-size: 17px;
    font-weight: 600;
    margin-bottom: 15px;
}

.sort-buttons-b {
    display: flex;
    gap: 15px;
}

.sort-btn-b {
    display: flex;
    align-items: center;
    gap: 3px;
    padding: 5px 10px;
    background-color: rgba(168, 206, 210, 0.1);
    border: none;
    border-radius: 6px;
    color: rgba(168, 206, 210, 0.7);
    cursor: pointer;
    font-size: 13px;
}

.sort-btn-b:hover {
    background-color: rgba(168, 206, 210, 0.15);
}

.favorites-list-b {
    margin-top: 2px;
}

.favorite-item-b {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 8px 0;
    border-bottom: 1px solid rgba(168, 206, 210, 0.05);
}

.Star-icon-b {
    width: 28px;
    height: 28px;
    color: rgba(196, 61, 149, 0.818);
    background-color: rgba(168, 206, 210, 0.1);
    border-radius: 50%;
}

.favorite-info-b {
    flex: 1;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.token-name-b {
    font-size: 14px;
    color: rgb(189, 208, 211);
    font-weight: 500;
}

.market-cap-b {
    display: block;
    text-align: right;
    color: rgba(168, 206, 210, 0.6);
    font-size: 14px;
}

.change-b {
    font-size: 13px;
}

.change-b.positive-b {
    color: rgba(213, 252, 255, 0.6);
}

.change-b.negative-b {
    color: rgba(172, 35, 124);
}
</style>