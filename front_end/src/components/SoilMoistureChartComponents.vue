<script>
import 'echarts'
import VChart from 'vue-echarts'
import TimeIcon from "@/components/icon/TimeIcon.vue";
import WeatherIcon from "@/components/icon/WeatherIcon.vue";
import gptIcon from "@/components/icon/gptIcon.vue";
import axios from "axios";

export default {
  components: {
    gptIcon,
    VChart,
    TimeIcon,
    WeatherIcon
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
          precipitation: 0
        }
      }
    }
  },
  watch: {
    historyData: {
      handler: function (newVal, oldVal) {
        // 提取historyData中的时间，土壤湿度，降雨量，湿度
        const time = [];
        const soilHumidity = [];
        const precipitation = [];
        const humidity = [];
        // 间隔 length/200 取一个值
        for (let i = 0; i < newVal.length; i += Math.ceil(newVal.length / 50)) {
          time.push(newVal[i].timestamp);
          soilHumidity.push(newVal[i].soilHumidity);
          precipitation.push(newVal[i].precipitation);
          humidity.push(newVal[i].humidity);
        }
        // 更新图表
        this.options.xAxis.data = time;
        this.options.series[0].data = soilHumidity;
        this.options.series[1].data = humidity;
        this.options.series[2].data = precipitation;

      },
      deep: true
    },
    nowData: {
      handler: function (newVal, oldVal) {
        // 如果id模time.length/200不为0，则不添加数据
        if (this.historyData.length % Math.ceil(this.historyData.length / 50) !== 0) {
          return;
        }
        // 提取nowData中的时间，土壤湿度，降雨量，湿度
        const time = newVal.timestamp;
        const soilHumidity = newVal.soilHumidity;
        const precipitation = newVal.precipitation;
        const humidity = newVal.humidity;
        // 将当前数据添加到图表中
        this.options.xAxis.data.push(time);
        this.options.series[0].data.push(soilHumidity);
        this.options.series[1].data.push(humidity);
        this.options.series[2].data.push(precipitation);
        // 如果数据长度超过2000，则删除最早的数据
        if (this.options.xAxis.data.length > 2000) {
          this.options.xAxis.data.shift();
          this.options.series[0].data.shift();
          this.options.series[1].data.shift();
          this.options.series[2].data.shift();
        }
      },
      deep: true
    }
  },

  data() {
    return {
      nowTime: '',
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
                <span style="color: ${params[0].color}; margin-left:5px;">土壤湿度: ${params[0].value}%</span>
              </div>
              <div>
                <span style="display:inline-block; width:10px; height:10px; background-color:${params[1].color}; border-radius:50%;"></span>
                <span style="color: ${params[1].color}; margin-left:5px;">湿度: ${params[1].value}%</span>
              </div>
              <div>
                <span style="display:inline-block; width:10px; height:10px; background-color:${params[2].color}; border-radius:50%;"></span>
                <span style="color: ${params[2].color}; margin-left:5px;">降水量: ${params[2].value}%</span>
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
            show : false
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
            color: '#FFA500',
            // 隐藏数据点
            showSymbol: false,
            connectNulls: true,
            lineStyle: {
              color: '#FFA500',
              width: 2
            },
            emphasis: {
              focus: 'series', // 用于突出显示数据
              itemStyle: {
                color: '#ae7405'
              },
              label: {
                show: true, // 显示悬停时的数据标签
                formatter: function(params) {
                  return params.value; // 可自定义显示的内容
                }
              },
            },
          },
          {
            data: [],
            type: 'line',
            smooth: 0.5,
            symbol: 'circle',
            // 浅蓝色
            color: '#00BFFF',
            // 隐藏数据点
            showSymbol: false,
            connectNulls: true,
            lineStyle: {
              color: '#00BFFF',
              width: 2
            },
            emphasis: {
              focus: 'series', // 用于突出显示数据
              itemStyle: {
                color: '#00688B'
              },
              label: {
                show: true, // 显示悬停时的数据标签
                formatter: function(params) {
                  return params.value; // 可自定义显示的内容
                }
              },
            },
          },
          {
            data: [],
            type: 'line',
            smooth: 0.5,
            symbol: 'circle',
            // 浅蓝色
            color: '#016a8d',
            // 隐藏数据点
            showSymbol: false,
            connectNulls: true,
            lineStyle: {
              color: '#016a8d',
              width: 2
            },
            emphasis: {
              focus: 'series', // 用于突出显示数据
              itemStyle: {
                color: '#012836'
              },
              label: {
                show: true, // 显示悬停时的数据标签
                formatter: function(params) {
                  return params.value; // 可自定义显示的内容
                }
              },
            },
          }
        ]
      }
    }
  },
  methods: {
    async gptButton() {
      const res = await axios.get('http://101.43.162.244:8080/user/gpt?temperature=' + this.nowData.temperature + '&humidity=' + this.nowData.humidity + '&soilHumidity=' + this.nowData.soilHumidity + '&precipitation=' + this.nowData.precipitation);
      this.gptContent = res.data.data;
      console.log(this.gptContent);
    },
  },
}




</script>

<template>
  <div class="soilMoistureChartComponents">
    <div class="soilMoistureChartComponentsTitle">
      湿度、土壤湿度、降水量曲线图
    </div>
    <div class="soilMoistureChartComponentsChart">
      <v-chart :option="options" class="chart"></v-chart>
    </div>
    <div class="gptComponents">
      <div class="gptIcon">
        <gptIcon></gptIcon>
      </div>
      <div class="gptComponentsTitle">
        智能分析
      </div>
      <pre class="gptComponentsContent">{{this.gptContent}}</pre>
      <div class="gptButton" @click="gptButton">
      </div>
    </div>
  </div>
</template>

<style scoped>
.soilMoistureChartComponentsTitle {
  /* 设置字体大小 */
  font-size: 20px;
  /* 设置字体加粗 */
  font-weight: bold;
}
.soilMoistureChartComponents {
  /* 设置大小 */
  height: 100%;
  width: 100%;
  /* 设置中心位置 */
  position: absolute;
}
.soilMoistureChartComponentsChart {
  /* 设置大小 */
  height: 80%;
  width: 72%;
  /* 设置颜色为白色 */
  background-color: white;
  /* 设置圆角 */
  border-radius: 20px;
  /* 放到中间 */
  position: absolute;
  top: 60%;
  transform: translateY(-50%);
}
.chart {
  height: 130%;
  width: 100%;
  /* 放到中间 */
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.gptComponents {
  /* 设置大小 */
  height: 191%;
  width: 23%;
  /* 设置颜色为白色 */
  background-color: white;
  /* 设置圆角 */
  border-radius: 20px;
  /* 放到右边 */
  position: absolute;
  top: 115%;
  left: 76.5%;
  transform: translateY(-50%);
}
.gptIcon {
  /* 设置大小 */
  height: 5%;
  width: 13%;
  /* 放到左上角 */
  position: absolute;
  top: 7%;
  left: 23%;
  transform: translate(-50%, -50%);
}
.gptComponentsTitle {
  /* 设置字体大小 */
  font-size: 15px;
  /* 放到中间 */
  position: absolute;
  top: 7%;
  left: 53%;
  transform: translate(-50%, -50%);
}

.gptComponentsContent {
  /* 设置字体大小 */
  font-size: 14px;
  position: absolute;
  top: 8%;
  left: 5%;
  /* 设置边距 */
  padding: 0 5%;
  white-space: pre-wrap; /* 自动换行并保留换行符和空白 */
  width: 80%;
  height: 70%;
  overflow: auto; /* 滚动条设置 */
}

.gptButton {
  /* 设置背景图片 */
  background-image: url(../assets/analyze.png);
  /* 设置大小 */
  height: 40px;
  width: 120px;
  /* 显示完全 */
  background-size: 100% 100%;
  /* 放到正下方 */
  position: absolute;
  top: 90%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置鼠标悬停时的样式 */
  cursor: pointer;
  /* 设置过渡效果 */
  transition: all 0.5s;
  /* 设置半透明 */
  opacity: 0.5;
}

.gptButton:hover {
  /* 变大 */
  transform: translate(-50%, -50%) scale(1.1);
  opacity: 1;
}

</style>