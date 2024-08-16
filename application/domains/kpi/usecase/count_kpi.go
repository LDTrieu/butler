package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func (u *usecase) CountKpi(ctx context.Context, date string, env string) error {
	dates := getDatesFromStartMonthTillNow(date)
	if len(dates) > 31 {
		logrus.Errorf("getDatesFromStartMonthTillNow: %v", dates)
		return fmt.Errorf("some date is invalid")
	}
	for _, d := range dates {
		if env == "prod" {
			if err := u.lib.KafkaPublisherProd.WriteByKey(CountKpiMessage{
				Date: d,
			}, d, "qc.wms.queue.count_kpi"); err != nil {
				logrus.Errorf("KafkaPublisherProd.WriteByKey: %v", err)
				return err
			}
		} else {
			if err := u.lib.KafkaPublisherQc.WriteByKey(CountKpiMessage{
				Date: d,
			}, d, "qc.wms.queue.count_kpi"); err != nil {
				logrus.Errorf("KafkaPublisherQc.WriteByKey: %v", err)
				return err
			}
		}
	}
	return nil
}

type CountKpiMessage struct {
	Date string
}

func getDatesFromStartMonthTillNow(date string) []string {
	if _, err := time.Parse(time.DateOnly, date); err == nil {
		return []string{date}
	}

	timeRequest, err := time.Parse("2006-01", "2024-06")
	if err != nil {
		return nil
	}
	res := make([]string, 0)
	for i := timeRequest; i.Before(time.Now().AddDate(0, 0, -1)) && i.Month() == timeRequest.Month(); i = i.AddDate(0, 0, 1) {
		res = append(res, i.Format("2006-01-02"))
	}

	return res
}
