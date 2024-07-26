 
#include <WiFi.h>
#include <ArduinoJson.h>
#include <huaweiyun_mqtt.h>
#include <DHT.h>
#include<ArduinoJson.h>
#include <SHA256.h>
#include <DHT_U.h>
#include <MQUnifiedsensor.h>
#include "EmonLib.h"  
#include <FlowSensor.h>
#include<math.h>


EnergyMonitor emon1;
#define CURRENT_CALIBRATION 111.1
 
//上云时间现态
unsigned long lastSend = 0;
//开启计时标志
bool countingFlag = false;
//计时时间
unsigned long doorOpenTime = 0;
 
//你的wifi
#define WIFI_SSID "IAmGonnaTakeOff"
#define WIFI_PASSWD "y359nvdu"

//华为云上云条件
#define server_ip "120ee0b540.st1.iotda-device.cn-north-4.myhuaweicloud.com"
#define device_id "6663d8537dbfd46fabbf54b9_device_"
#define device_secret "12345678"
#define this_time "2021020403"
#define product_id "6663d8537dbfd46fabbf54b9"
#define project_id "a129d472438743eda97676e5a925f87d"
#define instance_id "6044ea76-7cc7-4236-ac36-1bdfb8ab5528"

//属性上报Topic---发布
#define HUAWEI_TOPIC_PROP_POST "$oc/devices/" device_id "/sys/properties/report"
//平台下发命令给设备的频道
#define HUAWEI_TOPIC_PROP_CMD "$oc/devices/" device_id "/sys/commands/#"
//设备上报命令执行结果的频道  $oc/devices/{device_id}/sys/commands/response/request_id={request_id}
#define HUAWEI_TOPIC_PROP_CMD_RESP "$oc/devices/" device_id "/sys/commands/response"

// 继电器控制输出端
#define INDOOR_LIGHT_SWITCH_PIN 23
#define BUZZER_SWITCH_PIN 22
#define AIR_CONDITIONER_SWITCH_PIN 21
#define ACCESS_CONTROL_SWITCH_PIN 16
#define WINDOW_SWITCH_PIN 18
#define PUMP_SWITCH_PIN 5
#define HUMIDIFIER_SWITCH_PIN 17
#define EXTERNAL_LIGHT_SWITCH_PIN 19

// 传感器输入
#define TEMP_SENSOR_PIN 33
#define RAIN_SENSOR_PIN 25
#define FLAME_SENSOR_PIN 26
#define LIGHT_SENSOR_PIN 34
#define WATER_FLOW_SENSOR_PIN 35
#define CURRENT_SENSOR_PIN 32
#define SOIL_MOISTURE_SENSOR_PIN 36
#define AIR_QUALITY_SENSOR_PIN 39
#define HUMAN_SENSOR_PIN 27


#define DHTTYPE DHT11   // 定义传感器类型为DHT11

#define placa "ESP32"
#define Voltage_Resolution 3.3 // ESP32 的 ADC 电压分辨率是 3.3V
#define pin 39// 使用 ESP32 的 GPIO35 作为模拟输入
#define type "MQ-135"
#define ADC_Bit_Resolution 12 // ESP32 的 ADC 位分辨率是 12 位
#define RatioMQ135CleanAir 3.6 // RS / R0 = 3.6 ppm

MQUnifiedsensor MQ135(placa, Voltage_Resolution, ADC_Bit_Resolution, pin, type);

#define FLOW_TYPE YFS201
FlowSensor flowsensor(FLOW_TYPE,WATER_FLOW_SENSOR_PIN);
unsigned long reset = 0;

// Uncomment if use ESP8266 and ESP32
void IRAM_ATTR count()
{
  flowsensor.count();
}
// void count()
// {
// 	flowsensor.count();
// }



//创建DHT实例
DHT_Unified dht(TEMP_SENSOR_PIN, DHTTYPE);


//创建WiFiClient实例
WiFiClient espClient;
//创建MqttClient实例
PubSubClient mqttClient(espClient);

float input_to_lux(uint16_t x) {
    //float e = exp(1); // 自然对数的底 e
    return (40000) * pow(x, -0.6021);
}

//连接Wifi
void initWifi(const char *ssid, const char *password)
{
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
 
  while (WiFi.status() != WL_CONNECTED)
  {
    Serial.println("WiFi does not connect, try again ...");
    delay(1000);
  }
 
  Serial.println("Wifi is connected.");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
}


void init_mq135()
{
  MQ135.setRegressionMethod(1); // 设置回归模型用于计算 PPM 浓度

  MQ135.init(); // 初始化传感器
  Serial.print("Calibrating please wait.");
  float calcR0 = 0;
  for (int i = 1; i <= 10; i++) {
    MQ135.update(); // 更新数据，ESP32 将从模拟引脚读取电压
    calcR0 += MQ135.calibrate(RatioMQ135CleanAir);
    Serial.print(".");
  }
  MQ135.setR0(calcR0 / 10);
  Serial.println("  done!");

  if (isinf(calcR0)) {
    Serial.println("Warning: Connection issue, R0 is infinite (Open circuit detected) please check your wiring and supply");
    while (1);
  }
  if (calcR0 == 0) {
    Serial.println("Warning: Connection issue found, R0 is zero (Analog pin shorts to ground) please check your wiring and supply");
    while (1);
  }

  Serial.println("** Values from MQ-135 ****");
  Serial.println("|    CO   |  Alcohol |   CO2  |  Toluen  |  NH4  |  Aceton  |");
}
 
 
//连接Mqtt并订阅属性设置Topic
bool mqttCheckConnect()
{
  bool connected = connectHuaweiyunMQTT(mqttClient,server_ip,device_id,device_secret,this_time);

  if (connected)
  {
    Serial.println("MQTT connect succeed!");
    //订阅属性设置Topic,订阅云端下发指令频道
    Serial.println("subscribe done");
    return true;
  }
  else
  {
    Serial.println("MQTT connect failed!");
    return false;
  }
}

// 上报属性Topic数据
void mqttIntervalPost(float temperature,float humidity,int air_quality,bool rainfall
,bool windows,bool pump,float soil_moisture,float electric_current,
float water_discharge,bool air_conditioner,int sunlight,bool external_light,
bool indoor_light,bool human_existence,bool fire_occurence,bool humidifier,float CO,float Alcohol,float CO2,float Toluen,float NH4,float Aceton,float total) 
{
  StaticJsonDocument<768> doc;
  char Output[768];

  // 创建services数组
  JsonArray services = doc.createNestedArray("services");

  // 向services数组添加对象
  JsonObject service1 = services.createNestedObject();
  service1["service_id"] = "atmospheric_environment";
  JsonObject properties1 = service1.createNestedObject("properties");
  properties1["temperature"] = temperature;
  properties1["humidity"] = humidity;
  properties1["air_quality"] = air_quality;
  properties1["rainfall"] = rainfall? "true" : "false";
  properties1["windows"] = windows? "true" : "false";
  properties1["pump"] =   pump? "true" : "false";
  properties1["soil_moisture"] = soil_moisture;
  properties1["co"] = CO;
  properties1["alcohol"] = Alcohol;
  properties1["co2"] = CO2;
  properties1["toluen"] = Toluen;
  properties1["nh4"] = NH4;
  properties1["aceton"] = Aceton;


  JsonObject service2 = services.createNestedObject();
  service2["service_id"] = "park_energy";
  JsonObject properties2 = service2.createNestedObject("properties");
  properties2["electric_current"] =  electric_current;
  properties2["water_discharge"] = water_discharge;
  properties2["air_conditioner"] =  air_conditioner? "true" : "false";
  properties2["total_water_usage"] = total;

  JsonObject service3 = services.createNestedObject();
  service3["service_id"] = "park_lighting";
  JsonObject properties3 = service3.createNestedObject("properties");
  properties3["sunlight"] =   sunlight;
  properties3["external_light"] =   external_light? "true" : "false";
  properties3["indoor_light"] =   indoor_light? "true" : "false";

  JsonObject service4 = services.createNestedObject();
  service4["service_id"] = "personnal_access";
  JsonObject properties4 = service4.createNestedObject("properties");
  properties4["human_existence"] =   human_existence? "true" : "false";

  JsonObject service5 = services.createNestedObject();
  service5["service_id"] = "park_fire_protection";
  JsonObject properties5 = service5.createNestedObject("properties");
  properties5["fire_occurence"] =   fire_occurence? "true" : "false";
  properties5["humidifier"] =   humidifier? "true" : "false";

  // 将JSON文档输出到串行监视器
  serializeJson(doc, Output);
  //serializeJsonPretty(doc, Serial);

  //Serial.println(Output);
  mqttClient.publish(HUAWEI_TOPIC_PROP_POST,Output); 
}





void processCommand(StaticJsonDocument<256> &doc, std::string requestId)
{
  // StaticJsonDocument<256> paras;
  // deserializeJson(paras, doc["paras"]);

  JsonObject paras = doc["paras"];
        //判断指令名称
  if (strcmp(doc["command_name"],"atmospheric_environment_commands")==0)
  {
          //判断指令参数
    if (paras.containsKey("window_switch")) {
      bool windowSwitch = paras["window_switch"].as<bool>();
      // Do something with windowSwitch
      digitalWrite(WINDOW_SWITCH_PIN, windowSwitch ? LOW : HIGH);
    }
    if (paras.containsKey("pump_switch")) {
      bool pumpSwitch = paras["pump_switch"].as<bool>();
      // Do something with pumpSwitch
      digitalWrite(PUMP_SWITCH_PIN, pumpSwitch ? LOW : HIGH);
    }
  }
  if (strcmp(doc["command_name"],"air_conditioner_commands")==0)
  {
    //判断指令参数
    if (paras.containsKey("air_conditioner_switch"))
    {
      bool airConditionerSwitch = paras["air_conditioner_switch"].as<bool>();
      // Do something with airConditionerSwitch
      digitalWrite(AIR_CONDITIONER_SWITCH_PIN, airConditionerSwitch ? LOW : HIGH);
    }
  }
  
  if (strcmp(doc["command_name"],"light_switch_commands")==0)
  {
    //判断指令参数
    //Serial.println("light_switch_commands!!!!!!!!!!!!!");
    //serializeJsonPretty(paras, Serial);
    if (paras.containsKey("external_light_switch"))
    {
      bool externalLightSwitch = paras["external_light_switch"].as<bool>();
      // Do something with externalLightSwitch
      digitalWrite(EXTERNAL_LIGHT_SWITCH_PIN, externalLightSwitch ? LOW : HIGH);
    }
    if (paras.containsKey("indoor_light_switch"))
    {
      bool indoorLightSwitch = paras["indoor_light_switch"].as<bool>();
      Serial.println(indoorLightSwitch);
      // Do something with indoorLightSwitch
      digitalWrite(INDOOR_LIGHT_SWITCH_PIN, indoorLightSwitch ? LOW : HIGH);
    }
  }
  if (strcmp(doc["command_name"],"access_control_commands")==0)
  {
    //判断指令参数
    if (paras.containsKey("access_control_switch"))
    {
      bool accessControlSwitch = paras["access_control_switch"].as<bool>();
      // Do something with accessControlSwitch
      digitalWrite(ACCESS_CONTROL_SWITCH_PIN, accessControlSwitch ? LOW : HIGH);
    }
  }
  if (strcmp(doc["command_name"],"park_fire_protection_commands")==0)
  {
    //判断指令参数
    if (paras.containsKey("buzzer_switch"))
    {
      bool buzzerSwitch = paras["buzzer_switch"].as<bool>();
      // Do something with buzzerSwitch
      digitalWrite(BUZZER_SWITCH_PIN, buzzerSwitch ? LOW : HIGH);
    }
    if (paras.containsKey("humidifier_switch"))
    {
      bool humidifierSwitch = paras["humidifier_switch"].as<bool>();
      // Do something with humidifierSwitch
      digitalWrite(HUMIDIFIER_SWITCH_PIN, humidifierSwitch ? LOW : HIGH);
    }
  }
}


void sendCommandResponse(std::string requestId)
{
  std::string responseTopic = "$oc/devices/" + std::string(device_id) + "/sys/commands/response/request_id=" + requestId;
  std::string responsePayload = R"({
    "result_code": 0,
    "response_name": "COMMAND_RESPONSE",
    "paras": {
        "result": "success"
    }
  })";
  mqttClient.publish(responseTopic.c_str(), responsePayload.c_str());
}
 
//监听云端下发指令并处理 
void callback(char *topic, byte *payload, unsigned int length)
{
    
  Serial.println();
  Serial.println();
  Serial.print("Message arrived [");
  Serial.print(topic);
  Serial.print("] ");
  Serial.println();

  payload[length] = '\0';
  Serial.println((char *)payload);

  //json解析payload
  StaticJsonDocument<256> doc;
  DeserializationError error = deserializeJson(doc, payload);
  std::string requestId;

  //判断json解析是否成功
  if (!error){
  std::string topicStr(topic);
  std::string baseSearchString = "$oc/devices/" + std::string(device_id) + "/sys/commands/request_id=";
  size_t requestIdPos = topicStr.find(baseSearchString);
  if (requestIdPos != std::string::npos)
  {
    // 提取request_id
    requestId = topicStr.substr(requestIdPos + baseSearchString.length());

    //判断是否为云端下发的指令
    if (doc.containsKey("command_name"))
    {
      //一个处理指令函数
      processCommand(doc,requestId);
      //发送响应报文
      sendCommandResponse(requestId);
    }
    else{
        Serial.println("报文格式错误");}
    }
    else{
      Serial.println("无法提取request_id");}
  }
  else{
      Serial.println("json解析失败");}
}

void timing_attribute_reporting()
{
  // 读取传感器数据
  float temperature ;
  float humidity ;
  sensors_event_t event;
  dht.temperature().getEvent(&event);
  if (isnan(event.temperature)) {
    Serial.println(F("Error reading temperature!"));
  }
  else {
    temperature = event.temperature;
  }
  dht.humidity().getEvent(&event);
  if (isnan(event.relative_humidity)) {
    Serial.println(F("Error reading humidity!"));
  }
  else {
    humidity = event.relative_humidity;
  }

  //空气质量传感器读数
  MQ135.update(); // 更新数据，ESP32 将从模拟引脚读取电压

  MQ135.setA(605.18); MQ135.setB(-3.937); // 设置计算 CO 浓度的方程
  float CO = MQ135.readSensor(); // 读取 CO 浓度

  MQ135.setA(77.255); MQ135.setB(-3.18); // 设置计算酒精浓度的方程
  float Alcohol = MQ135.readSensor(); // 读取酒精浓度

  MQ135.setA(110.47); MQ135.setB(-2.862); // 设置计算 CO2 浓度的方程
  float CO2 = MQ135.readSensor()+400; // 读取 CO2 浓度

  MQ135.setA(44.947); MQ135.setB(-3.445); // 设置计算甲苯浓度的方程
  float Toluen = MQ135.readSensor(); // 读取甲苯浓度

  MQ135.setA(102.2); MQ135.setB(-2.473); // 设置计算氨气浓度的方程
  float NH4 = MQ135.readSensor(); // 读取氨气浓度

  MQ135.setA(34.668); MQ135.setB(-3.369); // 设置计算丙酮浓度的方程
  float Aceton = MQ135.readSensor(); // 读取丙酮浓度

  Serial.print("|   "); Serial.print(CO);
  Serial.print("   |   "); Serial.print(Alcohol);
  Serial.print("   |   "); Serial.print(CO2); // 添加 400 PPM 的偏移量到 CO2 值上
  Serial.print("   |   "); Serial.print(Toluen);
  Serial.print("   |   "); Serial.print(NH4);
  Serial.print("   |   "); Serial.print(Aceton);
  Serial.println("   |");
  float air_quality = Alcohol;
  //float air_quality = analogRead(AIR_QUALITY_SENSOR_PIN);
  //雨滴是没雨的时候为高电压，有雨的时候为低电压
  bool rainfall = !digitalRead(RAIN_SENSOR_PIN);
  bool windows = !digitalRead(WINDOW_SWITCH_PIN);
  bool pump = !digitalRead(PUMP_SWITCH_PIN);

  int soil_moisture_analread = analogRead(SOIL_MOISTURE_SENSOR_PIN);
  int min=1400;
  int max=4095;
  float soil_moisture = (max-soil_moisture_analread)*100/(max-min);
  Serial.println(soil_moisture);

  float  electric_current = emon1.calcIrms(1480);  // Calculate Irms only
  
  flowsensor.read();
  float water_discharge=flowsensor.getFlowRate_m();
  Serial.print("Water discharge: ");
  Serial.println(water_discharge);
  float total = flowsensor.getVolume();
  Serial.print("Total: ");
  Serial.println(total);
  
  bool air_conditioner = !digitalRead(AIR_CONDITIONER_SWITCH_PIN);
  uint16_t sunlight_input = analogRead(LIGHT_SENSOR_PIN);
  float sunlight = input_to_lux(sunlight_input);
  Serial.print("Sunlight: ");
  Serial.println(sunlight);
  //大于2000开灯
  //int sunlight = digitalRead(LIGHT_SENSOR_PIN);
  bool external_light = !digitalRead(EXTERNAL_LIGHT_SWITCH_PIN);
  bool indoor_light = !digitalRead(INDOOR_LIGHT_SWITCH_PIN);
  bool human_existence = digitalRead(HUMAN_SENSOR_PIN);
  bool fire_occurence = !digitalRead(FLAME_SENSOR_PIN);
  bool humidifier = !digitalRead(HUMIDIFIER_SWITCH_PIN);


  // 上报属性数据
  mqttIntervalPost(temperature, humidity, air_quality, rainfall, windows,
   pump, soil_moisture, electric_current, water_discharge, air_conditioner, 
   sunlight, external_light, indoor_light, human_existence, fire_occurence,
    humidifier, CO, Alcohol, CO2, Toluen, NH4, Aceton, total);
}
 
void setup() 
{

  Serial.begin(9600);
  dht.begin();
  sensor_t sensor;
  dht.temperature().getSensor(&sensor);
  dht.humidity().getSensor(&sensor);
  
  flowsensor.begin(count);

  init_mq135();
  //初始化继电器输出端
  pinMode(INDOOR_LIGHT_SWITCH_PIN, OUTPUT);
  pinMode(BUZZER_SWITCH_PIN, OUTPUT);
  pinMode(AIR_CONDITIONER_SWITCH_PIN, OUTPUT);
  pinMode(ACCESS_CONTROL_SWITCH_PIN, OUTPUT);
  pinMode(WINDOW_SWITCH_PIN, OUTPUT);
  pinMode(PUMP_SWITCH_PIN, OUTPUT);
  pinMode(HUMIDIFIER_SWITCH_PIN, OUTPUT);
  pinMode(EXTERNAL_LIGHT_SWITCH_PIN, OUTPUT);

  //初始化传感器输入端
  pinMode(TEMP_SENSOR_PIN, INPUT);
  pinMode(RAIN_SENSOR_PIN, INPUT);
  pinMode(FLAME_SENSOR_PIN, INPUT);
  pinMode(LIGHT_SENSOR_PIN, INPUT);
  //pinMode(WATER_FLOW_SENSOR_PIN, INPUT);

  emon1.current(CURRENT_SENSOR_PIN, CURRENT_CALIBRATION); 

  pinMode(SOIL_MOISTURE_SENSOR_PIN, INPUT);
  pinMode(AIR_QUALITY_SENSOR_PIN, INPUT);
  pinMode(HUMAN_SENSOR_PIN, INPUT);


  //让继电器的外部灯一直亮
  digitalWrite(EXTERNAL_LIGHT_SWITCH_PIN, LOW);
  //设置门禁关闭
  digitalWrite(ACCESS_CONTROL_SWITCH_PIN, HIGH);
  //设置加湿器关闭
  digitalWrite(HUMIDIFIER_SWITCH_PIN, HIGH);
  //设置空调关闭  
  digitalWrite(AIR_CONDITIONER_SWITCH_PIN, HIGH);
  //设置窗户关闭
  digitalWrite(WINDOW_SWITCH_PIN, HIGH);
  //设置水泵关闭
  digitalWrite(PUMP_SWITCH_PIN, HIGH);
  //设置室内灯关闭
  digitalWrite(INDOOR_LIGHT_SWITCH_PIN, HIGH);
  //设置蜂鸣器关闭
  digitalWrite(BUZZER_SWITCH_PIN, HIGH);


  initWifi(WIFI_SSID, WIFI_PASSWD); 

  Serial.println();

  if(mqttCheckConnect()) 
  {
    mqttClient.subscribe(HUAWEI_TOPIC_PROP_CMD); // 订阅云端下发指令频道
  }

  Serial.println();
  mqttClient.setCallback(callback); // 回调，监听云端下发指令，当ESP8266收到订阅Topic后调用callback函数
}
 
 

void loop()
{
  //每隔5s尝试连接一次云
  if (millis() - lastSend >= 5000)
  {
    if (!mqttClient.connected())
    {
      mqttCheckConnect();
    }
    timing_attribute_reporting();
    lastSend = millis();
  }

  //保持门打开3秒
  if (digitalRead(ACCESS_CONTROL_SWITCH_PIN) == LOW&&digitalRead(FLAME_SENSOR_PIN)==HIGH) {
    //此时已经开始计时
    if(countingFlag) {
      if(millis() - doorOpenTime >= 3000) {
        countingFlag = false;
        doorOpenTime = 0;
        digitalWrite(ACCESS_CONTROL_SWITCH_PIN, HIGH);
      }
    }
    else {
      countingFlag = true;
      doorOpenTime = millis();
    }
  }

  mqttClient.loop();
}
 