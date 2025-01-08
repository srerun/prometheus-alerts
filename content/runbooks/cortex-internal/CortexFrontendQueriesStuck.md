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

{{% rule "cortex/cortex-internal.yml" "CortexFrontendQueriesStuck" %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/cortex-internal/CortexFrontendQueriesStuck.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule "CortexFrontendQueriesStuck":

## Meaning

The "CortexFrontendQueriesStuck" alert is triggered when there are queued up queries in the Cortex query-frontend, indicating that the frontend is unable to process queries in a timely manner. This can lead to delays in query execution and impact the overall performance of the system.

## Impact

The impact of this alert is high, as it can cause:

* Delays in query execution, leading to poor user experience
* Increased latency in data ingestion and processing
* Potential losses in data fidelity and accuracy
* Increased load on the system, leading to potential resource exhaustion

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cortex query-frontend logs for errors or unusual behavior
2. Verify that the query-frontend is properly configured and running
3. Check the Cortex cluster's resource utilization (CPU, memory, disk space) to ensure it is within acceptable limits
4. Investigate any recent changes to the system or configuration that may have caused the issue
5. Review the query-frontend queue metrics to understand the volume and type of queries being executed

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Cortex query-frontend service to clear any stuck queries
2. Investigate and address any underlying issues causing the query-frontend to become stuck (e.g. resource utilization, configuration errors)
3. Consider scaling up the Cortex cluster to handle increased load
4. Implement query queuing and batch processing to reduce the load on the query-frontend
5. Consider implementing query timeouts and retries to prevent queries from getting stuck in the queue

Remember to follow established change management procedures when making any changes to the system.