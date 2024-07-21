<script>
import 'echarts'
import VChart from 'vue-echarts'
import gptIcon from "@/components/icon/gptIcon.vue";

export default {
  components: {
    VChart,
    gptIcon
  },
  props: {
    historyData: {
      type: Array,
      default: () => {
        return [
          {
            timestamp: '0',
            temperature: 0,
            humidity: 0,
            soilHumidity: 0,
            precipitation: 0
          },
        ]
      }
    },
    nowData: {
      type: Object,
      default: () => {
        return {
          temperature: 0,
          humidity: 0,
          soilHumidity: 0,
          precipitation: 0,
          gptContent: ''
        }
      }
    }
  },
  watch: {
    historyData: {
      handler: function (newVal, oldVal) {
        // 提取historyData中的时间，温度
        const time = [];
        const temperature = [];
        // 间隔 length/200 取一个值
        for (let i = 0; i < newVal.length; i += Math.ceil(newVal.length / 50)) {
          time.push(newVal[i].timestamp);
          temperature.push(newVal[i].temperature);
        }
        // 更新图表
        this.options.xAxis.data = time;
        this.options.series[0].data = temperature;
      },
      deep: true
    },
    nowData: {
      handler: function (newVal, oldVal) {
        //如果id模time.length/200不为0，则不添加数据
        if (this.historyData.length % Math.ceil(this.historyData.length / 50) !== 0) {
          return;
        }
        // 将当前的时间和温度添加到图表中
        this.options.xAxis.data.push(newVal.timestamp);
        this.options.series[0].data.push(newVal.temperature);
      },
      deep: true
    },
  },

  data() {
    return {
      gptContent: '',
      options: {
        tooltip: {
          trigger: 'axis',
          formatter: function(params) {
            // 获取当前的时间值
            const time = params[0].axisValue;
            return `
              <div>时间: ${time}</div>
              <div>
                <span style="display:inline-block; width:10px; height:10px; background-color:${params[0].color}; border-radius:50%;"></span>
                <span style="color: ${params[0].color}; margin-left:5px;">温度: ${params[0].value}℃</span>
              </div>
            `;
          },
          axisPointer: {
            type: 'none'
          }
        },
        xAxis: {
          type: 'category',
          data: [],
          axisTick: {
            show: false
          },
          axisLabel: {
            show: false
          }
        },
        yAxis: {
          type: 'value',
          splitNumber: 3
        },
        series: [
          {
            data: [],
            type: 'line',
            smooth: 0.5,
            symbol: 'circle',
            color: '#f48888',
            // 隐藏数据点
            showSymbol: false,
            connectNulls: true,
            lineStyle: {
              color: '#f48888',
              width: 2
            },
            emphasis: {
              focus: 'series', // 用于突出显示数据
              itemStyle: {
                color: '#ec1f1f'
              },
              label: {
                show: true, // 显示悬停时的数据标签
                formatter: function(params) {
                  return params.value; // 可自定义显示的内容
                }
              },
            },
          },
        ]
      }
    }
  }
}




</script>

<template>
  <div class="temperatureChartComponents">
    <div class="temperatureChartComponentsTitle">
      温度图
    </div>
    <div class="temperatureChartComponentsChart">
      <v-chart :option="options" class="temperatureChart"></v-chart>
    </div>
  </div>
</template>

<style scoped>
.temperatureChartComponentsTitle {
  /* 设置字体大小 */
  font-size: 20px;
  /* 设置字体加粗 */
  font-weight: bold;
}
.temperatureChartComponents {
  /* 设置大小 */
  height: 100%;
  width: 100%;
  /* 设置中心位置 */
  position: absolute;
  /* 不可遮挡到其他组件 */
  z-index: 0;
}
.temperatureChartComponentsChart {
  /* 设置大小 */
  height: 80%;
  width: 100%;
  /* 设置颜色为白色 */
  background-color: white;
  /* 设置圆角 */
  border-radius: 20px;
  /* 放到中间 */
  position: absolute;
  top: 60%;
  transform: translateY(-50%);
  /* 不可遮挡到其他组件 */
  z-index: 0;
}
.temperatureChart {
  height: 130%;
  width: 100%;
  /* 放到中间 */
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.canvas {
  z-index: 0;
}


</style>