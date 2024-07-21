<script>
export default {
  name: "PrecipitationIcon"
}
</script>

<template>
  <div class="precipitationLayer">
    <div class="precipitationIconComponents">
    </div>
  </div>

</template>

<style scoped>
.precipitationLayer {
  /* 设置背景为深蓝色 */
  background-color: #1e90ff;
  /* 设置圆角 */
  border-radius: 30%;
  /* 设置大小 */
  height: 100%;
  width: 100%;
}
.precipitationIconComponents {
  background-image: url('../../assets/precipitation.png');
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