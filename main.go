package main

import (
	"time"

	"github.com/minilikmila/tg_bot.git/config"
	_ "github.com/minilikmila/tg_bot.git/helper/env"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func main() {

	logrus.Infof(`WELCOME TO BUILD BOT - WITH - GO `)
 
	// router := gin.Default()
	
	// gin.SetMode(gin.DebugMode)
	// START TELEBOT HERE
	pref := tele.Settings{
		Token: config.TOKEN,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	 bot, err := tele.NewBot(pref)
	 if err != nil {
		panic("error when init bot.")
	}

	bot.Handle("/hello", func(ctx tele.Context) error {
		return ctx.Send("Hello from GO/Telebot")
	})

	bot.Handle("/aboutus", func(ctx tele.Context) error {
		return ctx.Send(`This is about section`)
	})

	// Init bot 
	bot.Start()



	// logrus.Infoln("Gin mode ", gin.Mode())


	// var port = fmt.Sprintf(`:%s`, config.PORT )
	//  logrus.Infoln(router.Run(port))
	//  err != nil {
	// 	logrus.Errorf(`Error encountered when starting the server at port - %s`, port)
	// }
	
}