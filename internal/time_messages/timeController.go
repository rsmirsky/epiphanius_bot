package time_messages

import (
	"context"
	"epiphanius_bot/pkg/storage"
	"epiphanius_bot/pkg/types"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stephenafamo/kronika"
)

func GenerateMessages() ([]string, error) {

	var GenerateMessage string

	var userMessages = []string{}

	ctx := context.Background()

	start, err := time.Parse(
		"2006-01-02 15:04",
		"2022-06-21 12:53", //-3 hours beetwen messageSendTime
	) // any time in the past works but it should be on the hour
	if err != nil {
		panic(err)
	}

	fmt.Println()
	interval := time.Second * 60 // * 24 // 1 hour
	fmt.Println(interval)

	// get messages
	messages, _ := storage.GetMessages()

	for t := range kronika.Every(ctx, start, interval) {

		for _, message := range messages {

			if message.Type == types.MessageTypeToday {

				GenerateMessage = fmt.Sprintf("Вітаємо! Сьогодні %s \n %s \n %d", message.Name, message.Description)
				//return GenerateMessage, err

				userMessages = append(userMessages, GenerateMessage)

			}
			if message.Type == types.MessageTypeInAdvance {

				GenerateMessage = fmt.Sprintf("Привіт! Завтра %s \n Не забули?)", message.Name)
				log.Info(t)
				userMessages = append(userMessages, GenerateMessage)
				fmt.Println(userMessages)
				//return GenerateMessage, err
			}
		}
		//os.Exit(1)
	}
	return userMessages, err
}
