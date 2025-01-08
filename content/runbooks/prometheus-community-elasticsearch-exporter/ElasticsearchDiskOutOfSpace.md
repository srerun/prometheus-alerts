---
title: ElasticsearchDiskOutOfSpace
description: Troubleshooting for alert ElasticsearchDiskOutOfSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchDiskOutOfSpace

The disk usage is over 90%

<details>
  <summary>Alert Rule</summary>

{{% rule "elasticsearch/prometheus-community-elasticsearch-exporter.yml" "ElasticsearchDiskOutOfSpace" %}}

{{% comment %}}

```yaml
alert: ElasticsearchDiskOutOfSpace
expr: elasticsearch_filesystem_data_available_bytes / elasticsearch_filesystem_data_size_bytes * 100 < 10
for: 0m
labels:
    severity: critical
annotations:
    summary: Elasticsearch disk out of space (instance {{ $labels.instance }})
    description: |-
        The disk usage is over 90%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-community-elasticsearch-exporter/ElasticsearchDiskOutOfSpace.md

```

{{% /comment %}}

</details>


Here is a runbook for the ElasticsearchDiskOutOfSpace alert:

## Meaning

The ElasticsearchDiskOutOfSpace alert is triggered when the available disk space on an Elasticsearch node falls below 10% of the total disk capacity. This indicates that the disk is running out of space, which can cause Elasticsearch to become unstable, slow, or even crash.

## Impact

If left unaddressed, a disk running out of space can have severe consequences, including:

* Elasticsearch cluster instability or downtime
* Data loss or corruption
* Increased latency or timeouts for searches and indexing operations
* Potential data loss or inconsistency due to incomplete writes
* Increased risk of node crashes or failures

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the disk usage of the affected Elasticsearch node using the `df -h` command or a similar tool.
2. Verify that the disk usage is indeed above 90% by checking the `elasticsearch_filesystem_data_available_bytes` and `elasticsearch_filesystem_data_size_bytes` metrics.
3. Check the Elasticsearch logs for any errors or warnings related to disk space issues.
4. Review the indexing and search patterns to identify any unusual activity that may be contributing to the disk space issue.

## Mitigation

To mitigate the issue, follow these steps:

1. **Free up disk space**: Immediately free up disk space by deleting unnecessary files, logs, or indices.
2. **Add more disk space**: Consider adding more disk space to the affected node or distributing the data across multiple nodes.
3. **Optimize indexing and searching**: Optimize indexing and searching patterns to reduce disk usage.
4. **Monitor disk usage**: Regularly monitor disk usage to catch potential issues before they escalate.
5. **Consider cluster rebalancing**: If the issue persists, consider rebalancing the Elasticsearch cluster to distribute data more evenly across nodes.

Remember to follow your organization's specific procedures and guidelines for resolving disk space issues in your Elasticsearch cluster.