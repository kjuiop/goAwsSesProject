package main

import (
	"gopkg.in/gomail.v2"
	"sync"
)

type SesInfo struct {
	Sender     string // 보내는 메일 계정, Amazon SES 인증을 받아야 함
	SenderName string
	SmtpUser   string // Amazon SES 아이디
	SmtpPass   string // Amazon SES 비밀번호
	ConfigSet  string // 설정파일 이름
	Host       string
	Port       int
	CharSet    string
}

var sesInfo SesInfo
var once sync.Once
var e = SesInfo{}

func init() {
	once.Do(func() {
		e = SesInfo{
			Sender:     "master@example.com",
			SenderName: "Display name",
			SmtpUser:   "AWS SES ID",
			SmtpPass:   "AWS SES Password",
			Host:       "email-smtp.us-west-2.amazonaws.com",
			Port:       587,
		}
	})
}

func Send(subject, body string, to []string) error {
	m := gomail.NewMessage()
	// HTML 형식의 메시지
	m.SetBody("text/html", body)
	// 메시지 헤더 구성
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(e.Sender, e.SenderName)},
		"To":      to,
		"Subject": {subject},
		// 다음 두 항목은 Optional 항목으로 사용하지 않을 경우 제거
		"X-SES-CONFIGURATION-SET": {e.ConfigSet},
		"X-SES-MESSAGE-TAGS":      {"genre=test,genre2=test2"},
	})

	// 메일 발송
	d := gomail.NewDialer(e.Host, e.Port, e.SmtpUser, e.SmtpPass)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
