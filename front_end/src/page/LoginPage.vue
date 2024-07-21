<script>
import InputField from '../components/InputFieldComponents.vue';
import axios from "axios";
export default {
  components: {
    InputField,
  },
  data() {
    return {
      username: '',
      password: '',
      show: false,
      clientX: 0,
      loading: false,
      errorMessage: "", // 存储错误消息
    };
  },

  methods: {
    async login() {
      console.log(this.username); // 打印 username 的值
      console.log(this.password); // 打印 password 的值
      try {
        const response = await axios.post("http://localhost:3070/login", {
          username: this.username,
          password: this.password,
        });

        console.log(response.data);


        if (response.data && response.data.token) {
          this.setCookie("token", response.data.token, 7); // 设置 cookie 有效期为 7 天
          // 登录成功后跳转到首页
          this.$router.push("/index");
        } else {
          // 设置错误消息
          this.errorMessage = response.data.message || "登录失败，请重试。";
        }
      } catch (error) {
        console.error(error);
        // 设置错误消息
        this.errorMessage =
            (error.response &&
                error.response.data &&
                error.response.data.message) ||
            "网络错误，请重试。";
      }
    },
    setCookie(cname, cvalue, exdays) {
      const d = new Date();
      d.setTime(d.getTime() + exdays * 24 * 60 * 60 * 1000);
      const expires = "expires=" + d.toUTCString();
      document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
    },
  },
  mounted() {
    setTimeout(() => {
      this.show = true;
    }, 0);
  },
}
</script>

<template>
  <div class="background" >
    <div class="layer">
      <transition>
        <div v-show="show">
          <div class="logo"></div>
          <div class="text">
            <h1>登录</h1>
            <p>你好！通过登录去查看数据！</p>
          </div>
          <div class="username">
            <InputField
                type="text"
                label="账户"
                @inputEvent="this.username = $event"
            />
          </div>
          <div class="password">
            <InputField
                type="password"
                label="密码"
                @inputEvent="this.password = $event"
            />
          </div>
          <img
              class="login"
              src="../assets/log_in_button1.png"
              alt="Log In"
              @click="login"
          />
          <p class="register">没有账号？
            <router-link to="/register">注册</router-link>
          </p>
        </div>
      </transition>
    </div>
  </div>
</template>

<style scoped>
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
/* 设置背景图片 */
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
  height: 50vh;
  width: 50vw;
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
/* 设置文字 */
.text {
  /* 设置位置 */
  position: absolute;
  top: 25%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置大小 */
  width: 50%; /* 设置宽度为100% */
  height: 30%; /* 设置高度为100% */
  /* 设置颜色 */
  color: rgba(0, 0, 0, 0.5);
  /* 设置文字大小 */
  font-size: 1.2rem;
  /* 设置文字居中 */
  text-align: center;
}

h1 {
  /* 设置颜色 */
  color: rgba(0, 0, 0, 0.5);
  /* 设置文字大小 */
  font-size: 1.2rem;
  /* 设置文字居中 */
  text-align: center;
}

p {
  /* 设置颜色 */
  color: rgba(0, 0, 0, 0.5);
  /* 设置文字大小 */
  font-size: 0.8rem;
  /* 设置文字居中 */
  text-align: center;
}

/* 设置输入框 */
.username {
  width: 400px;
  height: 10px;
  /* 设置位置 */
  position: absolute;
  top: 40%;
  left: 50%;
  transform: translate(-50%, -50%);
}

/* 设置输入框 */
.password {
  width: 400px;
  height: 10px;
  /* 设置位置 */
  position: absolute;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -50%);
}


/* 设置登录按钮 */
.login {
  /* 设置位置 */
  transition: all 0.3s ease; /* 这将使标签的所有属性都以0.3秒的时间平滑过渡 */
  position: absolute;
  top: 75%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置大小 */
  width: 120px; /* 设置宽度为100% */
  height: 45px; /* 设置高度为30px */
  cursor: pointer; /* 更改鼠标光标的样式，表明这是一个可点击的元素 */
}

.login:hover {
  /* 尺寸变大 */
  transform: translate(-50%, -50%) scale(1.1);
}

/* 设置注册按钮 */
.register {
  /* 设置位置 */
  position: absolute;
  top: 85%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置大小 */
  width: 100%; /* 设置宽度为100% */
  height: 20px; /* 设置高度为30px */
  /* 设置深蓝色 */
  color: rgba(0, 0, 0, 0.5);
  /* 设置文字大小 */
  font-size: 0.8rem;
  /* 设置文字居中 */
  text-align: center;
}

</style>
