package sulat

import (
	"errors"
	"strings"
	"sync"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridConfig collects required settings for SendGrid SMTP server
type SendGridConfig struct {
	SendGridAPIKey, SendGridEndPoint, SendGridHost string
	mu                                             sync.Mutex
}

// SGC short hand for using the 'SendGridConfig' struct methods and initializes
// the default values for the following required configurations
var SGC = SendGridConfig{
	SendGridAPIKey:   "",
	SendGridEndPoint: "",
	SendGridHost:     "",
}

// SetSGC initializes the SendGrid SMTP required configurations
func (s *SendGridConfig) SetSGC(sgc *SendGridConfig) *SendGridConfig {
	s.mu.Lock()
	defer s.mu.Unlock()
	s = sgc
	return s
}

// MailOptions holds standard email options prior to sending an email
func (s *SendMail) MailOptions(sm *SendMail) []byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	s = sm

	// Set the required options prior to sending an email
	from := mail.NewEmail(s.From.Name, s.From.Address)
	to := mail.NewEmail(s.To.Name, s.To.Address)
	content := mail.NewContent("text/html", s.HTMLBody)
	m := mail.NewV3MailInit(from, s.Subject, to, content)

	// Optional configs
	if len(strings.TrimSpace(s.CC.Address)) == 0 {
		cc := mail.NewEmail(s.CC.Name, s.CC.Address)
		m.Personalizations[0].AddCCs(cc)
	}
	if len(strings.TrimSpace(s.BCC.Address)) == 0 {
		bcc := mail.NewEmail(s.BCC.Name, s.BCC.Address)
		m.Personalizations[0].AddBCCs(bcc)
	}
	return mail.GetRequestBody(m)
}

// Send will send the new email
func (s *SendMail) Send(sgc *SendGridConfig) (bool, error) {
	// Check the required SendGrid API information
	if len(strings.TrimSpace(sgc.SendGridAPIKey)) == 0 {
		return false, errors.New("sendgrid api key is required")
	}
	if len(strings.TrimSpace(sgc.SendGridEndPoint)) == 0 {
		return false, errors.New("sendgrid endpoint is required")
	}
	if len(strings.TrimSpace(sgc.SendGridHost)) == 0 {
		return false, errors.New("sendgrid host is required")
	}

	request := sendgrid.GetRequest(sgc.SendGridAPIKey, sgc.SendGridEndPoint, sgc.SendGridHost)
	request.Method = "POST"
	var Body = s.MailOptions(s)
	request.Body = Body

	_, err := sendgrid.API(request)
	if err != nil {
		return false, err
	}
	return true, nil
}