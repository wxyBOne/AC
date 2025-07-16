<template>
    <div class="mini-chart-b" ref="chartContainer"></div>
</template>

<script setup>
import { onMounted, ref, onUnmounted } from 'vue';
import * as echarts from 'echarts';

const props = defineProps({
    color: {
        type: String,
        default: '#ff69b4'
    }
});

const chartContainer = ref(null);
let chart = null;

onMounted(() => {
    chart = echarts.init(chartContainer.value);

    const option = {
        backgroundColor: 'transparent',
        grid: {
            left: 0,
            right: 0,
            top: 0,
            bottom: 0
        },
        xAxis: {
            type: 'category',
            show: false,
            data: Array(20).fill('')
        },
        yAxis: {
            type: 'value',
            show: false
        },
        series: [{
            data: generateRandomData(),
            type: 'line',
            smooth: true,
            symbol: 'none',
            lineStyle: {
                color: 'rgba(49, 224, 255, 0.2)',
                width: 1
            },
            areaStyle: {
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                    offset: 0,
                    color: 'rgba(49, 224, 255, 0.2)'
                }, {
                    offset: 1,
                    color: 'rgba(128, 220, 236, 0.08)'
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

function generateRandomData() {
    return Array(20).fill(0).map(() => Math.random() * 100);
}
</script>

<style scoped>
.mini-chart-b {
    width: 100px;
    height: 40px;
}
</style>