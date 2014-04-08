package main

import (
  "fmt"
  "net/http"
  "bytes"
  "flag"
  "os"
)

const header = `
___________.__
\_   _____/|  | ___.__.
 |    __)  |  |<   |  |
 |     \   |  |_\___  |
 \___  /   |____/ ____|
     \/         \/
`

var message, to, from, twilio_auth_token, twilio_account string
var other_details = true

func main() {

  setupFlags()
  validateEnvs()

  if flag.NFlag() == 0 || !other_details {
    printUsage()
  } else if message != "" && other_details {
    sendSms()
  } else {
    panic("Ohh :(")
  }

}

func sendSms() {

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
  flag.StringVar(&to, "t", "", "number you will send to / FLY_TO env")
  flag.StringVar(&from, "f", "", "number you will send from / FLY_FROM env")
  flag.StringVar(&twilio_account, "a", "", "Twilio account number / FLY_TWILIO_ACCOUNT env")
  flag.StringVar(&twilio_auth_token, "o", "", "Twilio auth token / FLY_TWILIO_AUTH_TOKEN env")
  flag.Parse()
}

func validateEnvs() {

  if to == "" && os.Getenv("FLY_TO") != "" {
    to = os.Getenv("FLY_TO")
  } else {
    other_details = false
  }

  if from == "" && os.Getenv("FLY_FROM") != "" {
    from = os.Getenv("FLY_FROM")
  } else {
    other_details = false
  }

  if twilio_auth_token == "" && os.Getenv("FLY_TWILIO_AUTH_TOKEN") != "" {
    twilio_auth_token = os.Getenv("FLY_TWILIO_AUTH_TOKEN")
  } else {
    other_details = false
  }

  if twilio_account == "" && os.Getenv("FLY_TWILIO_ACCOUNT") != "" {
    twilio_account = os.Getenv("FLY_TWILIO_ACCOUNT")
  } else {
    other_details = false
  }

}

func printUsage() {
  red()
  fmt.Println(header)
  reset()

  blue()
  fmt.Print("Usage")
  reset()
  fmt.Println(":")

  fmt.Println()
  flag.PrintDefaults()
  fmt.Println()

  blue()
  fmt.Print("Note")
  reset()
  fmt.Println(": All options must be set...")
  fmt.Println()
}

func red() {
  fmt.Print("\x1b[31;1m")
}

func blue() {
  fmt.Print("\x1b[34;1m")
}

func reset() {
  fmt.Print("\x1b[0m")
}
