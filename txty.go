package main

import (
  "fmt"
  "net/http"
  "bytes"
  "flag"
  "os"
)

const header = `
___________         __          
\__    ___/__  ____/  |_ ___.__.
  |    |  \  \/  /\   __<   |  |
  |    |   >    <  |  |  \___  |
  |____|  /__/\_ \ |__|  / ____|
                \/       \/     
`

var message, to, from, twilio_auth_token, twilio_account string
var other_details = true

func main() {

  setupFlags()
  validateEnvs()

  if flag.NFlag() == 0 || !other_details {
    fmt.Print("\x1b[31;1m") //red
    fmt.Println(header)
    fmt.Print("\x1b[0m") // reset

    fmt.Println("\x1b[34;1mUsage\x1b[0m:")
    fmt.Println()
    flag.PrintDefaults()
    fmt.Println()
    fmt.Println("\x1b[34;1mNote\x1b[0m: All options must be set...")
    fmt.Println()
  } else {
    if message != "" && other_details {
      sendSms(message)
    }
  }

}

func sendSms(message string) {

  client := &http.Client{}
  postData := "Body="+message+"&To="+to+"&From="+from
  
  req, _ := http.NewRequest("POST", apiAddress(), bytes.NewReader([]byte(postData)))
  req.SetBasicAuth(twilio_account, twilio_auth_token)
  
  req.ContentLength = int64(len(postData))
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  res, err := client.Do(req)

  defer res.Body.Close()

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Message sent :D!")
  }

}

func apiAddress() string {
  return "https://api.twilio.com/2010-04-01/Accounts/"+twilio_account+"/SMS/Messages.json"
}

func setupFlags() {
  flag.StringVar(&message, "m", "", "message you want to send")
  flag.StringVar(&to, "t", "", "number you will send to / TXTY_TO env")
  flag.StringVar(&from, "f", "", "number you will send from / TXTY_FROM env")
  flag.StringVar(&twilio_account, "a", "", "Twilio account number / TXTY_TWILIO_ACCOUNT env")
  flag.StringVar(&twilio_auth_token, "o", "", "Twilio auth token / TXTY_TWILIO_AUTH_TOKEN env")
  flag.Parse()
}

func validateEnvs() {
  
  if to == "" && os.Getenv("TXTY_TO") != "" {
    to = os.Getenv("TXTY_TO")
  } else {
    other_details = false
  }

  if from == "" && os.Getenv("TXTY_FROM") != "" {
    from = os.Getenv("TXTY_FROM")
  } else {
    other_details = false
  }

  if twilio_auth_token == "" && os.Getenv("TXTY_TWILIO_AUTH_TOKEN") != "" {
    twilio_auth_token = os.Getenv("TXTY_TWILIO_AUTH_TOKEN")
  } else {
    other_details = false
  }

  if twilio_account == "" && os.Getenv("TXTY_TWILIO_ACCOUNT") != "" {
    twilio_account = os.Getenv("TXTY_TWILIO_ACCOUNT")
  } else {
    other_details = false
  }

}