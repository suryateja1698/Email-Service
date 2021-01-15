package models

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

/*

  "from_email":"suryatejam98@gmail.com",
    "name_from":"surya",
    "to_email":"suryatejam98@gmail.com",
    "name_to":"teja",
    "cc_email":"msd07surya@gmail.com",

    "sub":"This is the test",
    "text":"Hope it will work"
 "name_cc":"sample",
    "bcc_email":"suryateja1605@gmail.com",
	"namme_bcc":"Random",

*/
