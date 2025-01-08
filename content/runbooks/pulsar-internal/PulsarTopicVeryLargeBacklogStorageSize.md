---
title: PulsarTopicVeryLargeBacklogStorageSize
description: Troubleshooting for alert PulsarTopicVeryLargeBacklogStorageSize
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarTopicVeryLargeBacklogStorageSize

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The topic backlog storage size is over 20 GB

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarTopicVeryLargeBacklogStorageSize" %}}

<!-- Rule when generated

```yaml
alert: PulsarTopicVeryLargeBacklogStorageSize
expr: sum(pulsar_storage_size > 20*1024*1024*1024) by (topic)
for: 1h
labels:
    severity: critical
annotations:
    summary: Pulsar topic very large backlog storage size (instance {{ $labels.instance }})
    description: |-
        The topic backlog storage size is over 20 GB
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarTopicVeryLargeBacklogStorageSize.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
