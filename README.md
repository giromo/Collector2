# 🚀 Collector2

[![Go](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-Custom-lightgrey.svg)]()
[![Last Commit](https://img.shields.io/github/last-commit/giromo/Collector2?logo=github)](https://github.com/giromo/Collector2/commits/main)

---

## درباره پروژه

**Collector2** یک ابزار قدرتمند برای جمع‌آوری، تست سرعت و پردازش کانفیگ‌های مختلف (مانند vmess, vless, trojan و...) است که بیشتر برای شبکه‌های خصوصی مجازی (VPN) و کاربردهای تست سرور‌ها استفاده می‌شود. این پروژه با زبان Go توسعه داده شده و اسکریپت‌ها و ابزارهایی برای اتوماسیون فرآیند تست و آنالیز کانفیگ‌ها فراهم می‌کند.

---

## ✨ ویژگی‌ها

- 🚦 تست سرعت اتوماتیک سرورها با ابزار LiteSpeedTest و v2ray-core
- 📦 تبدیل فرمت‌های مختلف کانفیگ و خروجی به صورت Base64 یا متنی
- 📝 گزارش‌گیری و ثبت خروجی تست‌ها به صورت فایل JSON و لاگ
- 🔍 پردازش و فیلترینگ کانفیگ‌ها برای بهبود سرعت و کیفیت شبکه
- 🎯 اسکریپت‌های Bash و Python برای اجرای سریع عملیات

---

## 📂 ساختار پوشه‌ها

- `utils/speedtest/` : اسکریپت‌ها و ابزارهای تست سرعت
- `sub/` : فایل‌های کانفیگ تجمیع‌شده و پردازش‌شده
- `logs/` : گزارش اجرای تست‌ها
- سایر فایل‌ها و ابزارها برای پردازش و تبدیل کانفیگ‌ها

---

## ⚡ نحوه استفاده

1. ابزار را کلون کنید:
   ```bash
   git clone https://github.com/giromo/Collector2.git
   cd Collector2
   ```

2. اجرای اسکریپت تست سرعت:
   ```bash
   bash utils/speedtest/speedtest2.sh
   ```

3. پردازش خروجی با پایتون:
   ```bash
   python3 utils/speedtest/output.py
   ```

4. فایل‌های خروجی را در مسیرهای `logs/` و `utils/speedtest/` بیابید.

---

## 🛠️ پیش‌نیازها

- **Go** (برای توسعه)
- **Python 3** (برای پردازش خروجی)
- **wget**, **unzip**, **bash**

---

## 🧑‍💻 مشارکت

خوشحال می‌شویم اگر باگ، ایده یا پیشنهاد خود را در [Issues](https://github.com/giromo/Collector2/issues) ثبت کنید!

---

## 📜 لایسنس

این پروژه تحت لایسنس سفارشی منتشر شده است.

---

<div align="center">
با Collector2 سرعت و کیفیت سرویس‌های شبکه خود را به سطح جدیدی ببرید! 🚀💡
</div>
