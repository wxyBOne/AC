<template>
    <div class="block-details">
        <!-- 头部 -->
        <button class="back-btn" @click="handleClose">返回</button>

        <!-- 基本信息 -->
        <div id="block-info">
            <div class="info-row">
                <h1 class="blocks-title">区块 {{ block.number }}</h1>
                <div class="info-group">
                    <div class="info-label">Gas 消耗</div>
                    <div class="info-value">{{ block.gasUsed?.toString() }}</div>
                </div>
                <div class="info-group">
                    <div class="info-label">Gas 限制</div>
                    <div class="info-value">{{ block.gasLimit?.toString() }}</div>
                </div>
            </div>

            <!-- 时间信息 -->
            <div class="info-section">
                <div class="info-label">开采时间</div>
                <div class="info-value">{{ formatDate(block.timestamp) }}</div>
            </div>

            <!-- 区块哈希 -->
            <div class="info-section">
                <div class="info-label">区块哈希</div>
                <div class="info-value hash">{{ block.hash }}</div>
            </div>

            <!-- 交易列表 -->
            <div class="transactions-section" v-if="block.transactions?.length">
                <div class="transaction-item" v-for="tx in block.transactions" :key="tx">
                    <div class="info-label">交易哈希</div>
                    <div class="info-value hash">{{ tx }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const props = defineProps({
    block: {
        type: Object,
        required: true
    }
});

const emit = defineEmits(['close']);

const handleClose = () => {
    emit('close');
};

// 格式化时间
const formatDate = (timestamp) => {
    if (!timestamp) return '';
    const date = new Date(timestamp * 1000);
    return date.toLocaleString('en-US', {
        month: '2-digit',
        day: '2-digit',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false
    });
};
</script>

<style scoped>
@import url(../css/BlockDetails.css);
</style>