---
title: NetdataMdMismatchCntUnsynchronizedBlocks
description: Troubleshooting for alert NetdataMdMismatchCntUnsynchronizedBlocks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataMdMismatchCntUnsynchronizedBlocks

RAID Array have unsynchronized blocks

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataMdMismatchCntUnsynchronizedBlocks" %}}

{{% comment %}}

```yaml
alert: NetdataMdMismatchCntUnsynchronizedBlocks
expr: netdata_md_mismatch_cnt_unsynchronized_blocks_average > 1024
for: 2m
labels:
    severity: warning
annotations:
    summary: Netdata MD mismatch cnt unsynchronized blocks (instance {{ $labels.instance }})
    description: |-
        RAID Array have unsynchronized blocks
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataMdMismatchCntUnsynchronizedBlocks.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `NetdataMdMismatchCntUnsynchronizedBlocks`:

## Meaning

The `NetdataMdMismatchCntUnsynchronizedBlocks` alert is triggered when the average number of unsynchronized blocks in the Netdata metadata (MD) exceeds 1024 for a period of 2 minutes. This indicates that there are inconsistencies in the RAID array, which can lead to data corruption or loss.

## Impact

If left unaddressed, this issue can result in:

* Data corruption or loss
* RAID array degradation or failure
* System downtime or instability
* Data availability and integrity issues

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Netdata dashboard to verify the inconsistency in the MD metadata.
2. Review the RAID array configuration and ensure that all disks are properly connected and configured.
3. Check the system logs for any errors or warnings related to the RAID array or disk errors.
4. Run a thorough disk check and RAID array validation to identify any inconsistencies.
5. Verify that the Netdata agent is properly configured and running correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately stop any write operations to the RAID array to prevent further data corruption.
2. Run a RAID array resynchronization or rebuild operation to fix the inconsistencies.
3. Verify that the Netdata agent is properly configured and running correctly.
4. Implement redundancy and backup mechanisms to ensure data availability and integrity.
5. Schedule regular disk checks and RAID array validations to prevent future occurrences.

Note: The steps above are general guidelines and may vary depending on the specific system configuration and setup. It is recommended to consult the system administrators and experts for specific guidance on resolving the issue.