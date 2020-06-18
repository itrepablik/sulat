
![sulat](https://user-images.githubusercontent.com/58651329/82276988-5dcdf680-99b9-11ea-8ba4-f0264d7cbf72.png)
The **sulat** package is an easier way of sending an automatic email notification from your Go's project, currently supported the SendGrid SMTP API and the Gomail classic SMTP server, in the future we can add more to cater to your needs.

# Installation
```
go get -u github.com/itrepablik/sulat
```

# Usage using SendGrid
This is the sample used for the sulat package to send email notification from your Go's application using the SendGrid API key.
```
package main

import (
	"fmt"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/sulat"
)

// HTMLHeader is the HTML skeletal framework head section of the standard HTML structure that serves as an email content
const HTMLHeader = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
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

// HTMLFooter is the HTML skeletal framework footer section of the standard HTML structure that serves as an email content
const HTMLFooter = `<div style="text-align: center; font-size: 12px; color: #b2b2b5; margin-top: 20px">
<p> Powered by ITRepablik.com
	<br>
	<a href="javascript: void(0);" style="color: #b2b2b5; text-decoration: underline;">Unsubscribe</a>
</p>
</div>
</div>
</div>
</body>
</html>`

// bodyHTML is your custom email contents, this is just an example of a password reset auto email from your Go's project
var bodyHTML = `<div style="padding: 40px; background: #fff;">
<table border="0" cellpadding="0" cellspacing="0" style="width: 100%;">
	<tbody>
		<tr>
			<td style="border-bottom:1px solid #f6f6f6;">
				<h1 style="font-size:14px; font-family:arial; margin:0px; font-weight:bold;">Hi politz,</h1>
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
				<a href="#">https://itrepablik.com/activate/abcde12345</a><br/>
				If you continue to have problems, please feel free to contact us at <a href="mailto:support@itrepablik.com">support@itrepablik.com</a>
			</td>
		</tr>
	</tbody>
</table>
</div>`

// FullHTML use this when you preferred the full HTML template as your content
var FullHTML = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
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
                                <a href="#">https://itrepablik.com/activate/hello-world</a><br/>
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

// SGC initialize this variable globally sulat.SendGridConfig{}
var SGC = sulat.SGC{}

func init() {
	// Start sending the email.
	SGC = sulat.SGC{
		SendGridAPIKey:   "YOUR_SENDGRID_API_KEY",
		SendGridEndPoint: "/v3/mail/send",
		SendGridHost:     "https://api.sendgrid.com",
	}
}

func main() {
	// Prepare the HTML email content
	mailOpt := &sulat.SendMail{
		Subject: "Inquiry for the new ITR Sulat package",
		From:    sulat.NewEmail("ITR Support", "support@itrepablik.com"),
		To:      sulat.NewEmail("Politz", "politz@live.com"),
	}

	//*****************************************************************************
	// Use only either 'Method 1' or 'Method 2' for your email HTML content
	//*****************************************************************************

	// Method 1: Set full HTML template as your email content.
	// e.g email marketing campaign template
	htmlContent, err := sulat.SetHTML(&sulat.EmailHTMLFormat{
		IsFullHTML:       true,
		FullHTMLTemplate: FullHTML,
	})

	// Method 2: Set this standard HTML header and footer but with different HTML body
	// this is usually use when you've fixed header and footer content
	// e.g standard email notifications such as password reset, email confirmation, etc.
	htmlContent, err = sulat.SetHTML(&sulat.EmailHTMLFormat{
		IsFullHTML: false,
		HTMLHeader: HTMLHeader,
		HTMLBody:   bodyHTML,
		HTMLFooter: HTMLFooter,
	})

	isSend, err := sulat.SendEmailSG(mailOpt, htmlContent, &SGC)
	if err != nil {
		itrlog.Fatal(err)
	}
	fmt.Println("isSend: ", isSend)
}
```

# Usage using Gomail
This is how you can use the classic SMTP server to send email notifications from your Go's application.
```
package main

import (
	"fmt"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/sulat"
)

// SMTPCon initialize this variable globally sulat.SMTPConfig{} for the 'SMTP' Classic
var SMTPCon = sulat.SMTPConfig{}

func init() {
	// Initialize the 'SMTP' classic
	SMTPCon = sulat.SMTPConfig{
		Host:     "smtp.host.com",
		Port:     25,
		UserName: "your@email.com",
		Password: "your_smtp_password",
	}
}

func main() {
	// This is how to use the 'sulat' package using 'SendGrid' API key
	// Prepare the HTML email content
	mailOpt := &sulat.MailClassicHeader{
		Subject: "Inquiry for the new ITR Sulat package",
		From:    "support@itrepablik.com",
		To:      []string{"email1@mail.com", "email2@mail.com"},
		CC:      []string{"email3@mail.com", "email4@mail.com"},
		BCC:     []string{"email5@mail.com", "email6@mail.com"},
	}

	//*****************************************************************************
	// Use only either 'Method 1' or 'Method 2' for your email HTML content
	//*****************************************************************************

	// Method 1: Set full HTML template as your email content.
	// e.g email marketing campaign template
	htmlContent, err := sulat.SetHTML(&sulat.EmailHTMLFormat{
		IsFullHTML:       true,
		FullHTMLTemplate: FullHTML,
	})

	// Method 2: Set this standard HTML header and footer but with different HTML body
	// this is usually use when you've fixed header and footer content
	// e.g standard email notifications such as password reset, email confirmation, etc.
	htmlContent, err = sulat.SetHTML(&sulat.EmailHTMLFormat{
		IsFullHTML: false,
		HTMLHeader: HTMLHeader,
		HTMLBody:   bodyHTML,
		HTMLFooter: HTMLFooter,
	})

	// Send email using 'SMTP' classic method
	isSend, err := sulat.SendEmailSMTP(mailOpt, htmlContent, &SMTPCon)
	if err != nil {
		itrlog.Fatal(err)
	}
	fmt.Println("isSend: ", isSend)
}
```

# License
Code is distributed under MIT license, feel free to use it in your proprietary projects as well.
