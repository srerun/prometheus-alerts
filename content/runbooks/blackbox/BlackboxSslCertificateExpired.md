---
title: BlackboxSslCertificateExpired
description: Troubleshooting for alert BlackboxSslCertificateExpired
#published: true
date: 2023-12-13T19:41:13.737Z
tags: 
editor: markdown
dateCreated: 2023-12-13T16:44:44.229Z
---

# BlackboxSslCertificateExpired

## Meaning
[//]: # "Short paragraph that explains what the alert means"
SSL certificate has expired already 


<details>
  <summary>Alert Rule</summary>

```yaml
alert: BlackboxSslCertificateExpired
expr: round((last_over_time(probe_ssl_earliest_cert_expiry[10m]) - time()) / 86400, 0.1) < 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Blackbox SSL certificate expired (instance {{ $labels.instance }})
    description: |-
        SSL certificate has expired already
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/BlackboxSslCertificateExpired

```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
