package main

import (
	"purchase-notifi/pkg/entity"
	pkglogger "purchase-notifi/pkg/logger"
	"purchase-notifi/pkg/notification"
	"purchase-notifi/pkg/operators"
	"flag"
	"go.uber.org/zap"
)

var (
	err          error
	logger       *zap.Logger
	operatorType = flag.String("o", "email", "operator for send notification")
)

func main() {
	flag.Parse()

	logger, _ := pkglogger.NewZapLogger()
	defer logger.Sync()

	// purchase initialization
	purchase := entity.NewPurchase(
		17,
		entity.NewCustomer(7, "customer@info.com", "123456789"),
		entity.NewProduct("iPhone 12", 699),
	)

	// notification initialization
	notifi := notification.New(operators.NewEmail(), operators.NewSms())

	// send a purchase notification
	if err := notifi.Purchase(*operatorType, purchase); err != nil {
		logger.Error("error sending purchase notification", zap.Error(err))
		return
	}

	logger.Info("purchase notification sent successfully", zap.String("operator", *operatorType))
}
