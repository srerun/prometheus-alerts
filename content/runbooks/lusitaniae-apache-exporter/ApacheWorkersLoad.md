---
title: ApacheWorkersLoad
description: Troubleshooting for alert ApacheWorkersLoad
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApacheWorkersLoad

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Apache workers in busy state approach the max workers count 80% workers busy on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "apache/lusitaniae-apache-exporter.yml" "ApacheWorkersLoad" %}}

{{% comment %}}

```yaml
alert: ApacheWorkersLoad
expr: (sum by (instance) (apache_workers{state="busy"}) / sum by (instance) (apache_scoreboard) ) * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: Apache workers load (instance {{ $labels.instance }})
    description: |-
        Apache workers in busy state approach the max workers count 80% workers busy on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/lusitaniae-apache-exporter/ApacheWorkersLoad.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
