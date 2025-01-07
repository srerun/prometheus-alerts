---
title: PulsarSubscriptionHighNumberOfBacklogEntries
description: Troubleshooting for alert PulsarSubscriptionHighNumberOfBacklogEntries
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarSubscriptionHighNumberOfBacklogEntries

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The number of subscription backlog entries is over 5k

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PulsarSubscriptionHighNumberOfBacklogEntries
expr: sum(pulsar_subscription_back_log) by (subscription) > 5000
for: 1h
labels:
    severity: warning
annotations:
    summary: Pulsar subscription high number of backlog entries (instance {{ $labels.instance }})
    description: |-
        The number of subscription backlog entries is over 5k
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PulsarSubscriptionHighNumberOfBacklogEntries

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
