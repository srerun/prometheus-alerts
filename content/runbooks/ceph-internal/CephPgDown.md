---
title: CephPgDown
description: Troubleshooting for alert CephPgDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephPgDown

Some Ceph placement groups are down. Please ensure that all the data are available.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephPgDown" %}}

{{% comment %}}

```yaml
alert: CephPgDown
expr: ceph_pg_down > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Ceph PG down (instance {{ $labels.instance }})
    description: |-
        Some Ceph placement groups are down. Please ensure that all the data are available.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephPgDown.md

```

{{% /comment %}}

</details>


## Meaning

The CephPgDown alert is triggered when one or more Ceph placement groups (PGs) are in a down state. This indicates that data availability is at risk, as PGs are responsible for storing and managing data in a Ceph cluster.

## Impact

The impact of this alert is high, as it can lead to data unavailability and potential data loss. If left unaddressed, this issue can cause:

* Data inaccessibility
* Application downtime
* Data loss or corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Ceph cluster's overall health using the Ceph dashboard or command-line tools.
2. Identify the specific PGs that are down using the Ceph command `ceph pg dump`.
3. Check the Ceph node's system logs for any error messages related to the down PGs.
4. Verify that all Ceph nodes are running and reachable.
5. Check the network connectivity between Ceph nodes.

Look for common causes of PG downtime, such as:

* Node failures or restarts
* Network connectivity issues
* Disk failures or corruption
* Configuration errors

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the root cause of the PG downtime and address it accordingly.
2. If a node is down, restart it or replace it if necessary.
3. If network connectivity is the issue, restore connectivity between nodes.
4. If disk failures or corruption are the cause, replace the faulty disks and run `ceph pg repair` to recover data.
5. If configuration errors are the cause, correct the configuration and restart the affected nodes.
6. Once the root cause is addressed, run `ceph pg repair` to recover the down PGs.
7. Monitor the Ceph cluster's health and verify that all PGs are up and running.
8. Perform a data scrub to ensure data integrity.

Remember to follow proper change management procedures and communicate with stakeholders before making any changes to the Ceph cluster.