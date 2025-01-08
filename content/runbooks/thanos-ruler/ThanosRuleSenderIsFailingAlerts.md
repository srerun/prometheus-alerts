---
title: ThanosRuleSenderIsFailingAlerts
description: Troubleshooting for alert ThanosRuleSenderIsFailingAlerts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleSenderIsFailingAlerts

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Rule {{$labels.instance}} is failing to send alerts to alertmanager.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleSenderIsFailingAlerts" %}}

<!-- Rule when generated

```yaml
alert: ThanosRuleSenderIsFailingAlerts
expr: sum by (job, instance) (rate(thanos_alert_sender_alerts_dropped_total{job=~".*thanos-rule.*"}[5m])) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Rule Sender Is Failing Alerts (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.instance}} is failing to send alerts to alertmanager.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleSenderIsFailingAlerts.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
