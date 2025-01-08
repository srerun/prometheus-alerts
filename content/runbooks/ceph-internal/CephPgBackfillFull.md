---
title: CephPgBackfillFull
description: Troubleshooting for alert CephPgBackfillFull
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephPgBackfillFull

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Ceph placement groups are located on full Object Storage Daemon on cluster. Those PGs can be unavailable shortly. Please check OSDs, change weight or reconfigure CRUSH rules.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephPgBackfillFull" %}}

{{% comment %}}

```yaml
alert: CephPgBackfillFull
expr: ceph_pg_backfill_toofull > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph PG backfill full (instance {{ $labels.instance }})
    description: |-
        Some Ceph placement groups are located on full Object Storage Daemon on cluster. Those PGs can be unavailable shortly. Please check OSDs, change weight or reconfigure CRUSH rules.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephPgBackfillFull.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
