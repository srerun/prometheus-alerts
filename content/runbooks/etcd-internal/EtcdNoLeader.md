---
title: EtcdNoLeader
description: Troubleshooting for alert EtcdNoLeader
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdNoLeader

Etcd cluster have no leader

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdNoLeader" %}}

{{% comment %}}

```yaml
alert: EtcdNoLeader
expr: etcd_server_has_leader == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Etcd no Leader (instance {{ $labels.instance }})
    description: |-
        Etcd cluster have no leader
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdNoLeader.md

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
