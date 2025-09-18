#!/bin/bash
# ایجاد دایرکتوری‌های لازم
mkdir -p utils/speedtest logs

wget -O utils/speedtest/lite-linux-amd64.gz https://github.com/xxf098/LiteSpeedTest/releases/download/v0.14.1/lite-linux-amd64-v0.14.1.gz
gzip -d utils/speedtest/lite-linux-amd64.gz
wget -O utils/speedtest/lite_config.json https://raw.githubusercontent.com/giromo/Collector2/main/utils/speedtest/lite_config.json

# اجرای LiteSpeedTest با محدودیت زمانی
chmod +x utils/speedtest/lite-linux-amd64
timeout 300 utils/speedtest/lite-linux-amd64 --config utils/speedtest/lite_config.json --test ./bulk/merge3.txt > utils/speedtest/speedtest_output.log 2>&1

# انتقال out.json (در صورت تولید در ریشه)
[ -f out.json ] && mv out.json utils/speedtest/out.json
