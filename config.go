package sulat

import (
	"errors"
	"strings"
	"sync"
)

// Email holds email name and address info
type Email struct {
	Name, Address string
}

// SendMail holds a common email information
type SendMail struct {
	Subject, HTMLBody string
	From, To, CC, BCC Email
	mu                sync.Mutex
}

// NewEmailContent combines the new HTML email content from the header, body and footer sections
func NewEmailContent(headerHTML, bodyHTML, footerHTML string) (string, error) {
	// Check the required HTML entities
	if len(strings.TrimSpace(headerHTML)) == 0 {
		return "", errors.New("html header is required")
	}
	if len(strings.TrimSpace(bodyHTML)) == 0 {
		return "", errors.New("html body content is required")
	}
	if len(strings.TrimSpace(footerHTML)) == 0 {
		return "", errors.New("html footer is required")
	}
	content := headerHTML + bodyHTML + footerHTML
	return content, nil
}

// NewEmail holds an email address with name and email address.
func NewEmail(name string, address string) Email {
	return Email{
		Name:    name,
		Address: address,
	}
}
