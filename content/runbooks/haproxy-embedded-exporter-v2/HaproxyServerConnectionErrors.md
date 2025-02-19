---
title: HaproxyServerConnectionErrors
description: Troubleshooting for alert HaproxyServerConnectionErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyServerConnectionErrors

Too many connection errors to {{ $labels.server }} server (> 100 req/s). Request throughput may be too high.

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-embedded-exporter-v2.yml" "HaproxyServerConnectionErrors" %}}

{{% comment %}}

```yaml
alert: HaproxyServerConnectionErrors
expr: (sum by (proxy) (rate(haproxy_server_connection_errors_total[1m]))) > 100
for: 0m
labels:
    severity: critical
annotations:
    summary: HAProxy server connection errors (instance {{ $labels.instance }})
    description: |-
        Too many connection errors to {{ $labels.server }} server (> 100 req/s). Request throughput may be too high.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/HaproxyServerConnectionErrors.md

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
