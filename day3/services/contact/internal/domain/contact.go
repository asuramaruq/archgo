package domain

type Contact struct {
	ID       int
	FullName struct {
		FirstName  string
		LastName   string
		MiddleName string
	}
	PhoneNumber string
}

func NewContact(id int, lastName string, firstName string, middleName string, phoneNumber string) *Contact {
	contact := &Contact{
		ID: id,
	}
	contact.SetPhoneNumber(phoneNumber)
	contact.FullName = struct {
		FirstName  string
		LastName   string
		MiddleName string
	}{FirstName: firstName, LastName: lastName, MiddleName: middleName}
	return contact
}

func (c *Contact) GetPhone() string {
	return c.PhoneNumber
}

func (c *Contact) GetFullName() string {
	return c.FullName.LastName + " " + c.FullName.FirstName + " " + c.FullName.MiddleName
}

func (c *Contact) SetPhoneNumber(phoneNumber string) {
	var cleanedNumber string
	for _, char := range phoneNumber {
		if char >= '0' && char <= '9' {
			cleanedNumber += string(char)
		}
	}
	c.PhoneNumber = cleanedNumber
}
