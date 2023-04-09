package monitor

import (
	"fmt"
	"github.com/huantt/mac-keeper/pkg/macos"
	_ "github.com/spf13/cobra"
	"sync"
	"time"
)

type PowerMonitor struct {
	alertFunctions []Alert
}

func NewPowerMonitor() *PowerMonitor {
	return &PowerMonitor{}
}

type Alert func(message string) error

func (p *PowerMonitor) AddAlertFunc(alertFunc Alert) *PowerMonitor {
	p.alertFunctions = append(p.alertFunctions, alertFunc)
	return p
}

func (p *PowerMonitor) Run() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		if powerSource, err := macos.GetPowerSource(); err != nil {
			fmt.Println(fmt.Sprintf("Failed to get power source. Error: %s", err.Error()))
		} else if *powerSource != macos.ACPowerSource {
			fmt.Println("Alerting...")
			wg := sync.WaitGroup{}
			wg.Add(len(p.alertFunctions))
			for _, alertFunc := range p.alertFunctions {
				go func(alert Alert) {
					defer wg.Done()
					err := alert(fmt.Sprintf("Your macos has been unplugged!"))
					if err != nil {
						fmt.Println(err)
					}
				}(alertFunc)
			}
			wg.Wait()
		} else {
			fmt.Println(fmt.Sprintf("Using %s", *powerSource))
		}
	}
}
