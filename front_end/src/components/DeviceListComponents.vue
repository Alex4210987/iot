<template>
  <div class="deviceListComponents">
    <div class="deviceListTitle">
      设备列表
    </div>
    <div class="deviceListComponentsLayer">
      <div class="deviceList"
           v-for="device in devices"
           :key="device.id"
           @click="selectDevice(device.id)"
           :class="{'selectedDevice': user_mcu_id === device.id}">
        <div class="deviceIcon">
          <DeviceIcon />
        </div>
        <span class="deviceName">{{ device.name }}</span>
      </div>
    </div>
  </div>
</template>

<script>
import DeviceIcon from "@/components/icon/DeviceIcon.vue";
import axios from "axios";
import { sharedState } from "@/sharedState";

export default {
  name: "DeviceListComponents",
  components: {
    DeviceIcon,
  },
  data() {
    return {
      devices: [],
    };
  },
  computed: {
    user_mcu_id() {
      return sharedState.user_mcu_id;
    },
  },
  mounted() {
    this.getDevices(); // 初始加载设备列表
    this.intervalId = setInterval(this.getDevices, 10000); // 每10秒刷新
  },
  beforeDestroy() {
    clearInterval(this.intervalId); // 清除定时器
  },
  methods: {
    async getDevices() {
      try {
        const res = await axios.get('http://localhost:8080/devices');
        this.devices = res.data.data;
        console.log(this.devices);
      } catch (error) {
        console.error("Error fetching device data:", error);
      }
    },
    selectDevice(id) {
      sharedState.user_mcu_id = id;
    },
  },
};
</script>

<style scoped>
.deviceListComponents {
  height: 100%;
  width: 100%;
}

.deviceListTitle {
  position: absolute;
  top: 0;
  left: 10%;
  font-size: 20px;
  color: #000000;
  font-weight: bold;
}

.deviceListComponentsLayer {
  background-color: rgba(255, 255, 255, 0.5);
  border-radius: 20px;
  height: 80%;
  width: 90%;
  position: absolute;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -50%);
  overflow-y: scroll;
}

.deviceList {
  display: flex;
  align-items: center;
  padding: 10px;
  cursor: pointer;
  opacity: 0.7;
  transition: all 0.3s;
}

.deviceList:hover {
  opacity: 1;
}

.deviceIcon {
  width: 24px;
  height: 24px;
  margin-right: 10px;
}

.deviceName {
  font-size: 18px;
}

.selectedDevice {
  opacity: 1;
}
</style>
