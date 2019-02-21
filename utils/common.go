package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"time"
	"wyimserver/common"
)

func SendRequestToWy(url, method string, byteArr []byte) ([]byte, int, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(byteArr))
	if err != nil {
		logs.Error("请求构造消息体错误:%v", err)
		return nil, 0, err
	}
	var nonce = GenerateRandomString(20)
	var nowStr = fmt.Sprintf("%d", int(time.Now().Unix()))

	req.Header.Set("AppKey", common.APPKEY)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("CurTime", nowStr)
	req.Header.Set("CheckSum", CalculationCheckSum(nonce, nowStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		logs.Error("请求错误:%v", err)
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, resp.StatusCode, nil
}

func CalculationCheckSum(nonce, curTime string) string {
	var str = fmt.Sprintf("%s%s%s", common.APPSECRET, nonce, curTime)
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func GenerateRandomString(length int) string {
	// 36进制，正好是 [0-9]+[a-z] 的长度
	const base = 36

	// 用于每次读一个新的随机数出来
	size := big.NewInt(base)

	n := make([]byte, length)
	for i, _ := range n {
		c, _ := rand.Int(rand.Reader, size)
		// 把小于等于36的数字 按照36进制恰好可以转换成数字或者a-z的字符
		str := strconv.FormatInt(c.Int64(), base)
		n[i] = str[0]
	}
	return string(n)
}
