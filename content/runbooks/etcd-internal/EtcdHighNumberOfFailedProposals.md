---
title: EtcdHighNumberOfFailedProposals
description: Troubleshooting for alert EtcdHighNumberOfFailedProposals
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighNumberOfFailedProposals

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Etcd server got more than 5 failed proposals past hour

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighNumberOfFailedProposals" %}}

{{% comment %}}

```yaml
alert: EtcdHighNumberOfFailedProposals
expr: increase(etcd_server_proposals_failed_total[1h]) > 5
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high number of failed proposals (instance {{ $labels.instance }})
    description: |-
        Etcd server got more than 5 failed proposals past hour
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighNumberOfFailedProposals.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
