---
title: CephPgUnavailable
description: Troubleshooting for alert CephPgUnavailable
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephPgUnavailable

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Ceph placement groups are unavailable.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephPgUnavailable" %}}

{{% comment %}}

```yaml
alert: CephPgUnavailable
expr: ceph_pg_total - ceph_pg_active > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Ceph PG unavailable (instance {{ $labels.instance }})
    description: |-
        Some Ceph placement groups are unavailable.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephPgUnavailable.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
