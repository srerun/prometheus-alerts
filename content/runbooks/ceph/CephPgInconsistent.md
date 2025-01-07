---
title: CephPgInconsistent
description: Troubleshooting for alert CephPgInconsistent
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephPgInconsistent

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Ceph placement groups are inconsistent. Data is available but inconsistent across nodes.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CephPgInconsistent
expr: ceph_pg_inconsistent > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Ceph PG inconsistent (instance {{ $labels.instance }})
    description: |-
        Some Ceph placement groups are inconsistent. Data is available but inconsistent across nodes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/CephPgInconsistent

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
