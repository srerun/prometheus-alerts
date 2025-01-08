---
title: HostClockNotSynchronising
description: Troubleshooting for alert HostClockNotSynchronising
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostClockNotSynchronising

Clock not synchronising. Ensure NTP is configured on this host.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostClockNotSynchronising" %}}

{{% comment %}}

```yaml
alert: HostClockNotSynchronising
expr: (min_over_time(node_timex_sync_status[1m]) == 0 and node_timex_maxerror_seconds >= 16) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host clock not synchronising (instance {{ $labels.instance }})
    description: |-
        Clock not synchronising. Ensure NTP is configured on this host.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostClockNotSynchronising.md

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
