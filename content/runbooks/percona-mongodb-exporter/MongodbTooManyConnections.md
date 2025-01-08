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

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbTooManyConnections" %}}

{{% comment %}}

```yaml
alert: MongodbTooManyConnections
expr: avg by(instance) (rate(mongodb_ss_connections{conn_type="current"}[1m])) / avg by(instance) (sum (mongodb_ss_connections) by (instance)) * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB too many connections (instance {{ $labels.instance }})
    description: |-
        Too many connections (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbTooManyConnections.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `MongodbTooManyConnections`:

## Meaning

The `MongodbTooManyConnections` alert is triggered when the average rate of current MongoDB connections exceeds 80% of the total available connections on a specific instance. This alert indicates that the MongoDB instance is experiencing a high connection load, which can lead to performance issues, slow queries, and even crashes.

## Impact

* Slow query performance
* Increased latency
* MongoDB instance crashes or becomes unresponsive
* Decreased overall system performance
* Potential data loss or corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check the MongoDB instance's connection metrics**: Review the `mongodb_ss_connections` metric to determine the current connection count and the rate of new connections.
2. **Investigate recent changes**: Look for recent changes to the application or database configuration that may be contributing to the high connection load.
3. **Verify MongoDB instance resources**: Check the MongoDB instance's resource utilization (e.g., CPU, memory, disk space) to ensure it is not under-provisioned.
4. **Review MongoDB logs**: Analyze the MongoDB logs to identify any errors or warnings related to connections or performance issues.

## Mitigation

To mitigate the issue, follow these steps:

1. **Investigate and address the root cause**: Identify the underlying reason for the high connection load and address it accordingly (e.g., optimize application code, adjust database configuration, etc.).
2. **Scale up the MongoDB instance**: Consider increasing the MongoDB instance's resources (e.g., add more nodes, increase instance size) to handle the increased connection load.
3. **Implement connection pooling**: Configure connection pooling to reduce the number of new connections and improve overall performance.
4. **Monitor MongoDB performance**: Closely monitor the MongoDB instance's performance and adjust as necessary to prevent future issues.

Remember to update this runbook with specific details about your environment and procedures to make it more effective.