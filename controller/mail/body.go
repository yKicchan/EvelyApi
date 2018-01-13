package mail

import (
    . "EvelyApi/config"
)

func makeHtmlBody(title, subtitle, message, url, button string) string {
    return `
<html>
<head>
<title>` + title + `</title>
</head>
<body>
<table width="100%" height="100%" style="min-width:348px;" border="0" cellspacing="0" cellpadding="0">
    <tbody>
        <tr align="center">
            <td>
                <table border="0" cellspacing="0" cellpadding="0" style="padding-bottom:20px;max-width:500px;min-width:220px;border-top:4px solid #ffcd8e;border-right:1px solid #dcdcdc;border-bottom:1px solid #dcdcdc;border-left:1px solid #dcdcdc;border-radius:4px">
                    <tr>
                        <td>
                            <div style="font-family:Arial,sans-serif;padding:30px 20px 10px 20px;color:rgba(0,0,0,0.87);font-size:20px;text-align:center;">
                                <div>
                                    ` + title + `<br>
                                    <span style="color:rgba(0,0,0,0.87);font-size:0.85em;text-decoration:none">` + subtitle + `</span>
                                </div>
                            </div>
                            <div style="border-top:1px solid #e8e8e8;width:100%;"></div>
                            <div style="font-family:Arial,sans-serif;font-size:0.8em;color:rgba(0,0,0,0.87);line-height:1.6em;padding:30px 20px 0px;">
                                <div>
                                    ` + message + `
                                    <div style="padding-top:24px;text-align:center">
                                        <a href="` + url + `" style="display:inline-block;text-decoration:none;" target="_blank">
                                            <table border="0" style="background-color:#4184f3;border-radius:3px;width:80px">
                                                <tr>
                                                    <td style="padding-left:7px;padding-right:7px;text-align:center">
                                                        <a href="` + url + `" style="font-family:Arial,sans-serif;color:#ffffff;text-decoration:none;font-size:0.85em;line-height:1.8em;">` + button + `</a>
                                                    </td>
                                                </tr>
                                            </table>
                                        </a>
                                    </div>
                                </div>
                            </div>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </tbody>
</table>
</body>
</html>
`
}

func makeSignUpBody(email, url string) string {
    return makeHtmlBody(SIGNUP_BODY_TITLE, email, SIGNUP_BODY_MESSAGE, url, SIGNUP_BODY_BUTTON)
}
