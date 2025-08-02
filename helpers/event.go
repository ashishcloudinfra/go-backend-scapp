package helpers

import (
	"fmt"
	"time"

	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func CreateScheduledEvents(eventId string, eventReq types.EventFormValues) error {
	startDate, _ := time.Parse("2006-01-02", eventReq.StartDate)
	endDate, _ := time.Parse("2006-01-02", eventReq.EndDate)

	if !eventReq.IsRecurring {
		// Single event for non-recurring
		return insertScheduledEvent(eventId, startDate, endDate, eventReq.StartTime, eventReq.EndTime)
	}

	// Generate events for recurring
	switch types.RecurrenceType(eventReq.RecurrenceType) {
	case types.Daily:
		return handleDailyRecurrence(eventId, startDate, endDate, eventReq)
	case types.Weekly:
		return handleWeeklyRecurrence(eventId, startDate, endDate, eventReq)
	case types.Monthly:
		return handleMonthlyRecurrence(eventId, startDate, endDate, eventReq)
	case types.Yearly:
		return handleYearlyRecurrence(eventId, startDate, endDate, eventReq)
	default:
		return fmt.Errorf("unsupported recurrence type: %s", eventReq.RecurrenceType)
	}
}

func handleDailyRecurrence(eventId string, startDate, endDate time.Time, eventReq types.EventFormValues) error {
	repeatEvery := eventReq.Metadata.Daily.RepeatEvery
	if repeatEvery == nil {
		return fmt.Errorf("missing repeatEvery for daily recurrence")
	}

	currentDate := startDate
	for !currentDate.After(endDate) {
		if err := insertScheduledEvent(eventId, currentDate, currentDate, eventReq.StartTime, eventReq.EndTime); err != nil {
			return err
		}
		currentDate = currentDate.AddDate(0, 0, *repeatEvery)
	}
	return nil
}

func handleWeeklyRecurrence(eventId string, startDate, endDate time.Time, eventReq types.EventFormValues) error {
	repeatEvery := eventReq.Metadata.Weekly.RepeatEvery
	repeatOn := eventReq.Metadata.Weekly.RepeatOn
	if repeatEvery == nil || len(repeatOn) == 0 {
		return fmt.Errorf("missing repeatEvery or repeatOn for weekly recurrence")
	}

	currentDate := startDate
	for !currentDate.After(endDate) {
		dayOfWeek := currentDate.Weekday().String()
		for _, targetDay := range repeatOn {
			if dayOfWeek == targetDay {
				if err := insertScheduledEvent(eventId, currentDate, currentDate, eventReq.StartTime, eventReq.EndTime); err != nil {
					return err
				}
			}
		}
		currentDate = currentDate.AddDate(0, 0, 1) // Move to next day
	}
	return nil
}

func handleMonthlyRecurrence(eventId string, startDate, endDate time.Time, eventReq types.EventFormValues) error {
	repeatEvery := eventReq.Metadata.Monthly.RepeatEvery
	monthDate := eventReq.Metadata.Monthly.MonthDate
	weekNumber := eventReq.Metadata.Monthly.MonthWeekNumber
	selectedWeekDay := eventReq.Metadata.Monthly.SelectedWeekDay
	selectorType := eventReq.Metadata.Monthly.SelectorType

	if repeatEvery == nil {
		return fmt.Errorf("missing repeatEvery for monthly recurrence")
	}

	currentDate := startDate
	for !currentDate.After(endDate) {
		if *selectorType == "day" && monthDate != nil {
			scheduledDate := time.Date(currentDate.Year(), currentDate.Month(), *monthDate, 0, 0, 0, 0, currentDate.Location())
			if !scheduledDate.After(endDate) && scheduledDate.Month() == currentDate.Month() {
				if err := insertScheduledEvent(eventId, scheduledDate, scheduledDate, eventReq.StartTime, eventReq.EndTime); err != nil {
					return err
				}
			}
		} else if *selectorType == "weekday" && weekNumber != nil && selectedWeekDay != nil {
			// Specific weekday of the month (e.g., first Monday)
			scheduledDate := findNthWeekday(currentDate.Year(), currentDate.Month(), *weekNumber, *selectedWeekDay)
			if !scheduledDate.IsZero() && !scheduledDate.After(endDate) {
				if err := insertScheduledEvent(eventId, scheduledDate, scheduledDate, eventReq.StartTime, eventReq.EndTime); err != nil {
					return err
				}
			}
		}

		// Move to the next month based on repeatEvery
		currentDate = currentDate.AddDate(0, *repeatEvery, 0)
	}
	return nil
}

func handleYearlyRecurrence(eventId string, startDate, endDate time.Time, eventReq types.EventFormValues) error {
	month := eventReq.Metadata.Yearly.Every
	monthDate := eventReq.Metadata.Yearly.MonthDate
	weekNumber := eventReq.Metadata.Yearly.MonthWeekNumber
	selectedWeekDay := eventReq.Metadata.Yearly.SelectedWeekDay
	selectorType := eventReq.Metadata.Yearly.SelectorType

	if month == nil || (*selectorType == "day" && monthDate == nil) || (*selectorType == "weekday" && (weekNumber == nil || selectedWeekDay == nil)) {
		return fmt.Errorf("missing metadata for yearly recurrence")
	}

	currentDate := startDate
	for !currentDate.After(endDate) {
		monthIndex := parseMonthName(*month)
		if monthIndex == 0 {
			return fmt.Errorf("invalid month name: %s", *month)
		}

		if *selectorType == "day" {
			// Specific day of the month (e.g., January 1st)
			scheduledDate := time.Date(currentDate.Year(), time.Month(monthIndex), *monthDate, 0, 0, 0, 0, currentDate.Location())
			if !scheduledDate.After(endDate) {
				if err := insertScheduledEvent(eventId, scheduledDate, scheduledDate, eventReq.StartTime, eventReq.EndTime); err != nil {
					return err
				}
			}
		} else if *selectorType == "weekday" {
			// Specific weekday of the month (e.g., first Monday of January)
			scheduledDate := findNthWeekday(currentDate.Year(), time.Month(monthIndex), *weekNumber, *selectedWeekDay)
			if !scheduledDate.IsZero() && !scheduledDate.After(endDate) {
				if err := insertScheduledEvent(eventId, scheduledDate, scheduledDate, eventReq.StartTime, eventReq.EndTime); err != nil {
					return err
				}
			}
		}

		// Move to the next year
		currentDate = currentDate.AddDate(1, 0, 0)
	}
	return nil
}

func findNthWeekday(year int, month time.Month, weekNumber string, weekday string) time.Time {
	dayOfWeek := parseWeekdayName(weekday)
	if dayOfWeek == -1 {
		return time.Time{} // Invalid weekday
	}

	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	weekdayOffset := (int(dayOfWeek) - int(firstDay.Weekday()) + 7) % 7
	firstWeekday := firstDay.AddDate(0, 0, weekdayOffset)

	switch weekNumber {
	case "first":
		return firstWeekday
	case "second":
		return firstWeekday.AddDate(0, 0, 7)
	case "third":
		return firstWeekday.AddDate(0, 0, 14)
	case "fourth":
		return firstWeekday.AddDate(0, 0, 21)
	case "last":
		lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
		return findNthWeekdayReverse(lastDay, dayOfWeek)
	default:
		return time.Time{} // Invalid week number
	}
}

func findNthWeekdayReverse(date time.Time, weekday time.Weekday) time.Time {
	for date.Weekday() != weekday {
		date = date.AddDate(0, 0, -1)
	}
	return date
}

func parseWeekdayName(weekday string) time.Weekday {
	switch weekday {
	case "Sunday":
		return time.Sunday
	case "Monday":
		return time.Monday
	case "Tuesday":
		return time.Tuesday
	case "Wednesday":
		return time.Wednesday
	case "Thursday":
		return time.Thursday
	case "Friday":
		return time.Friday
	case "Saturday":
		return time.Saturday
	default:
		return -1
	}
}

func parseMonthName(month string) int {
	switch month {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	default:
		return 0
	}
}

func insertScheduledEvent(eventId string, startDate, endDate time.Time, startTime, endTime string) error {
	query := fmt.Sprintf(`INSERT INTO %s (startDate, endDate, startTime, endTime, eventId) VALUES ($1, $2, $3, $4, $5)`, database.SCHEDULED_EVENTS_TABLE_NAME)
	_, err := database.Execute(query, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), startTime, endTime, eventId)
	return err
}
