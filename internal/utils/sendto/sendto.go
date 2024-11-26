package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/shinkaym/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPUsername = "khoapham1405@gmail.com"
	SMTPPassword = "jnlo nehe kstv tnvr"
)

func BuildMessage(mail Mail) string {
	msg := fmt.Sprintf("From %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To %s\r\n", strings.Join(mail.To, ","))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "OTP Service"},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	}

	messageEmail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Email send failed with net/smtp:", zap.Error(err))
		return err
	}

	return nil
}

func SendTemplateEmail(to []string, from string, nameTemplate string, dataTemplate map[string]interface{}) error {
	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		return err
	}

	return send(to, from, htmlBody)
}

func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates-email/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "OTP Service"},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	emailHeaders := []string{
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=UTF-8",
		"Content-Transfer-Encoding: quoted-printable",
	}

	fullMessage := append(emailHeaders, BuildMessage(contentEmail))

	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, to, []byte(strings.Join(fullMessage, "\r\n")))
	if err != nil {
		global.Logger.Error("Email send failed with net/smtp:", zap.Error(err))
		return err
	}

	return nil
}
