---
title: BlackboxSslCertificateWillExpireSoon
description: Troubleshooting for alert BlackboxSslCertificateWillExpireSoon
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxSslCertificateWillExpireSoon

## Meaning
[//]: # "Short paragraph that explains what the alert means"
SSL certificate expires in less than 20 days

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: BlackboxSslCertificateWillExpireSoon
expr: 3 <= round((last_over_time(probe_ssl_earliest_cert_expiry[10m]) - time()) / 86400, 0.1) < 20
for: 0m
labels:
    severity: warning
annotations:
    summary: Blackbox SSL certificate will expire soon (instance {{ $labels.instance }})
    description: |-
        SSL certificate expires in less than 20 days
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/BlackboxSslCertificateWillExpireSoon

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
