---
title: MongodbCursorsTimeouts
description: Troubleshooting for alert MongodbCursorsTimeouts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbCursorsTimeouts

Too many cursors are timing out

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbCursorsTimeouts" %}}

{{% comment %}}

```yaml
alert: MongodbCursorsTimeouts
expr: increase(mongodb_ss_metrics_cursor_timedOut[1m]) > 100
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB cursors timeouts (instance {{ $labels.instance }})
    description: |-
        Too many cursors are timing out
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbCursorsTimeouts.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "MongodbCursorsTimeouts":

## Meaning

The "MongodbCursorsTimeouts" alert is triggered when the number of MongoDB cursors timing out exceeds 100 in a 1-minute window. This indicates that a significant number of operations are waiting for an excessive amount of time, leading to performance issues and potential data inconsistencies.

## Impact

The impact of this alert is high, as it can cause:

* Slow query performance, leading to delays in data retrieval and processing
* Increased latency, affecting the overall responsiveness of the application
* Potential data inconsistencies, as operations may not complete successfully
* Downtime or unavailability of critical services, depending on the application's dependence on MongoDB

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the MongoDB server logs for any errors or warnings related to cursor timeouts
2. Verify that the MongoDB instance is properly configured, including cursor timeouts and batch sizes
3. Investigate the MongoDB query patterns and optimize any inefficient queries
4. Check for any resource constraints, such as CPU, memory, or disk usage, that may be contributing to the timeouts
5. Verify that the Percona MongoDB Exporter is properly configured and collecting metrics correctly

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and optimize inefficient queries that may be contributing to the cursor timeouts
2. Increase the cursor timeout value in the MongoDB configuration, if necessary
3. Implement query caching or other performance optimization techniques to reduce the load on the MongoDB instance
4. Consider scaling up the MongoDB instance or adding additional resources to alleviate resource constraints
5. Monitor the situation closely and adjust the alert threshold or configuration as needed to prevent false positives or false negatives

Remember to refer to the Percona MongoDB Exporter documentation and MongoDB best practices for more detailed guidance on troubleshooting and optimizing MongoDB performance.