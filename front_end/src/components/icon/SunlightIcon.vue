<script>
export default {
  name: "SunlightIcon"
}
</script>

<template>
  <div class="sunlightLayer">
    <div class="sunlightIconComponents">
    </div>
  </div>

</template>

<style scoped>
.sunlightLayer {
  /* 设置背景为浅蓝色 */
  background-color: #DEAA76;
  /* 设置圆角 */
  border-radius: 30%;
  /* 设置大小 */
  height: 100%;
  width: 100%;
}
.sunlightIconComponents {
  background-image: url('../../assets/sunlight.png');
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