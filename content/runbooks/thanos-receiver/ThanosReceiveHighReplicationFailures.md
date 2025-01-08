---
title: ThanosReceiveHighReplicationFailures
description: Troubleshooting for alert ThanosReceiveHighReplicationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHighReplicationFailures

Thanos Receive {{$labels.job}} is failing to replicate {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveHighReplicationFailures" %}}

{{% comment %}}

```yaml
alert: ThanosReceiveHighReplicationFailures
expr: thanos_receive_replication_factor > 1 and ((sum by (job) (rate(thanos_receive_replications_total{result="error", job=~".*thanos-receive.*"}[5m])) / sum by (job) (rate(thanos_receive_replications_total{job=~".*thanos-receive.*"}[5m]))) > (max by (job) (floor((thanos_receive_replication_factor{job=~".*thanos-receive.*"}+1)/ 2)) / max by (job) (thanos_receive_hashring_nodes{job=~".*thanos-receive.*"}))) * 100
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Receive High Replication Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} is failing to replicate {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveHighReplicationFailures.md

```

{{% /comment %}}

</details>


Here is the runbook for the ThanosReceiveHighReplicationFailures alert:

## Meaning

The ThanosReceiveHighReplicationFailures alert is triggered when a Thanos Receive instance is experiencing a high rate of replication failures. This alert indicates that the instance is failing to replicate a significant percentage of requests, which can lead to data loss and inconsistencies.

## Impact

The impact of this alert is high, as it can result in:

* Data loss and inconsistencies
* Downtime or unavailability of the Thanos Receive instance
* Impact on dependent services and applications

## Diagnosis

To diagnose the root cause of the issue, follow these steps:

1. Check the Thanos Receive instance logs for errors and exceptions related to replication.
2. Verify the replication factor and hashring configuration for the instance.
3. Check the network connectivity and latency between the Thanos Receive instance and its dependencies.
4. Investigate any recent changes to the instance configuration or deployed code.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve any underlying issues causing the replication failures.
2. Verify that the replication factor and hashring configuration are correct and optimal for the instance.
3. Consider temporarily increasing the replication factor to reduce the risk of data loss.
4. Implement additional logging and monitoring to detect and alert on replication failures more quickly.
5. Restart the Thanos Receive instance if necessary.

Note: This runbook is a starting point and may need to be tailored to your specific environment and use case.