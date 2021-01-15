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
