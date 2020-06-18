package sulat

import (
	"crypto/tls"
	"strings"
	"sync"

	"github.com/itrepablik/gomail"
)

// SMTPConfig is the collection of classic SMTP required information
type SMTPConfig struct {
	Host               string
	Port               int
	UserName, Password string
}

// MailClassicHeader holds a common email information header configurations
type MailClassicHeader struct {
	From, Subject, HTMLBody string
	To, CC, BCC             []string
	mu                      sync.Mutex
}

// SendEmailSMTP uses the classic SMTP local hosted server
func SendEmailSMTP(s *MailClassicHeader, emf *EmailHTMLFormat, sc *SMTPConfig) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var err error
	m := gomail.NewMessage()
	m.SetHeader("From", s.From)
	m.SetHeader("To", s.To...)

	// Append the following extra email addresses if existed
	if len(s.CC) > 0 {
		m.SetHeader("Cc", s.CC...)
	}
	if len(s.BCC) > 0 {
		m.SetHeader("Bcc", s.BCC...)
	}
	m.SetHeader("Subject", s.Subject)

	emailContent := ""
	if emf.IsFullHTML {
		emailContent, err = NewFullHTMLContent(emf.FullHTMLTemplate)
	} else {
		emailContent, err = NewEmailContent(emf.HTMLHeader, emf.HTMLBody, emf.HTMLFooter)
	}
	if err != nil {
		return false, err
	}
	m.SetBody("text/html", emailContent)

	// Check if SMTP with authentication or not
	if len(strings.TrimSpace(sc.UserName)) > 0 && len(strings.TrimSpace(sc.Password)) > 0 {
		e := gomail.NewDialer(sc.Host, sc.Port, sc.UserName, sc.Password)
		e.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		if err := e.DialAndSend(m); err != nil {
			return false, err
		}
	} else {
		d := gomail.Dialer{Host: sc.Host, Port: sc.Port} // Without auth
		if err := d.DialAndSend(m); err != nil {
			return false, err
		}
	}
	return true, nil
}
