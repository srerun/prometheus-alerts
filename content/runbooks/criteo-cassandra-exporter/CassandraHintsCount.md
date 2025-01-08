---
title: CassandraHintsCount
description: Troubleshooting for alert CassandraHintsCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraHintsCount

Cassandra hints count has changed on {{ $labels.instance }} some nodes may go down

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraHintsCount" %}}

{{% comment %}}

```yaml
alert: CassandraHintsCount
expr: changes(cassandra_stats{name="org:apache:cassandra:metrics:storage:totalhints:count"}[1m]) > 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra hints count (instance {{ $labels.instance }})
    description: |-
        Cassandra hints count has changed on {{ $labels.instance }} some nodes may go down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraHintsCount.md

```

{{% /comment %}}

</details>


Here is a runbook for the CassandraHintsCount alert:

## Meaning

The CassandraHintsCount alert is triggered when the number of hints in a Cassandra cluster changes rapidly, indicating potential issues with data consistency and node availability.

## Impact

If left unaddressed, this issue can lead to:

* Data inconsistencies across nodes in the Cassandra cluster
* Node failures or crashes, resulting in downtime and data loss
* Performance degradation and increased latency for applications relying on Cassandra

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cassandra cluster's overall health and performance using metrics such as node up/down status, CPU usage, and disk usage.
2. Investigate the specific node(s) where the hints count changed rapidly using tools like `nodetool` or the Cassandra GUI.
3. Review the Cassandra logs for any error messages or warnings related to hints or node communication.
4. Verify that the cluster is properly configured and that there are no network connectivity issues.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and fix any underlying issues causing nodes to go down or hints to build up, such as:
	* Network connectivity problems
	* Disk space issues
	* High CPU usage
	* Misconfigured Cassandra settings
2. Run `nodetool repair` to ensure data consistency across nodes.
3. Consider increasing the `phi_convict_threshold` value to prevent nodes from being incorrectly marked as down.
4. Monitor the cluster closely for any further issues and adjust the `cassandra.yaml` configuration as needed to prevent similar issues in the future.

Remember to check the provided links for more information on Cassandra configuration and troubleshooting:

* [Cassandra documentation](https://cassandra.apache.org/doc/latest/)
* [Cassandra troubleshooting guide](https://docs.datastax.com/en/troubleshooting_guide/doc/troubleshooting_guide/index.html)

By following these steps, you should be able to diagnose and mitigate the issue causing the CassandraHintsCount alert.