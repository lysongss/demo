package main

import (
	"fmt"
	"time"

	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/api"
)

func main() {
	println("main")
	driver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{
		"--start-maximized",
		"--disable-infobars",
		"--webkit-text-size-adjust"}))
	driver.Start()
	var err error
	var page *agouti.Page
	page, err = driver.NewPage()
	if err != nil {
		println(err.Error())
		return
	}
	page.Navigate("https://www.xin.com/chongqing")
	// html, err := page.HTML()
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }

	keyword := "宝马"
	binputSearch := page.FindByID("MD-home-banner-brandSearch")
	//page.Find("#MD-home-banner-brandSearch xx xx")
	binputSearch.Fill(keyword)

	time.Sleep(time.Second * 3)

	li, _ := page.Find(".MD-common-header-nav a").Elements()
	println(fmt.Sprintf("a元素个数 %d", len(li)))

	li, _ = page.Find("#ui-id-1").All("li").Elements()

	println(fmt.Sprintf("li元素个数 %d", len(li)))

	li[3].Click()
	time.Sleep(time.Second * 5)

	li, _ = page.Find("#search_container div ul").All("li.con").Elements()
	println(fmt.Sprintf("search_container li元素个数 %d", len(li)))
	if len(li) > 4 {
		aimg, _ := li[3].GetElement(api.Selector{"css selector", ".aimg"})
		href, _ := aimg.GetAttribute("href")
		println("href：" + href)
		if href != "" {
			page.Navigate(href)
			time.Sleep(time.Second * 5)
			text, _ := page.Find(".cd_m_h_zjf").FirstByClass("cd_m_h_tit").Text()
			println(text)
		}
	}
	println("end")
	for {
		time.Sleep(time.Second * 5)
	}
}
