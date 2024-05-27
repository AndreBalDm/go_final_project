package nextdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/AndreBalDm/go_final_project/params"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	dateInTimeFormat, err := time.Parse(params.DFormat, date)
	if err != nil {
		return "", fmt.Errorf("format err parsing giv date: %w", err)
	}
	if repeat == "" {
		return "", fmt.Errorf("format err repeat, empty")
	}
	switch repeat[0] {
	case 'd':
		return addDays(now, dateInTimeFormat, repeat)
	case 'y':
		return addYear(now, dateInTimeFormat)
	case 'w':
		return addWeek(now, dateInTimeFormat, repeat)
	case 'm':
		return addMonth(now, dateInTimeFormat, repeat)
	default:
		return "", fmt.Errorf("format err repeat, fist symbol no correct")
	}
}

// search number to transfer task, dates of the month
func addMonth(now time.Time, dateInTimeFormat time.Time, repeat string) (string, error) {
	monthSlice := strings.Split(repeat, " ")
	var monthNumDaySlice []string
	var monthNumSlice []string
	var validDays map[int]string
	var validMonths map[int]string
	var err error
	// parsing map
	if len(monthSlice) == 2 {
		monthNumDaySlice = strings.Split(monthSlice[1], ",")
		validDays, err = mapValidDays(monthNumDaySlice)
		if err != nil {
			return "", fmt.Errorf("not valid days %v", err)
		}
	} else if len(monthSlice) == 3 {
		monthNumDaySlice = strings.Split(monthSlice[1], ",")
		validDays, err = mapValidDays(monthNumDaySlice)
		if err != nil {
			return "", fmt.Errorf("not valid days %v", err)
		}
		monthNumSlice = strings.Split(monthSlice[2], ",")
		validMonths, err = mapValidMonth(monthNumSlice)
		if err != nil {
			return "", fmt.Errorf("not valid days %v", err)
		}
	} else {
		return "", fmt.Errorf("unknow err repeat month")
	}

	//check date today
	if dateInTimeFormat.Format(params.DFormat) < now.Format(params.DFormat) {
		dateInTimeFormat = now
	}

	if len(validMonths) == 0 {
		return setDateFromCurrentMonth(validDays, dateInTimeFormat)
	} else {
		return setDateForSpecificMonths(validDays, validMonths, dateInTimeFormat)
	}
}

// parsing number month map
func mapValidMonth(monthNumSlice []string) (map[int]string, error) {
	validMonth := make(map[int]string)
	for _, month := range monthNumSlice {
		monthInt, err := strconv.Atoi(month)
		if err != nil {
			return nil, fmt.Errorf("format err repeat, no number month")
		}
		if monthInt > 12 || monthInt < 1 {
			return nil, fmt.Errorf("format err repeat number no correct")
		}
		validMonth[monthInt] = month
	}
	return validMonth, nil
}

// parsing number day map
func mapValidDays(monthNumDaySlice []string) (map[int]string, error) {
	validDays := make(map[int]string)
	for _, day := range monthNumDaySlice {
		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return nil, fmt.Errorf("format err repeat, no number week")
		}
		if dayInt > 31 || dayInt < -2 || dayInt == 0 {
			return nil, fmt.Errorf("format err repeat number no correct")
		}
		validDays[dayInt] = day
	}
	return validDays, nil
}

// search date if not month
func setDateFromCurrentMonth(validDays map[int]string, dateInTimeFormat time.Time) (string, error) {
	amountOfDaysInMonth := daysInMonth(dateInTimeFormat)
	for i := 1; i <= 31; i++ {
		_, ok := validDays[i]
		if ok {
			if (i <= int(dateInTimeFormat.Day())) || (i > amountOfDaysInMonth) {
				varDate := dateInTimeFormat.AddDate(0, 1, 0)
				validDays[i] = varDate.Format("200101") + fmt.Sprintf("%02d", i)
			} else {
				validDays[i] = dateInTimeFormat.Format("200101") + fmt.Sprintf("%02d", i)
			}
		}
	}
	if _, ok := validDays[-1]; ok {
		validDays[-1] = dateInTimeFormat.AddDate(0, 1, -dateInTimeFormat.Day()).Format(params.DFormat)
	}
	if _, ok := validDays[-2]; ok {
		validDays[-2] = dateInTimeFormat.AddDate(0, 1, -dateInTimeFormat.Day()-1).Format(params.DFormat)
	}
	//back date map
	return minDate(validDays), nil
}

// search date in month
func setDateForSpecificMonths(validDays map[int]string, validMonths map[int]string, dateInTimeFormat time.Time) (string, error) {
	resultDate := make(map[int]string)
	var addYear bool
	for j, validMonth := range validMonths {
		if fmt.Sprintf("%02s", validMonth) < dateInTimeFormat.Format("01") {
			for i, validDay := range validDays {
				addYear = true
				resultDate[i+j*100] = dateFormat(dateInTimeFormat, validMonth, validDay, addYear)
			}
		} else if fmt.Sprintf("%02s", validMonth) == dateInTimeFormat.Format("01") {
			for i, validDay := range validDays {
				if fmt.Sprintf("%02s", validDay) > dateInTimeFormat.Format("02") {
					addYear = false
					resultDate[i+j*100] = dateFormat(dateInTimeFormat, validMonth, validDay, addYear)
				} else {
					addYear = true
					resultDate[i+j*100] = dateFormat(dateInTimeFormat, validMonth, validDay, addYear)
				}
			}
		} else {
			for i, validDay := range validDays {
				addYear = false
				resultDate[i+j*100] = dateFormat(dateInTimeFormat, validMonth, validDay, addYear)
			}
		}

	}
	return minDate(resultDate), nil
}

// search suitable dates in map
func minDate(validDays map[int]string) string {
	targetDay := time.Now().AddDate(2, 0, 0).Format(params.DFormat)
	for _, validDay := range validDays {
		if (validDay < targetDay) && (validDay != "") {
			targetDay = validDay
		}
	}
	return targetDay
}

// calculation number last days in month
func daysInMonth(date time.Time) int {
	return int(date.AddDate(0, 1, -date.Day()).Day())
}

// formated date
func dateFormat(date time.Time, month string, day string, addYear bool) string {
	if addYear {
		return date.AddDate(1, 0, 0).Format("2001") + fmt.Sprintf("%02s", month) + fmt.Sprintf("%02s", day)
	}
	return date.Format("2001") + fmt.Sprintf("%02s", month) + fmt.Sprintf("%02s", day)
}

// search number to transfer task, dates of the days
func addDays(now time.Time, dateInTimeFormat time.Time, repeat string) (string, error) {
	daySlice := strings.Split(repeat, " ")
	if len(daySlice) != 2 {
		return "", fmt.Errorf("format err repeat number no correct")
	}
	dayCount, err := strconv.Atoi(daySlice[1])
	if err != nil {
		return "", fmt.Errorf("format err repeat days no number: %w", err)
	}
	if dayCount > 400 {
		return "", fmt.Errorf("format err repeat number>400")
	}

	//take new date task
	dateInTimeFormat = dateInTimeFormat.AddDate(0, 0, dayCount)
	for dateInTimeFormat.Format(params.DFormat) <= now.Format(params.DFormat) {
		dateInTimeFormat = dateInTimeFormat.AddDate(0, 0, dayCount)
	}

	return dateInTimeFormat.Format(params.DFormat), nil
}

// search number to transfer task, dates of the days week
func addWeek(now time.Time, dateInTimeFormat time.Time, repeat string) (string, error) {
	daySlice := strings.Split(repeat, " ")
	if len(daySlice) != 2 {
		return "", fmt.Errorf("format err repeat number week no correct")
	}
	weekDaySlice := strings.Split(daySlice[1], ",")
	validDays := make(map[int]string)
	for _, day := range weekDaySlice {
		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return "", fmt.Errorf("format err repeat days no number")
		}
		if dayInt > 7 || dayInt < 1 {
			return "", fmt.Errorf("format err repeat days no number 1 to 7")
		}
		validDays[dayInt] = day
	}

	//check date today
	if dateInTimeFormat.Format(params.DFormat) < now.Format(params.DFormat) {
		dateInTimeFormat = now
	}

	//map key
	for i := 1; i <= 7; i++ {
		dateInTimeFormat = dateInTimeFormat.AddDate(0, 0, 1)
		weekDay := weekDayNumber(dateInTimeFormat)
		_, ok := validDays[weekDay]
		if ok {
			validDays[weekDay] = dateInTimeFormat.Format(params.DFormat)
		}
	}
	//search date in map
	targetDay := dateInTimeFormat.Format(params.DFormat)
	for _, validDay := range validDays {
		if validDay < targetDay {
			targetDay = validDay
		}
	}
	return targetDay, nil
}

// days  week take number 1 to 7
func weekDayNumber(day time.Time) int {
	weekDay := day.Weekday()
	if int(weekDay) == 0 {
		return 7
	}
	return int(weekDay)
}

// search number to year
func addYear(now time.Time, dateInTimeFormat time.Time) (string, error) {
	//new date task
	dateInTimeFormat = dateInTimeFormat.AddDate(1, 0, 0)
	for dateInTimeFormat.Format(params.DFormat) <= now.Format(params.DFormat) {
		dateInTimeFormat = dateInTimeFormat.AddDate(1, 0, 0)
	}
	return dateInTimeFormat.Format(params.DFormat), nil
}
