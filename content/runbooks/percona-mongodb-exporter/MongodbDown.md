---
title: MongodbDown
description: Troubleshooting for alert MongodbDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
MongoDB instance is down

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: MongodbDown
expr: mongodb_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB Down (instance {{ $labels.instance }})
    description: |-
        MongoDB instance is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/MongodbDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
