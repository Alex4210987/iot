<script>
export default {
  name: "TemperatureIcon"
}
</script>

<template>
  <div class="temperatureLayer">
    <div class="temperatureIconComponents">
    </div>
  </div>

</template>

<style scoped>
.temperatureLayer {
  /* 设置背景为浅红色 */
  background-color: #f48888;
  /* 设置圆角 */
  border-radius: 30%;
  /* 设置大小 */
  height: 100%;
  width: 100%;
}
.temperatureIconComponents {
  background-image: url('../../assets/temperature.png');
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