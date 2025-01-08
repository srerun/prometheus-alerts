---
title: CephPgActivationLong
description: Troubleshooting for alert CephPgActivationLong
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephPgActivationLong

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Ceph placement groups are too long to activate.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephPgActivationLong" %}}

<!-- Rule when generated

```yaml
alert: CephPgActivationLong
expr: ceph_pg_activating > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph PG activation long (instance {{ $labels.instance }})
    description: |-
        Some Ceph placement groups are too long to activate.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephPgActivationLong.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
