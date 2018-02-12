package member

// Member สมาชิกผู้ถือบัตร เป็นบุคคลธรรมดา ซึ่งอาจมีหลาย "นิติบุคคล" ในสังกัด
type Member struct {
	ID           uint64 `json:"id"`
	MobileNumber string `json:"mobile_number"`
	*People
	Companies []*Company `json:"companies"`
}

// NewMember initiate Member ใหม่ ต้องกำหนดค่า People และ Company
func NewMember(n string, p *People, c ...*Company) *Member {
	return &Member{MobileNumber: n, People: p, Companies: c}
}

// People บุคคลธรรมดา
type People struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CitizenID string `json:"citizen_id"`
}

// NewPeople initiate People ใหม่
func NewPeople(f, l, c string) *People {
	return &People{
		FirstName: f,
		LastName:  l,
		CitizenID: c,
	}
}

// Company นิติบุคคล บริษัท กิจการ จดทะเบียนภาษีตามกฏหมาย
type Company struct {
	ID    uint64
	Name  string
	TaxID string
}

// NewCompany initiate บริษัทใหม่
func NewCompany(n, t string) *Company {
	return &Company{Name:  n, TaxID: t}
}

// Merchant ผู้ค่าที่ออกบัตรเครดิต
type Merchant struct {	
	ID uint64
	*People
	*Company
}

// NewMerchant Initiate new merchant
func NewMerchant(p *People, c *Company) *Merchant {
	return &Merchant{People: p, Company: c}
}
