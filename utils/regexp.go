package utils

import (
	"regexp"
	"strings"
)

func MatchEmailFormat(email string) bool {
	reg := regexp.MustCompile(`^[0-9a-z][0-9a-z-_.]{0,35}@([0-9a-z][0-9a-z-]{0,35}[0-9a-z]\.){1,5}[a-z]{2,4}$`)
	return reg.MatchString(email)
}

func CheckName(name string) bool {
	if len(name) < 3 || len(name) > 20*3 {
		return false
	}
	space := regexp.MustCompile(" ")
	special := regexp.MustCompile("[-_+=^a-zA-Z0-9]")
	if space.MatchString(name) {
		return false
	}
	r := special.ReplaceAllString(name, "")
	r = strings.Map(func(c rune) rune {
		if c >= 0x4E00 && c <= 0x9FA5 { // 常用汉字范围
			return -1 // 忽略
		}
		return c
	}, r)
	return r == ""
}

// CheckPasswordStrength 检测密码强度
func CheckPasswordStrength(password string) bool {
	if len(password) < 6 || len(password) > 64*3 {
		return false // 长度大于 6
	}
	A := regexp.MustCompile(`[A-Z]`)
	a := regexp.MustCompile(`[a-z]`)
	figure := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[!@#$%^&*()\-_+=\\|\[\]{}:'",<.>/?]`) // 除掉 ; sql注入常用符号
	if !A.MatchString(password) {
		return false // 必须要有大写字母
	}
	if !a.MatchString(password) {
		return false // 必须要有小写字母
	}
	if !figure.MatchString(password) {
		return false // 必须要有数字
	}
	if !special.MatchString(password) {
		return false // 必须要有特殊字符
	}
	if len(special.ReplaceAll(figure.ReplaceAll(a.ReplaceAll(A.ReplaceAll([]byte(password), []byte("")), []byte("")), []byte("")), []byte(""))) != 0 {
		return false // 不能有其他字符
	}
	return true
}
