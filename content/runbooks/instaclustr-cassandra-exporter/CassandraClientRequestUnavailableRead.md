---
title: CassandraClientRequestUnavailableRead
description: Troubleshooting for alert CassandraClientRequestUnavailableRead
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraClientRequestUnavailableRead

Some Cassandra client requests are unavailable to read - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraClientRequestUnavailableRead" %}}

{{% comment %}}

```yaml
alert: CassandraClientRequestUnavailableRead
expr: changes(cassandra_client_request_unavailable_exceptions_total{operation="read"}[1m]) > 0
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra client request unavailable read (instance {{ $labels.instance }})
    description: |-
        Some Cassandra client requests are unavailable to read - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraClientRequestUnavailableRead.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule:

## Meaning

The CassandraClientRequestUnavailableRead alert indicates that some Cassandra client requests are unable to read data from the Cassandra cluster. This can lead to application errors, data loss, and poor system performance.

## Impact

The impact of this alert can be significant, as it may cause:

* Application errors and failures
* Data loss or inconsistencies
* Poor system performance and slow response times
* Downtime and revenue loss

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the Cassandra cluster status and node health using the Cassandra UI or `nodetool` command.
2. Verify that the Cassandra client configuration is correct and up-to-date.
3. Check the Cassandra system logs for errors related to read operations.
4. Verify that the Cassandra cluster has sufficient resources (e.g., CPU, memory, disk space) to handle read requests.
5. Investigate recent changes to the Cassandra schema, data model, or application code that may be causing read issues.

## Mitigation

To mitigate this alert, follow these steps:

1. Restart the Cassandra client or application to reconnect to the Cassandra cluster.
2. Verify that the Cassandra cluster is properly configured and resourced to handle read requests.
3. Investigate and fix any underlying issues causing read errors, such as data inconsistencies or schema changes.
4. Consider implementing load balancing or queuing mechanisms to handle high volumes of read requests.
5. Monitor the Cassandra cluster and client performance to ensure that the issue does not recur.

Remember to follow the link to the runbook provided in the alert annotation for more detailed steps and guidelines.