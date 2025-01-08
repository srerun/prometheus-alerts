---
title: CephMonitorClockSkew
description: Troubleshooting for alert CephMonitorClockSkew
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephMonitorClockSkew

Ceph monitor clock skew detected. Please check ntp and hardware clock settings

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephMonitorClockSkew" %}}

{{% comment %}}

```yaml
alert: CephMonitorClockSkew
expr: abs(ceph_monitor_clock_skew_seconds) > 0.2
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph monitor clock skew (instance {{ $labels.instance }})
    description: |-
        Ceph monitor clock skew detected. Please check ntp and hardware clock settings
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephMonitorClockSkew.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the CephMonitorClockSkew alert:

## Meaning

The CephMonitorClockSkew alert is triggered when the clock skew between Ceph monitors exceeds 0.2 seconds. This can cause issues with data consistency and cluster stability.

## Impact

If left unaddressed, clock skew can lead to:

* Inconsistent data replication
* Cluster instability
* Performance degradation
* Potential data loss

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NTP (Network Time Protocol) configuration and make sure it is correctly set up and synchronized across all nodes in the cluster.
2. Verify the hardware clock settings on each node to ensure they are accurate and consistent.
3. Review the Ceph monitor logs for any errors or warnings related to clock skew.
4. Use the `ceph mon dump` command to check the current clock skew values.

## Mitigation

To mitigate the issue, follow these steps:

1. Adjust the NTP configuration to ensure accurate time synchronization across all nodes in the cluster.
2. Correct any hardware clock settings that are found to be inaccurate or inconsistent.
3. Restart the Ceph monitor service on each node to ensure that the new clock settings take effect.
4. Monitor the cluster for any further issues related to clock skew and adjust the NTP configuration and hardware clock settings as needed.

Additional resources:

* [Ceph documentation on clock synchronization](https://docs.ceph.com/en/latest/rados/troubleshooting/clock-skew/)
* [NTP configuration guide](https://www.ntp.org/ntpfaq/)