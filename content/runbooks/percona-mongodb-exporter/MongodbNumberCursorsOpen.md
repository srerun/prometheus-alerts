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

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbNumberCursorsOpen" %}}

{{% comment %}}

```yaml
alert: MongodbNumberCursorsOpen
expr: mongodb_ss_metrics_cursor_open{csr_type="total"} > 10 * 1000
for: 2m
labels:
    severity: warning
annotations:
    summary: MongoDB number cursors open (instance {{ $labels.instance }})
    description: |-
        Too many cursors opened by MongoDB for clients (> 10k)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbNumberCursorsOpen.md

```

{{% /comment %}}

</details>


## Meaning

The `MongodbNumberCursorsOpen` alert is triggered when the number of open cursors in a MongoDB instance exceeds 10,000. This can indicate a potential performance issue, as a high number of open cursors can lead to increased memory usage and slower query performance.

## Impact

If left unchecked, a high number of open cursors can cause:

* Increased memory usage, leading to performance degradation and potential OOM errors
* Slower query performance, impacting application responsiveness and user experience
* Potential crashes or instability in the MongoDB instance

## Diagnosis

To diagnose the root cause of the alert, follow these steps:

1. Check the MongoDB logs for any errors or warnings related to cursors or memory usage
2. Verify that the `mongodb_ss_metrics_cursor_open` metric is correctly configured and reporting accurate data
3. Investigate the application(s) connected to the MongoDB instance, as excessive cursor creation can be caused by inefficient query patterns or poor application design
4. Review the MongoDB instance configuration, including cursor timeout settings and resource allocation

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and optimize the application(s) responsible for creating excessive cursors
2. Implement cursor timeouts or limits to prevent prolonged cursor maintenance
3. Adjust MongoDB instance configuration to optimize cursor management and memory allocation
4. Consider implementing query optimization techniques, such as indexing or caching, to reduce query latency and cursor creation
5. Monitor the `mongodb_ss_metrics_cursor_open` metric and adjust the alert threshold as necessary to ensure timely detection of potential issues.

For additional guidance, refer to the [Percona MongoDB Exporter Runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbNumberCursorsOpen.md) for more detailed steps and best practices.