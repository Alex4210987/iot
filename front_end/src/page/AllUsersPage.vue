<script>

import axios from "axios";

export default {
  name: "allUsersPage",
  components: {
  },
  data() {
    return {
      allUsersRecords: [
      ],
    }
  },
  methods: {
    async getUsers() {
      // 携带token请求
      const res = await axios.get("http://101.43.162.244:8080/user/userList", {
        headers: {
          token: this.token,
        },
      });
      this.allUsersRecords = res.data.data;
    },
    getCookie(name) {
      const value = `; ${document.cookie}`;
      const parts = value.split(`; ${name}=`);
      if (parts.length === 2) {
        return parts.pop().split(";").shift();
      }
    },
  },
  computed: {
    token() {
      return this.getCookie('token');
    },
  },
  mounted() {
    this.getUsers();
  },

}
</script>

<template>
  <div class="background">
    <div class="layer">
      <div class="allUsersPage">
        <div class="allUsersTitle">
          用户列表
        </div>
        <div class="allUsersListComponentsLayer">
          <div class="allUsersListHeader">
            <span class="headerItem">用户名</span>
            <span class="headerItem">状态</span>
            <span class="headerItem">最后登录时间</span>
          </div>
          <!-- <div class="allUsersList" v-for="record in allUsersRecords" :key="record.id">
            <span class="allUsersItem">{{record.username}}</span>
            <span class="allUsersItem">
              <span v-if="record.status">
                <img src="../assets/green.png" alt="在线" class="onlineIcon">
              </span>
              <span v-else>
                <img src="../assets/gray.png" alt="离线" class="offlineIcon">
              </span>
            </span>
            <span class="allUsersItem">{{record.last_login_time}}</span>
          </div> -->
          <div class="allUsersList">
            <span class="allUsersItem">username</span>
            <span class="allUsersItem">
              <span v-if="true">
                <img src="../assets/green.png" alt="在线" class="onlineIcon">
              </span>
              <span v-else>
                <img src="../assets/gray.png" alt="离线" class="offlineIcon">
              </span>
            </span>
            <span class="allUsersItem">2024-07-14 16:01</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.background {
  background-image: url('../assets/background.png');
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  height: 100vh;
  width: 100vw;
}
.layer {
  background-color: rgba(255, 255, 255, 0.7);
  height: 90vh;
  width: 90vw;
  /* 放到中间 */
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置圆角 */
  border-radius: 15px;
  /* 设置阴影 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
}

.allUsersPage {
  /* 设置大小 */
  height: 100%;
  width: 68%;
  /* 移动到右边 */
  position: absolute;
  top: 50%;
  left: 40%;
  transform: translate(-50%, -50%);
}

.allUsersTitle {
  /* 设置左上角 */
  position: absolute;
  top: 6.5%;
  left: 6%;
  /* 设置字体大小 */
  font-size: 30px;
  /* 设置字体颜色 */
  color: #000000;
  /* 设置字体粗细 */
  font-weight: bold;
}

.allUsersListComponentsLayer {
  /* 设置背景为白色半透明 */
  background-color: rgba(255, 255, 255, 0.7);
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

.allUsersList {
  padding: 10px;
}

.allUsersRecord {
  font-size: 18px;
}

/* .allUsersListHeader {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-between;
} */
 .allUsersListHeader {
  display: flex;
  justify-content: space-between;
  padding: 10px;
  background-color: #f1f1f1;
  border-bottom: 1px solid #ccc;
}
.headerItem {
  flex: 1; /* 分配相同的空间 */

  text-align: center;

  /* 设置字体大小 */
  font-size: 20px;
  /* 设置字体颜色 */
  color: #000000;
  /* 加粗 */
  font-weight: bold;
}

.allUsersItem {
  flex: 1; /* 分配相同的空间 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 15px;
  /* 设置字体颜色 */
  color: #000000;
}

.allUsersList {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-around;
  padding: 10px;
  /* 设置边框 */
  border: 3px solid rgba(0, 0, 0, 0.5);
  /* 设置圆角 */
  border-radius: 30px;
  /* 设置左右边距 */
  margin: 2% 2%;
}

.onlineIcon {
  height: 20px;
  width: 20px;
}

.offlineIcon {
  height: 20px;
  width: 20px;
}


</style>