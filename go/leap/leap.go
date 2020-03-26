// Package leap provides details about leap years in the Gregorian calendar
package leap

/*
IsLeapYear will return a boolean indicating whether a given year(int) is a leap year or not
Leap years occur on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
*/
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}
