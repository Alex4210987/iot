<script>
export default {
  name: "UserIcon"
}
</script>

<template>
  <div class="userIconLayer">
    <div class="userIconComponents"></div>
  </div>
</template>

<style scoped>
.userIconLayer {
  /* 设置背景为浅绿色 */
  background-color: #abe7b8;
  /* 设置圆角 */
  border-radius: 30%;
  /* 设置大小 */
  height: 100%;
  width: 100%;
}
.userIconComponents {
  background-image: url('../../assets/user.png');
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