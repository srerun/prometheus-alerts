---
title: EtcdHighNumberOfLeaderChanges
description: Troubleshooting for alert EtcdHighNumberOfLeaderChanges
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighNumberOfLeaderChanges

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Etcd leader changed more than 2 times during 10 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighNumberOfLeaderChanges" %}}

{{% comment %}}

```yaml
alert: EtcdHighNumberOfLeaderChanges
expr: increase(etcd_server_leader_changes_seen_total[10m]) > 2
for: 0m
labels:
    severity: warning
annotations:
    summary: Etcd high number of leader changes (instance {{ $labels.instance }})
    description: |-
        Etcd leader changed more than 2 times during 10 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighNumberOfLeaderChanges.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
