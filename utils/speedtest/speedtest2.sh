#!/bin/bash
mkdir -p utils/speedtest logs

wget -O utils/speedtest/lite-linux-amd64.gz https://github.com/xxf098/LiteSpeedTest/releases/download/v0.14.1/lite-linux-amd64-v0.14.1.gz
gzip -d utils/speedtest/lite-linux-amd64.gz
wget -O utils/speedtest/lite_config.json https://raw.githubusercontent.com/giromo/Collector2/main/utils/speedtest/lite_config.json

chmod +x utils/speedtest/lite-linux-amd64
timeout 300 utils/speedtest/lite-linux-amd64 --config utils/speedtest/lite_config.json --test ./bulk/merge2.txt > utils/speedtest/speedtest_output.log 2>&1

# اضافه کردن تست با v2ray-core برای vmess و vless
wget -O utils/speedtest/v2ray-linux-64.zip https://github.com/v2fly/v2ray-core/releases/latest/download/v2ray-linux-64.zip
unzip -o utils/speedtest/v2ray-linux-64.zip -d utils/speedtest/v2ray
chmod +x utils/speedtest/v2ray/v2ray
utils/speedtest/v2ray/v2ray test -c ./bulk/merge3.txt >> utils/speedtest/speedtest_output.log 2>&1

[ -f out.json ] && mv out.json utils/speedtest/out.json
