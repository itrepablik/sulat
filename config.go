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

// EmailHTMLFormat is an HTML email content format
type EmailHTMLFormat struct {
	HTMLHeader, HTMLBody, HTMLFooter string
	FullHTMLTemplate                 string
	IsFullHTML                       bool
	mu                               sync.Mutex
}

// HTMLEmailHeader is the HTML skeletal framework head section of the standard HTML structure that serves as an email content
// This is just a default value for the HTML header, you can always override this setting before sending out an email
var HTMLEmailHeader = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html dir="ltr" xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta name="viewport" content="width=device-width" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <link rel="icon" type="image/ico" sizes="16x16" href="https://itrepablik.com/static/assets/images/favicon.ico">
    <title>Email Notifications</title>
</head>
<body style="margin:0px; background: #f8f8f8; ">
    <div width="100%" style="background: #f8f8f8; padding: 0px 0px; font-family:arial; line-height:28px; height:100%;  width: 100%; color: #514d6a;">
        <div style="max-width: 700px; padding:50px 0;  margin: 0px auto; font-size: 14px">
            <table border="0" cellpadding="0" cellspacing="0" style="width: 100%; margin-bottom: 20px">
                <tbody>
                    <tr>
                        <td style="vertical-align: top; padding-bottom:30px;" align="center">
                            <a href="https://itrepablik.com" target="_blank">
                                <img src="https://itrepablik.com/static/assets/images/ITRepablik_top_logo.png" style="width:230px; height:auto;" alt="xtreme admin" style="border:none">
                            </a>
                        </td>
                    </tr>
                </tbody>
            </table>`

// HTMLEmailFooter is the HTML skeletal framework footer section of the standard HTML structure that serves as an email content
// This is just a default value for the HTML footer, you can always override this setting before sending out an email
var HTMLEmailFooter = `<div style="text-align: center; font-size: 12px; color: #b2b2b5; margin-top: 20px">
<p> Powered by ITRepablik.com
	<br>
	<a href="javascript: void(0);" style="color: #b2b2b5; text-decoration: underline;">Unsubscribe</a>
</p>
</div>
</div>
</div>
</body>
</html>`

// HTMLEmailBody is your custom email contents, this is just an example of a password reset auto email from your Go's project
// This is just a default value for the HTML email content, you can always override this setting before sending out an email
var HTMLEmailBody = `<div style="padding: 40px; background: #fff;">
<table border="0" cellpadding="0" cellspacing="0" style="width: 100%;">
	<tbody>
		<tr>
			<td style="border-bottom:1px solid #f6f6f6;">
				<h1 style="font-size:14px; font-family:arial; margin:0px; font-weight:bold;">Hi UserName,</h1>
			</td>
		</tr>
		<tr>
			<td style="padding:10px 0 30px 0;">
				<p>A request to reset your password has been made. If you did not make this request, simply ignore this email. If you did make this request, please reset your password:</p>
				<center>
				<a href="#" style="display: inline-block; padding: 11px 30px; margin: 20px 0px 30px; font-size: 15px; color: #fff; background: #4fc3f7; border-radius: 60px; text-decoration:none;">Reset Password</a>
				</center>
				<b>- Thanks (ITRepablik.com Team)</b>
			</td>
		</tr>
		<tr>
			<td style="border-top:1px solid #f6f6f6; padding-top:20px; color:#777">
				If the button above does not work, try copying and pasting the URL into your browser.<br/>
				<a href="#">https://itrepablik.com/activate/1234567890</a><br/>
				If you continue to have problems, please feel free to contact us at <a href="mailto:support@itrepablik.com">support@itrepablik.com</a>
			</td>
		</tr>
	</tbody>
</table>
</div>`

// FullHTMLContent use this when you pref
var FullHTMLContent = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html dir="ltr" xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta name="viewport" content="width=device-width" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <link rel="icon" type="image/ico" sizes="16x16" href="https://itrepablik.com/static/assets/images/favicon.ico">
    <title>Email Notifications</title>
</head>
<body style="margin:0px; background: #f8f8f8; ">
    <div width="100%" style="background: #f8f8f8; padding: 0px 0px; font-family:arial; line-height:28px; height:100%;  width: 100%; color: #514d6a;">
        <div style="max-width: 700px; padding:50px 0;  margin: 0px auto; font-size: 14px">
            <table border="0" cellpadding="0" cellspacing="0" style="width: 100%; margin-bottom: 20px">
                <tbody>
                    <tr>
                        <td style="vertical-align: top; padding-bottom:30px;" align="center">
                            <a href="https://itrepablik.com" target="_blank">
                                <img src="https://itrepablik.com/static/assets/images/ITRepablik_top_logo.png" style="width:230px; height:auto;" alt="xtreme admin" style="border:none">
                            </a>
                        </td>
                    </tr>
                </tbody>
            </table>

            <div style="padding: 40px; background: #fff;">
                <table border="0" cellpadding="0" cellspacing="0" style="width: 100%;">
                    <tbody>
                        <tr>
                            <td style="border-bottom:1px solid #f6f6f6;">
                                <h1 style="font-size:14px; font-family:arial; margin:0px; font-weight:bold;">Hi UserName,</h1>
                            </td>
                        </tr>
                        <tr>
                            <td style="padding:10px 0 30px 0;">
                                <p>A request to reset your password has been made. If you did not make this request, simply ignore this email. If you did make this request, please reset your password:</p>
                                <center>
                                <a href="#" style="display: inline-block; padding: 11px 30px; margin: 20px 0px 30px; font-size: 15px; color: #fff; background: #4fc3f7; border-radius: 60px; text-decoration:none;">Reset Password</a>
                                </center>
                                <b>- Thanks (ITRepablik.com Team)</b>
                            </td>
                        </tr>
                        <tr>
                            <td style="border-top:1px solid #f6f6f6; padding-top:20px; color:#777">
                                If the button above does not work, try copying and pasting the URL into your browser.<br/>
                                <a href="#">https://itrepablik.com/activate/1234567890</a><br/>
                                If you continue to have problems, please feel free to contact us at <a href="mailto:support@itrepablik.com">support@itrepablik.com</a>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>

            <div style="text-align: center; font-size: 12px; color: #b2b2b5; margin-top: 20px">
                <p> Powered by ITRepablik.com
                    <br>
                    <a href="javascript: void(0);" style="color: #b2b2b5; text-decoration: underline;">Unsubscribe</a>
                </p>
            </div>
        </div>
    </div>
</body>
</html>`

// SGCon initialize this variable globally for sulat.SendGridConfig{} struct
var SGCon = SGC{}

func init() {
	// Initialize the 'SendGrid' API
	SGCon = SGC{
		SendGridAPIKey:   "YOUR_SEND_GRID_API_KEY",
		SendGridEndPoint: "/v3/mail/send",
		SendGridHost:     "https://api.sendgrid.com",
	}
}

// SetHTML email content in HTML format
func SetHTML(e *EmailHTMLFormat) (*EmailHTMLFormat, error) {
	if e.IsFullHTML {
		if len(strings.TrimSpace(e.FullHTMLTemplate)) == 0 {
			return &EmailHTMLFormat{}, errors.New("full html content is required")
		}
		return &EmailHTMLFormat{
			HTMLHeader:       "",
			HTMLBody:         "",
			HTMLFooter:       "",
			FullHTMLTemplate: e.FullHTMLTemplate,
			IsFullHTML:       true,
		}, nil
	}
	if len(strings.TrimSpace(e.HTMLHeader)) == 0 {
		e.HTMLHeader = HTMLEmailHeader
	}
	if len(strings.TrimSpace(e.HTMLBody)) == 0 {
		e.HTMLBody = HTMLEmailBody
	}
	if len(strings.TrimSpace(e.HTMLFooter)) == 0 {
		e.HTMLFooter = HTMLEmailFooter
	}
	return &EmailHTMLFormat{
		HTMLHeader:       e.HTMLHeader,
		HTMLBody:         e.HTMLBody,
		HTMLFooter:       e.HTMLFooter,
		FullHTMLTemplate: "",
		IsFullHTML:       false,
	}, nil
}

// NewEmailContent combines the new HTML email content from the header, body and footer sections
func NewEmailContent(headerHTML, bodyHTML, footerHTML string) (string, error) {
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

// NewFullHTMLContent is a full HTML template return as a string HTML format
func NewFullHTMLContent(fullHTML string) (string, error) {
	if len(strings.TrimSpace(fullHTML)) == 0 {
		return "", errors.New("full html template is required")
	}
	return fullHTML, nil
}

// NewEmail holds an email address with name and email address.
func NewEmail(name string, address string) Email {
	return Email{
		Name:    name,
		Address: address,
	}
}
