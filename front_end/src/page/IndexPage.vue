<template>
  <div class="background">
    <div class="layer">
        <div class="IndexPage">
        <div class="panel">
          <PanelComponents :nowData="nowData" />
        </div>
        <div class="soilMoistureChart">
          <SoilMoistureChart :historyData="historyData" :nowData="nowData" />
        </div>
        <div class="TemperatureChart">
          <TemperatureChart :historyData="historyData" :nowData="nowData" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import PanelComponents from "@/components/PanelComponents.vue";
import SoilMoistureChart from "@/components/SoilMoistureChartComponents.vue";
import TemperatureChart from "@/components/TemperatureChartComponents.vue";
import axios from "axios";

export default {
   name: "Index",
   components: {
     PanelComponents,
     SoilMoistureChart,
     TemperatureChart
   },
  data() {
    return {
      nowData : {
        timestamp: 0,
        temperature: 0,
        humidity: 0,
        soilHumidity: 0,
        waterCurrent: 0,
        electricCurrent: 0,
      },
      historyData : [],
    }
  },
  mounted() {
    this.fetchData();
    this.gethistoryData();
    this.realtimeIntervalId = setInterval(this.fetchData, 10000);
    this.historyIntervalId = setInterval(this.gethistoryData, 10000);
  },
  beforeDestroy() {
    if (this.realtimeIntervalId) {
      clearInterval(this.realtimeIntervalId);
    }
    if (this.historyIntervalId) {
      clearInterval(this.historyIntervalId);
    }
  },
  methods: {
    async fetchData() {
      try {
        const response = await axios.get('http://localhost:3070/user/data/realtime');
        const data = response.data;
        this.nowData = {
          timestamp: new Date().toISOString(),
          temperature: data.temperature,
          humidity: data.humidity,
          soilHumidity: data.soilHumidity,
          waterCurrent: data.waterCurrent,
          electricCurrent: data.electricCurrent
        };
        this.historyData.push(this.nowData);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
    async gethistoryData() {
      try {
        const response = await axios.get('http://localhost:3070/data/history');
        const historydata = response.data;
        this.historyData = historydata;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
  }
}
</script>

<style scoped>
.background {
  background-image: url('../assets/background.png');
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  height: 100vh;
  width: 100vw;
}
.layer {
  background-color: rgba(255, 255, 255, 0.7);
  height: 90vh;
  width: 90vw;
  /* 放到中间 */
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置圆角 */
  border-radius: 15px;
  /* 设置阴影 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
}
.panel {
  /* 放到左上角 */
  position: absolute;
  top: 5%;
  left: 10%;
  /* 设置大小 */
  height: 25%;
  width: 60%;
}
.soilMoistureChart {
  /* 放到左上角 */
  position: absolute;
  top: 30%;
  left: 10%;
  /* 设置大小 */
  height: 30%;
  width: 60%;
}
.TemperatureChart {
  /* 放到左下角 */
  position: absolute;
  top: 63%;
  left: 10%;
  /* 设置大小 */
  height: 30%;
  width: 43.5%;
}
</style>
