package models

type statusType string

const (
	PROCESSING statusType = "PROCESSING"
	READY      statusType = "READY"
	SENT       statusType = "SENT"
)

type Order struct {
	BaseModel

	Country string     `gorm:"not null;type:varchar(150)"`
	City    string     `gorm:"not null;type:varchar(150)"`
	Address string     `gorm:"not null;type:varchar(150)"`
	Status  statusType `gorm:"type:varchar(150);default:'PROCESSING'"`

	UserID   string    `gorm:"index;type:uuid"`
	Products []Product `gorm:"many2many:order_product"`
}

func (o Order) String() string {
	return o.Address
}
