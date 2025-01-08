---
title: EtcdMemberCommunicationSlow
description: Troubleshooting for alert EtcdMemberCommunicationSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdMemberCommunicationSlow

Etcd member communication slowing down, 99th percentile is over 0.15s

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdMemberCommunicationSlow" %}}

{{% comment %}}

```yaml
alert: EtcdMemberCommunicationSlow
expr: histogram_quantile(0.99, rate(etcd_network_peer_round_trip_time_seconds_bucket[1m])) > 0.15
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd member communication slow (instance {{ $labels.instance }})
    description: |-
        Etcd member communication slowing down, 99th percentile is over 0.15s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdMemberCommunicationSlow.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "EtcdMemberCommunicationSlow":

## Meaning

The EtcdMemberCommunicationSlow alert is triggered when the 99th percentile of etcd network peer round trip time exceeds 0.15 seconds over a 1-minute period, indicating that etcd member communication is slowing down. This could be a sign of network issues, high latency, or resource constraints affecting etcd performance.

## Impact

Slowed etcd member communication can lead to:

* Delayed writes and reads to etcd
* Increased latency in distributed systems that rely on etcd
* Potential for data inconsistency or loss
* Increased risk of etcd cluster instability or split-brain scenarios

## Diagnosis

To diagnose the root cause of the slow etcd member communication:

1. Check the etcd cluster logs for any errors or warnings related to network communication or resource constraints.
2. Investigate the network infrastructure and connectivity between etcd members to identify any issues or bottlenecks.
3. Verify that etcd members have sufficient resources (CPU, memory, disk space) to operate efficiently.
4. Check for any recent changes to the etcd configuration, network topology, or system updates that may be contributing to the slow communication.
5. Use tools like `etcdctl` or `curl` to test the communication between etcd members and verify the round-trip time.

## Mitigation

To mitigate the effects of slow etcd member communication:

1. Identify and address any network issues or bottlenecks, such as packet loss, high latency, or congestion.
2. Optimize etcd configuration to reduce the load on the network, such as increasing the `sync-interval` or `send-queue-size`.
3. Consider upgrading etcd members to improve performance or adding more resources (e.g., increasing CPU or memory).
4. Implement retries and timeouts in applications that interact with etcd to improve resilience to slow communication.
5. Consider implementing etcd clustering features, such as leader election or distributed locks, to reduce the impact of slow communication.

Remember to investigate and address the root cause of the issue to prevent future occurrences of slow etcd member communication.