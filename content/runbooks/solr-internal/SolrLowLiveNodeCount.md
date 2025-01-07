---
title: SolrLowLiveNodeCount
description: Troubleshooting for alert SolrLowLiveNodeCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SolrLowLiveNodeCount

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Solr collection {{ $labels.collection }} has less than two live nodes for replica {{ $labels.replica }} on {{ $labels.base_url }}.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: SolrLowLiveNodeCount
expr: solr_collections_live_nodes < 2
for: 0m
labels:
    severity: critical
annotations:
    summary: Solr low live node count (instance {{ $labels.instance }})
    description: |-
        Solr collection {{ $labels.collection }} has less than two live nodes for replica {{ $labels.replica }} on {{ $labels.base_url }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/SolrLowLiveNodeCount

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
