<?php
function decode_vmess($vmess_config)
{
    $vmess_data = substr($vmess_config, 8); // remove "vmess://"
    $decoded_data = json_decode(base64_decode($vmess_data), true);
    return $decoded_data;
}

function encode_vmess($config)
{
    $encoded_data = base64_encode(json_encode($config));
    $vmess_config = "vmess://" . $encoded_data;
    return $vmess_config;
}

function remove_duplicate_vmess($input)
{
    $array = explode("\n", $input);
    $result = [];
    foreach ($array as $item) {
        $parts = decode_vmess($item);
        if ($parts !== NULL) {
            $part_ps = $parts["ps"];
            unset($parts["ps"]);
            if (count($parts) >= 3) {
                ksort($parts);
                $part_serialize = base64_encode(serialize($parts));
                $result[$part_serialize][] = $part_ps ?? "";
            }
        }
    }
    $finalResult = [];
    foreach ($result as $serial => $ps) {
        $partAfterHash = $ps[0] ?? "";
        $part_serialize = unserialize(base64_decode($serial));
        $part_serialize["ps"] = $partAfterHash;
        $finalResult[] = encode_vmess($part_serialize);
    }
    $output = "";
    foreach ($finalResult as $config) {
        $output .= $output == "" ? $config : "\n" . $config;
    }
    return $output;
}

// Fetch VMess configurations from URL
$vmess_url = "https://raw.githubusercontent.com/Argh94/V2RayAutoConfig/refs/heads/main/configs/Vmess.txt";
$ch = curl_init($vmess_url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_FAILONERROR, true);
$vmess_data = curl_exec($ch);
if ($vmess_data === false) {
    file_put_contents("logs/php_error.log", "cURL Error: " . curl_error($ch));
    curl_close($ch);
    exit(1);
}
curl_close($ch);
