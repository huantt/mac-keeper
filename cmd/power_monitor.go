package cmd

import (
	"fmt"
	"github.com/huantt/mac-keeper/impl/alerter"
	"github.com/huantt/mac-keeper/monitor"
	"github.com/spf13/cobra"
)

const defaultSOSSound = "data/sos_sound.mp3"

func MonitorPowerCommand() *cobra.Command {
	var sosSound string
	var slackWebhookURL string
	var slackChannel string

	cmd := &cobra.Command{
		Use: "monitor-power",
		Run: func(cmd *cobra.Command, args []string) {
			powerMonitor := monitor.NewPowerMonitor()
			fmt.Println(fmt.Sprintf("Using sos sound: %s", sosSound))
			soundPlayer := alerter.NewSoundPlayer(sosSound)
			if slackWebhookURL != "" && slackChannel != "" {
				fmt.Println(fmt.Sprintf("Added Slack alerting channel: %s", slackChannel))
				slackService := alerter.NewSlackService(slackWebhookURL, slackChannel)
				powerMonitor.AddAlertFunc(slackService.Alert)
			}
			powerMonitor.AddAlertFunc(soundPlayer.Alert)
			powerMonitor.Run()
		},
	}

	cmd.Flags().StringVar(&sosSound, "sos-sound", defaultSOSSound, "Path to sound file")
	cmd.Flags().StringVar(&slackWebhookURL, "slack-webhook-url", "", "Slack webhook URL")
	cmd.Flags().StringVar(&slackChannel, "slack-channel", "", "Slack channel")
	cmd.MarkFlagsRequiredTogether("slack-webhook-url", "slack-channel")
	return cmd
}
