// Package weather cung cấp các công cụ để dự báo điều kiện thời tiết.
package weather

// CurrentCondition lưu trữ điều kiện thời tiết hiện tại của một vị trí.
var CurrentCondition string

// CurrentLocation lưu trữ tên thành phố hoặc địa điểm hiện tại.
var CurrentLocation string

// Forecast trả về chuỗi mô tả dự báo thời tiết cho một thành phố cụ thể.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}