---
title: CloudflareHttp5xxErrorRate
description: Troubleshooting for alert CloudflareHttp5xxErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CloudflareHttp5xxErrorRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cloudflare high HTTP 5xx error rate (> 5% for domain {{ $labels.zone }})

<details>
  <summary>Alert Rule</summary>

{{% rule "cloudflare/lablabs-cloudflare-exporter.yml" "CloudflareHttp5xxErrorRate" %}}

<!-- Rule when generated

```yaml
alert: CloudflareHttp5xxErrorRate
expr: (sum by (zone) (rate(cloudflare_zone_requests_status{status=~"^5.."}[5m])) / on (zone) sum by (zone) (rate(cloudflare_zone_requests_status[5m]))) * 100 > 5
for: 0m
labels:
    severity: critical
annotations:
    summary: Cloudflare http 5xx error rate (instance {{ $labels.instance }})
    description: |-
        Cloudflare high HTTP 5xx error rate (> 5% for domain {{ $labels.zone }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/lablabs-cloudflare-exporter/CloudflareHttp5xxErrorRate.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
