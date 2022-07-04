<?php

// 进制转换
function baseconvert($str, $frombase = 10, $tobase = 36)
{
    $str = trim($str);
    if (intval($frombase) != 10) {
        $len = strlen($str);
        $q = 0;
        for ($i = 0; $i < $len; $i++) {
            $r = base_convert($str[$i], $frombase, 10);
            $q = bcadd(bcmul($q, $frombase), $r);
        }
    } else $q = $str;

    if (intval($tobase) != 10) {
        $s = '';
        while (bccomp($q, '0', 0) > 0) {
            $r = intval(bcmod($q, $tobase));
            $s = base_convert($r, 10, $tobase) . $s;
            $q = bcdiv($q, $tobase, 0);
        }
    } else $s = $q;

    return $s;
}

// 字符串转36进制字符
function s2h($str)
{
    return baseconvert(bin2hex($str), 16, 36);
}

// 36进制字符转字符串
function h2s($hex)
{
    return hex2bin(baseconvert($hex, 36, 16));
}

// 字符串转10进制
function s2n($str)
{
    return h2n(s2h($str));
}

// 10进制转字符串
function n2s($num)
{
    return h2s(n2h($num));
}

// 36进制转10进制
function h2n($hex)
{
    return baseconvert($hex, 36, 10);
}

// 10进制转36进制
function n2h($num)
{
    return baseconvert($num, 10, 36);
}

##############  rsa ############

// 公钥加密
function encrypt_rsa($m, $e, $n)
{
    $m = s2n($m);
    return n2h(bcpowmod($m, $e, $n)); // 私钥加密是  d,n
}

// 私钥解密
function decrypt_rsa($c, $d, $n)
{
    $c = h2n($c);
    return n2s(bcpowmod($c, $d, $n)); // 公钥解密是 e,n
}

###### 示例 #####

$n = h2n('92hjwbx7cc6i2cur0p27es6v5s9d0y21');
$e = h2n('1ekh');
$d = h2n('3foj2s1eglg0uv17bzksnzwccwhp1dnh');
$c = encrypt_rsa('xdrsec_wlza', $e, $n);
$m = decrypt_rsa($c, $d, $n);
echo '密文：' . $c;
echo '明文：' . $m;