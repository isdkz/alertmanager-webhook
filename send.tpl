------------------------------
     中远测试环境告警来了
------------------------------
状态: {{.Status}}
告警名字: {{.Labels.alertname}}
告警实例: {{.Labels.instance}}
告警等级: {{.Labels.severity}}
告警描述: {{.Annotations.description}}
开始时间: {{.StartsAt.Format "2006-01-02 15:04:05"}}
结束时间: {{.EndsAt.Format "2006-01-02 15:04:05"}}
------------------------------
