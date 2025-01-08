---
title: CephPgIncomplete
description: Troubleshooting for alert CephPgIncomplete
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephPgIncomplete

Some Ceph placement groups are incomplete. Please ensure that all the data are available.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephPgIncomplete" %}}

{{% comment %}}

```yaml
alert: CephPgIncomplete
expr: ceph_pg_incomplete > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Ceph PG incomplete (instance {{ $labels.instance }})
    description: |-
        Some Ceph placement groups are incomplete. Please ensure that all the data are available.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephPgIncomplete.md

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
