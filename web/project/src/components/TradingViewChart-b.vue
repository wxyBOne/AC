<template>
    <!-- 图表容器 -->
    <div class="chart-wrapper-b">
        <!-- 图表头部 -->
        <div class="chart-header-b">
            <!-- 左侧信息区 -->
            <div class="left-section">
                <!-- 交易对选择器 -->
                <div class="pair-selector">
                    今日流水 <el-icon>
                        <ArrowDown />
                    </el-icon>
                </div>
                <!-- 价格信息区 -->
                <div class="price-section">
                    <div class="price">¥{{ formatAmount(totalAmount) }}<span class="unit">RMB</span></div>
                    <div class="change-badge">{{ formatChange(priceChange) }}</div>
                </div>
            </div>
            <!-- 时间过滤器 -->
            <div class="time-filters">
                <span v-for="period in periods" :key="period" :class="['filter-item', { active: activePeriod === period }]"
                    @click="activePeriod = period">
                    {{ period }}
                </span>
            </div>
        </div>
        <!-- ECharts 图表容器 -->
        <div class="chart-container" ref="chartContainer"></div>
    </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import * as echarts from 'echarts';
import { ArrowDown } from '@element-plus/icons-vue';
import axios from 'axios';

// 图表容器引用
const chartContainer = ref(null);
let chart = null;

// 数据状态
const totalAmount = ref(0);
const priceChange = ref(0);

// 格式化金额
const formatAmount = (amount) => {
    return (amount || 0).toFixed(2);
};

// 格式化变化百分比
const formatChange = (change) => {
    const value = change || 0;
    return `${value >= 0 ? '+' : ''}${value}%`;
};

// 时间周期选项
const periods = ['1H', '4H', '1D', '1W', '1M', '6M'];
const activePeriod = ref('1W');

// 获取交易数据
const fetchTransactionData = async () => {
    try {
        const token = localStorage.getItem('adminToken');
        const response = await axios.get('http://localhost:8001/admin/transaction/summary', {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        // 更新总额，确保有值
        totalAmount.value = response.data.totalAmount || 0;

        // 更新图表数据
        const transactions = response.data.recentTransactions || [];
        if (transactions.length > 0) {
            updateChartWithTransactions(transactions);

            // 计算价格变化
            if (transactions.length >= 2) {
                const latestPrice = transactions[0].price || 0;
                const previousPrice = transactions[1].price || 0;
                if (previousPrice !== 0) {
                    priceChange.value = ((latestPrice - previousPrice) / previousPrice * 100).toFixed(2);
                } else {
                    priceChange.value = 0;
                }
            }
        }
    } catch (error) {
        console.error('获取交易数据失败:', error);
        totalAmount.value = 0;
        priceChange.value = 0;
    }
};

// 更新图表数据（最多10条数据）
const updateChartWithTransactions = (transactions) => {
    if (!chart || !transactions.length) return;

    const times = transactions.map(tx => {
        const date = new Date(tx.created_at);
        return date.toLocaleTimeString('zh-CN', {
            hour: '2-digit',
            minute: '2-digit',
            hour12: false
        });
    }).reverse();

    const prices = transactions.map(tx => tx.price || 0).reverse();
    const dates = transactions.map(tx => tx.created_at).reverse();

    const option = {
        backgroundColor: 'transparent',
        animation: false,
        grid: [{
            left: 46,
            right: 5,
            top: '2.5%',
            height: '72%'
        }, {
            left: 46,
            right: 5,
            bottom: '8%',
            height: '8%'
        }],
        xAxis: [{
            type: 'category',
            data: times,
            axisLine: { lineStyle: { color: 'rgba(168, 206, 210, 0.05)' } },
            axisTick: { show: false },
            axisLabel: {
                color: 'rgba(168, 206, 210, 0.5)',
                fontSize: 11,
                margin: 12,
                align: 'center',
                interval: 'auto'
            },
            splitLine: {
                show: true,
                lineStyle: { color: 'rgba(168, 206, 210, 0.05)' }
            }
        }, {
            type: 'category',
            gridIndex: 1,
            data: times,
            axisLine: { lineStyle: { color: 'rgba(168, 206, 210, 0.05)' } },
            axisTick: { show: false },
            axisLabel: {
                show: true,
                color: 'rgba(168, 206, 210, 0.5)',
                fontSize: 11,
                margin: 12,
                align: 'center',
                interval: 'auto'
            }
        }],
        yAxis: [{
            scale: true,
            position: 'left',
            splitLine: {
                show: true,
                lineStyle: { color: 'rgba(168, 206, 210, 0.05)' }
            },
            axisLine: {
                show: true,
                lineStyle: { color: 'rgba(168, 206, 210, 0.05)' }
            },
            axisTick: { show: false },
            axisLabel: {
                color: 'rgba(168, 206, 210, 0.5)',
                fontSize: 11,
                margin: 12,
                formatter: value => value.toFixed(2)
            },
            min: 0,  // 设置最小值
            max: 'dataMax',  // 根据数据自动设置最大值
            splitNumber: 8,  // 设置分割段数
            interval: 100    // 设置间隔
        }, {
            scale: true,
            gridIndex: 1,
            splitNumber: 2,
            show: false
        }],
        series: [{
            name: '价格',
            type: 'line',
            data: prices,
            areaStyle: {
                opacity: 0.2,
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                    offset: 0,
                    color: 'rgba(11, 72, 255)'  // 顶部颜色
                }, {
                    offset: 1,
                    color: 'rgba(196, 61, 149, 0.5)'  // 底部颜色
                }])
            },
            lineStyle: {
                width: 1,
                color: 'rgba(128, 220, 236, 0.7)'  // 线条颜色
            },
            symbol: 'circle',        // 数据点样式
            symbolSize: 6,          // 数据点大小
            itemStyle: {
                color: 'rgba(196, 61, 149, 0.818)',  // 数据点颜色
                borderColor: 'rgba(128, 220, 236, 0.7)', // 数据点边框颜色
                borderWidth: 3      // 数据点边框宽度
            },
            smooth: true            // 平滑曲线
        }, {
            name: '成交量',
            type: 'bar',
            xAxisIndex: 1,
            yAxisIndex: 1,
            data: prices,
            itemStyle: {
                color: function (params) {
                    return {
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                            offset: 0,
                            color: 'rgba(168, 206, 210, 0.2)'
                        }, {
                            offset: 1,
                            color: 'rgba(168, 206, 210, 0.1)'
                        }])
                    };
                }
            },
            barWidth: '80%'
        }],
        tooltip: {
            trigger: 'axis',
            axisPointer: {
                type: 'cross',
                crossStyle: { color: 'rgba(168, 206, 210, 0.2)' }
            },
            backgroundColor: '#1a1b23',
            borderColor: 'rgba(168, 206, 210, 0.1)',
            textStyle: { color: 'rgb(189, 208, 211)' },
            formatter: function (params) {
                if (!params || !params.length) return '';
                const price = params[0].data || 0;  // 直接使用 data
                const index = params[0].dataIndex;
                const date = new Date(dates[index]);

                const formattedTime = date.toLocaleTimeString('zh-CN', {
                    hour: '2-digit',
                    minute: '2-digit',
                    hour12: false
                });

                return `
            <div style="font-size: 12px; padding: 4px;">
                <div>Time: ${formattedTime}</div>
                <div>Price: ¥${price.toFixed(2)}</div>
            </div>
        `;
            }
        }
    };

    chart.setOption(option);
};

onMounted(() => {
    chart = echarts.init(chartContainer.value);
    fetchTransactionData();
    window.addEventListener('resize', handleResize);
});

function handleResize() {
    chart && chart.resize();
}

onUnmounted(() => {
    if (chart) {
        chart.dispose();
        window.removeEventListener('resize', handleResize);
    }
});
</script>

<style scoped>
/* 图表外层容器 */
.chart-wrapper-b {
    width: 100%;
    background-color: rgba(68, 68, 98, 0);
    border-radius: 12px;
    padding: 8px 0;
}

/* 图表头部样式 */
.chart-header-b {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    margin-bottom: 10px;
    padding: 0 8px;
}

/* 左侧信息区样式 */
.left-section {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

/* 交易对选择器样式 */
.pair-selector {
    display: flex;
    align-items: center;
    gap: 6px;
    color: rgb(160, 194, 198);
    font-size: 15px;
    cursor: pointer;
}

/* 价格信息区样式 */
.price-section {
    width: 60px;
    display: flex;
    flex-direction: column;
    gap: 4px;
}

/* 价格样式 */
.price {
    font-size: 24px;
    color: rgb(189, 208, 211);
    font-weight: 600;
}

/* 单位样式 */
.unit {
    margin-left: 4px;
    color: rgb(189, 208, 211, 0.9);
    font-size: 14px;
}

/* 涨跌幅标签样式 */
.change-badge {
    display: inline-block;
    width: 72px;
    text-align: center;
    margin: 2px 0;
    padding: 3px 0;
    background-color: rgba(213, 252, 255, 0.1);
    color: rgba(213, 252, 255, 0.6);
    border-radius: 4px;
    font-size: 12px;
}

/* 时间过滤器样式 */
.time-filters {
    display: flex;
    gap: 16px;
}

/* 过滤器选项样式 */
.filter-item {
    color: rgba(168, 206, 210, 0.5);
    font-size: 13px;
    cursor: pointer;
    transition: color 0.3s;
}

.filter-item:hover {
    color: rgb(189, 208, 211);
}

.filter-item.active {
    color: rgb(189, 208, 211);
}

/* ECharts 容器样式 */
.chart-container {
    margin-left: 10px;
    width: 100%;
    height: 260px;
}
</style>