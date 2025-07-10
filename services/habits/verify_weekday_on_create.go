package habits

import (
	"reflect"
	"strings"
	"time"

	"github.com/dev-Gois/habbits-api/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func VerifyWeekdayOnCreate(habit *models.Habit) error {
	today := time.Now()
	weekday := today.Weekday().String()

	fieldName := cases.Title(language.English).String(strings.ToLower(weekday))

	val := reflect.ValueOf(habit).Elem().FieldByName(fieldName)

	if val.IsValid() && val.Kind() == reflect.Bool && val.Bool() {
		habitCheck := &models.HabitCheck{
			HabitID: habit.ID,
			Done:    false,
			Date:    today,
		}

		if err := habitCheck.Create(); err != nil {
			return err
		}
	}

	return nil
}
