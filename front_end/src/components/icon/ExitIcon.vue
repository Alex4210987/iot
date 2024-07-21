<script>
export default {
  name: "ExitIcon"
}
</script>

<template>
  <div class="exitIcon">
    </div>
</template>

<style scoped>
.exitIcon {
  background-image: url('../../assets/exit.png');
  /* 设置大小比例为原图的大小 */
  background-size: 50% 100%;
  /* 设置不重复 */
  background-repeat: no-repeat;
  /* 设置大小 */
  height: 100%;
  width: 100%;
  /* 放到中间 */
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>