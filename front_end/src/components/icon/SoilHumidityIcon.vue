<script>
export default {
  name: "SoilHumidityIcon"
}
</script>

<template>
  <div class="soilHumidityLayer">
    <div class="soilHumidityIconComponents">
    </div>
  </div>

</template>

<style scoped>
.soilHumidityLayer {
  /* 设置背景为土色 */
  background-color: #d2b48c;
  /* 设置圆角 */
  border-radius: 30%;
  /* 设置大小 */
  height: 100%;
  width: 100%;
}
.soilHumidityIconComponents {
  background-image: url('../../assets/soilHumidity.png');
  /* 设置大小比例为原图的大小 */
  background-size: 100% 100%;
  /* 设置大小 */
  height: 80%;
  width: 80%;
  /* 放到中间 */
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

</style>