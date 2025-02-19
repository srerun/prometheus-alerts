---
title: SolrLowLiveNodeCount
description: Troubleshooting for alert SolrLowLiveNodeCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SolrLowLiveNodeCount

Solr collection {{ $labels.collection }} has less than two live nodes for replica {{ $labels.replica }} on {{ $labels.base_url }}.

<details>
  <summary>Alert Rule</summary>

{{% rule "solr/solr-internal.yml" "SolrLowLiveNodeCount" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/solr-internal/SolrLowLiveNodeCount.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
