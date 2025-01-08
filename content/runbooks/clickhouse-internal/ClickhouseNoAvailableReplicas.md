---
title: ClickhouseNoAvailableReplicas
description: Troubleshooting for alert ClickhouseNoAvailableReplicas
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseNoAvailableReplicas

No available replicas in ClickHouse.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseNoAvailableReplicas" %}}

{{% comment %}}

```yaml
alert: ClickhouseNoAvailableReplicas
expr: ClickHouseErrorMetric_NO_AVAILABLE_REPLICA == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: ClickHouse No Available Replicas (instance {{ $labels.instance }})
    description: |-
        No available replicas in ClickHouse.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseNoAvailableReplicas.md

```

{{% /comment %}}

</details>


Here is the runbook for the ClickhouseNoAvailableReplicas alert:

## Meaning
The ClickhouseNoAvailableReplicas alert is triggered when there are no available replicas in a ClickHouse instance. This means that the instance is not able to provide redundancy and high availability, which can lead to data loss and downtime in case of a failure.

## Impact
The impact of this alert is high, as it can lead to:

* Data loss in case of a failure
* Downtime and unavailability of the ClickHouse instance
* Increased risk of data corruption and inconsistencies
* Potential loss of business critical data

## Diagnosis
To diagnose the root cause of this alert, follow these steps:

* Check the ClickHouse instance logs for any errors or warnings related to replica availability
* Verify the replica configuration and ensure that it is correct and up-to-date
* Check the status of the replicas and ensure that they are online and available
* Verify that there are no network connectivity issues between the replicas and the master node
* Check the disk usage and verify that there is enough free space available on the replicas

## Mitigation
To mitigate this alert, follow these steps:

* Immediately investigate and resolve any issues related to replica availability
* Ensure that the replica configuration is correct and up-to-date
* Restart the ClickHouse instance and verify that the replicas are online and available
* Implement a backup and restore process to ensure business continuity in case of a failure
* Consider adding additional replicas to ensure high availability and redundancy
* Verify that the ClickHouse instance is properly configured for high availability and redundancy

It is recommended to closely monitor the ClickHouse instance after resolving the issue to ensure that the replicas remain available and that the instance remains stable.