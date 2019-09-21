## Quick Start

### 发送邮件

邮件发信人收信人以及邮箱配置，均在email.SendEmailExample()中, 请输入你进行配置后再运行example.
```go
import (
	"fmt"
	"github.com/kushao1267/go-tools/email"
)

func SendEmailExample() {
	mailUsername := "foo@bar.com"
	mailPassword := "password"
	mailAddr := "mail.example.com:smtp" // the mailAddr must include a port
	mailFrom := mailUsername
	mailTo := []string{"jianliu001922@gmail.com"}

	mailSender := email.NewAuth(mailAddr, mailUsername, mailPassword)

	subject := "邮件主题"
	content := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>%s</title>
</head>
<body>
<h1>Hello World!</h1>
</body>
</html>`, subject)

	if err := mailSender.SendEmail(subject, mailFrom, mailTo, email.HtmlType, content, true); err != nil {
		fmt.Println("发送邮件失败.")
	} else {
		fmt.Println("发送邮件成功.")
	}

}
```

输出结果:
```
发送邮件成功.
```
