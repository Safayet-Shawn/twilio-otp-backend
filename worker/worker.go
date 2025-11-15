package worker

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/safayet-shawn/twilio-otp-app/model"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func StartWork(jobs <-chan model.Task) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromPhone := os.Getenv("TWILIO_PHONE_NUMBER")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	fmt.Println("=======>", accountSid, authToken, fromPhone)
	go func() {
		for job := range jobs {
			param := &openapi.CreateMessageParams{}
			param.SetTo(job.Phone)
			param.SetFrom("+13158093313")

			// param.SetTo(fromPhone)
			param.SetBody(fmt.Sprintf("Your OTP is %s", job.Otp))

			_, err := client.Api.CreateMessage(param)
			if err != nil {
				log.Printf("Failed to send otp", err)
			} else {
				log.Println("Successfully send the otp")
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
