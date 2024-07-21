<template>
  <div class="userListComponents">
    <div class="userListTitle">
      用户列表
    </div>
    <div class="userListComponentsLayer">
      <div class="userList" v-for="user in users" :key="user">
        <div class="userIcon">
          <UserIcon />
        </div>
        <span class="userName">{{ user }}</span>
      </div>
    </div>
  </div>
</template>

<script>
import UserIcon from "@/components/icon/UserIcon.vue";
import axios from "axios";

export default {
  name: "UserListComponents",
  components: {
    UserIcon,
  },
  data() {
    return {
      users: [],
      intervalId: null,
    };
  },
  mounted() {
    this.getUsers(); // 初始加载用户列表
    this.intervalId = setInterval(this.getUsers, 10000); // 每10秒刷新
  },
  beforeDestroy() {
    clearInterval(this.intervalId); // 清除定时器
  },
  methods: {
    async getUsers() {
      try {
        const res = await axios.get("http://localhost:8080/user/username");
        this.users = res.data.data; // 直接使用返回的数组
        console.log(this.users);
      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    },
  }
};
</script>

<style scoped>
.userListComponents {
  /* 设置大小 */
  height: 100%;
  width: 100%;
}

.userListTitle {
  /* 设置左上角 */
  position: absolute;
  top: 0;
  left: 10%;
  /* 设置字体大小 */
  font-size: 20px;
  /* 设置字体颜色 */
  color: #000000;
  /* 设置字体粗细 */
  font-weight: bold;
}

.userListComponentsLayer {
  /* 设置背景为白色半透明 */
  background-color: rgba(255, 255, 255, 0.5);
  /* 设置圆角 */
  border-radius: 20px;
  /* 设置大小 */
  height: 80%;
  width: 90%;
  /* 移动到中间 */
  position: absolute;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置滚动 */
  overflow-y: scroll;
}

.userList {
  display: flex;
  align-items: center;
  padding: 10px;
}

.userIcon {
  width: 24px;
  height: 24px;
  margin-right: 10px;
}

.userName {
  font-size: 18px;
}
</style>
