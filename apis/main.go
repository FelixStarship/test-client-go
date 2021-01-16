package main

import (
	"github.com/jakecoffman/cron"
	//"github.com/robfig/cron"
	"log"
)

func main() {
	log.Println("Starting...")
	c := cron.New()
	c.AddFunc("*/5 * * * * SAT,SUN", print5, "cron")
	c.Start()
	c.RemoveJob("cron")
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//
	//	c := cron.New()
	//	c.Stop()
	//	cron := r.Header.Get("cron")
	//	log.Println(cron)
	//	//0 0 0 ? * SAT,SUN  周六，周日 21：58分执行
	//	//*/5 * * * * SAT,SUN
	//
	//	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") }, "Less Frequent")
	//
	//	c.AddFunc("*/5 * * * * SAT,SUN", print5,"cron")
	//
	//	// 定时15秒，每5秒执行print5
	//	//c.AddFunc("*/15 * * * * *", print5)
	//	// 开始
	//	c.Start()
	//
	//
	//	if cron!="*/5 * * * * SAT,SUN" {
	//
	//		c.RemoveJob("cron")
	//
	//        log.Println("stop")
	//		c.AddFunc(cron,print5,"test")
	//
	//	}
	//	c.Stop()
	//})
	//
	//http.ListenAndServe(":8080", nil)
	select {}
}

func print5() {
	log.Println("Run 5s cron")
}
