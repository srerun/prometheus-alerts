---
title: CassandraClientRequestUnavailableRead
description: Troubleshooting for alert CassandraClientRequestUnavailableRead
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraClientRequestUnavailableRead

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Cassandra client requests are unavailable to read - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/instaclustr-cassandra-exporter.yml" "CassandraClientRequestUnavailableRead" %}}

{{% comment %}}

```yaml
alert: CassandraClientRequestUnavailableRead
expr: changes(cassandra_client_request_unavailable_exceptions_total{operation="read"}[1m]) > 0
for: 2m
labels:
    severity: critical
annotations:
    summary: Cassandra client request unavailable read (instance {{ $labels.instance }})
    description: |-
        Some Cassandra client requests are unavailable to read - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/instaclustr-cassandra-exporter/CassandraClientRequestUnavailableRead.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
