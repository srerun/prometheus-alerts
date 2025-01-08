---
title: NatsHighRoutesCount
description: Troubleshooting for alert NatsHighRoutesCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighRoutesCount

High number of NATS routes ({{ $value }}) for {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighRoutesCount" %}}

{{% comment %}}

```yaml
alert: NatsHighRoutesCount
expr: gnatsd_varz_routes > 10
for: 3m
labels:
    severity: warning
annotations:
    summary: Nats high routes count (instance {{ $labels.instance }})
    description: |-
        High number of NATS routes ({{ $value }}) for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighRoutesCount.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "NatsHighRoutesCount":

## Meaning

The "NatsHighRoutesCount" alert is triggered when the number of NATS routes exceeds 10. This indicates that the NATS server is handling an unusually high number of routes, which can lead to performance issues and increased latency.

## Impact

A high number of NATS routes can cause:

* Increased memory usage and CPU utilization on the NATS server
* Slower message processing and delivery times
* Increased risk of server crashes or timeouts
* Potential disruption to applications depending on the NATS server

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NATS server logs for any errors or warnings related to route management
2. Verify that the NATS configuration is correct and optimized for the current workload
3. Review the application logs to identify any issues with message processing or delivery
4. Use the `nats_varz_routes` metric to monitor the number of routes over time and identify trends or patterns

## Mitigation

To mitigate the issue, follow these steps:

1. Review and adjust the NATS server configuration to optimize route management and reduce memory usage
2. Identify and address any application-level issues causing an excessive number of routes to be created
3. Consider implementing route filtering or aggregation to reduce the number of routes
4. Monitor the NATS server metrics closely to catch any signs of route growth and take proactive measures to prevent issues.

Additional resources:

* [NATS documentation: Route Management](https://docs.nats.io/running-a-nats-service/configuration/route_management)
* [NATS exporter documentation: gnatsd_varz_routes](https://github.com/nats-io/prometheus-nats-exporter/blob/master/docs/metrics.md#gnatsd_varz_routes)