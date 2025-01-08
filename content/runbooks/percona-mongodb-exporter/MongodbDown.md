---
title: MongodbDown
description: Troubleshooting for alert MongodbDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbDown

MongoDB instance is down

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbDown" %}}

{{% comment %}}

```yaml
alert: MongodbDown
expr: mongodb_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB Down (instance {{ $labels.instance }})
    description: |-
        MongoDB instance is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbDown.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `MongodbDown`:

## Meaning

The `MongodbDown` alert is triggered when the `mongodb_up` metric returns a value of 0, indicating that the MongoDB instance is not responding or is down. This alert is critical, as it can cause data loss and affect the overall performance of the application.

## Impact

* Data loss or inconsistencies due to the unavailability of the MongoDB instance
* Application downtime or degradation, leading to a poor user experience
* Increased latency or errors in the application, causing revenue loss or customer dissatisfaction
* Potential security risks due to the exposure of sensitive data

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the MongoDB instance logs for errors or exceptional conditions that may be causing the downtime.
2. Verify that the MongoDB instance is properly configured and running with the correct credentials.
3. Check the system resources (CPU, memory, disk space) to ensure they are not depleted.
4. Verify that the network connection to the MongoDB instance is stable and not blocked by firewalls or other security measures.
5. Check the Percona MongoDB Exporter logs for any errors or issues that may be preventing it from reporting the MongoDB status accurately.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the MongoDB instance and verify that it is running correctly.
2. Check and adjust the MongoDB configuration to ensure it is optimal for the current workload.
3. Increase system resources (CPU, memory, disk space) if necessary to prevent resource depletion.
4. Verify that the network connection to the MongoDB instance is stable and not blocked by firewalls or other security measures.
5. Verify that the Percona MongoDB Exporter is correctly configured and reporting the MongoDB status accurately.

Additional steps may be necessary depending on the specific cause of the issue. Please refer to the Percona MongoDB Exporter documentation and MongoDB documentation for further troubleshooting and resolution steps.