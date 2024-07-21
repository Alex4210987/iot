<template>
  <div class="background">
    <div class="layer">
      <div class="allDevicePage">
        <div class="allDeviceTitle">
          历史数据列表
        </div>
        <div class="allDeviceListComponentsLayer">
          <div class="allDeviceListHeader">
            <span class="headerItem">时间戳</span>
            <span class="headerItem">温度</span>
            <span class="headerItem">湿度</span>
            <span class="headerItem">土壤湿度</span>
            <span class="headerItem">水流量</span>
            <span class="headerItem">电流量</span>
            <span class="headerItem">光照强度</span>
          </div>
          <div class="historyList" v-for="record in historyRecords" :key="record.timestamp">
            <span class="historyItem">{{ record.timestamp }}</span>
            <span class="historyItem">{{ record.temperature }}</span>
            <span class="historyItem">{{ record.humidity }}</span>
            <span class="historyItem">{{ record.soil_moisture }}</span>
            <span class="historyItem">{{ record.water_discharge }}</span>
            <span class="historyItem">{{ record.electric_current }}</span>
            <span class="historyItem">{{ record.sunlight }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "AllDevicePage",
  data() {
    return {
      historyRecords: [],
    }
  },
  methods: {
    async getHistoryRecords() {
      try {
        const res = await axios.get("http://localhost:3070/data/history");
        this.historyRecords = res.data.historyRecords;
        console.log(this.historyRecords)
      } catch (error) {
        console.error("Error fetching history records:", error);
      }
    },
  },
  mounted() {
    this.getHistoryRecords();
  }
}
</script>

<style scoped>
/* Add your styles here */
</style>
