<script setup>
import { ref, watchEffect } from 'vue';
import { useRoute } from 'vue-router';
import SidebarComponents from "@/components/SidebarComponents.vue";
import RightSidebarComponents from "@/components/RightSidebarComponents.vue";

const route = useRoute();
const showSidebar = ref(true);

watchEffect(() => {
  showSidebar.value = route.path !== '/login' && route.path !== '/register';
});

</script>

<template>
  <div id="app" class="app">
    <router-view />
  </div>
  <div class="sideBar" v-if="showSidebar">
    <SidebarComponents />
  </div>
  <div class="rightSidebar" v-if="showSidebar">
    <RightSidebarComponents />
  </div>
</template>

<style>
body, html, #app {
  margin: 0;
  padding: 0;
}
.sideBar {
  /* 放到左上角 */
  position: absolute;
  top: 10%;
  left: 8%;
  /* 设置大小 */
  height: 80%;
  width: 7%;
}
.rightSidebar {
  /* 放到右上角 */
  position: absolute;
  top: 10%;
  right: 5%;
  /* 设置大小 */
  height: 80%;
  width: 23%;
}
/* 设置所有字体 */
* {
  font-family: 'roboto';
}
</style>
