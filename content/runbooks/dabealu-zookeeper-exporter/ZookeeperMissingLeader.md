---
title: ZookeeperMissingLeader
description: Troubleshooting for alert ZookeeperMissingLeader
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZookeeperMissingLeader

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Zookeeper cluster has no node marked as leader

<details>
  <summary>Alert Rule</summary>

{{% rule "zookeeper/dabealu-zookeeper-exporter.yml" "ZookeeperMissingLeader" %}}

<!-- Rule when generated

```yaml
alert: ZookeeperMissingLeader
expr: sum(zk_server_leader) == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Zookeeper missing leader (instance {{ $labels.instance }})
    description: |-
        Zookeeper cluster has no node marked as leader
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dabealu-zookeeper-exporter/ZookeeperMissingLeader.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
