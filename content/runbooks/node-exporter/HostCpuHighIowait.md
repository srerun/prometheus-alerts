---
title: HostCpuHighIowait
description: Troubleshooting for alert HostCpuHighIowait
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostCpuHighIowait

## Meaning
[//]: # "Short paragraph that explains what the alert means"
CPU iowait > 10%. A high iowait means that you are disk or network bound.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostCpuHighIowait" %}}

{{% comment %}}

```yaml
alert: HostCpuHighIowait
expr: (avg by (instance) (rate(node_cpu_seconds_total{mode="iowait"}[5m])) * 100 > 10) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host CPU high iowait (instance {{ $labels.instance }})
    description: |-
        CPU iowait > 10%. A high iowait means that you are disk or network bound.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostCpuHighIowait.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
