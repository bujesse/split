package jobs

import (
	"split/repositories"
	"split/services"
	"time"

	"github.com/go-co-op/gocron"
)

func SchedulerInit(
	expenseRepo repositories.ExpenseRepository,
	currencyRepo repositories.CurrencyRepository,
	fxRateRepo repositories.FxRateRepository,
) {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(1).Day().At("05:00").Do(func() {
		ProcessRecurringExpenses(expenseRepo)
		services.FetchAndStoreFxRates(currencyRepo, fxRateRepo)
	})
	// scheduler.Every(1).Minute().Do(func() { ProcessRecurringExpenses(expenseRepo) })
	scheduler.StartAsync()
}
