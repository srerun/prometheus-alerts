---
title: MongodbTooManyConnections
description: Troubleshooting for alert MongodbTooManyConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbTooManyConnections

Too many connections (> 80%)

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbTooManyConnections" %}}

{{% comment %}}

```yaml
alert: MongodbTooManyConnections
expr: avg by(instance) (rate(mongodb_connections{state="current"}[1m])) / avg by(instance) (sum (mongodb_connections) by (instance)) * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB too many connections (instance {{ $labels.instance }})
    description: |-
        Too many connections (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbTooManyConnections.md

```

{{% /comment %}}

</details>


## Meaning

The `MongodbTooManyConnections` alert is triggered when the average rate of current MongoDB connections per instance exceeds 80% of the total available connections. This indicates that the MongoDB instance is experiencing a high load, which can lead to performance issues, increased latency, and potentially even crashes.

## Impact

If left unchecked, a high number of connections to MongoDB can cause:

* Performance degradation: As the number of connections increases, MongoDB may struggle to handle the load, leading to slow query response times and decreased overall performance.
* Increased latency: Higher connection numbers can result in increased latency for applications relying on MongoDB, affecting user experience and overall system responsiveness.
* Crash or restart: In extreme cases, excessive connections can cause MongoDB to crash or restart, leading to data loss, corruption, or unavailability.

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. **Check MongoDB connection metrics**: Review the `mongodb_connections` metric to identify the instance(s) with high connection rates.
2. **Investigate connection patterns**: Analyze the connection patterns to determine if there are any unusual or unexpected spikes in connections.
3. **Verify application behavior**: Check application logs and metrics to ensure that the application is not experiencing any issues that might be contributing to the high connection count.
4. **Review MongoDB configuration**: Verify that the MongoDB configuration is optimal for the current workload, including settings such as connection limits, timeout values, and resource allocation.
5. **Check for any recent changes**: Investigate if there have been any recent changes to the application, infrastructure, or MongoDB configuration that might be contributing to the high connection count.

## Mitigation

To mitigate the `MongodbTooManyConnections` alert, follow these steps:

1. **Identify and optimize resource-intensive queries**: Analyze MongoDB query logs to identify resource-intensive queries and optimize them to reduce the load on the instance.
2. **Implement connection pooling**: Consider implementing connection pooling mechanisms to reduce the number of connections to MongoDB.
3. **Adjust MongoDB configuration**: Adjust MongoDB configuration settings, such as connection limits, timeout values, and resource allocation, to optimize performance and reduce connection counts.
4. **Scale MongoDB instance**: Consider scaling the MongoDB instance to handle the increased load, either by adding more resources or by distributing the load across multiple instances.
5. **Monitor and adjust**: Continuously monitor MongoDB performance and adjust the mitigation steps as needed to ensure the connection count remains within a healthy range.