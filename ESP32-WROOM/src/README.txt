音频开发板:
LED4:IO19
LED5:IO22

      drawRGBBitmap(int16_t x, int16_t y, const uint16_t bitmap[], int16_t w,
                    int16_t h),
      drawRGBBitmap(int16_t x, int16_t y, uint16_t *bitmap, int16_t w,
                    int16_t h),
tft.drawRGBBitmap(0,0,cat,100,100);
  delay(1000);

  u8g2.firstPage();
    do{
    car_distance = float(hcsr04.distanceInMillimeters())/10;
    u8g2.setFont(u8g2_font_wqy14_t_gb2312a);
    u8g2.setCursor(0,16-1);
    u8g2.print("距离:");
    u8g2.setCursor(48,16-1);
    u8g2.print(car_distance);
    u8g2.drawStr(88,16-1,"cm");

    u8g2.setCursor(0,16*2-1);
    u8g2.print("左轮速度:");
    u8g2.setCursor(72,16*2-1);
    u8g2.print(left_speed);
    u8g2.drawStr(107,16*2-1,"rad");

    u8g2.setCursor(0,16*3-1);
    u8g2.print("右轮速度:");
    u8g2.setCursor(72,16*3-1);
    u8g2.print(right_speed);
    u8g2.drawStr(107,16*3-1,"rad");

    u8g2.setCursor(0,16*4-1);
    u8g2.print("电源电压:");
    u8g2.setCursor(72,16*4-1);
    u8g2.print((((analogRead(POWER_ADC_PIN)*3.3)/4095)*910)/220);
    u8g2.drawStr(104,16*4-1,"V");
    //u8g2.drawCircle(64-((64*MPU6050_NUM.angleX)/90),32-((32*MPU6050_NUM.angleY)/90),3);
  }while (u8g2.nextPage());