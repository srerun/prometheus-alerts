---
title: PrometheusNotificationsBacklog
description: Troubleshooting for alert PrometheusNotificationsBacklog
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusNotificationsBacklog

The Prometheus notification queue has not been empty for 10 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusNotificationsBacklog" %}}

{{% comment %}}

```yaml
alert: PrometheusNotificationsBacklog
expr: min_over_time(prometheus_notifications_queue_length[10m]) > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus notifications backlog (instance {{ $labels.instance }})
    description: |-
        The Prometheus notification queue has not been empty for 10 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusNotificationsBacklog.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
