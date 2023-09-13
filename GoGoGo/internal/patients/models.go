package patients

type Patient struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	LastName      string `json:"last_name"`
	Address       int    `json:"address" `
	IdCredential  string `json:"code_value"`
	DischargeDate bool   `json:"discharge_date"`
}
