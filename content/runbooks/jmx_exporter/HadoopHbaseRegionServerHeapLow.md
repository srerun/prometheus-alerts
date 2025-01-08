---
title: HadoopHbaseRegionServerHeapLow
description: Troubleshooting for alert HadoopHbaseRegionServerHeapLow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopHbaseRegionServerHeapLow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HBase Region Servers are running low on heap space.

<details>
  <summary>Alert Rule</summary>

{{% rule "hadoop/jmx_exporter.yml" "HadoopHbaseRegionServerHeapLow" %}}

{{% comment %}}

```yaml
alert: HadoopHbaseRegionServerHeapLow
expr: hadoop_hbase_region_server_heap_bytes / hadoop_hbase_region_server_max_heap_bytes < 0.2
for: 10m
labels:
    severity: critical
annotations:
    summary: Hadoop HBase Region Server Heap Low (instance {{ $labels.instance }})
    description: |-
        HBase Region Servers are running low on heap space.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/jmx_exporter/HadoopHbaseRegionServerHeapLow.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
