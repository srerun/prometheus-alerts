---
title: HaproxyBackendDown
description: Troubleshooting for alert HaproxyBackendDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyBackendDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HAProxy backend is down

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyBackendDown" %}}

{{% comment %}}

```yaml
alert: HaproxyBackendDown
expr: haproxy_backend_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: HAProxy backend down (instance {{ $labels.instance }})
    description: |-
        HAProxy backend is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyBackendDown.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
