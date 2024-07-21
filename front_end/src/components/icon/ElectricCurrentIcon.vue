<script>
export default {
  name: "ElectricCurrentIcon"
}
</script>

<template>
  <div class="electricCurrentLayer">
    <div class="electricCurrentIconComponents">
    </div>
  </div>

</template>

<style scoped>
.electricCurrentLayer {
  /* 设置背景为浅蓝色 */
  background-color: #03E4D8;
  /* 设置圆角 */
  border-radius: 30%;
  /* 设置大小 */
  height: 100%;
  width: 100%;
}
.electricCurrentIconComponents {
  background-image: url('../../assets/electricCurrent.png');
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