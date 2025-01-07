---
title: RedisMissingMaster
description: Troubleshooting for alert RedisMissingMaster
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisMissingMaster

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Redis cluster has no node marked as master.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: RedisMissingMaster
expr: (count(redis_instance_info{role="master"}) or vector(0)) < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis missing master (instance {{ $labels.instance }})
    description: |-
        Redis cluster has no node marked as master.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RedisMissingMaster

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
