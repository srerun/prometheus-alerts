---
title: HostConntrackLimit
description: Troubleshooting for alert HostConntrackLimit
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostConntrackLimit

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The number of conntrack is approaching limit

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostConntrackLimit
expr: (node_nf_conntrack_entries / node_nf_conntrack_entries_limit > 0.8) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 5m
labels:
    severity: warning
annotations:
    summary: Host conntrack limit (instance {{ $labels.instance }})
    description: |-
        The number of conntrack is approaching limit
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostConntrackLimit

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
