---
title: NatsHighConnectionCount
description: Troubleshooting for alert NatsHighConnectionCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighConnectionCount

High number of NATS connections ({{ $value }}) for {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighConnectionCount" %}}

{{% comment %}}

```yaml
alert: NatsHighConnectionCount
expr: gnatsd_varz_connections > 100
for: 3m
labels:
    severity: warning
annotations:
    summary: Nats high connection count (instance {{ $labels.instance }})
    description: |-
        High number of NATS connections ({{ $value }}) for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighConnectionCount.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `NatsHighConnectionCount`:

## Meaning

The `NatsHighConnectionCount` alert is triggered when the number of connections to a NATS instance exceeds 100. This alert is categorized as a warning, indicating a potential issue that may impact system performance.

## Impact

A high number of connections to a NATS instance can lead to:

* Increased memory usage, potentially causing performance degradation or even crashes.
* Slowdowns in message processing, leading to delayed or lost messages.
* Increased latency and decreased overall system responsiveness.

## Diagnosis

To diagnose the root cause of the high connection count, perform the following steps:

1. Check the NATS instance's configuration to ensure it is properly tuned for the expected load.
2. Investigate the applications or services connected to the NATS instance to identify any issues or misconfigurations.
3. Review the NATS instance's logs to identify any errors or warnings related to connection handling.
4. Use the `gnatsd_varz` metrics to monitor the connection count and identify any trends or patterns.

## Mitigation

To mitigate the high connection count, perform the following steps:

1. Identify and disconnect any unnecessary or idle connections to the NATS instance.
2. Optimize the NATS instance's configuration to handle the current load, if necessary.
3. Implement connection pooling or other load-reducing measures to minimize the number of connections.
4. Consider increasing the resources allocated to the NATS instance, such as increasing the instance size or adding more nodes to the cluster.

Remember to monitor the NATS instance's performance and adjust the mitigation steps as necessary to ensure the connection count returns to a normal level.