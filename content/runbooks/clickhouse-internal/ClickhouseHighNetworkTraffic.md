---
title: ClickhouseHighNetworkTraffic
description: Troubleshooting for alert ClickhouseHighNetworkTraffic
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseHighNetworkTraffic

Network traffic is unusually high, may affect cluster performance.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseHighNetworkTraffic" %}}

{{% comment %}}

```yaml
alert: ClickhouseHighNetworkTraffic
expr: ClickHouseMetrics_NetworkSend > 250 or ClickHouseMetrics_NetworkReceive > 250
for: 5m
labels:
    severity: warning
annotations:
    summary: ClickHouse High Network Traffic (instance {{ $labels.instance }})
    description: |-
        Network traffic is unusually high, may affect cluster performance.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseHighNetworkTraffic.md

```

{{% /comment %}}

</details>


## Meaning

The ClickhouseHighNetworkTraffic alert is triggered when the ClickHouse instance is experiencing unusually high network traffic, with either the send or receive rates exceeding 250 bytes per second. This alert indicates a potential performance issue that may impact the overall cluster performance.

## Impact

If left unaddressed, high network traffic can lead to:

* Slow query performance
* Increased latency
* Decreased cluster responsiveness
* Potential out-of-memory errors
* Impaired data ingestion and processing capabilities

## Diagnosis

To diagnose the root cause of the high network traffic, perform the following steps:

1. **Check ClickHouse logs**: Review the ClickHouse logs to identify any errors, warnings, or unusual patterns that may be contributing to the high network traffic.
2. **Analyze query patterns**: Use the ClickHouse query log to identify any resource-intensive or long-running queries that may be causing the high network traffic.
3. **Investigate data ingestion**: Check the data ingestion pipeline to ensure that it is functioning correctly and not overwhelming the ClickHouse instance with high volumes of data.
4. **Verify network configuration**: Confirm that the network configuration is correctly set up and not causing any bottlenecks or connectivity issues.
5. **Check for resource shortages**: Verify that the ClickHouse instance has sufficient resources (CPU, memory, disk space) to handle the current workload.

## Mitigation

To mitigate the high network traffic, perform the following steps:

1. **Optimize queries**: Optimize resource-intensive queries to reduce their impact on the network traffic.
2. **Implement query throttling**: Throttle queries to prevent excessive resource usage and reduce network traffic.
3. **Adjust data ingestion rates**: Adjust the data ingestion rate to prevent overwhelming the ClickHouse instance with high volumes of data.
4. **Upgrade instance resources**: Consider upgrading the ClickHouse instance resources (CPU, memory, disk space) to better handle the current workload.
5. **Implement traffic shaping**: Consider implementing traffic shaping or quality of service (QoS) policies to prioritize critical traffic and reduce the impact of high network traffic.