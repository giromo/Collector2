import json
import base64
import os
import time
from datetime import datetime

out_json = './utils/speedtest/out.json'
sub_all_base64 = "./sub/sub_merge_base64.txt"
sub_all = "./sub/sub_merge.txt"
Base64_file_base64 = "./Base64.txt"
V2_file = "./V2.txt"
Allconfig_Base = "./Allconfig"
splitted_output = "./sub/splitted/"

def read_json(file):
    if not os.path.isfile(file):
        with open('logs/output_py.log', 'w') as f:
            f.write("Error: out.json not found")
        exit(1)
    with open(file, 'r', encoding='utf-8') as f:
        proxies_all = json.load(f)["nodes"]
        f.close()
    return proxies_all

def output(list, num):
    def arred(x, n): return x * (10 ** n) // 1 / (10 ** n)
    output_list = []

    current_date_time = datetime.now()
    current_month = current_date_time.strftime("%b")
    current_day = current_date_time.strftime("%d")
    updated_hour = current_date_time.strftime("%H")
    updated_minute = current_date_time.strftime("%M")
    final_string = f"{current_month}-{current_day} | {updated_hour}:{updated_minute}"
    final_others_string = f"{current_month}-{current_day}"

    for item in list:
        info = "id: %s | remarks: %s | protocol: %s | ping: %s MS | avg_speed: %s MB | max_speed: %s MB | Link: %s\n" % (
            str(item["id"]), item["remarks"], item["protocol"], str(item["ping"]),
            str(arred(item["avg_speed"] * 0.00000095367432, 3)), str(arred(item["max_speed"] * 0.00000095367432, 3)), item["link"])
        output_list.append(info)

    with open('./LogInfo.txt', 'w') as f1:
        f1.writelines(output_list)
        f1.close()

    modified_output_list = []
    for i, item in enumerate(list):
        link = item["link"].split("#")[0]
        if i == 0:
            config_string = f"#馃寪 亘賴 乇賵夭乇爻丕賳蹖 卮丿賴 丿乇 {final_string} | 賴乇 2 爻丕毓鬲 讴丕賳賮蹖讴 噩丿蹖丿 丿丕乇蹖賲"
        else:
            config_string = f"#馃寪爻乇賵乇 {i} | {final_others_string} | MTSRVRS"
        modified_link = link + config_string
        modified_output_list.append(modified_link)

    content = '\n'.join(modified_output_list)
    content_base64 = base64.b64encode(content.encode('utf-8')).decode('ascii')
    content_base64_part = base64.b64encode('\n'.join(modified_output_list[0:num]).encode('utf-8')).decode('ascii')

    os.makedirs(splitted_output, exist_ok=True)
    vmess_outputs = []
    vless_outputs = []
    trojan_outputs = []
    ssr_outputs = []
    ss_outputs = []
    tuic_outputs = []
    hysteria2_outputs = []
    wireguard_outputs = []

    for output in modified_output_list:
        if str(output).startswith("vmess://"):
            vmess_outputs.append(output)
        elif str(output).startswith("vless://"):
            vless_outputs.append(output)
        elif str(output).startswith("trojan://"):
            trojan_outputs.append(output)
        elif str(output).startswith("ssr://"):
            ssr_outputs.append(output)
        elif str(output).startswith(("ss://", "shadowsocks://")):
            ss_outputs.append(output)
        elif str(output).startswith(("tuic://", "tuic5://")):
            tuic_outputs.append(output)
        elif str(output).startswith(("hy2://", "hysteria2://")):
            hysteria2_outputs.append(output)
        elif str(output).startswith("wireguard://"):
            wireguard_outputs.append(output)

    with open(os.path.join(splitted_output, "vmess.txt"), 'w') as f:
        f.write("\n".join(vmess_outputs))
        f.close()

    with open(os.path.join(splitted_output, "vless.txt"), 'w') as f:
        f.write("\n".join(vless_outputs))
        f.close()

    with open(os.path.join(splitted_output, "trojan.txt"), 'w') as f:
        f.write("\n".join(trojan_outputs))
        f.close()

    with open(os.path.join(splitted_output, "ssr.txt"), 'w') as f:
        f.write("\n".join(ssr_outputs))
        f.close()

    with open(os.path.join(splitted_output, "ss.txt"), 'w') as f:
        f.write("\n".join(ss_outputs))
        f.close()

    with open(os.path.join(splitted_output, "tuic.txt"), 'w') as f:
        f.write("\n".join(tuic_outputs))
        f.close()

    with open(os.path.join(splitted_output, "hysteria2.txt"), 'w') as f:
        f.write("\n".join(hysteria2_outputs))
        f.close()

    with open(os.path.join(splitted_output, "wireguard.txt"), 'w') as f:
        f.write("\n".join(wireguard_outputs))
        f.close()

    with open(sub_all_base64, 'w+', encoding='utf-8') as f:
        f.write(content_base64)
        f.close()

    with open(Base64_file_base64, 'w+', encoding='utf-8') as f:
        f.write(content_base64_part)
        f.close()

    with open(sub_all, 'w') as f:
        f.write(content)
        f.close()

    with open(Allconfig_Base, 'w') as f:
        f.write(content)
        f.close()

    with open(V2_file, 'w') as f:
        f.write('\n'.join(modified_output_list[0:num]))
        f.close()

    return content

if __name__ == '__main__':
    num = 200
    value = read_json(out_json)
    output(value, value.__len__() if value.__len__() <= num else num)
