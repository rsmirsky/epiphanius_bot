package telebot

import (
	"context"
	"epiphanius_bot/pkg/controller"
	"epiphanius_bot/pkg/storage"
	"fmt"

	"github.com/stephenafamo/kronika"

	"strconv"

	"epiphanius_bot/pkg/types"

	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(d *gorm.DB) {
	db = d
}

func Telebot() {
	//telebot
	pref := tele.Settings{
		Token:  viper.GetString("TOKEN"), //"5556619444:AAG0jojLq5XbZhw9AuPf6TBRtAOYaKmQutE", //os.Getenv("TOKEN"), //5556619444:AAG0jojLq5XbZhw9AuPf6TBRtAOYaKmQutE
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("error connect to telebot", err)
	}

	//keyboard
	
	var (
		// Universal markup builders.
		menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
		selector = &tele.ReplyMarkup{}

		// Reply buttons.
		btnHelp = menu.Text("ℹ НАЙБЛИЖЧЕ СВЯТО")

		// Inline buttons.
		//
		// Pressing it will cause the client to
		// send the bot a callback.
		//
		// Make sure Unique stays unique as per button kind
		// since it's required for callback routing to work.
		//
		btnPrev = selector.Data("⬅", "prev", "ede")
		btnNext = selector.Data("➡", "next", "tut")
	)

	menu.Reply(
		menu.Row(btnHelp),
	)
	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)
	

	//get next holiday

	nextHoliday, err := storage.NextHoliday()
	if err != nil {
		log.Info("error get next holiday")
	}

	var NextMessage string
	for _, next := range nextHoliday {
		NextMessage = fmt.Sprintf("%s\n%s", next.Name, next.Date)
	}
	b.Handle(&btnHelp, func(c tele.Context) error {
		return c.Send(NextMessage)
	})

	//create users

	b.Handle("/start", func(c tele.Context) error {

		var (
			userId   = c.Sender().ID
			userName = c.Sender().Username
			//text = c.Text()
		)
		controller.CreateUser(userId, userName)
		log.Info("create new user")

		return c.Send("Hello!", menu)
		//controller.SendUserMessages()
	})
	

	go b.Start()

	//send to users

	var GenerateMessage string
	var GenerateMessage2 string
	//var GenerateMessageNotification string

	ctx := context.Background()

	// use time.Date()
	start, err := time.Parse(
		"2006-01-02 15:04",
		"2022-06-21 12:00", //-3 hours beetwen messageSendTime
	) // any time in the past works but it should be on the hour
	if err != nil {
		panic(err)
	}

	interval := time.Hour * 24 // * 24 // 1 hour
	fmt.Println(interval)

	// get messages
	messages, _ := storage.GetMessages()

	users, _ := storage.GetUsers()

	for t := range kronika.Every(ctx, start, interval) {

		for _, message := range messages {

			if message.Type == types.MessageTypeToday {

				GenerateMessage = fmt.Sprintf("Вітаємо! Сьогодні %s \n %s \n %d", message.Name, message.Description)

				for _, result := range users {

					TelegramUser := types.TelegramUser{
						UserID: strconv.Itoa(int(result.UserID)),
					}

					b.Send(TelegramUser, GenerateMessage)

				}
			}

			if message.Type == types.MessageTypeInAdvance {

				GenerateMessage2 = fmt.Sprintf("Привіт! Завтра %s \n Не забули?)", message.Name)
				log.Info(t)

				for _, result := range users {

					TelegramUser := types.TelegramUser{
						UserID: strconv.Itoa(int(result.UserID)),
					}

					b.Send(TelegramUser, GenerateMessage2, menu)

				}
			}
			//os.Exit(1)
		}

	}
}
