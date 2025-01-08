---
title: HadoopYarnContainerAllocationFailures
description: Troubleshooting for alert HadoopYarnContainerAllocationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopYarnContainerAllocationFailures

There is a significant number of YARN container allocation failures.

<details>
  <summary>Alert Rule</summary>

{{% rule "hadoop/jmx_exporter.yml" "HadoopYarnContainerAllocationFailures" %}}

{{% comment %}}

```yaml
alert: HadoopYarnContainerAllocationFailures
expr: hadoop_yarn_container_allocation_failures_total > 10
for: 10m
labels:
    severity: warning
annotations:
    summary: Hadoop YARN Container Allocation Failures (instance {{ $labels.instance }})
    description: |-
        There is a significant number of YARN container allocation failures.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/jmx_exporter/HadoopYarnContainerAllocationFailures.md

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
