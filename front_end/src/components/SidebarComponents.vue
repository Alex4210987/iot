<script>
import HomePageIcon from "@/components/icon/HomePageIcon.vue";
import AllUsersIcon from './icon/AllUsersIcon.vue'
import AllDeviceIcon from './icon/AllDeviceIcon.vue'
import HistoryIcon from './icon/HistoryIcon.vue'
import ExitIcon from './icon/ExitIcon.vue'
import MainIcon from "@/components/icon/MainIcon.vue";
import FaceIdentifyIcon from './icon/FaceIdentifyIcon.vue'; 
import {useRouter} from "vue-router";
import {ref} from "vue";
export default {
  name: "SidebarComponents",
  components: {
    MainIcon,
    HomePageIcon,
    AllUsersIcon,
    AllDeviceIcon,
    HistoryIcon,
    ExitIcon,
    FaceIdentifyIcon
  },
  setup() {
    const router = useRouter();
    const currentSelection = ref(''); // 用来存储当前选择的模块

    const navigateTo = (path, selection) => {
      currentSelection.value = selection;
      router.push(path);
    };

    return {
      navigateTo,
      currentSelection,
    };
  },
  methods: {
    ToMainPage() {
      this.navigateTo('/index', 'index');
    },
    ToAllUsersPage() {
      this.navigateTo('/allUsers', 'allUsers');
    },
    ToAllDevicePage() {
      this.navigateTo('/allDevice', 'allDevice');
    },
    ToHistoryPage() {
      this.navigateTo('/history', 'history');
    },
    ToExitPage() {
      this.navigateTo('/login', 'exit');
    },
    ToFaceIdentifyPage() {
      this.navigateTo('/faceIdentify', 'faceIdentify');
    }
  }
}
</script>

<template>
  <div class="sidebarComponents">
    <div class="home">
      <HomePageIcon />
    </div>
    <div class="main"
         @click=ToMainPage
          :class="{selected: currentSelection === 'index' || currentSelection === ''}"
    >
      <MainIcon />
    </div>
    <div class="allUsers"
         @click=ToAllUsersPage
          :class="{selected: currentSelection === 'allUsers'}"
    >
      <AllUsersIcon />
    </div>
    <div class="allDevice"
         @click=ToAllDevicePage
          :class="{selected: currentSelection === 'allDevice'}"
    >
      <AllDeviceIcon />
    </div>
    <div class="history"
         @click=ToHistoryPage
          :class="{selected: currentSelection === 'history'}"
    >
      <HistoryIcon />
    </div>
    <div class="faceIdentify"
         @click="ToFaceIdentifyPage"
         :class="{selected: currentSelection === 'faceIdentify'}"
    >
      <FaceIdentifyIcon />
    </div>
    <div class="exit"
         @click=ToExitPage
          :class="{selected: currentSelection === 'exit'}"
    >
      <ExitIcon />
    </div>
  </div>
</template>

<style scoped>
.sidebarComponents {
  /* 设置大小 */
  height: 100%;
  width: 100%;

}
.home {
  /* 放到左上角 */
  position: absolute;
  top: 2%;
  left: 5%;
  /* 设置大小比例为原图的大小 */
  background-size: 50% 100%;
  /* 设置不重复 */
  background-repeat: no-repeat;
  /* 设置大小 */
  height: 30px;
  width: 60px;
}
.main {
  /* 放到左上角 */
  position: absolute;
  top: 25%;
  left: 5%;
  /* 设置大小比例为原图的大小 */
  background-size: 50% 100%;
  /* 设置不重复 */
  background-repeat: no-repeat;
  /* 设置大小 */
  height: 30px;
  width: 60px;
  /* 设置半透明 */
  opacity: 0.5;
  /* 设置过渡效果 */
  transition: all 0.5s;
  /* 设置鼠标悬停时的样式 */
  cursor: pointer;
}
.allUsers {
  /* 放到左上角 */
  position: absolute;
  top: 35%;
  left: 5%;
  /* 设置大小比例为原图的大小 */
  background-size: 50% 100%;
  /* 设置不重复 */
  background-repeat: no-repeat;
  /* 设置大小 */
  height: 30px;
  width: 60px;
  /* 设置半透明 */
  opacity: 0.5;
  transition: all 0.5s;
  /* 设置鼠标悬停时的样式 */
  cursor: pointer;
}
.allDevice {
  /* 放到左上角 */
  position: absolute;
  top: 45%;
  left: 5%;
  /* 设置大小比例为原图的大小 */
  background-size: 50% 100%;
  /* 设置不重复 */
  background-repeat: no-repeat;
  /* 设置大小 */
  height: 30px;
  width: 60px;
  /* 设置半透明 */
  opacity: 0.5;
  transition: all 0.5s;
  /* 设置鼠标悬停时的样式 */
  cursor: pointer;
}
.history {
  /* 放到左上角 */
  position: absolute;
  top: 55%;
  left: 5%;
  /* 设置大小比例为原图的大小 */
  background-size: 50% 100%;
  /* 设置不重复 */
  background-repeat: no-repeat;
  /* 设置大小 */
  height: 30px;
  width: 60px;
  /* 设置半透明 */
  opacity: 0.5;
  transition: all 0.5s;
  /* 设置鼠标悬停时的样式 */
  cursor: pointer;
}
.faceIdentify {
  position: absolute;
  top: 65%;
  left: 5%;
  background-size: 50% 100%;
  background-repeat: no-repeat;
  height: 30px;
  width: 60px;
  opacity: 0.5;
  transition: all 0.5s;
  cursor: pointer;
}
.exit {
  /* 放到左上角 */
  position: absolute;
  top: 90%;
  left: 5%;
  /* 设置大小比例为原图的大小 */
  background-size: 50% 100%;
  /* 设置不重复 */
  background-repeat: no-repeat;
  /* 设置大小 */
  height: 30px;
  width: 60px;
  /* 设置半透明 */
  opacity: 0.5;
  transition: all 0.5s;
  /* 设置鼠标悬停时的样式 */
  cursor: pointer;
}


.home:hover,
.main:hover,
.allUsers:hover,
.allDevice:hover,
.history:hover,
.exit:hover,
.faceIdentify:hover {
  /* 设置不透明 */
  opacity: 1;
}

.selected {
  /* 变大 */
  transform: scale(1.5);
  /* 设置不透明 */
  opacity: 1;
}

</style>