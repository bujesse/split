package jobs

import (
	"split/repositories"
	"time"

	"github.com/go-co-op/gocron"
)

func SchedulerInit(expenseRepo repositories.ExpenseRepository) {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(1).Day().At("05:00").Do(func() { ProcessRecurringExpenses(expenseRepo) })
	// scheduler.Every(1).Minute().Do(func() { ProcessRecurringExpenses(expenseRepo) })
	scheduler.StartAsync()
}
