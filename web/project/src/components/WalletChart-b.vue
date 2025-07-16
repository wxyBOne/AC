<template>
    <div class="wallet-chart-b">
        <div class="chart-scale-b">
            <div class="scale-item-b">50%</div>
            <div class="scale-item-b">10%</div>
        </div>
        <div class="chart-container-b" ref="chartContainer"></div>
        <div class="time-marker-b">
            Wed, 07:30 PM
        </div>
    </div>
</template>

<script setup>
import { onMounted, ref, onUnmounted } from 'vue';
import * as echarts from 'echarts';

const chartContainer = ref(null);
let chart = null;

onMounted(() => {
    chart = echarts.init(chartContainer.value);

    const option = {
        backgroundColor: 'transparent',
        grid: {
            left: '12%',
            right: '12%',
            top: '10%',
            bottom: '15%',
            containLabel: false // 重要：不包含坐标轴标签
        },
        xAxis: {
            type: 'category',
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri'],
            axisLine: {
                show: true,
                lineStyle: {
                    color: 'rgba(168, 206, 210, 0.05)'
                }
            },
            axisTick: {
                show: false
            },
            axisLabel: {
                color: 'rgba(168, 206, 210, 0.6)',
                fontSize: 11,
                margin: 12
            }
        },
        yAxis: {
            type: 'value',
            min: 0,
            max: 60,
            interval: 20,
            show: false, // 隐藏内置的y轴
            splitLine: {
                show: true,
                lineStyle: {
                    color: 'rgba(168, 206, 210, 0.05)',
                    width: 1
                }
            }
        },
        series: [{
            data: [30, 45, 35, 50, 40],
            type: 'line',
            smooth: true,
            symbol: 'none',
            lineStyle: {
                color: 'rgba(255, 67, 145, 0.3)',
                width: 3
            },
            areaStyle: {
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                    offset: 0,
                    color: 'rgba(255, 67, 145, 0.3)'
                }, {
                    offset: 1,
                    color: 'rgba(128, 220, 236, 0.1)'
                }])
            }
        }]
    };

    chart.setOption(option);
    window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
    if (chart) {
        chart.dispose();
        window.removeEventListener('resize', handleResize);
    }
});

function handleResize() {
    chart && chart.resize();
}
</script>

<style scoped>
.wallet-chart-b {
    position: relative;
    width: 100%;
    height: 150px;
    margin-top: 20px;
}

.chart-container-b {
    width: 100%;
    height: 100%;
}

.chart-scale-b {
    position: absolute;
    left: 0;
    /* 调整位置 */
    top: 0;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 15px 0;
    /* 调整上下内边距 */
    z-index: 1;
}

.scale-item-b {
    color: rgba(168, 206, 210, 0.6);
    font-size: 12px;
    padding-left: 20px;
    /* 添加左内边距 */
}

.time-marker-b {
    position: absolute;
    right: 20px;
    top: 50%;
    transform: translateY(-50%);
    background-color: rgba(168, 206, 210, 0.1);
    padding: 4px 8px;
    border-radius: 4px;
    color: rgba(168, 206, 210, 0.6);
    font-size: 12px;
    z-index: 1;
}
</style>