---
title: PulsarSubscriptionVeryHighNumberOfBacklogEntries
description: Troubleshooting for alert PulsarSubscriptionVeryHighNumberOfBacklogEntries
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarSubscriptionVeryHighNumberOfBacklogEntries

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The number of subscription backlog entries is over 100k

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarSubscriptionVeryHighNumberOfBacklogEntries" %}}

{{% comment %}}

```yaml
alert: PulsarSubscriptionVeryHighNumberOfBacklogEntries
expr: sum(pulsar_subscription_back_log) by (subscription) > 100000
for: 1h
labels:
    severity: critical
annotations:
    summary: Pulsar subscription very high number of backlog entries (instance {{ $labels.instance }})
    description: |-
        The number of subscription backlog entries is over 100k
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarSubscriptionVeryHighNumberOfBacklogEntries.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
