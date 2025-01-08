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


Here is a sample runbook for the CephPgUnavailable alert:

## Meaning

The CephPgUnavailable alert indicates that one or more Ceph placement groups (PGs) are currently unavailable. This is a critical alert as it can lead to data unavailability and potential data loss.

## Impact

The impact of this alert is high, as it affects the availability of data stored in Ceph. If left unresolved, this issue can lead to:

* Data unavailability
* Data loss
* Downtime for critical systems and applications
* Revenue loss due to system unavailability

## Diagnosis

To diagnose the root cause of this issue, follow these steps:

1. Check the Ceph cluster status using `ceph -s` or `ceph status` command.
2. Verify the PG status using `ceph pg dump` command.
3. Check the Ceph logs for any errors or warnings related to PG unavailability.
4. Identify the specific PGs that are unavailable using `ceph pg ls` command.
5. Verify the health of the Ceph nodes and OSDs using `ceph osd ls` and `ceph osd df` commands.

## Mitigation

To mitigate this issue, follow these steps:

1. Identify the root cause of the PG unavailability (e.g. node failure, network issues, disk errors).
2. Resolve the underlying issue (e.g. replace failed node, fix network issues, repair disk errors).
3. Run `ceph pg repair` command to repair the affected PGs.
4. Verify the PG status using `ceph pg dump` command to ensure that all PGs are now available.
5. Monitor the Ceph cluster status and PG status to ensure that the issue is resolved and does not reoccur.

Remember to also update the alert status in Prometheus to reflect the resolution of the issue.