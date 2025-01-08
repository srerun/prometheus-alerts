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

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbCursorsTimeouts" %}}

{{% comment %}}

```yaml
alert: MongodbCursorsTimeouts
expr: increase(mongodb_metrics_cursor_timed_out_total[1m]) > 100
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB cursors timeouts (instance {{ $labels.instance }})
    description: |-
        Too many cursors are timing out
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbCursorsTimeouts.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `MongodbCursorsTimeouts`:

## Meaning

This alert is triggered when the number of MongoDB cursor timeouts exceeds 100 in a 1-minute interval, indicating that MongoDB is experiencing performance issues.

## Impact

* Slow query performance: Cursor timeouts can lead to slow query performance, which can impact application responsiveness and user experience.
* Data inconsistencies: In extreme cases, cursor timeouts can cause data inconsistencies or even data loss.
* System instability: Prolonged cursor timeouts can lead to system instability, as resources are wasted on waiting for unresponsive queries.

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check MongoDB logs**: Review MongoDB logs for any error messages related to cursor timeouts, slow queries, or other performance issues.
2. **Analyze query performance**: Use tools like MongoDB's built-in `explain` method or third-party tools like MongoDB Compass or MongoDB Atlas to analyze query performance and identify slow queries.
3. **Verify cursor configuration**: Check the cursor configuration to ensure it is set up correctly and not causing unnecessary timeouts.
4. **Check system resources**: Verify that the system has sufficient resources (CPU, memory, disk space) to handle the workload.

## Mitigation

To mitigate the issue, follow these steps:

1. **Optimize slow queries**: Identify and optimize slow queries to reduce their execution time and prevent cursor timeouts.
2. **Adjust cursor configuration**: Adjust the cursor configuration to allow for more time to execute or to increase the batch size to reduce the number of cursors.
3. **Scale MongoDB resources**: Scale up MongoDB resources (e.g., increase instance size, add more nodes) to handle the workload.
4. **Implement query retries**: Implement query retries to prevent data inconsistencies in case of timeouts.

Note: Refer to the MongoDB documentation and your organization's best practices for specific guidance on resolving cursor timeouts.