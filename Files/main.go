package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "regexp"
    "strconv"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

var client = &http.Client{}

func main() {
    channels := []string{
        "https://t.me/s/MsV2ray",
        "https://t.me/s/DailyV2RY",
        "https://t.me/s/FreakConfig",
        "https://t.me/s/EliV2ray",
        "https://t.me/s/ServerNett",
        "https://t.me/s/v2rayngvpn",
        "https://t.me/s/custom_14",
        "https://t.me/s/vpnaloo",
        "https://t.me/s/v2ray_outlineir",
        "https://t.me/s/vmess_vless_v2rayng",
        "https://t.me/s/PrivateVPNs",
        "https://t.me/s/freeland8",
        "https://t.me/s/vmessiran",
        "https://t.me/s/Outline_Vpn",
        "https://t.me/s/V2rayNG3",
        "https://t.me/s/ShadowsocksM",
        "https://t.me/s/VmessProtocol",
        "https://t.me/s/Easy_Free_VPN",
        "https://t.me/s/V2RAY_VMESS_free",
        "https://t.me/s/free4allVPN",
        "https://t.me/s/configV2rayForFree",
        "https://t.me/s/configV2rayNG",
        "https://t.me/s/XvProxy",
        "https://t.me/s/vpnmasi",
        "https://t.me/s/ViPVpn_v2ray",
        "https://t.me/s/ultrasurf_12",
        "https://t.me/s/Awlix_ir",
        "https://t.me/s/V2ray_Alpha",
        "https://t.me/s/An0nymousTeam",
        "https://t.me/s/CaV2ray",
        "https://t.me/s/proxy_kafee",
        "https://t.me/s/proxy_mtm",
        "https://t.me/s/server_nekobox",
        "https://t.me/s/v2raying",
        "https://t.me/s/V2RAYROZ",
        "https://t.me/s/sinavm",
        "https://t.me/s/customv2ray",
        "https://t.me/s/fastkanfig",
        "https://t.me/s/Hope_Net",
        "https://t.me/s/Free_Vpn_for_All_of_Us",
        "https://t.me/s/NoForcedHeaven",
        "https://t.me/s/freedomofinfor",
        "https://t.me/s/Vpnstable",
        "https://t.me/s/orange_vpns",
        "https://t.me/s/FreeVPNHomesConfigs",
        "https://t.me/s/Ahmedhamoomi_Servers",
        "https://t.me/s/ISVvpn",
        "https://t.me/s/proSSH",
        "https://t.me/s/GrizzlyVPN",
        "https://t.me/s/v2Source",
        "https://t.me/s/GetConfigIR",
        "https://t.me/s/Parsashonam",
        "https://t.me/s/ArV2ray",
        "https://t.me/s/configfa",
        "https://t.me/s/napsternetv",
        "https://t.me/s/Alpha_V2ray_Iran",
        "https://t.me/s/vpn_proxy666",
        "https://t.me/s/xsfilternet",
        "https://t.me/s/Hiidify_V2ray",
        "https://t.me/s/ConfigWireguard",
        "https://t.me/s/WireVpnGuard",
        "https://t.me/s/LonUp_M",
        "https://t.me/s/v2rayfree",
        "https://t.me/s/v2ray_free_conf",
        "https://t.me/s/IP_CF_Config",
        "https://t.me/s/shadowproxy66",
        "https://t.me/s/OutlineReleasedKey",
        "https://t.me/s/prrofile_purple",
        "https://t.me/s/meli_proxyy",
        "https://t.me/s/DirectVPN",
        "https://t.me/s/ViProxys",
        "https://t.me/s/heyatserver",
        "https://t.me/s/vpnfail_vless",
        "https://t.me/s/AzadNet",
        "https://t.me/s/golestan_vpn",
        "https://t.me/s/xroVPN",
        "https://t.me/s/viturey",
        "https://t.me/s/Ln2Ray",
        "https://t.me/s/ZeusVPN_Net",
        "https://t.me/s/v2nodes",
        "https://t.me/s/v2box_free",
        "https://t.me/s/VlessConfig",
        "https://t.me/s/oneclickvpnkeys",
        "https://t.me/s/SafeNet_Server",
        "https://t.me/s/lrnbymaa",
    }

    configs := map[string]string{
        "ss":         "",
        "vmess":      "",
        "trojan":     "",
        "vless":      "",
        "mixed":      "",
        "tuic":       "",
        "hysteria2":  "",
        "wireguard":  "",
    }

    myregex := map[string]string{
        "ss":         `(.{3})ss:\/\/|shadowsocks:\/\/`,
        "vmess":      `vmess:\/\/`,
        "trojan":     `trojan:\/\/`,
        "vless":      `vless:\/\/`,
        "tuic":       `tuic:\/\/|tuic5:\/\/`,
        "hysteria2":  `hy2:\/\/|hysteria2:\/\/`,
        "wireguard":  `wireguard:\/\/`,
    }

    for i := 0; i < len(channels); i++ {
        all_messages := false
        if strings.Contains(channels[i], "{all_messages}") {
            all_messages = true
            channels[i] = strings.Split(channels[i], "{all_messages}")[0]
        }

        req, err := http.NewRequest("GET", channels[i], nil)
        if err != nil {
            log.Fatalf("Error When requesting to: %s Error : %s", channels[i], err)
        }

        resp, err1 := client.Do(req)
        if err1 != nil {
            log.Fatal(err1)
        }
        defer resp.Body.Close()

        doc, err := goquery.NewDocumentFromReader(resp.Body)
        if err != nil {
            log.Fatal(err)
        }

        messages := doc.Find(".tgme_widget_message_wrap").Length()
        link, exist := doc.Find(".tgme_widget_message_wrap .js-widget_message").Last().Attr("data-post")
        if messages < 100 && exist == true {
            number := strings.Split(link, "/")[1]
            fmt.Println(number)

            doc = GetMessages(10, doc, number, channels[i])
        }

        if all_messages {
            fmt.Println(doc.Find(".js-widget_message_wrap").Length())
            doc.Find(".tgme_widget_message_text").Each(func(j int, s *goquery.Selection) {
                message_text := s.Text()
                lines := strings.Split(message_text, "\n")
                for a := 0; a < len(lines); a++ {
                    for _, regex_value := range myregex {
                        re := regexp.MustCompile(regex_value)
                        lines[a] = re.ReplaceAllStringFunc(lines[a], func(match string) string {
                            return "\n" + match
                        })
                    }
                    for proto := range configs {
                        if strings.Contains(lines[a], proto) {
                            configs["mixed"] += "\n" + lines[a] + "\n"
                        }
                    }
                }
            })
        } else {
            doc.Find("code,pre").Each(func(j int, s *goquery.Selection) {
                message_text := s.Text()
                lines := strings.Split(message_text, "\n")
                for a := 0; a < len(lines); a++ {
                    for proto_regex, regex_value := range myregex {
                        re := regexp.MustCompile(regex_value)
                        lines[a] = re.ReplaceAllStringFunc(lines[a], func(match string) string {
                            if proto_regex == "ss" {
                                if match[:3] == "vme" {
                                    return "\n" + match
                                } else if match[:3] == "vle" {
                                    return "\n" + match
                                } else {
                                    return "\n" + match
                                }
                            } else {
                                return "\n" + match
                            }
                        })

                        if len(strings.Split(lines[a], "\n")) > 1 {
                            myconfigs := strings.Split(lines[a], "\n")
                            for i := 0; i < len(myconfigs); i++ {
                                if myconfigs[i] != "" {
                                    re := regexp.MustCompile(regex_value)
                                    myconfigs[i] = strings.ReplaceAll(myconfigs[i], " ", "")
                                    match := re.FindStringSubmatch(myconfigs[i])
                                    if len(match) >= 1 {
                                        if proto_regex == "ss" {
                                            if match[1][:3] == "vme" {
                                                configs["vmess"] += "\n" + myconfigs[i] + "\n"
                                            } else if match[1][:3] == "vle" {
                                                configs["vless"] += "\n" + myconfigs[i] + "\n"
                                            } else {
                                                configs["ss"] += "\n" + myconfigs[i][3:] + "\n"
                                            }
                                        } else {
                                            configs[proto_regex] += "\n" + myconfigs[i] + "\n"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            })
        }
    }

    for proto, configcontent := range configs {
        WriteToFile(RemoveDuplicate(configcontent), proto+"_iran.txt")
    }
}

func WriteToFile(fileContent string, filePath string) {
    if _, err := os.Stat(filePath); err == nil {
        err = ioutil.WriteFile(filePath, []byte{}, 0644)
        if err != nil {
            fmt.Println("Error clearing file:", err)
            return
        }
    } else if os.IsNotExist(err) {
        _, err = os.Create(filePath)
        if err != nil {
            fmt.Println("Error creating file:", err)
            return
        }
    } else {
        fmt.Println("Error checking file:", err)
        return
    }

    err := ioutil.WriteFile(filePath, []byte(fileContent), 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }

    fmt.Println("File written successfully")
}

func load_more(link string) *goquery.Document {
    req, _ := http.NewRequest("GET", link, nil)
    fmt.Println(link)
    resp, _ := client.Do(req)
    doc, _ := goquery.NewDocumentFromReader(resp.Body)
    return doc
}

func GetMessages(length int, doc *goquery.Document, number string, channel string) *goquery.Document {
    x := load_more(channel + "?before=" + number)
    html2, _ := x.Html()
    reader2 := strings.NewReader(html2)
    doc2, _ := goquery.NewDocumentFromReader(reader2)
    doc.Find("body").AppendSelection(doc2.Find("body").Children())
    newDoc := goquery.NewDocumentFromNode(doc.Selection.Nodes[0])
    messages := newDoc.Find(".js-widget_message_wrap").Length()
    if messages > length {
        return newDoc
    } else {
        num, _ := strconv.Atoi(number)
        n := num - 21
        if n > 0 {
            ns := strconv.Itoa(n)
            GetMessages(length, newDoc, ns, channel)
        } else {
            return newDoc
        }
    }
    return newDoc
}

func RemoveDuplicate(config string) string {
    lines := strings.Split(config, "\n")
    uniqueLines := make(map[string]bool)
    for _, line := range lines {
        if len(line) > 0 {
            uniqueLines[line] = true
        }
    }
    uniqueString := strings.Join(getKeys(uniqueLines), "\n")
    return uniqueString
}

func getKeys(m map[string]bool) []string {
    keys := make([]string, len(m))
    i := 0
    for k := range m {
        keys[i] = k
        i++
    }
    return keys
}

func getTimestamp(message string) int64 {
    var data map[string]interface{}
    if err := json.Unmarshal([]byte(message), &data); err != nil {
        return 0
    }
    timestamp, _ := data["date"].(int64)
    return timestamp
}
