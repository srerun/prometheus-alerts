---
title: ConsulServiceHealthcheckFailed
description: Troubleshooting for alert ConsulServiceHealthcheckFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ConsulServiceHealthcheckFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Service: `{{ $labels.service_name }}` Healthcheck: `{{ $labels.service_id }}`

<details>
  <summary>Alert Rule</summary>

{{% rule "consul/consul-exporter.yml" "ConsulServiceHealthcheckFailed" %}}

<!-- Rule when generated

```yaml
alert: ConsulServiceHealthcheckFailed
expr: consul_catalog_service_node_healthy == 0
for: 1m
labels:
    severity: critical
annotations:
    summary: Consul service healthcheck failed (instance {{ $labels.instance }})
    description: |-
        Service: `{{ $labels.service_name }}` Healthcheck: `{{ $labels.service_id }}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/consul-exporter/ConsulServiceHealthcheckFailed.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
