---
title: HostPhysicalComponentTooHot
description: Troubleshooting for alert HostPhysicalComponentTooHot
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostPhysicalComponentTooHot

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Physical hardware component too hot

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostPhysicalComponentTooHot" %}}

{{% comment %}}

```yaml
alert: HostPhysicalComponentTooHot
expr: ((node_hwmon_temp_celsius * ignoring(label) group_left(instance, job, node, sensor) node_hwmon_sensor_label{label!="tctl"} > 75)) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 5m
labels:
    severity: warning
annotations:
    summary: Host physical component too hot (instance {{ $labels.instance }})
    description: |-
        Physical hardware component too hot
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostPhysicalComponentTooHot.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
