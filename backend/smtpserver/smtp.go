package smptserver

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

// Object for SMPT Server
type Server struct {
	Email    string // Email address
	Host     string // SMTP host
	Password string // Email password
	Port     string // SMTP port
}

// Object for SMTP Auth
type Sender struct {
	auth smtp.Auth // SMTP Auth
}

// Object for Email Message
type Message struct {
	To          []string          // Send email to
	Subject     string            // Email subject
	Body        string            // Email Body
	Attachments map[string][]byte // Excel file to attach
}

// New returns a SMTP Auth object
func New(svr *Server) *Sender {
	return &Sender{smtp.PlainAuth("", svr.Email, svr.Password, svr.Host)}
}

// Send sends the email message using SMPT Server
func (s *Sender) Send(m *Message, svr *Server) error {
	return smtp.SendMail(fmt.Sprintf("%s:%s", svr.Host, svr.Port), s.auth, svr.Email, m.To, m.ToBytes())
}

// NewMessage returns an email message object
func NewMessage(s, b string) *Message {
	return &Message{Subject: s, Body: b, Attachments: make(map[string][]byte)}
}

// AttachFile attachs the excel file into the email request
func (m *Message) AttachFile(src string) error {
	// Read file
	b, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	// Attach file
	m.Attachments["report.xlsx"] = b
	return nil
}

// ToBytes converts the request in bytest
func (m *Message) ToBytes() []byte {
	// Create buffer and write message
	buf := bytes.NewBuffer(nil)
	// Write subject of email
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	// Write send to of email
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	// Look for attachments
	withAttachments := len(m.Attachments) > 0
	// Write headers
	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	}
	// Write body of request and attach file
	buf.WriteString(m.Body)
	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}
