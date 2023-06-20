package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/smtp"
	"os"
	"path/filepath"
)

// SendEmail sends an email with an image attachment
func SendEmail(templatePath string, data interface{}, to string, subject string) error {
	// Get email credentials from environment variables
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Create a new template
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Create a new buffer
	buf := new(bytes.Buffer)

	// Execute the template, and pass the data
	err = t.Execute(buf, data)
	if err != nil {
		return err
	}

	// Create a new message
	message := new(bytes.Buffer)
	message.WriteString(fmt.Sprintf("From: %s\r\n", email))
	message.WriteString(fmt.Sprintf("To: %s\r\n", to))
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	message.WriteString("MIME-Version: 1.0\r\n")
	message.WriteString("Content-Type: multipart/mixed; boundary=boundary1234\r\n")
	message.WriteString("\r\n")

	// Add the text part (email body)
	message.WriteString("--boundary1234\r\n")
	message.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	message.WriteString("Content-Transfer-Encoding: 7bit\r\n")
	message.WriteString("\r\n")
	message.WriteString(buf.String())
	message.WriteString("\r\n")

	// Add the image attachment
	imagePath := "./../avatar.png" // Replace with the actual path to your image
	attachment, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer attachment.Close()

	_, attachmentFileName := filepath.Split(imagePath)
	message.WriteString("--boundary1234\r\n")
	message.WriteString("Content-Type: image/jpeg\r\n")
	message.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", attachmentFileName))
	message.WriteString("Content-Transfer-Encoding: base64\r\n")
	message.WriteString("\r\n")
	encoder := base64.NewEncoder(base64.StdEncoding, message)
	io.Copy(encoder, attachment)
	encoder.Close()
	message.WriteString("\r\n")

	message.WriteString("--boundary1234--\r\n")

	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", email, password, smtpHost)

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{to}, message.Bytes())
	if err != nil {
		return err
	}

	return nil
}



























// package utils

// import (
// 	"bytes"
// 	"fmt"
// 	"html/template"
// 	"net/smtp"
// 	"os"
// )

// // SendEmail sends an email
// func SendEmail(templatePath string, data interface{}, to string, subject string) error {
// 	// Get email credentials from environment variables
// 	email := os.Getenv("EMAIL")
// 	password := os.Getenv("PASSWORD")
// 	smtpHost := os.Getenv("SMTP_HOST")
// 	smtpPort := os.Getenv("SMTP_PORT")

// 	// Create a new template
// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		return err
// 	}

// 	// Create a new buffer
// 	buf := new(bytes.Buffer)

// 	// Execute the template, and pass the data
// 	err = t.Execute(buf, data)
// 	if err != nil {
// 		return err
// 	}

// 	// Create a new message
// 	message := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s", email, to, subject, buf.String()))

// 	// Authenticate with the SMTP server
// 	auth := smtp.PlainAuth("", email, password, smtpHost) // the "" is the identity, it can be anything, it's not used here

// 	// Send the email
// 	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{to}, message)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
