---
title: ClickhouseNoLiveReplicas
description: Troubleshooting for alert ClickhouseNoLiveReplicas
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseNoLiveReplicas

There are too few live replicas available, risking data loss and service disruption.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseNoLiveReplicas" %}}

{{% comment %}}

```yaml
alert: ClickhouseNoLiveReplicas
expr: ClickHouseErrorMetric_TOO_FEW_LIVE_REPLICAS == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: ClickHouse No Live Replicas (instance {{ $labels.instance }})
    description: |-
        There are too few live replicas available, risking data loss and service disruption.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseNoLiveReplicas.md

```

{{% /comment %}}

</details>


Here is a runbook for the ClickhouseNoLiveReplicas alert rule:

## Meaning

The ClickhouseNoLiveReplicas alert is triggered when there are too few live replicas available in a ClickHouse cluster, indicating a high risk of data loss and service disruption. This alert is critical and requires immediate attention.

## Impact

The impact of this alert is significant, as it can lead to:

* Data loss: With too few live replicas, ClickHouse may not be able to maintain data durability and consistency, resulting in potential data loss.
* Service disruption: The lack of live replicas can cause queries to fail, leading to service disruptions and impacting the availability of critical applications.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ClickHouse cluster status: Verify the current status of the ClickHouse cluster, including the number of live replicas and any error messages.
2. Investigate node status: Check the status of individual nodes in the cluster to identify any issues or failures that may be contributing to the lack of live replicas.
3. Review recent changes: Check the recent changes made to the ClickHouse configuration, node deployments, or application code to see if any changes may have caused the issue.
4. Check system resources: Verify that the underlying system resources, such as CPU, memory, and disk space, are sufficient to support the ClickHouse cluster.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve underlying issues: Address any underlying issues found during diagnosis, such as node failures or configuration errors.
2. Increase replica count: Temporarily increase the replica count to ensure data durability and consistency.
3. Implement rolling restarts: Perform rolling restarts of ClickHouse nodes to ensure that all nodes are healthy and live.
4. Monitor cluster status: Closely monitor the ClickHouse cluster status to ensure that the issue has been resolved and live replicas are available.

By following these steps, you should be able to diagnose and mitigate the ClickhouseNoLiveReplicas issue, restoring data durability and consistency to the ClickHouse cluster.