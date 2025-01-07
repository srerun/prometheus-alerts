---
title: MeilisearchIndexIsEmpty
description: Troubleshooting for alert MeilisearchIndexIsEmpty
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MeilisearchIndexIsEmpty

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Meilisearch instance is down

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: MeilisearchIndexIsEmpty
expr: meilisearch_index_docs_count == 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Meilisearch index is empty (instance {{ $labels.instance }})
    description: |-
        Meilisearch instance is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/MeilisearchIndexIsEmpty

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
