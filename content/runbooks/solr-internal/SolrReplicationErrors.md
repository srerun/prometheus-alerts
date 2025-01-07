---
title: SolrReplicationErrors
description: Troubleshooting for alert SolrReplicationErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SolrReplicationErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Solr collection {{ $labels.collection }} has failed updates for replica {{ $labels.replica }} on {{ $labels.base_url }}.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: SolrReplicationErrors
expr: increase(solr_metrics_core_errors_total{category="REPLICATION"}[1m]) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Solr replication errors (instance {{ $labels.instance }})
    description: |-
        Solr collection {{ $labels.collection }} has failed updates for replica {{ $labels.replica }} on {{ $labels.base_url }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/SolrReplicationErrors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
