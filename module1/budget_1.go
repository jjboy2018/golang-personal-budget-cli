package module1

import (
	"fmt"
)

// Budget stores budget information
Budget struct {
	Max float32 
	Items []Item
}

// Item stores item information
Item struct{
	Description string
	Price float32
}
