---
title: HaproxyServerHealthcheckFailure
description: Troubleshooting for alert HaproxyServerHealthcheckFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyServerHealthcheckFailure

Some server healthcheck are failing on {{ $labels.server }}

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/embedded-exporter-v2.yml" "HaproxyServerHealthcheckFailure" %}}

{{% comment %}}

```yaml
alert: HaproxyServerHealthcheckFailure
expr: increase(haproxy_server_check_failures_total[1m]) > 0
for: 1m
labels:
    severity: warning
annotations:
    summary: HAProxy server healthcheck failure (instance {{ $labels.instance }})
    description: |-
        Some server healthcheck are failing on {{ $labels.server }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/HaproxyServerHealthcheckFailure.md

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
