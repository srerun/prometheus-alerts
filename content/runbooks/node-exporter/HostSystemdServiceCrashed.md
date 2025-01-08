---
title: HostSystemdServiceCrashed
description: Troubleshooting for alert HostSystemdServiceCrashed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostSystemdServiceCrashed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
systemd service crashed

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostSystemdServiceCrashed" %}}

<!-- Rule when generated

```yaml
alert: HostSystemdServiceCrashed
expr: (node_systemd_unit_state{state="failed"} == 1) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host systemd service crashed (instance {{ $labels.instance }})
    description: |-
        systemd service crashed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostSystemdServiceCrashed.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
