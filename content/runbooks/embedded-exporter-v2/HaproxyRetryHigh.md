---
title: HaproxyRetryHigh
description: Troubleshooting for alert HaproxyRetryHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyRetryHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
High rate of retry on {{ $labels.proxy }} - {{ $value | printf "%.2f"}}

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/embedded-exporter-v2.yml" "HaproxyRetryHigh" %}}

{{% comment %}}

```yaml
alert: HaproxyRetryHigh
expr: sum by (proxy) (rate(haproxy_backend_retry_warnings_total[1m])) > 10
for: 2m
labels:
    severity: warning
annotations:
    summary: HAProxy retry high (instance {{ $labels.instance }})
    description: |-
        High rate of retry on {{ $labels.proxy }} - {{ $value | printf "%.2f"}}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/HaproxyRetryHigh.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
