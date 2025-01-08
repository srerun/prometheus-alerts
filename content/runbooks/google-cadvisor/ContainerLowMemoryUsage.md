---
title: ContainerLowMemoryUsage
description: Troubleshooting for alert ContainerLowMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerLowMemoryUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Container Memory usage is under 20% for 1 week. Consider reducing the allocated memory.

<details>
  <summary>Alert Rule</summary>

{{% rule "docker-containers/google-cadvisor.yml" "ContainerLowMemoryUsage" %}}

<!-- Rule when generated

```yaml
alert: ContainerLowMemoryUsage
expr: (sum(container_memory_working_set_bytes{name!=""}) BY (instance, name) / sum(container_spec_memory_limit_bytes > 0) BY (instance, name) * 100) < 20
for: 7d
labels:
    severity: info
annotations:
    summary: Container Low Memory usage (instance {{ $labels.instance }})
    description: |-
        Container Memory usage is under 20% for 1 week. Consider reducing the allocated memory.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerLowMemoryUsage.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
