<script>
import TemperatureIcon from "@/components/icon/TemperatureIcon.vue";
import HumidityIcon from "@/components/icon/HumidityIcon.vue";
import SoilHumidityIcon from "@/components/icon/SoilHumidityIcon.vue";
import PrecipitationIcon from "@/components/icon/PrecipitationIcon.vue";
import ElectricCurrentIcon  from "@/components/icon/ElectricCurrentIcon.vue";
import SunlightIcon from "@/components/icon/SunlightIcon.vue";


import axios from "axios";
export default {
  components: {
    TemperatureIcon,
    HumidityIcon,
    SoilHumidityIcon,
    PrecipitationIcon,
    ElectricCurrentIcon,
    SunlightIcon
  },
  props: {
    nowData: {
      type: Object,
      default: () => ({
        temperature: 0,
        humidity: 0,
        soilHumidity: 0,
        waterCurrent: 0,
        electricCurrent: 0,
        sunlight: 0
      })
    }
  },
  data() {
    return {
      temperature: 0,
      humidity: 0,
      soilHumidity: 0,
      waterCurrent: 0,
      electricCurrent: 0,
      sunlight: 0,
      intervalId: null // 定时器ID
    }
  },
  watch: {
    nowData: {
      handler(newVal) {
        this.temperature = newVal.temperature;
        this.humidity = newVal.humidity;
        this.soilHumidity = newVal.soilHumidity;
        this.waterCurrent = newVal.waterCurrent;
        this.electricCurrent = newVal.electricCurrent;
        this.sunlight = newVal.sunlight;
      },
      deep: true
    }
  },
  
};
</script>

<template>
  <div class="panelComponents">
    <div class="panelComponentsTitle">
      监测面板
    </div>
    <div class="temperaturePanel">
      <div class="panelLayer">
        <div class="temperatureIcon">
          <TemperatureIcon />
        </div>
        <div class="temperatureTitle">
          温度
        </div>
        <div class="temperatureValue">
          {{ temperature }}℃
        </div>
      </div>
    </div>
    <div class="humidityPanel">
      <div class="panelLayer">
        <div class="humidityIcon">
          <HumidityIcon />
        </div>
        <div class="humidityTitle">
          湿度
        </div>
        <div class="humidityValue">
          {{ humidity }}%
        </div>
      </div>
    </div>
    <div class="soilHumidityPanel">
      <div class="panelLayer">
        <div class="soilHumidityIcon">
          <SoilHumidityIcon />
        </div>
        <div class="soilHumidityTitle">
          土壤湿度
        </div>
        <div class="soilHumidityValue">
          {{ soilHumidity }}%
        </div>
      </div>
    </div>
    <div class="waterCurrentPanel">
      <div class="panelLayer">
        <div class="waterCurrentIcon">
          <PrecipitationIcon />
        </div>
        <div class="waterCurrentTitle">
          水流量
        </div>
        <div class="waterCurrentValue">
          {{ waterCurrent }}mm
        </div>
      </div>
    </div>

    <div class="electricCurrentPanel">
      <div class="panelLayer">
        <div class="electricCurrentIcon">
          <ElectricCurrentIcon />
        </div>
        <div class="electricCurrentTitle">
          电流量
        </div>
        <div class="electricCurrentValue">
          {{ electricCurrent }}A
        </div>
      </div>
    </div>

    <div class="sunlightPanel">
      <div class="panelLayer">
        <div class="sunlightIcon">
          <SunlightIcon />
        </div>
        <div class="sunlightTitle">
          光照强度
        </div>
        <div class="sunlightValue">
          {{ sunlight }}lux
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped>
.panelComponentsTitle {
  font-size: 20px;
  font-weight: bold;
  margin-top: 10px;
  margin-bottom: 10px;
  /* 放到左上角 */
  position: absolute;
  top: 0;
  left: 0;
}
.panelLayer {
  background-color: rgba(255, 255, 255, 1);
  height: 70%;
  width: 90%;
  /* 设置圆角 */
  border-radius: 15px;
  /* 设置阴影 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
  /* 设置位置 */
  position: absolute;
  top: 50%;
  transform: translate(0, -50%);
}
.temperaturePanel {
  /* 放到左侧 */
  position: absolute;
  top: 20%;
  left: 0;
  /* 设置大小 */
  height: 80%;
  width: 16.7%;
}
.humidityPanel {
  /* 放到左中侧 */
  position: absolute;
  top: 20%;
  left: 16.7%;
  /* 设置大小 */
  height: 80%;
  width: 16.7%;
}
.soilHumidityPanel {
  /* 放到右中侧 */
  position: absolute;
  top: 20%;
  left: 33.3%;
  /* 设置大小 */
  height: 80%;
  width: 16.7%;
}
.waterCurrentPanel {
  /* 放到右侧 */
  position: absolute;
  top: 20%;
  left: 50%;
  /* 设置大小 */
  height: 80%;
  width: 16.7%;
}
.electricCurrentPanel {
  /* 放到右侧 */
  position: absolute;
  top: 20%;
  left: 66.7%;
  /* 设置大小 */
  height: 80%;
  width: 16.7%;
}
.sunlightPanel {
  /* 放到右侧 */
  position: absolute;
  top: 20%;
  left: 83.3%;
  /* 设置大小 */
  height: 80%;
  width: 16.7%;
}
.humidityIcon {
  /* 放到左上角 */
  position: absolute;
  top: 15%;
  left: 10%;
  /* 设置大小 */
  height: 20%;
  width: 13%;
}
.humidityTitle {
  /* 放到中上 */
  position: absolute;
  top: 14%;
  left: 30%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 15px;
}
.humidityValue {
  /* 放到中间 */
  position: absolute;
  top: 40%;
  left: 10%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 30px;
  /* 设置字体粗 */
  font-weight: bold;
}
.temperatureIcon {
  /* 放到左上角 */
  position: absolute;
  top: 15%;
  left: 10%;
  /* 设置大小 */
  height: 20%;
  width: 13%;
}
.temperatureTitle {
  /* 放到中上 */
  position: absolute;
  top: 14%;
  left: 30%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 15px;
}
.temperatureValue {
  /* 放到中间 */
  position: absolute;
  top: 40%;
  left: 10%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 30px;
  /* 设置字体粗 */
  font-weight: bold;
}
.waterCurrentIcon {
  /* 放到左上角 */
  position: absolute;
  top: 15%;
  left: 10%;
  /* 设置大小 */
  height: 20%;
  width: 13%;
}
.waterCurrentTitle {
  /* 放到中上 */
  position: absolute;
  top: 14%;
  left: 30%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 15px;
}
.waterCurrentValue {
  /* 放到中间 */
  position: absolute;
  top: 40%;
  left: 10%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 30px;
  /* 设置字体粗 */
  font-weight: bold;
}
.soilHumidityIcon {
  /* 放到左上角 */
  position: absolute;
  top: 15%;
  left: 10%;
  /* 设置大小 */
  height: 20%;
  width: 13%;
}
.soilHumidityTitle {
  /* 放到中上 */
  position: absolute;
  top: 14%;
  left: 30%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 15px;
}
.soilHumidityValue {
  /* 放到中间 */
  position: absolute;
  top: 40%;
  left: 10%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 30px;
  /* 设置字体粗 */
  font-weight: bold;
}
.electricCurrentIcon {
  /* 放到左上角 */
  position: absolute;
  top: 15%;
  left: 10%;
  /* 设置大小 */
  height: 20%;
  width: 13%;
}
.electricCurrentTitle {
  /* 放到中上 */
  position: absolute;
  top: 14%;
  left: 30%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 15px;
}
.electricCurrentValue {
  /* 放到中间 */
  position: absolute;
  top: 40%;
  left: 10%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 30px;
  /* 设置字体粗 */
  font-weight: bold;
}
.sunlightIcon {
  /* 放到左上角 */
  position: absolute;
  top: 15%;
  left: 10%;
  /* 设置大小 */
  height: 20%;
  width: 13%;
}
.sunlightTitle {
  /* 放到中上 */
  position: absolute;
  top: 14%;
  left: 30%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 15px;
}
.sunlightValue {
  /* 放到中间 */
  position: absolute;
  top: 40%;
  left: 10%;
  /* 文字居中 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 30px;
  /* 设置字体粗 */
  font-weight: bold;
}
</style>