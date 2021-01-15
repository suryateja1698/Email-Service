# Email-Service
This is a Email Service application based on Golang, REST, Mux, Mailjet


 * Here we need to install mailjet in our local, you can read more about 
 [mailjet](https://github.com/mailjet/mailjet-apiv3-go) 
 
 After installing `mailjet` package

 1. ## main.go
   
   ``` go
   router := handler.NewRouter()
	http.Handle("/", router)
	//	os.Setenv("PORT", "8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	n := negroni.Classic()

	n.UseHandler(handlers.CORS()(router))

	fmt.Println("Listening on port\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), n)
	if err != nil {
		panic(err.Error())
	}


   ```

2. ## handler/handler.go

   In `handler.go` we can give our handler functions, here we are using `mux`

```go
    router.HandleFunc("/send-email", email.SendOne).Methods("POST")
```	

	



3. ## models/email.go

In `models.email.go` we define our request structure for sending mail 

 ```go
 type From struct {
	FromEmail string `json:"from_email"`
	FromName  string `json:"Name_from"`
	ToEmail   string `json:"to_email"`
	ToName    string `json:"name_to"`
	CcName    string `json:"name_cc"`
	CcEmail   string `json:"cc_email"`
	BccEmail  string `json:"bcc_email"`
	BccName   string `json:"name_bcc"`
	Subject   string `json:"sub"`
	TextPart  string `json:"text"`
}
```

4. ## email/email.go

In `emai/email/go` we send our request

```go
var req models.From

	err := json.NewDecoder(r.Body).Decode(&req)

	resp, err := mailjetmail.SendEmailWithMailjet(r.Context(), req)
``` 

5. ## mailjetmail/mailjet.go

In `mailjetmail/mailjet.go` we use `mailjet` package and send our request to them, 

```go
mailjetClient.SendMailV31(&messages)
```
Here we need to pass API keys 

mailjetClient := mailjet.NewMailjetClient("", "")

After use `POSTMAN` and do `POST` request and send request like this 

```go
  "from_email":"s@gmail.com",
    "name_from":"name1",
    "to_email":"m98@gmail.com",
    "name_to":"name2",
    "cc_email":"msa@gmail.com",

    "sub":"This is the test",
    "text":"Hope it will work"
 "name_cc":"sample",
    "bcc_email":"tea1@gmail.com",
	"namme_bcc":"Random",

```
**You will successfully send email**

 
