---
title: MongodbNumberCursorsOpen
description: Troubleshooting for alert MongodbNumberCursorsOpen
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbNumberCursorsOpen

Too many cursors opened by MongoDB for clients (> 10k)

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbNumberCursorsOpen" %}}

{{% comment %}}

```yaml
alert: MongodbNumberCursorsOpen
expr: mongodb_metrics_cursor_open{state="total_open"} > 10000
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB number cursors open (instance {{ $labels.instance }})
    description: |-
        Too many cursors opened by MongoDB for clients (> 10k)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbNumberCursorsOpen.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule "MongodbNumberCursorsOpen":

## Meaning

This alert is triggered when the number of open cursors in a MongoDB instance exceeds 10,000. A cursor is a control structure that enables traversal over the records in a database. When a cursor is open, it holds a reference to the data and consumes system resources. A high number of open cursors can lead to performance degradation, increased memory usage, and even crashes.

## Impact

If left unchecked, a high number of open cursors can cause:

* Performance degradation: MongoDB will spend more time and resources managing the cursors, leading to slower query execution and increased latency.
* Increased memory usage: Each open cursor consumes memory, which can lead to memory exhaustion and crashes.
* Stability issues: In extreme cases, an excessive number of open cursors can cause MongoDB to crash or become unresponsive.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the MongoDB logs for any errors or warnings related to cursors.
2. Use the MongoDB shell or a GUI tool to inspect the current cursor statistics.
3. Identify the applications or users that are opening an excessive number of cursors.
4. Verify that the MongoDB instance has sufficient resources (CPU, memory, and disk space) to handle the workload.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and optimize the applications or users that are opening an excessive number of cursors.
2. Implement cursor timeouts or limits to prevent applications from holding onto cursors for too long.
3. Increase the MongoDB instance's resources (CPU, memory, and disk space) to handle the workload.
4. Consider implementing connection pooling or other optimization techniques to reduce the number of open cursors.
5. Monitor the MongoDB instance's performance and adjust the configuration as needed to prevent similar issues in the future.

Remember to investigate and address the root cause of the issue to prevent it from happening again in the future.