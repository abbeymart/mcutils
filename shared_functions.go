// Shared functions for the mc-applications
package utils

func LeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
