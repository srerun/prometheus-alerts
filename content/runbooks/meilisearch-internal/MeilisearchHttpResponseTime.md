---
title: MeilisearchHttpResponseTime
description: Troubleshooting for alert MeilisearchHttpResponseTime
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MeilisearchHttpResponseTime

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Meilisearch http response time is too high

<details>
  <summary>Alert Rule</summary>

{{% rule "meilisearch/meilisearch-internal.yml" "MeilisearchHttpResponseTime" %}}

<!-- Rule when generated

```yaml
alert: MeilisearchHttpResponseTime
expr: meilisearch_http_response_time_seconds > 0.5
for: 0m
labels:
    severity: warning
annotations:
    summary: Meilisearch http response time (instance {{ $labels.instance }})
    description: |-
        Meilisearch http response time is too high
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/meilisearch-internal/MeilisearchHttpResponseTime.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
