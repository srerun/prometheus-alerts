---
title: ApacheRestart
description: Troubleshooting for alert ApacheRestart
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApacheRestart

Apache has just been restarted.

<details>
  <summary>Alert Rule</summary>

{{% rule "apache/lusitaniae-apache-exporter.yml" "ApacheRestart" %}}

{{% comment %}}

```yaml
alert: ApacheRestart
expr: apache_uptime_seconds_total / 60 < 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Apache restart (instance {{ $labels.instance }})
    description: |-
        Apache has just been restarted.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/lusitaniae-apache-exporter/ApacheRestart.md

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
