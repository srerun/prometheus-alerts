---
title: HostNodeOvertemperatureAlarm
description: Troubleshooting for alert HostNodeOvertemperatureAlarm
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostNodeOvertemperatureAlarm

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Physical node temperature alarm triggered

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostNodeOvertemperatureAlarm
expr: ((node_hwmon_temp_crit_alarm_celsius == 1) or (node_hwmon_temp_alarm == 1)) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: critical
annotations:
    summary: Host node overtemperature alarm (instance {{ $labels.instance }})
    description: |-
        Physical node temperature alarm triggered
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostNodeOvertemperatureAlarm

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
