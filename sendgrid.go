package sulat

import (
	"errors"
	"strings"
	"sync"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SGC collects required configurations for SendGrid SMTP server
type SGC struct {
	SendGridAPIKey, SendGridEndPoint, SendGridHost string
	mu                                             sync.Mutex
}

// SM short hand for using the 'SendMail' struct methods
var SM = SendMail{}

// Options holds standard email options prior to sending an email
func (s *SendMail) Options(sm *SendMail) []byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	s = sm

	// Set the required options prior to sending an email
	from := mail.NewEmail(s.From.Name, s.From.Address)
	to := mail.NewEmail(s.To.Name, s.To.Address)
	content := mail.NewContent("text/html", s.HTMLBody)
	m := mail.NewV3MailInit(from, s.Subject, to, content)

	// Optional configs
	if len(strings.TrimSpace(s.CC.Address)) > 0 {
		cc := mail.NewEmail(s.CC.Name, s.CC.Address)
		m.Personalizations[0].AddCCs(cc)
	}
	if len(strings.TrimSpace(s.BCC.Address)) > 0 {
		bcc := mail.NewEmail(s.BCC.Name, s.BCC.Address)
		m.Personalizations[0].AddBCCs(bcc)
	}
	return mail.GetRequestBody(m)
}

// Send will send the new email
func (s *SendMail) Send(byteMailOpt []byte, sgc *SGC) (bool, error) {
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
	var Body = byteMailOpt
	request.Body = Body

	_, err := sendgrid.API(request)
	if err != nil {
		return false, err
	}
	return true, nil
}
