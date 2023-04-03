package services

import (
	"WEB/internal/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

func ECPay(r models.Payment) string {

	// 建立 ECPay 付款請求資訊
	req := make(map[string]string)
	req["MerchantTradeNo"] = fmt.Sprintf("%d", time.Now().UnixNano())
	req["MerchantTradeDate"] = time.Now().Format("2006/01/02 15:04:05")
	req["TotalAmount"] = fmt.Sprint(r.Amount)
	req["TradeDesc"] = r.Desc
	req["ItemName"] = r.Desc
	req["ExpireDate"] = "3"
	req["Email"] = r.Email
	req["ReturnURL"] = "https://e479-220-129-179-212.jp.ngrok.io/callback/ecpay"
	req["OrderResultURL"] = "https://e479-220-129-179-212.jp.ngrok.io/pay-success"
	req["ClientRedirectURL"] = "https://e479-220-129-179-212.jp.ngrok.io/make-reservation"
	req["MerchantID"] = "2000132"
	req["PaymentType"] = "aio"
	req["ChoosePayment"] = "ALL"
	req["EncryptType"] = "1"

	// 依照規定進行參數排序
	keys := make([]string, 0, len(req))
	for k := range req {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// 將排序後的參數串接成字串
	var buf strings.Builder
	for _, k := range keys {
		buf.WriteString(k + "=" + req[k] + "&")
	}

	s := fmt.Sprintf("HashKey=%s&%sHashIV=%s", "5294y06JbISpM5x9", buf.String(), "v77hoKGq4kWxNNIS")

	s = url.QueryEscape(s)

	checkMacValue := EncryptSHA256(strings.ToLower(s))

	fmt.Println(checkMacValue)

	// 將加密後的結果加入參數中
	req["CheckMacValue"] = strings.ToUpper(checkMacValue)

	formHtml := fmt.Sprint("<form id='payForm'  action='https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5' method='POST'>")

	// 將付款請求資訊轉換為表單格式
	form := url.Values{}
	for k, v := range req {
		form.Add(k, v)
		formHtml += fmt.Sprintf("<input type='hidden' name='%s' value='%s' />", k, v)
	}

	formHtml += "</form><script>document.getElementById(\"payForm\").submit();</script>"
	return formHtml
}

func EncryptSHA256(source string) string {
	fmt.Println(source)
	utf8Source := []byte(strings.TrimSpace(source))
	hash := sha256.New()
	hash.Write(utf8Source)
	md := hash.Sum(nil)
	macValue := hex.EncodeToString(md)

	return strings.ToUpper(macValue)
}
