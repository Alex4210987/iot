<script>
export default {
  name: "DeviceIcon"
}
</script>

<template>
  <div class="deviceLayer">
    <div class="deviceIconComponents"></div>
  </div>
</template>

<style scoped>
.deviceLayer {
  /* 设置背景为浅灰色 */
  background-color: #656464;
  /* 设置圆角 */
  border-radius: 30%;
  /* 设置大小 */
  height: 100%;
  width: 100%;
}
.deviceIconComponents {
  background-image: url('../../assets/device.png');
  /* 设置大小比例为原图的大小 */
  background-size: 100% 100%;
  /* 设置大小 */
  height: 80%;
  width: 80%;
  /* 放到中间 */
  position: relative;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

</style>