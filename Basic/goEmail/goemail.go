package main
import "gopkg.in/gomail.v2"
func main()  {
	// 初始化对象
	m := gomail.NewMessage()
	// 发件人地址
	//m.SetAddressHeader("From", "kill885560@outlook.com", "戴")
	m.SetHeader("From", m.FormatAddress("kill885560@outlook.com","戴"))
	//收件人地址
	m.SetHeader("To", "superempirenet@outlook.com", "paoliancoltd@outlook.com")
	//主题
	m.SetHeader("Subject", "Hello!")
	//内容
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")
	//配置stmp参数参数
	d := gomail.NewDialer("smtp-mail.outlook.com", 587, "kill885560@outlook.com", "Asdf3344")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
