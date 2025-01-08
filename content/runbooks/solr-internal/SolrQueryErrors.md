---
title: SolrQueryErrors
description: Troubleshooting for alert SolrQueryErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SolrQueryErrors

Solr has increased query errors in collection {{ $labels.collection }} for replica {{ $labels.replica }} on {{ $labels.base_url }}.

<details>
  <summary>Alert Rule</summary>

{{% rule "solr/solr-internal.yml" "SolrQueryErrors" %}}

{{% comment %}}

```yaml
alert: SolrQueryErrors
expr: increase(solr_metrics_core_errors_total{category="QUERY"}[1m]) > 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Solr query errors (instance {{ $labels.instance }})
    description: |-
        Solr has increased query errors in collection {{ $labels.collection }} for replica {{ $labels.replica }} on {{ $labels.base_url }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/solr-internal/SolrQueryErrors.md

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
