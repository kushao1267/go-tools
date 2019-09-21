package email

import (
	"fmt"
	"net/smtp"
	"strings"
)

// 邮件类型
const (
	HtmlType = "html"
	TextType = "text"
)

type Auth struct {
	SMTP     string // just like ip:port, e.g smtp.example.com:25
	Username string
	Password string
	auth     smtp.Auth
}

// NewAuth 新建认证实例
func NewAuth(addr string, username string, password string) *Auth {
	var domain string
	addrStrings := strings.Split(addr, ":")
	if len(addrStrings) > 0 {
		domain = addrStrings[0]
	}
	var auth smtp.Auth
	if password != "" {
		auth = smtp.PlainAuth("", username, password, domain)
	}
	return &Auth{
		SMTP:     addr,
		Username: username,
		Password: password,
		auth:     auth,
	}
}

// SendEmail 发送邮件
func (a *Auth) SendEmail(subject string, from string, to []string, mailType string, message string, tls bool) error {
	var contentType = "text/plain; charset=UTF-8"
	if mailType == HtmlType {
		contentType = "text/html; charset=UTF-8"
	}
	var msg = "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + contentType + "\r\n\r\n" +
		message + "\r\n"
	if tls == true {
		return smtp.SendMail(a.SMTP, a.auth, from, to, []byte(msg))
	}
	return SendMailWithoutTLS(a.SMTP, a.auth, from, to, []byte(msg))
}

// SendMailWithoutTLS 不加密传输的认证方式来发送邮件
func SendMailWithoutTLS(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Hello("localhost"); err != nil {
		return err
	}
	err = c.Auth(a)
	if err != nil {
		return err
	}

	if err = c.Mail(from); err != nil {
		fmt.Printf("mail\n")
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
