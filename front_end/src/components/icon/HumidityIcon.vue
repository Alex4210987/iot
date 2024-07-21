<script>
export default {
  name: "HumidityIcon"
}
</script>

<template>
  <div class="humidityLayer">
    <div class="humidityIconComponents">
    </div>
  </div>

</template>

<style scoped>
.humidityLayer {
  /* 设置背景为浅蓝色 */
  background-color: #58bfef;
  /* 设置圆角 */
  border-radius: 30%;
  /* 设置大小 */
  height: 100%;
  width: 100%;
}
.humidityIconComponents {
  background-image: url('../../assets/humidity.png');
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