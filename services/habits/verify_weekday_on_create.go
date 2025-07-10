package habits

import (
	"reflect"
	"strings"
	"time"

	"github.com/dev-Gois/habbits-api/models"
	"github.com/dev-Gois/habbits-api/services/habit_checks"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func VerifyWeekdayOnCreate(habit *models.Habit) error {
	today := time.Now()
	weekday := today.Weekday().String()

	fieldName := cases.Title(language.English).String(strings.ToLower(weekday))

	val := reflect.ValueOf(habit).Elem().FieldByName(fieldName)

	if val.IsValid() && val.Kind() == reflect.Bool && val.Bool() {
		if err := habit_checks.Create(habit.ID); err != nil {
			return err
		}
	}

	return nil
}
