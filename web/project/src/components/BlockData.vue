<template>
    <div class="panel-b assets-panel-b">
        <div class="assets-table-b">
            <div class="table-header-b">
                <div class="th-b block-col-b">区块号<el-icon>
                        <ArrowDown />
                    </el-icon></div>
                <div class="th-b token-col-b">开采时间<el-icon>
                        <ArrowDown />
                    </el-icon></div>
                <div class="th-b token-col-b">Gas 消耗<el-icon>
                        <ArrowDown />
                    </el-icon></div>
                <div class="th-b token-col-b">交易数量<el-icon>
                        <ArrowDown />
                    </el-icon></div>
                <div class="th-b token-col-b">操作</div>
            </div>

            <div class="table-body-b">
                <div v-for="block in blockchainMetrics" :key="block.value" class="table-row-b"
                    @click="getBlockDetails(block.value)">
                    <div class="td-b token-col-b">
                        <el-icon class="Trophy">
                            <Menu />
                        </el-icon>
                        <span>{{ block.value }}</span>
                    </div>
                    <div class="td-b change-col-b">{{ block.timestamp }}</div>
                    <div class="td-b price-col-b">{{ block.gasUsed }}</div>
                    <div class="td-b cap-col-b">{{ block.transactions }}</div>
                    <div class="td-b trade-col-b">
                        <button class="trade-btn-b">查看</button>
                    </div>
                </div>
            </div>
        </div>

        <el-dialog v-model="blockchainData.selectedBlock" :show-close="false" :modal="true" class="block-dialog">
            <BlockDetails v-if="blockchainData.selectedBlock" :block="blockchainData.selectedBlock"
                @close="blockchainData.selectedBlock = null" />
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, onMounted, computed, onUnmounted } from 'vue';
import { ethers } from 'ethers';
import {
    Grid, Menu, View, ArrowDown, CaretTop, CaretBottom, Trophy
} from '@element-plus/icons-vue';
import BlockDetails from './BlockDetails.vue';

// 创建 provider
const provider = new ethers.providers.JsonRpcProvider('http://localhost:7545');
const blockchainData = ref({
    latestBlock: 0,
    blocks: [],
    selectedBlock: null
});

// 获取区块数据
const fetchBlockData = async () => {
    try {
        // 获取最新区块号
        const latestBlockNumber = await provider.getBlockNumber();
        blockchainData.value.latestBlock = latestBlockNumber;

        // 获取最近的几个区块信息
        const blocks = [];
        for (let i = 0; i < 5; i++) {
            const block = await provider.getBlock(latestBlockNumber - i);
            if (block) {
                blocks.push({
                    number: block.number,
                    timestamp: new Date(block.timestamp * 1000),
                    gasUsed: block.gasUsed.toString(),
                    transactions: block.transactions.length
                });
            }
        }
        blockchainData.value.blocks = blocks;
    } catch (error) {
        console.error('获取区块数据失败:', error);
    }
};

// 获取区块详情
const getBlockDetails = async (blockNumber) => {
    try {
        const block = await provider.getBlock(blockNumber, true);
        blockchainData.value.selectedBlock = block;
    } catch (error) {
        console.error('获取区块详情失败:', error);
    }
};

// 格式化时间
const formatDate = (date) => {
    return date.toLocaleString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit'
    });
};

// 计算区块指标
const blockchainMetrics = computed(() =>
    blockchainData.value.blocks.map(block => ({
        token: `Block ${block.number}`,
        change: 0,
        value: block.number,
        gasUsed: block.gasUsed,
        timestamp: formatDate(block.timestamp),
        transactions: block.transactions,
        unit: 'Block'
    }))
);

onMounted(() => {
    fetchBlockData();
    // 定期更新数据
    const interval = setInterval(fetchBlockData, 15000);
    onUnmounted(() => clearInterval(interval));
});
</script>

<style scoped>
@import url(../css/overview.css);
</style>