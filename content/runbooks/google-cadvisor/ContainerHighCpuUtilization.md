---
title: ContainerHighCpuUtilization
description: Troubleshooting for alert ContainerHighCpuUtilization
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerHighCpuUtilization

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Container CPU utilization is above 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "docker-containers/google-cadvisor.yml" "ContainerHighCpuUtilization" %}}

{{% comment %}}

```yaml
alert: ContainerHighCpuUtilization
expr: (sum(rate(container_cpu_usage_seconds_total{container!=""}[5m])) by (pod, container) / sum(container_spec_cpu_quota{container!=""}/container_spec_cpu_period{container!=""}) by (pod, container) * 100) > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: Container High CPU utilization (instance {{ $labels.instance }})
    description: |-
        Container CPU utilization is above 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerHighCpuUtilization.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
