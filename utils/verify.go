package utils

import (
	"math/rand"
	"message-board/config"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

const (
	EmailTemplate = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>Message-Board 验证码</title>
		</head>
		<body>
			【Message-Board】：验证码为 {0} 您正在使用邮箱进行验证，请勿泄露哦！
		</body>
		</html>`
	Subject  = "验证码"
	StmpHost = "smtp.qq.com"
	StmpPort = "587"
)

var VerifyMap map[string]string

func VerifyInputCode(email, code string) error {
	if got, ok := VerifyMap[email]; ok && got == code {
		return nil
	}
	if got, ok := VerifyMap[email]; ok && got != code {
		return ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40001,
			Info:       "invalid verify code",
			Detail:     "验证码错误",
		}
	}
	if _, ok := VerifyMap[email]; !ok {
		return ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40001,
			Info:       "invalid verify code",
			Detail:     "验证码过期",
		}
	}
	return ServerInternalError
}

func SendVerifyCode(target, code string) {
	VerifyMap[target] = code

	go func() {
		err := SendEmail(code, target)
		if err != nil {
			LoggerWarning("发送失败")
			return
		}
	}()

	go func() {
		// 两分钟后过期
		<-time.NewTimer(2 * time.Minute).C
		delete(VerifyMap, target)
	}()
}

// SendRandomVerifyCode 发送随机验证码
func SendRandomVerifyCode(target string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sum = 0
	sum += 1 << 14 // 保证为5位数
	for i := 1; i <= 15; i++ {
		sum += r.Intn(2) << i
	}
	var vCode = strconv.Itoa(sum)
	SendVerifyCode(target, vCode)
	return vCode
}

// SendEmail
// 发送邮件
// 返回：
//   - nil						发送成功
//   - ServerInternalError		发送失败
func SendEmail(verifyCode string, email ...string) error {
	for _, addr := range email {
		body := strings.Replace(EmailTemplate, "{0}", verifyCode, 1)
		msg := []byte("To: " + addr + "\r\nFrom: " + config.EmailConfig.EmailAuthSender + "<" + config.EmailConfig.EmailAuthAccount + ">" + "\r\nSubject: " + Subject + "\r\n" + "Content-Type: text/" + "html" + "; charset=UTF-8" + "\r\n\r\n" + body + "\r\n")
		auth := smtp.PlainAuth("", config.EmailConfig.EmailAuthAccount, config.EmailConfig.EmailAuthPassword, StmpHost)
		err := smtp.SendMail(StmpHost+":"+StmpPort, auth, config.EmailConfig.EmailAuthAccount, []string{addr}, msg)
		if err != nil {
			return ServerInternalError
		}
	}
	return nil
}
