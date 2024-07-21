<script>

import InputField from "@/components/InputFieldComponents.vue";
import axios from "axios";

export default {
  name: "AllDevicePage",
  components: {
    InputField
  },
  data() {
    return {
      allDeviceRecords: [
      ],
      showModal: false, // 控制弹窗的显示与隐藏
      deviceName: "", // 存储输入的设备名称
      //authCode: "", // 存储输入的设备授权码
     status: true,
    }
  },
  methods: {
    // 处理表单提交的方法
    async addDevice() {
      // 发送请求
      const res = await axios.post("http://10.12.168.7:3070/bind_mcu", {
        name: this.deviceName,
        status: this.status,
      });
      // 关闭弹窗
      this.showModal = false;
      // 请求成功后，重新获取设备列表
      await this.getDeviceList();
    },
    // 获取设备列表
    async getDeviceList() {
      const res = await axios.get("http://10.12.168.7:3070/mcuList");
      console.log(res.data.data);
      this.allDeviceRecords = res.data.data;
    },
  },
  mounted() {
    this.getDeviceList();
  }

}
</script>

<template>
  <div class="background">
    <div class="layer">
      <div class="allDevicePage">
        <div class="allDeviceTitle">
          设备列表
        </div>
        <div @click="showModal = true" class="addDeviceButton">
          <img src="../assets/createDevice.png" alt="add-device" class="addDeviceImage" />
          <div class="addDeviceText">添加设备</div>
        </div>
        <div class="allDeviceListComponentsLayer">
          <div class="allDeviceListHeader">
            <span class="headerItem">设备名</span>
            <span class="headerItem">MAC地址</span>
            <span class="headerItem">状态</span>
            <span class="headerItem">添加时间</span>
          </div>
          <div class="allDeviceList" v-for="record in allDeviceRecords" :key="record.id">
            <div class="deviceIcon"></div>
            <div class="deviceItem">{{ record.name }}</div>
            <div class="deviceItem">{{ record.mac }}</div>
            <div class="deviceItem">
              <span v-if="record.status">
                <img src="../assets/green.png" alt="在线" class="onlineIcon">
              </span>
              <span v-else>
                <img src="../assets/gray.png" alt="离线" class="offlineIcon">
              </span>
            </div>
            <div class="deviceItem">{{ record.time }}</div>
          </div>
          <!-- <div class="allDeviceList" >
            <div class="deviceIcon"></div>
            <div class="deviceItem">iphone13 pro max</div>
            <div class="deviceItem">aafdfdf</div>
            <div class="deviceItem">
              <span v-if="status">
                <img src="../assets/green.png" alt="在线" class="onlineIcon">
              </span>
              <span v-else>
                <img src="../assets/gray.png" alt="离线" class="offlineIcon">
              </span>
            </div>
            <div class="deviceItem">1231243</div>
          </div> -->
        </div>
      </div>
    </div>
    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <span @click="showModal = false" class="close">&times;</span>
        <div class="modalTitle">添加设备</div>
        <div class="modalInputName">
          <InputField
              type="text"
              label="设备名称"
              @inputEvent="this.deviceName = $event"
          />
        </div>
        <div class="modalInputAuthCode">
          <InputField
              type="text"
              label="设备认证码"
              @inputEvent="this.authCode = $event"
          />
        </div>
        <div @click="addDevice" class="addDeviceButtonInModal"></div>
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

.allDevicePage {
  /* 设置大小 */
  height: 100%;
  width: 68%;
  /* 移动到右边 */
  position: absolute;
  top: 50%;
  left: 40%;
  transform: translate(-50%, -50%);
}

.allDeviceTitle {
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

.allDeviceListComponentsLayer {
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

.allDeviceList {
  padding: 10px;
}

.allDeviceRecord {
  font-size: 18px;
}

.addDeviceButton {
  /* 放到右边 */
  position: absolute;
  top: 7%;
  right: 5%;
  /* 设置大小 */
  height: 4%;
  width: 20%;
  /* 设置边框黑色半透明 */
  border: 2px solid rgba(0, 0, 0, 0.5);
  /* 设置圆角 */
  border-radius: 10px;
  /* 设置过渡效果 */
  transition: all 0.3s ease;
  /* 设置鼠标悬停时的样式 */
  cursor: pointer;
}
.addDeviceImage {
  /* 放到左边 */
  position: absolute;
  top: 50%;
  left: 15%;
  transform: translate(-50%, -50%);
  /* 设置大小 */
  height: 80%;
  width: 13%;
  /* 设置半透明 */
  opacity: 0.5;
}
.addDeviceText {
  /* 放到右边 */
  position: absolute;
  top: 14%;
  right: 30%;
  /* 设置字体大小 */
  font-size: 15px;
  /* 设置字体颜色为黑色半透明 */
  color: rgba(0, 0, 0, 0.5);
  /* 设置字体粗细 */
  font-weight: bold;
}
.addDeviceButton:hover {
  /* 变大 */
  transform: scale(1.1);
}
.modal {
  /* 设置背景为黑色半透明 */
  background-color: rgba(0, 0, 0, 0.5);
  /* 设置大小 */
  height: 100%;
  width: 100%;
  /* 设置位置 */
  position: fixed;
  top: 0;
  left: 0;
  /* 设置层级 */
  z-index: 1;
  /* 设置过渡效果 */
  transition: all 0.3s ease;
}
.modal-content {
  /* 设置背景为白色 */
  background-color: #ffffff;
  /* 设置大小 */
  height: 30%;
  width: 30%;
  /* 设置位置 */
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置圆角 */
  border-radius: 20px;
  /* 设置阴影 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
}
.modalInputName {
  /* 设置位置 */
  position: absolute;
  top: 35%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置大小 */
  height: 20%;
  width: 80%;
}
.modalInputAuthCode {
  /* 设置位置 */
  position: absolute;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置大小 */
  height: 20%;
  width: 80%;
}
.modalTitle {
  /* 设置位置 */
  position: absolute;
  top: 15%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置字体大小 */
  font-size: 20px;
  /* 设置字体颜色 */
  color: #000000;
  /* 设置字体粗细 */
  font-weight: bold;
}
.addDeviceButtonInModal {
  /* 设置背景图片 */
  background-image: url('../assets/createDeviceButton.png');
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;

  /* 设置位置 */
  position: absolute;
  top: 80%;
  left: 50%;
  transform: translate(-50%, -50%);
  /* 设置大小为背景图片大小 */
  height: 13%;
  width: 23%;
  /* 设置过渡效果 */
  transition: all 0.3s ease;
  /* 设置鼠标悬停时的样式 */
  cursor: pointer;
  /* 设置半透明 */
  opacity: 0.5;
}
.addDeviceButtonInModal:hover {
  /* 变大 */
  transform: translate(-50%, -50%) scale(1.1) ;
  /* 变为不透明 */
  opacity: 1;
}
.close {
  /* 设置位置 */
  position: absolute;
  top: 5%;
  right: 2%;
  /* 设置大小为背景图片大小 */
  height: 5%;
  width: 5%;
  /* 设置过渡效果 */
  transition: all 0.3s ease;
  /* 设置鼠标悬停时的样式 */
  cursor: pointer;
  /* 设置半透明 */
  opacity: 0.5;
  scale: 1.5;
}
.close:hover {
  /* 变为不透明 */
  opacity: 1;
  /* 变大 */
  transform: scale(1.5) ;
}
/* .allDeviceListHeader {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-between;

} */
 .allDeviceListHeader {
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

.deviceItem {
  flex: 1; /* 分配相同的空间 */
  text-align: center;
  /* 设置字体大小 */
  font-size: 15px;
  /* 设置字体颜色 */
  color: #000000;
}

.allDeviceList {
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
  /* 设置大小 */
  height: 20px;
  width: 20px;
}

.offlineIcon {
  /* 设置大小 */
  height: 20px;
  width: 20px;
}

</style>