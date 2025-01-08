---
title: ClickhouseZookeeperConnectionIssues
description: Troubleshooting for alert ClickhouseZookeeperConnectionIssues
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseZookeeperConnectionIssues

ClickHouse is experiencing issues with ZooKeeper connections, which may affect cluster state and coordination.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseZookeeperConnectionIssues" %}}

{{% comment %}}

```yaml
alert: ClickhouseZookeeperConnectionIssues
expr: avg(ClickHouseMetrics_ZooKeeperSession) != 1
for: 3m
labels:
    severity: warning
annotations:
    summary: ClickHouse ZooKeeper Connection Issues (instance {{ $labels.instance }})
    description: |-
        ClickHouse is experiencing issues with ZooKeeper connections, which may affect cluster state and coordination.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseZookeeperConnectionIssues.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the ClickhouseZookeeperConnectionIssues alert:

## Meaning

The ClickhouseZookeeperConnectionIssues alert is triggered when the average value of ClickHouseMetrics_ZooKeeperSession is not equal to 1, indicating issues with ClickHouse's connection to ZooKeeper. This alert is critical because ZooKeeper is responsible for maintaining the cluster state and coordination in ClickHouse. Any disruptions to this connection can lead to inconsistent data, errors, or even cluster failures.

## Impact

The impact of this alert can be significant, as it may lead to:

* Inconsistent cluster state
* Errors or failures in data processing and queries
* Downtime or instability of the ClickHouse cluster
* Loss of data or data inconsistency

## Diagnosis

To diagnose the root cause of the ClickhouseZookeeperConnectionIssues alert, follow these steps:

1. Check the ClickHouse logs for any error messages related to ZooKeeper connections.
2. Verify that the ZooKeeper service is running and healthy.
3. Check the network connectivity between ClickHouse and ZooKeeper nodes.
4. Review the ClickHouse configuration to ensure that the ZooKeeper connection settings are correct.
5. Run ClickHouse queries to verify that the cluster is in a healthy state.

## Mitigation

To mitigate the ClickhouseZookeeperConnectionIssues alert, follow these steps:

1. Restart the ClickHouse service to re-establish the ZooKeeper connection.
2. Verify that the ZooKeeper service is running and healthy.
3. Check and repair any network issues between ClickHouse and ZooKeeper nodes.
4. Update the ClickHouse configuration to ensure that the ZooKeeper connection settings are correct.
5. Perform a rolling restart of the ClickHouse cluster to ensure that all nodes are connected to ZooKeeper.
6. Monitor the ClickHouse cluster for any further issues and take corrective action as needed.