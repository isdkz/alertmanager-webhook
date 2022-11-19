# This project references [alertmanaer-dingtalk-webhook](https://github.com/yunlzheng/alertmanaer-dingtalk-webhook) 
## Alertmanager Webhook
Webhook service support send Prometheus 2 alert message to GroupRobot.

## How To Use

```
cd cmd/webhook
go build --ldflags="-w -s" webhook.go
webhook -u=your_robot_url -p webhook_port(default 8080) -f template_file_path(default send.tpl) 
```

```
go run webhook -u=your_robot_url -p webhook_port(default 8080) -f template_file_path(default send.tpl)
```

### use this commond to see the help
```
webhook -h
```

* -u: default webhook url, all notifaction from alertmanager will direct to this webhook address.
* -p: the webhook port, webhook server will run in(the default is 8080)
* -f: the template file path, the webhook server will use this message template to send the message(the default is send.tpl)
* -t: type of the push message, text and md are supprted(the default is text)

Or you can overwrite by add annotations to Prometheus alertrule to special the webhook for each alert rule.

```
groups:
- name: hostStatsAlert
  rules:
  - alert: hostCpuUsageAlert
    expr: sum(avg without (cpu)(irate(node_cpu{mode!='idle'}[5m]))) by (instance) > 0.85
    for: 1m
    labels:
      severity: page
    annotations:
      summary: "Instance {{ $labels.instance }} CPU usgae high"
      description: "{{ $labels.instance }} CPU usage above 85% (current value: {{ $value }})"
      Robot: "xxx"
```
