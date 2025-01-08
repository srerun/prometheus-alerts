---
title: ApacheDown
description: Troubleshooting for alert ApacheDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApacheDown

Apache down

<details>
  <summary>Alert Rule</summary>

{{% rule "apache/lusitaniae-apache-exporter.yml" "ApacheDown" %}}

{{% comment %}}

```yaml
alert: ApacheDown
expr: apache_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Apache down (instance {{ $labels.instance }})
    description: |-
        Apache down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/lusitaniae-apache-exporter/ApacheDown.md

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
