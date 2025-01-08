---
title: CephPgInconsistent
description: Troubleshooting for alert CephPgInconsistent
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephPgInconsistent

Some Ceph placement groups are inconsistent. Data is available but inconsistent across nodes.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephPgInconsistent" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephPgInconsistent.md

```

{{% /comment %}}

</details>


Here is a runbook for the CephPgInconsistent alert rule:

## Meaning

The CephPgInconsistent alert is triggered when there are inconsistent placement groups (PGs) in the Ceph cluster. This means that data is available, but it's inconsistent across nodes. This can lead to data discrepancies and potential data loss.

## Impact

The impact of this alert is moderate to high, as it can lead to:

* Data inconsistencies across nodes
* Potential data loss
* Performance degradation
* Increased risk of cluster instability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Ceph cluster's health using the `ceph health` command.
2. Identify the PGs that are inconsistent using the `ceph pg dump` command.
3. Investigate the nodes that are reporting inconsistencies.
4. Check the Ceph logs for any error messages related to the inconsistent PGs.
5. Verify that the network connectivity between nodes is stable and not experiencing any issues.

## Mitigation

To mitigate the issue, follow these steps:

1. Run `ceph pg repair` to repair the inconsistent PGs.
2. Verify that the repair was successful by running `ceph pg dump` again.
3. If the issue persists, consider rebalancing the PGs using `ceph pg rebalance`.
4. Check the Ceph cluster's configuration to ensure it's properly set up for data replication and consistency.
5. Consider adding more nodes or increasing the replication factor to improve data durability and reduce the risk of inconsistencies.

Remember to monitor the Ceph cluster's health and performance closely after mitigation to ensure the issue is fully resolved.