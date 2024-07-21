<template>
  <div class="background">
    <div class="layer">
        <div class="faceIdentifyPage">
          <div class="faceIdentifyTitle">
            人脸识别
          </div>
          <div @click="showModal = true" class="addFaceButton">
            <img src="../assets/createDevice.png" alt="add-face" class="addFaceImage" />
            <div class="addFaceText">添加人脸</div>
          </div>
          <div class="allFaceListComponentsLayer">
            <div class="allFaceListHeader">
              <span class="headerItem">编号</span>
              <span class="headerItem">人脸图片</span>
              <span class="headerItem">创建时间</span>
              <span class="headerItem">操作</span>
            </div>
            <!-- <div class="allFaceList"> -->
              <!-- <div v-for="record in allFaceRecords" :key="record.id" class="faceListItem">
                <span>{{ record.id }}</span>
                <img :src="record.image" alt="Face Image" class="faceImage" />
                <span>{{ record.createdAt }}</span>
                <button @click="deleteFace(record.id)">删除</button>
              </div> -->
               <div class="faceList">
                  <div class="faceListItem">
                    <div class="faceListItemCell">111</div>
                    <div class="faceListItemCell"><img src="@/assets/1.png" alt="Face Image" class="faceImage" /></div>
                    <div class="faceListItemCell">2024-07-13 17:13</div>
                    <div class="faceListItemCell"><button @click="deleteFace(111)" class="actionButton">删除</button></div>
                  </div>
              </div>
              <div class="faceList">
                  <div class="faceListItem">
                    <div class="faceListItemCell">222</div>
                    <div class="faceListItemCell"><img src="@/assets/2.png" alt="Face Image" class="faceImage" /></div>
                    <div class="faceListItemCell">2024-07-13 17:13</div>
                    <div class="faceListItemCell"><button @click="deleteFace(111)" class="actionButton">删除</button></div>
                  </div>
              </div>
            <!-- </div> -->
          </div>
        </div>
    </div>
  </div>
    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <span @click="showModal = false" class="close">&times;</span>
        <div class="modalTitle">添加人脸</div>
        <div class="modal-body">
          <input type="file" @change="onFileChange" class="modal-input"/>
          <button @click="uploadFace" class="modal-button">提交图片</button>
        </div>
      </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      showModal: false,
      file: null,
      records: []
    };
  },
  methods: {
    onFileChange(event) {
      this.file = event.target.files[0];
    },
    async uploadFace() {
      try {
        const formData = new FormData();
        formData.append('file', this.file);

        const response = await axios.post('http://10.12.168.7:3070/face/search', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        });

        console.log('上传成功:', response);
        alert('图片上传成功！');
        // 上传成功后更新记录
        this.fetchRecords();
        this.closeDialog(); // 关闭弹窗
      } catch (error) {
        console.error('上传失败:', error);
        alert('图片上传失败！');
      }
    },
    async fetchRecords() {
      try {
        const response = await axios.get('http://localhost:3070/face/records');
        this.records = response.data;
        console.log(response.data);
      } catch (error) {
        console.error('获取记录失败:', error);
      }
    },
    closeDialog() {
      this.showDialog = false;
      this.file = null; // 清空选择的文件
    },
    async deleteFace(faceId) {
      try {
        const response = await axios.delete(`http://localhost:3000/face/records/${faceId}`);
        if (response.status === 200) {
          this.records = this.records.filter(record => record.id !== faceId);
        } else {
          console.error('Failed to delete face:', response.statusText);
        }
      } catch (error) {
        console.error('Error deleting face:', error);
      }
    }
  },
  mounted() {
    this.fetchRecords();
    setInterval(this.fetchRecords, 10000); // 每10秒获取一次记录
  }
};
</script>

<style>
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
.faceIdentifyPage {
  /* 设置大小 */
  height: 100%;
  width: 68%;
  /* 移动到右边 */
  position: absolute;
  top: 50%;
  left: 40%;
  transform: translate(-50%, -50%);
}
/* .face-recognition {
  padding: 20px;
} */
.faceIdentifyTitle {
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
.addFaceButton {
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
.addFaceImage {
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
.addFaceText {
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
.addFaceButton:hover {
  /* 变大 */
  transform: scale(1.1);
}
.allFaceListComponentsLayer {
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
.allFaceList {
  padding: 10px;
}
.allFaceRecord {
  font-size: 18px;
}
 .allFaceListHeader {
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

.faceList {
  display: flex;
  flex-direction: column;
}

.faceListItem {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #ccc;
}

.faceListItemCell {
  flex: 1;
  text-align: center;
  display: flex;
  justify-content: center;
  align-items: center;
}

.faceImage {
  max-width: 40px;
  max-height: 40px;
  object-fit: contain;
}

.actionButton {
  padding: 5px 10px;
  background-color: #ff4d4d;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.actionButton:hover {
  background-color: #ff0000;
}

.modal {
  background-color: rgba(0, 0, 0, 0.5);
  height: 100%;
  width: 100%;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1;
  transition: all 0.3s ease;
}

.modalTitle {
  font-size: 20px;
  color: #000000;
  font-weight: bold;
  text-align: center;
  margin-bottom: 20px;
}

.modal-content {
  background-color: #ffffff;
  height: auto;
  width: 30%;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.modal-body {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.modal-input {
  width: 80%;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.modal-button {
  width: 50%;
  padding: 10px;
  margin-top: 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.modal-button:hover {
  background-color: #45a049;
}

.modal-button:hover {
  background-color: #45a049;
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
</style>