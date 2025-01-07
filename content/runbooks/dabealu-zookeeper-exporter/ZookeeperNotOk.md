---
title: ZookeeperNotOk
description: Troubleshooting for alert ZookeeperNotOk
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZookeeperNotOk

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Zookeeper instance is not ok

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ZookeeperNotOk
expr: zk_ruok == 0
for: 3m
labels:
    severity: warning
annotations:
    summary: Zookeeper Not Ok (instance {{ $labels.instance }})
    description: |-
        Zookeeper instance is not ok
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ZookeeperNotOk

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
