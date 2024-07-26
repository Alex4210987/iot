#include <esp_camera.h>
#include <esp_log.h>
#include <Arduino.h>
#include <WiFi.h>
#include <HTTPClient.h>

#define WIFI_SSID "IAmGonnaTakeOff"
#define WIFI_PASSWD "y359nvdu"

// 摄像头模块的引脚定义
#define CAM_PIN_PWDN    -1 // power down is not used
#define CAM_PIN_RESET   -1 // software reset will be performed
#define CAM_PIN_XCLK    21
#define CAM_PIN_SIOD    26 // SDA
#define CAM_PIN_SIOC    27 // SCL

#define CAM_PIN_D7      35
#define CAM_PIN_D6      34
#define CAM_PIN_D5      39
#define CAM_PIN_D4      36
#define CAM_PIN_D3      19
#define CAM_PIN_D2      18
#define CAM_PIN_D1       5
#define CAM_PIN_D0       4
#define CAM_PIN_VSYNC   25
#define CAM_PIN_HREF    23
#define CAM_PIN_PCLK    22

#define HUMAN_SENSOR 32

// HTTP server configuration
//define SERVER_IP "10.12.180.104"
#define SERVER_IP "xanax.top"
#define SERVER_PORT "8080"
#define SERVER_PATH "/face/search"

// Helper macro to create the full URL
#define SERVER_URL "http://" SERVER_IP ":" SERVER_PORT SERVER_PATH


void sendPhoto() {
  // Capture image
  camera_fb_t *pic = esp_camera_fb_get();
  if (!pic) {
    Serial.println("Failed to capture image");
    return;
  }

  // Configure HTTP client
  HTTPClient http;
  http.begin(SERVER_URL);
  http.addHeader("Content-Type", "image/jpeg");

  // Send POST request with image data
  int httpResponseCode = http.POST(pic->buf, pic->len);

  // Check response code
  if (httpResponseCode > 0) {
    String response = http.getString();
    Serial.printf("HTTP Response code: %d\n", httpResponseCode);
    Serial.println(response);
  } else {
    Serial.printf("Error code: %d\n", httpResponseCode);
  }

  // End HTTP client
  http.end();

  // Return the frame buffer back to the driver
  esp_camera_fb_return(pic);
}

void checkPsram() {
  if (psramFound()) {
    Serial.println("PSRAM found and initialized");
  } else {
    Serial.println("PSRAM not found or initialization failed");
  }
}

void initWifi(const char *ssid, const char *password)
{
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
 
  while (WiFi.status() != WL_CONNECTED)
  {
    Serial.println("WiFi does not connect, try again ...");
    delay(1000);
  }
}

void setup()
{
  Serial.begin(115200);
  pinMode(HUMAN_SENSOR, INPUT);
  esp_log_level_set("camera", ESP_LOG_VERBOSE); // 设置日志级别为详细

  //连接WIFI
  initWifi(WIFI_SSID, WIFI_PASSWD);
  checkPsram(); // 检查 PSRAM 是否可用
  camera_config_t config;
  config.ledc_channel = LEDC_CHANNEL_0;
  config.ledc_timer = LEDC_TIMER_0;
  config.pin_d0 = CAM_PIN_D0;
  config.pin_d1 = CAM_PIN_D1;
  config.pin_d2 = CAM_PIN_D2;
  config.pin_d3 = CAM_PIN_D3;
  config.pin_d4 = CAM_PIN_D4;
  config.pin_d5 = CAM_PIN_D5;
  config.pin_d6 = CAM_PIN_D6;
  config.pin_d7 = CAM_PIN_D7;
  config.pin_xclk = CAM_PIN_XCLK;
  config.pin_pclk = CAM_PIN_PCLK;
  config.pin_vsync = CAM_PIN_VSYNC;
  config.pin_href = CAM_PIN_HREF;
  config.pin_sccb_sda = CAM_PIN_SIOD;
  config.pin_sccb_scl = CAM_PIN_SIOC;
  config.pin_pwdn = CAM_PIN_PWDN;
  config.pin_reset = CAM_PIN_RESET;
  config.xclk_freq_hz = 20000000;
  config.pixel_format = PIXFORMAT_JPEG;

  // 简化配置参数
  config.frame_size = FRAMESIZE_HD;
  config.jpeg_quality = 12;
  config.fb_count = 1;
  config.fb_location = CAMERA_FB_IN_PSRAM;
  config.grab_mode = CAMERA_GRAB_WHEN_EMPTY;

  // 初始化摄像头
  esp_err_t err = esp_camera_init(&config);
  if (err != ESP_OK) {
    Serial.printf("摄像头初始化失败，错误代码: 0x%x\n", err);
    return;
  }
  Serial.println("摄像头初始化成功");
}


void loop() {
  int value = digitalRead(HUMAN_SENSOR);
  if (value == 0) {
    Serial.println("No human detected");
    delay(1000);
    return;
  } else {
    Serial.println("Human detected");
    sendPhoto();
    delay(5000); // 给出一些延迟，避免频繁捕获和发送图像
  }
}

