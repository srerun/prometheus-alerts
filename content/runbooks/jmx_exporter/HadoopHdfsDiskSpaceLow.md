---
title: HadoopHdfsDiskSpaceLow
description: Troubleshooting for alert HadoopHdfsDiskSpaceLow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopHdfsDiskSpaceLow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Available HDFS disk space is running low.

<details>
  <summary>Alert Rule</summary>

{{% rule "hadoop/jmx_exporter.yml" "HadoopHdfsDiskSpaceLow" %}}

{{% comment %}}

```yaml
alert: HadoopHdfsDiskSpaceLow
expr: (hadoop_hdfs_bytes_total - hadoop_hdfs_bytes_used) / hadoop_hdfs_bytes_total < 0.1
for: 15m
labels:
    severity: warning
annotations:
    summary: Hadoop HDFS Disk Space Low (instance {{ $labels.instance }})
    description: |-
        Available HDFS disk space is running low.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/jmx_exporter/HadoopHdfsDiskSpaceLow.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
