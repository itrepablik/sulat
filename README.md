
![sulat](https://user-images.githubusercontent.com/58651329/82276988-5dcdf680-99b9-11ea-8ba4-f0264d7cbf72.png)
The **sulat** package is an easier way of sending an automatic email notification from your Go's project, currently supported the SendGrid SMTP API, in the future we can add more to cater to your needs.

# Installation
```
go get -u github.com/itrepablik/sulat
```

# Usage
This is the sample usage for the sulat package.
```
package main

import (
	"fmt"
	"time"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/sulat"
	"github.com/itrepablik/timaan"
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

func main() {
	// Create timaan token for email confirmation
	rt := timaan.RandomToken()
	emailConfirmPayload := timaan.TP{
		"USERNAME": "politz",
		"EMAIL":    "politz@live.com",
	}
	tok := timaan.TK{
		TokenKey: rt,
		Payload:  emailConfirmPayload,
		ExpireOn: time.Now().Add(time.Minute * 30).Unix(),
	}
	newToken, err := timaan.GenerateToken(rt, tok)
	if err != nil {
		itrlog.Fatal(err)
	}
	confirmURL := "https://itrepablik.com/confirm/" + newToken
	fmt.Println(confirmURL)

	// Start sending the email.
	sgc := sulat.SendGridConfig{
		SendGridAPIKey:   "YOUR_SEND_GRID_API_KEY",
		SendGridEndPoint: "/v3/mail/send",
		SendGridHost:     "https://api.sendgrid.com",
	}

	// Prepare the HTML email content
	emailContent, err := sulat.NewEmailContent(HTMLHeader, bodyHTML, HTMLFooter)
	if err != nil {
		itrlog.Fatal(err)
	}

	mailOpt := sulat.SM.Options(&sulat.SendMail{
		Subject:  "Inquiry for the new ITR Sulat package",
		From:     sulat.NewEmail("ITR Support", "support@itrepablik.com"),
		To:       sulat.NewEmail("Politz", "politz@live.com"),
		HTMLBody: emailContent,
	})

	isSend, err := sulat.SM.Send(mailOpt, &sgc)
	if err != nil {
		itrlog.Fatal(err)
	}
	fmt.Println("isSend: ", isSend)
}
```
# License
Code is distributed under MIT license, feel free to use it in your proprietary projects as well.
