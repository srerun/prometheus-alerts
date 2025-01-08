---
title: CassandraNodeIsUnavailable
description: Troubleshooting for alert CassandraNodeIsUnavailable
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraNodeIsUnavailable

Cassandra Node is unavailable - {{ $labels.cassandra_cluster }} {{ $labels.exported_endpoint }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraNodeIsUnavailable" %}}

{{% comment %}}

```yaml
alert: CassandraNodeIsUnavailable
expr: sum(cassandra_endpoint_active) by (cassandra_cluster,instance,exported_endpoint) < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra Node is unavailable (instance {{ $labels.instance }})
    description: |-
        Cassandra Node is unavailable - {{ $labels.cassandra_cluster }} {{ $labels.exported_endpoint }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraNodeIsUnavailable.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "CassandraNodeIsUnavailable":

## Meaning

The CassandraNodeIsUnavailable alert is triggered when a Cassandra node in a cluster is unavailable, indicating a potential issue with data storage and retrieval. This alert is critical as it can impact the overall performance and availability of the system.

## Impact

The unavailability of a Cassandra node can lead to:

* Data loss or inconsistencies
* Increased latency or timeouts for read and write operations
* Reduced system performance and availability
* Potential data corruption or inconsistencies
* Impact on business operations and revenue

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cassandra node's status using the Cassandra CLI or a monitoring tool like Prometheus.
2. Verify that the node is not responding to requests or is showing high latency.
3. Review the Cassandra logs to identify any error messages or exceptions related to the node's unavailability.
4. Check the system resources (CPU, memory, disk space) to ensure they are not exhausted.
5. Verify that the node is properly configured and that there are no network connectivity issues.

## Mitigation

To mitigate the issue, follow these steps:

1. **Restart the Cassandra node**: If the node is not responding, try restarting it to see if it recovers.
2. **Investigate and resolve underlying issues**: Identify and resolve any underlying issues causing the node's unavailability, such as resource exhaustion, network connectivity problems, or configuration errors.
3. **Failover to another node**: If the node cannot be recovered, failover to another node in the cluster to minimize data loss and ensure system availability.
4. **Restore from backup**: If data loss has occurred, restore from a backup to ensure data integrity and consistency.
5. **Perform a rolling restart**: Perform a rolling restart of the Cassandra cluster to ensure all nodes are updated and healthy.

Remember to follow your organization's specific procedures and guidelines for incident response and resolution.