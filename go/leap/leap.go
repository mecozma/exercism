/*
Package leap checks if a year is leap.
*/
package leap

// IsLeapYear function return true if the year is leap and false if not.
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
