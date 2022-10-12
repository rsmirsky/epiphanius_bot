package time_messages

import (
	"fmt"
	"time"

	//log "github.com/sirupsen/logrus"
	"context"
	"epiphanius_bot/pkg/storage"
	"epiphanius_bot/pkg/types"

	//"os"

	log "github.com/sirupsen/logrus"
	"github.com/stephenafamo/kronika"
)

// const layout = "2006.01.02 15:04"

// func GenerateMessages() (messageDate types.Message, err error) {

// 	//var holidayDate Holidays

// 	now := time.Now()

// 	currentDate := now.Format("2006.01.02 15:04")

// 	notificationDate := now.AddDate(0, 0, -1).Format("2006.01.02 15:04")

// 	fmt.Println(currentDate)
// 	fmt.Println("______________")
// 	fmt.Println(notificationDate)

// 	messagesList, err := storage.GetMessages()
// 	if err != nil {
// 		log.Error("error getting list of holidays from db")
// 	}

// 	ctx := context.Background()

// 	// use time.Date()
// 	start, err := time.Parse(
// 		"2006-01-02 15:04",
// 		"2022-06-21 07:00", //3 hours beetwen
// 	) // any time in the past works but it should be on the hour
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(start)

// 	fmt.Println()
// 	interval := time.Second * 60 // * 24 // 1 hour
// 	fmt.Println(interval)

// 	// notification date
// 	t, err := time.Parse(layout, messageDate.Date)
// 	fmt.Println(t, err)

// 	for t := range kronika.Every(ctx, start, interval) {
// 		// Perform action here

// 		for _, messageDate := range messagesList {

// 			// notification date
			
// 			tm, _ := time.Parse(layout, messageDate.Date)

// 			notification := tm.AddDate(0, 0, -1).Format("2006.01.02 15:04")

// 			// время не нужно сравнивать, только дату
// 			if time.Now().Format("2006.01.02 15:04") == notification {

// 				fmt.Println("Notification")
// 				 	log.Println(t.Format("2006-01-02 15:04"))
//                 //continue
// 			}
			
// 			//today holiday
// 			 if time.Now().Format("2006.01.02 15:04") == messageDate.Date {

// 			 	fmt.Println("Today holiday")

// 			 	log.Println(t.Format("2006-01-02 15:04"))

// 		     }

// 		}

// 	}
// 	return types.Message{}, err
// }

 func GenerateMessages() ([]string, error) {

	var GenerateMessage string
	//var GenerateMessageNotification string
    var  userMessages = []string{} 



	ctx := context.Background()

	// use time.Date()
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
