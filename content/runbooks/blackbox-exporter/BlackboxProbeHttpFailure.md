---
title: BlackboxProbeHttpFailure
description: Troubleshooting for alert BlackboxProbeHttpFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxProbeHttpFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HTTP status code is not 200-399

<details>
  <summary>Alert Rule</summary>

{{% rule "blackbox/blackbox-exporter.yml" "BlackboxProbeHttpFailure" %}}

<!-- Rule when generated

```yaml
alert: BlackboxProbeHttpFailure
expr: probe_http_status_code <= 199 OR probe_http_status_code >= 400
for: 0m
labels:
    severity: critical
annotations:
    summary: Blackbox probe HTTP failure (instance {{ $labels.instance }})
    description: |-
        HTTP status code is not 200-399
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/blackbox-exporter/BlackboxProbeHttpFailure.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
