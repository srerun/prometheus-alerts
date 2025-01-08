---
title: SolrUpdateErrors
description: Troubleshooting for alert SolrUpdateErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SolrUpdateErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Solr collection {{ $labels.collection }} has failed updates for replica {{ $labels.replica }} on {{ $labels.base_url }}.

<details>
  <summary>Alert Rule</summary>

{{% rule "solr/solr-internal.yml" "SolrUpdateErrors" %}}

{{% comment %}}

```yaml
alert: SolrUpdateErrors
expr: increase(solr_metrics_core_update_handler_errors_total[1m]) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Solr update errors (instance {{ $labels.instance }})
    description: |-
        Solr collection {{ $labels.collection }} has failed updates for replica {{ $labels.replica }} on {{ $labels.base_url }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/solr-internal/SolrUpdateErrors.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
