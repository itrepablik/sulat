package sulat

import "sync"

// Email holds email name and address info
type Email struct {
	Name, Address string
}

// EmailHTMLFormat email HTML content structure
type EmailHTMLFormat struct {
	EmailHTMLHeader, EmailHTMLFooter string
}

// SendMail holds a common email information
type SendMail struct {
	Subject, HTMLBody string
	From, To, CC, BCC Email
	HTMLHeaderFooter  EmailHTMLFormat
	mu                sync.Mutex
}
