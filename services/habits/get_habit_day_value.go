package habits

import (
	"reflect"
	"strings"
	"time"

	"github.com/dev-Gois/habbits-api/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetHabitDayValue(habit *models.Habit, date time.Time) bool {
	weekday := date.Weekday().String()

	fieldName := cases.Title(language.English).String(strings.ToLower(weekday))

	val := reflect.ValueOf(habit).Elem().FieldByName(fieldName)

	return val.Bool()
}
