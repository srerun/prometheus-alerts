---
title: EtcdInsufficientMembers
description: Troubleshooting for alert EtcdInsufficientMembers
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdInsufficientMembers

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Etcd cluster should have an odd number of members

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdInsufficientMembers" %}}

{{% comment %}}

```yaml
alert: EtcdInsufficientMembers
expr: count(etcd_server_id) % 2 == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Etcd insufficient Members (instance {{ $labels.instance }})
    description: |-
        Etcd cluster should have an odd number of members
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdInsufficientMembers.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
