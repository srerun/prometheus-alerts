---
title: CassandraStorageExceptions
description: Troubleshooting for alert CassandraStorageExceptions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraStorageExceptions

Something is going wrong with cassandra storage - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraStorageExceptions" %}}

{{% comment %}}

```yaml
alert: CassandraStorageExceptions
expr: changes(cassandra_storage_exceptions_total[1m]) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra storage exceptions (instance {{ $labels.instance }})
    description: |-
        Something is going wrong with cassandra storage - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraStorageExceptions.md

```

{{% /comment %}}

</details>


Here is a runbook for the CassandraStorageExceptions alert:

## Meaning

The CassandraStorageExceptions alert is triggered when there is an increase in Cassandra storage exceptions over a 1-minute period. This indicates that something is going wrong with Cassandra storage, which can lead to data inconsistency, unavailability, or even complete cluster failure.

## Impact

The impact of this alert can be severe, as it may cause:

* Data loss or corruption
* Decreased performance and availability of Cassandra cluster
* Increased latency and timeouts for applications dependent on Cassandra
* Potential cascading failures in upstream services

## Diagnosis

To diagnose the root cause of the CassandraStorageExceptions alert, follow these steps:

1. Check the Cassandra cluster logs for errors related to storage exceptions
2. Verify that the Cassandra node(s) are running and healthy
3. Check disk usage and free space on Cassandra nodes
4. Verify that there are no network connectivity issues between Cassandra nodes
5. Check for any recent changes to Cassandra configuration or schema
6. Review Cassandra metrics, such as disk usage, read/write latency, and error rates, to identify any trends or patterns

## Mitigation

To mitigate the CassandraStorageExceptions alert, follow these steps:

1. Investigate and resolve the root cause of the storage exceptions
2. Restart the Cassandra node(s) if necessary
3. Free up disk space on Cassandra nodes if disk usage is high
4. Adjust Cassandra configuration or schema to prevent storage exceptions
5. Monitor Cassandra metrics and logs closely to ensure the issue is fully resolved
6. Consider implementing additional monitoring and alerting for Cassandra storage exceptions to catch issues earlier in the future