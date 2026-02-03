package purchase
import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    return kind == "car" || kind == "truck"
	panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    choice := option1
	if option2 < option1 {
		choice = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", choice)
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var percentage float64
	if age < 3 {
		percentage = 0.8 // 80% giá gốc
	} else if age >= 10 {
		percentage = 0.5 // 50% giá gốc
	} else {
		percentage = 0.7 // 70% giá gốc cho xe từ 3 đến dưới 10 tuổi
	}
	return originalPrice * percentage
	panic("CalculateResellPrice not implemented")
}
