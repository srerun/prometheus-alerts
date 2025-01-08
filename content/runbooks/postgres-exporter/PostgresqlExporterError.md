---
title: PostgresqlExporterError
description: Troubleshooting for alert PostgresqlExporterError
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlExporterError

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Postgresql exporter is showing errors. A query may be buggy in query.yaml

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlExporterError" %}}

{{% comment %}}

```yaml
alert: PostgresqlExporterError
expr: pg_exporter_last_scrape_error > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql exporter error (instance {{ $labels.instance }})
    description: |-
        Postgresql exporter is showing errors. A query may be buggy in query.yaml
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlExporterError.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
