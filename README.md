## Install
```shell
go install github.com/huantt/mac-keeper@latest
```
```shell
export PATH=$PATH:$(go env GOPATH)/bin
```

## Power monitor
Alerts to "Alerting channels" when your macbook is unplugged.
Run it, then feel free to go to the WC in the coffee shop!

### Supported alerting channels:
- [x] Sound `[default]`
- [x] Slack

### Command
```shell
Usage:
   mac-keeper monitor-power [flags]

Flags:
  -h, --help                       help for monitor-power
      --slack-channel string       Slack channel
      --slack-webhook-url string   Slack webhook URL
      --sos-sound string           Path to sound file (default "data/sos_sound.mp3")
```

#### Sample

```shell
mac-keeper monitor-power \
--slack-webhook-url='https://hooks.slack.com/services/X/Y/Z' \
--slack-channel='#JackTT'
```