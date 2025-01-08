---
title: PatroniHasNoLeader
description: Troubleshooting for alert PatroniHasNoLeader
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PatroniHasNoLeader

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A leader node (neither primary nor standby) cannot be found inside the cluster {{ $labels.scope }}

<details>
  <summary>Alert Rule</summary>

{{% rule "patroni/embedded-exporter-patroni.yml" "PatroniHasNoLeader" %}}

{{% comment %}}

```yaml
alert: PatroniHasNoLeader
expr: (max by (scope) (patroni_master) < 1) and (max by (scope) (patroni_standby_leader) < 1)
for: 0m
labels:
    severity: critical
annotations:
    summary: Patroni has no Leader (instance {{ $labels.instance }})
    description: |-
        A leader node (neither primary nor standby) cannot be found inside the cluster {{ $labels.scope }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-patroni/PatroniHasNoLeader.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
