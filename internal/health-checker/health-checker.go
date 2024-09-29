package health_checker

type HealthChecker struct {
	msg string
}

func New(msg string) (*HealthChecker, error) {
	return &HealthChecker{msg: msg}, nil
}

func (hc *HealthChecker) HealthCheck() (string, error) {
	return hc.msg, nil
}
