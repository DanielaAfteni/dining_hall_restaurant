// Packages are used to organize related Go source files together into a single unit
// The current go program is in package basicproces
package basicproces

// A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type.
// Any real-world entity which has some set of properties/fields can be represented as a struct.
// This concept is generally compared with the classes in object-oriented programming.

// Config - has the majority of the properties/fields as int (because are integer numbers)
// and which are taken from the json file
type Config struct {
	// This Config struct will hold all configuration variables of the application that we read from file
	TimeUnit               int     `json:"time_unit"`
	NrOfTables             int     `json:"nr_of_tables"`
	NrOfWaiters            int     `json:"nr_of_waiters"`
	MaxOrderItemsCount     int     `json:"max_order_items_count"`
	MaxTableAvailableTime  int     `json:"max_table_available_time"`
	MaxWaitTimeCoefficient float64 `json:"max_wait_time_coefficient"`
	MaxOrderId             int     `json:"max_order_id"`
	MaxPickupTime          int     `json:"max_pickup_time"`
	KitchenUrl             string  `json:"kitchen_url"`
}

// set for this configuration including all the variables of the application (7 integers, 1 float, 1 url)
var scfg Config = Config{
	TimeUnit:               1000,
	NrOfTables:             10,
	NrOfWaiters:            4,
	MaxOrderItemsCount:     5,
	MaxTableAvailableTime:  20,
	MaxWaitTimeCoefficient: 1.3,
	MaxOrderId:             1000,
	MaxPickupTime:          5,
	KitchenUrl:             "http://kitchen_restaurant:8081",
}

// funtion for setting the configuration
func SettingtheConfig(s Config) {
	scfg = s
}
