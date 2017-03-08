package handler

import (
	"fmt"
	"strings"
	"time"

	stdCheck "sensu-common/check"
	"sensu/check"
)

type Point struct {
	Name  string
	Value float64
}

type Metric struct {
	Points []*Point
}

func (p Point) Render() string {
	return fmt.Sprintf("%s %f %d", p.Name, p.Value, time.Now().Unix())
}

func (m *Metric) AddPoint(p *Point) {
	m.Points = append(m.Points, p)
}

func (m *Metric) Render() check.ExtensionCheckResult {
	output := []string{}

	for _, p := range m.Points {
		output = append(output, p.Render())
	}

	return check.ExtensionCheckResult{
		Status: stdCheck.Success,
		Output: strings.Join(output, "\n"),
	}
}
