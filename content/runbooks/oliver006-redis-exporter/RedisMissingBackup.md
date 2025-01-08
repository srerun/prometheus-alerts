---
title: RedisMissingBackup
description: Troubleshooting for alert RedisMissingBackup
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisMissingBackup

Redis has not been backuped for 24 hours

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisMissingBackup" %}}

{{% comment %}}

```yaml
alert: RedisMissingBackup
expr: time() - redis_rdb_last_save_timestamp_seconds > 60 * 60 * 24
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis missing backup (instance {{ $labels.instance }})
    description: |-
        Redis has not been backuped for 24 hours
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisMissingBackup.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "RedisMissingBackup":

## Meaning

The "RedisMissingBackup" alert is triggered when the Redis database has not been backed up for more than 24 hours. This is a critical alert because it indicates that the data in Redis is at risk of being lost in case of a failure.

## Impact

If this alert is not addressed, the consequences could be severe:

* Data loss: If the Redis instance fails or is restarted, all data stored in Redis will be lost.
* System downtime: The system may experience downtime or errors until the Redis instance is restored from a backup.
* Business impact: The loss of Redis data can have a significant impact on business operations, leading to revenue loss, customer dissatisfaction, and damage to reputation.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Redis exporter logs to see if there are any errors or issues related to backups.
2. Verify that the Redis instance is configured correctly and that backups are scheduled to run regularly.
3. Check the system resources (CPU, memory, disk space) to ensure that they are not causing issues with the backup process.
4. Review the Redis configuration to ensure that the backup settings are correct and that the backup directory is accessible.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately trigger a manual backup of the Redis instance to ensure that the data is safe.
2. Investigate the cause of the backup failure and fix the underlying issue.
3. Verify that the backup process is scheduled to run regularly and that the backups are being stored correctly.
4. Consider implementing additional monitoring and alerting for Redis backups to ensure that this issue is caught earlier in the future.
5. Review and update the Redis configuration to ensure that it is optimal for the system's requirements.

Remember to update the runbook with any specific details or procedures relevant to your organization's Redis setup and backup processes.