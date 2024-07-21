<script>
import InputField from '../components/InputFieldComponents.vue';
import axios from "axios";
export default {
  components: {
    InputField,
  },
  data() {
    return {
      show : false,
      clientX: 0,
      username: '',
      password: '',
      confirmPassword: '',
      errorMessage: "",
      loading: false,
    };
  },
  methods: {
    async register() {
      try {
        if (this.password !== this.confirmPassword) {
          this.errorMessage = "密码和确认密码不匹配，请重新输入。";
          return;
        }

        const response = await axios.post("http://101.43.162.244:8080/user/register", {
          username: this.username,
          password: this.password,
          password_repeat: this.confirmPassword,
        });

        if (response.data && response.data.token) {
          this.setCookie("token", response.data.token, 7);
          // 注册成功后跳转到首页
          this.$router.push("/index");
        } else {
          this.errorMessage = response.data.message || "注册失败，请重试。";
        }
      } catch (error) {
        console.error(error);
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
  }
}
</script>

<template>
  <div class="background">
    <div class="layer">
      <transition>
        <div v-show="show">
          <div class="logo"></div>
          <div class="text">
            <h1>注册</h1>
          </div>
          <div class="username">
            <InputField
                type="text"
                label="账户"
                @inputEvent="username = $event"
            />
          </div>
          <div class="password">
            <InputField
                type="password"
                label="密码"
                @inputEvent="password = $event"
            />
          </div>
          <div class="repeatPassword">
            <InputField
                type="password"
                label="重复密码"
                @inputEvent="confirmPassword = $event"
            />
          </div>
          <img
              class="register"
              src="../assets/register_button.png"
              alt="Register"
              @click="register"
          />
          <p class="login">已经有账号了？
            <router-link to="/login">登录</router-link>
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

/* 设置所有字体 */
* {
  font-family: 'Roboto', sans-serif;
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
  top: 30%;
  left: 50%;
  transform: translate(-50%, -50%);
}

/* 设置输入框 */
.password {
  width: 400px;
  height: 10px;
  /* 设置位置 */
  position: absolute;
  top: 45%;
  left: 50%;
  transform: translate(-50%, -50%);
}

/* 设置输入框 */
.repeatPassword {
  width: 400px;
  height: 10px;
  /* 设置位置 */
  position: absolute;
  top: 60%;
  left: 50%;
  transform: translate(-50%, -50%);
}

/* 设置注册按钮 */
.register {
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

.register:hover {
  /* 尺寸变大 */
  transform: translate(-50%, -50%) scale(1.1);
}

/* 设置注册按钮 */
.login {
  /* 设置位置 */
  position: absolute;
  top: 85%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置大小 */
  width: 120px; /* 设置宽度为100% */
  height: 20px; /* 设置高度为30px */
  /* 设置颜色 */
  color: rgba(0, 0, 0, 0.5);
  /* 设置文字大小 */
  font-size: 0.8rem;
  /* 设置文字居中 */
  text-align: center;
}

</style>
