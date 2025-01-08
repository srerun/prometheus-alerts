---
title: ZookeeperDown
description: Troubleshooting for alert ZookeeperDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZookeeperDown

Zookeeper down on instance {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "zookeeper/dabealu-zookeeper-exporter.yml" "ZookeeperDown" %}}

{{% comment %}}

```yaml
alert: ZookeeperDown
expr: zk_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Zookeeper Down (instance {{ $labels.instance }})
    description: |-
        Zookeeper down on instance {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dabealu-zookeeper-exporter/ZookeeperDown.md

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
