package cronexp

import (
	//"fmt"
	//"net/http"
	//"strings"

	"github.com/gorhill/cronexpr"
	log "github.com/sirupsen/logrus"
	"time"
)

type (
	// Service Cron Descriptor service
	Service interface {
		DescribeCronExpression(expression string) (string, error)
	}

	CronExpression struct {
	}
)

func NewCronExpression() Service {
	return &CronExpression{}
}

func (expr *CronExpression) DescribeCronExpression(expression string) (string, error) {
	log.WithField("expression", expression).Debug("describing cron expression")
	expression = handleSpecialSyntax(expression)

	parsed, err := cronexpr.Parse(expression)
	if err != nil {
		return "", err
	}

	nextTime := parsed.Next(time.Now()).String()
	return nextTime, nil
	//// descriptor response
	//type Result struct {
	//	Description string `json:"description"`
	//}
	//// descriptor parameters
	//type DescParams struct {
	//	Expression string `url:"expression,omitempty"`
	//	Locale     string `url:"locale,omitempty"`
	//}
	//params := &DescParams{Expression: expression, Locale: "en-US"}
	//// invoke hermes trigger
	//var res Result
	//resp, err := api.endpoint.New().Get("api/descriptor").QueryStruct(params).ReceiveSuccess(&res)
	//if err != nil {
	//	log.WithError(err).Error("failed to invoke Cron Descriptor API")
	//	return "", err
	//}
	//if resp.StatusCode != http.StatusOK {
	//	log.WithField("http status", resp.Status).Error("Cron Descriptor API failed")
	//	return "", fmt.Errorf("%s: error describing cron expression '%s'", resp.Status, expression)
	//}
	//// descriptor service returns 200 even for error and error details in string
	//if strings.HasPrefix(res.Description, "Error:") {
	//	log.Errorf("error describing cron expression '%s'", res.Description)
	//	return "", fmt.Errorf("error describing cron expression '%s'", res.Description)
	//}
	//
	//log.WithField("description", res.Description).Debug("successfully described cron expression")
	//return res.Description, nil
}

func handleSpecialSyntax(expression string) string {
	switch expression {
	case "@yearly", "@annually":
		return "0 0 0 1 1 * *"
	case "@monthly":
		return "0 0 0 1 * * *"
	case "@weekly":
		return "0 0 0 * * 0 *"
	case "@daily", "@midnight":
		return "0 0 0 * * * *"
	case "@hourly":
		return "0 0 * * * * *"
	default:
		return expression
	}
}
