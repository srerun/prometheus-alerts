---
title: CortexFrontendQueriesStuck
description: Troubleshooting for alert CortexFrontendQueriesStuck
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexFrontendQueriesStuck

There are queued up queries in query-frontend.

<details>
  <summary>Alert Rule</summary>

{{% rule "cortex/coretex-internal.yml" "CortexFrontendQueriesStuck" %}}

{{% comment %}}

```yaml
alert: CortexFrontendQueriesStuck
expr: sum by (job) (cortex_query_frontend_queue_length) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Cortex frontend queries stuck (instance {{ $labels.instance }})
    description: |-
        There are queued up queries in query-frontend.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/coretex-internal/CortexFrontendQueriesStuck.md

```

{{% /comment %}}

</details>


## Meaning

The `CortexFrontendQueriesStuck` alert is triggered when there are queued up queries in the Cortex query-frontend. This means that the query-frontend is not able to process queries in a timely manner, leading to a backlog of queries.

## Impact

This alert has a critical severity, indicating that it can have a significant impact on the system. The stuck queries can cause:

* Delays in query execution, leading to slower response times for users
* Increased latency in the system, affecting overall performance
* Potential loss of data or incomplete results if queries are not processed correctly

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cortex query-frontend logs for errors or warnings that may indicate the cause of the stuck queries.
2. Verify the query-frontend configuration and ensure it is correctly set up.
3. Check the system resources (CPU, memory, disk space) to ensure they are not overwhelmed.
4. Verify that there are no network connectivity issues between the query-frontend and the underlying storage systems.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the query-frontend service to clear the queued up queries.
2. Investigate and address the root cause of the issue, such as:
	* Fixing configuration errors
	* Resolving system resource issues
	* Addressing network connectivity problems
3. Consider scaling up or optimizing the query-frontend resources to handle the query load.
4. Implement query queuing limits or rate limiting to prevent similar issues in the future.

Note: For more detailed steps and specific solutions, refer to the runbook link provided in the alert rule: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/coretex-internal/CortexFrontendQueriesStuck.md