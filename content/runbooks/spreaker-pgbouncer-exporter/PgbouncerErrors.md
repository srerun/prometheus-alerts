---
title: PgbouncerErrors
description: Troubleshooting for alert PgbouncerErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PgbouncerErrors

PGBouncer is logging errors. This may be due to a a server restart or an admin typing commands at the pgbouncer console.

<details>
  <summary>Alert Rule</summary>

{{% rule "pgbouncer/spreaker-pgbouncer-exporter.yml" "PgbouncerErrors" %}}

{{% comment %}}

```yaml
alert: PgbouncerErrors
expr: increase(pgbouncer_errors_count{errmsg!="server conn crashed?"}[1m]) > 10
for: 0m
labels:
    severity: warning
annotations:
    summary: PGBouncer errors (instance {{ $labels.instance }})
    description: |-
        PGBouncer is logging errors. This may be due to a a server restart or an admin typing commands at the pgbouncer console.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/spreaker-pgbouncer-exporter/PgbouncerErrors.md

```

{{% /comment %}}

</details>


Here is a runbook for the PgbouncerErrors alert rule:

## Meaning

The PgbouncerErrors alert is triggered when the number of errors logged by Pgbouncer increases by more than 10 in a 1-minute period, excluding errors with the message "server conn crashed?".

## Impact

This alert indicates that Pgbouncer is experiencing errors, which may be caused by a server restart or an administrator interacting with the Pgbouncer console. If left unchecked, these errors could lead to connectivity issues or data loss.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Pgbouncer logs for error messages to identify the root cause of the errors.
2. Verify that no recent server restarts or maintenance activities have occurred.
3. Investigate if any administrative commands were executed on the Pgbouncer console.
4. Review the Pgbouncer configuration to ensure it is correct and up-to-date.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve the underlying cause of the errors, based on the diagnosis.
2. Restart the Pgbouncer service to clear any transient errors.
3. Verify that the Pgbouncer configuration is correct and up-to-date.
4. Monitor the Pgbouncer logs for any further errors.

Additional resources:

* [Pgbouncer documentation](https://pgbouncer.github.io/)
* [Pgbouncer exporter documentation](https://github.com/spreaker/pgbouncer_exporter)