---
title: MgobBackupFailed
description: Troubleshooting for alert MgobBackupFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MgobBackupFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
MongoDB backup has failed

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/stefanprodan-mgob-exporter.yml" "MgobBackupFailed" %}}

<!-- Rule when generated

```yaml
alert: MgobBackupFailed
expr: changes(mgob_scheduler_backup_total{status="500"}[1h]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Mgob backup failed (instance {{ $labels.instance }})
    description: |-
        MongoDB backup has failed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/stefanprodan-mgob-exporter/MgobBackupFailed.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
