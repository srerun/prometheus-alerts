---
title: HostClockSkew
description: Troubleshooting for alert HostClockSkew
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostClockSkew

Clock skew detected. Clock is out of sync. Ensure NTP is configured correctly on this host.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostClockSkew" %}}

{{% comment %}}

```yaml
alert: HostClockSkew
expr: ((node_timex_offset_seconds > 0.05 and deriv(node_timex_offset_seconds[5m]) >= 0) or (node_timex_offset_seconds < -0.05 and deriv(node_timex_offset_seconds[5m]) <= 0)) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 10m
labels:
    severity: warning
annotations:
    summary: Host clock skew (instance {{ $labels.instance }})
    description: |-
        Clock skew detected. Clock is out of sync. Ensure NTP is configured correctly on this host.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostClockSkew.md

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
