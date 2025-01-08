---
title: RedisTooManyMasters
description: Troubleshooting for alert RedisTooManyMasters
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisTooManyMasters

Redis cluster has too many nodes marked as master.

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisTooManyMasters" %}}

{{% comment %}}

```yaml
alert: RedisTooManyMasters
expr: count(redis_instance_info{role="master"}) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis too many masters (instance {{ $labels.instance }})
    description: |-
        Redis cluster has too many nodes marked as master.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisTooManyMasters.md

```

{{% /comment %}}

</details>


Here is the runbook for the RedisTooManyMasters alert:

## Meaning

RedisTooManyMasters alert is triggered when the number of Redis instances with the "master" role exceeds 1. This indicates that there are multiple nodes in the Redis cluster that are marked as master, which can lead to data inconsistency and other issues.

## Impact

Having multiple masters in a Redis cluster can cause:

* Data inconsistencies and losses
* Increased latency and decreased performance
* Potential data corruption
* Unstable cluster behavior

If left unaddressed, this issue can lead to severe consequences, including data loss, system downtime, and revenue loss.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Redis cluster configuration to identify the nodes marked as master.
2. Verify that the Redis exporter is correctly configured and reporting accurate metrics.
3. Investigate the Redis instance logs for any errors or warnings related to master election or node communication.
4. Use Redis CLI or other tools to inspect the cluster state and node roles.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the extra master node(s) and demote them to slave or remove them from the cluster.
2. Verify that the Redis cluster configuration is correct and consistent across all nodes.
3. Restart the Redis instances to ensure that the changes take effect.
4. Monitor the cluster for any further issues or errors.
5. Review and update the Redis exporter configuration to ensure accurate monitoring and alerting.

Remember to follow proper procedures and take necessary precautions when making changes to the Redis cluster to avoid further disruption or data loss.