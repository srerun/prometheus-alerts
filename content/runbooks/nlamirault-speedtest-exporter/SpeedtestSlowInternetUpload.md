---
title: SpeedtestSlowInternetUpload
description: Troubleshooting for alert SpeedtestSlowInternetUpload
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SpeedtestSlowInternetUpload

Internet upload speed is currently {{humanize $value}} Mbps.

<details>
  <summary>Alert Rule</summary>

{{% rule "speedtest/nlamirault-speedtest-exporter.yml" "SpeedtestSlowInternetUpload" %}}

{{% comment %}}

```yaml
alert: SpeedtestSlowInternetUpload
expr: avg_over_time(speedtest_upload[10m]) < 20
for: 0m
labels:
    severity: warning
annotations:
    summary: SpeedTest Slow Internet Upload (instance {{ $labels.instance }})
    description: |-
        Internet upload speed is currently {{humanize $value}} Mbps.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nlamirault-speedtest-exporter/SpeedtestSlowInternetUpload.md

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
