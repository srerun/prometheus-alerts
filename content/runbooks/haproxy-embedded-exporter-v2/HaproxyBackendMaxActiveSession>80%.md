---
title: HaproxyBackendMaxActiveSession>80%
description: Troubleshooting for alert HaproxyBackendMaxActiveSession>80%
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyBackendMaxActiveSession>80%

Session limit from backend {{ $labels.proxy }} to server {{ $labels.server }} reached 80% of limit - {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-embedded-exporter-v2.yml" "HaproxyBackendMaxActiveSession>80%" %}}

{{% comment %}}

```yaml
alert: HaproxyBackendMaxActiveSession>80%
expr: ((haproxy_server_max_sessions >0) * 100) / (haproxy_server_limit_sessions > 0) > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: HAProxy backend max active session > 80% (instance {{ $labels.instance }})
    description: |-
        Session limit from backend {{ $labels.proxy }} to server {{ $labels.server }} reached 80% of limit - {{ $value | printf "%.2f"}}%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/HaproxyBackendMaxActiveSession>80%.md

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
