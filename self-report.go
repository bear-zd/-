package main

import (
	"github.com/tebeka/selenium"
	"os"
	"time"
)

const (
	chromeDriverPath = "xxxxxx/chromedriver.exe"
	port             = 8080
	username         = "xxxxxxxx"
	password         = "xxxxxxx"
	// 正则表达式匹配字符串包含“补报)”
	regex = ".*已按时.*"
)

func main1() {
	// Start a WebDriver server instance
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(false)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err)
		return
	}
	defer func(service *selenium.Service) {
		err := service.Stop()
		if err != nil {
			panic(err)
			return
		}
	}(service)

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "http://localhost:8080/wd/hub")
	if err != nil {
		panic(err)
		return
	}
	defer func(wd selenium.WebDriver) {
		err := wd.Quit()
		if err != nil {
			panic(err)
			return
		}
	}(wd)

	// Navigate to the simple playground interface.
	if err := wd.Get("https://selfreport.shu.edu.cn/Default.aspx"); err != nil {
		panic(err)
	}
	// Find the text input element by its name.
	elem1, err := wd.FindElement(selenium.ByID, "username")
	if err != nil {
		panic(err)
	}
	// Send the keys to the element.
	if err := elem1.SendKeys(username); err != nil {
		panic(err)
	}
	elem2, err := wd.FindElement(selenium.ByID, "password")
	if err != nil {
		panic(err)
	}
	if err := elem2.SendKeys(password); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)
	// check the login status
	btn1, err := wd.FindElement(selenium.ByID, "submit-button")
	if err != nil {
		panic(err)
	}
	if err := btn1.Click(); err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	btn2, err := wd.FindElement(selenium.ByID, "lnkReportHistory")
	if err != nil {
		panic(err)
	}
	if err := btn2.Click(); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)
	// while ture
	for {
		if err = wd.Get("https://selfreport.shu.edu.cn/ViewDayReport.aspx?day=2022-07-25"); err != nil {
			panic(err)
		}
		time.Sleep(2 * time.Second)
		// find the element by id
		elem3, err := wd.FindElement(selenium.ByID, "fineui_0-inputEl-icon")
		if err != nil {
			panic(err)
		}
		err = elem3.Click()
		if err != nil {
			panic(err)
		}
		time.Sleep(2 * time.Second)
		elem4, err := wd.FindElement(selenium.ByID, "ctl03_ctl00_btnReturn")
		if err != nil {
			panic(err)
		}
		err = elem4.Click()
		if err != nil {
			panic(err)
		}
		time.Sleep(2 * time.Second)
	}

}
