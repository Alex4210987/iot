<script>
import UserList from "@/components/UserListComponents.vue";
import UserIcon from "@/components/icon/UserIcon.vue";
import DeviceList from "@/components/DeviceListComponents.vue";
import axios from "axios";

export default {
  name: "RightSidebarComponents",
  components: { DeviceList, UserIcon, UserList },
  data() {
    return {
      user: {
        name: "", // 用户名
        onlineUsers: {}, // 在线用户（如果需要）
      },
    };
  },
  computed: {
    token() {
      return this.getCookie("token");
    },
  },
  watch: {
    token: {
      immediate: true, // 立即触发
      handler: function (newToken, oldToken) {
        if (newToken !== oldToken) {
          this.getUser(); // 获取用户信息
        }
      },
    },
  },
  methods: {
    getCookie(name) {
      const value = `; ${document.cookie}`;
      const parts = value.split(`; ${name}=`);
      if (parts.length === 2) {
        return parts.pop().split(";").shift();
      }
    },
    async getUser() {
      try {
        // 请求获取当前用户信息
        const userRes = await axios.get("http://localhost:3070/current-user", {
          headers: {
            Authorization: `Bearer ${this.token}`, // 使用 Bearer 方案传递 token
          },
        });

        // 更新当前用户信息
        const currentUser = userRes.data.current_user;
        this.user.name = currentUser;

        // 请求获取在线用户列表
        const onlineRes = await axios.get("http://localhost:3070/online-users", {
          headers: {
            Authorization: `Bearer ${this.token}`, // 使用 Bearer 方案传递 token
          },
        });

        // 处理在线用户数据
        const onlineUsers = onlineRes.data.online_users;
        if (Object.keys(onlineUsers).length > 0) {
          // 获取第一个用户
          const firstUser = Object.keys(onlineUsers)[0];
          this.user.onlineUsers = { [firstUser]: onlineUsers[firstUser] }; // 设置第一个用户
        } else {
          this.user.onlineUsers = {}; // 如果没有在线用户
        }

      } catch (error) {
        console.error("Error fetching user or online users:", error);
      }
    },
  },
};
</script>

<template>
  <div class="rightSidebarComponents">
    <div class="verticalLine"></div>
    <div class="user">
      <div class="userIcon">
        <UserIcon />
      </div>
      <div class="userName">
        {{ user.name }}
      </div>
    </div>
    <div class="userList">
      <UserList />
    </div>
    <div class="deviceList">
      <DeviceList />
    </div>
  </div>
</template>

<style scoped>
.rightSidebarComponents {
  /* 设置大小 */
  height: 100%;
  width: 100%;
}

.verticalLine {
  position: absolute;
  height: 80%; /* 调整线的高度 */
  width: 2px; /* 设置线的宽度 */
  left: 0; /* 调整线的水平位置 */
  top: 15%; /* 调整线的垂直位置 */
  /* 设置黑色半透明 */
  background-color: rgba(0, 0, 0, 0.1);
}

.user {
  /* 放到右上角 */
  position: absolute;
  top: 5%;
  right: 50%;
  transform: translate(50%, 0%);
  /* 设置大小 */
  height: 5%;
  width: 80%;
  /* 设置不重复 */
  background-repeat: no-repeat;
}
.userIcon {
  /* 放到中间 */
  position: absolute;
  top: 50%;
  left: 20%;
  transform: translate(-50%, -50%);
  /* 设置大小 */
  height: 80%;
  width: 10%;
}
.userName {
  /* 放到中间 */
  position: absolute;
  top: 45%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置字体 */
  font-size: 20px;
  font-weight: bold;
}

.userList {
  /* 放到右上角 */
  position: absolute;
  top: 15%;
  right: 50%;
  transform: translate(50%, 0%);
  /* 设置大小 */
  height: 40%;
  width: 80%;
}
.deviceList {
  /* 放到右下角 */
  position: absolute;
  top: 55%;
  right: 50%;
  transform: translate(50%, 0%);
  /* 设置大小 */
  height: 40%;
  width: 80%;
}
</style>
