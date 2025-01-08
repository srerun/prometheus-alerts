---
title: BlackboxProbeSlowHttp
description: Troubleshooting for alert BlackboxProbeSlowHttp
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxProbeSlowHttp

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HTTP request took more than 1s

<details>
  <summary>Alert Rule</summary>

{{% rule "blackbox/blackbox-exporter.yml" "BlackboxProbeSlowHttp" %}}

{{% comment %}}

```yaml
alert: BlackboxProbeSlowHttp
expr: avg_over_time(probe_http_duration_seconds[1m]) > 1
for: 1m
labels:
    severity: warning
annotations:
    summary: Blackbox probe slow HTTP (instance {{ $labels.instance }})
    description: |-
        HTTP request took more than 1s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/blackbox-exporter/BlackboxProbeSlowHttp.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
