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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Redis has not been backuped for 24 hours

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisMissingBackup" %}}

<!-- Rule when generated

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

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
