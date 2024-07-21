<script>
export default {
  name: "gptIcon"
}
</script>

<template>
  <div class="gptLayer">
    <div class="gptIconComponents">
    </div>
  </div>

</template>

<style scoped>
.gptLayer {
  /* 设置背景为蓝色 */
  background-color: #00bfff;
  /* 设置圆角 */
  border-radius: 30%;
  /* 设置大小 */
  height: 100%;
  width: 100%;
}
.gptIconComponents {
  background-image: url('../../assets/gpt.png');
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