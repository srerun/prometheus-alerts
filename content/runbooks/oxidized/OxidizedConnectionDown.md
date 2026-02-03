---
title: OxidizedConnectionDown
description: Troubleshooting for alert OxidizedConnectionDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OxidizedConnectionDown

The prometheus exporter cannot communicate with Oxidized instance.

<details>
  <summary>Alert Rule</summary>

{{% rule "oxidized/oxidized.yml" "OxidizedConnectionDown" %}}

{{% comment %}}

```yaml
alert: OxidizedConnectionDown
expr: |
    oxidized_status == 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Oxidized exporter cannot connect to Oxidized
    description: |
        The prometheus exporter cannot communicate with Oxidized instance.
        oxidized_status = 0 (error)
        Check exporter logs, connectivity and Oxidized service status.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/oxidized/oxidizedconnectiondown/

```

{{% /comment %}}

</details>


## Meaning
The OxidizedConnectionDown alert is triggered when the Prometheus exporter is unable to connect to the Oxidized instance. This alert indicates a critical issue, as it prevents the exporter from collecting and reporting data from Oxidized. The `oxidized_status == 0` condition specifically indicates an error status, suggesting a problem with the connection or the Oxidized service itself.

## Impact
The impact of this alert is significant, as it affects the monitoring and visibility of network device configurations. If the Prometheus exporter cannot connect to Oxidized, it will not be able to collect data on device configurations, backups, and changes. This can lead to a lack of visibility into network device configurations, making it challenging to detect and respond to issues, such as configuration drift or unauthorized changes. Additionally, this can also affect the ability to track compliance with regulatory requirements or organizational policies.

## Diagnosis
To diagnose the issue, follow these steps:
1. **Check exporter logs**: Review the Prometheus exporter logs to identify any error messages related to the connection to Oxidized.
2. **Verify connectivity**: Check the network connectivity between the Prometheus exporter and the Oxidized instance, ensuring that there are no firewall rules or network issues blocking the connection.
3. **Check Oxidized service status**: Verify that the Oxidized service is running and functioning correctly, checking for any errors or issues in the Oxidized logs.
4. **Verify configuration**: Review the Prometheus exporter configuration to ensure that it is correctly configured to connect to the Oxidized instance.

## Mitigation
To mitigate the issue, follow these steps:
1. **Restart the Prometheus exporter**: Attempt to restart the Prometheus exporter to see if it can reestablish a connection to Oxidized.
2. **Check and fix network connectivity issues**: If network connectivity issues are found, work with the network team to resolve them.
3. **Restart the Oxidized service**: If the Oxidized service is not running or is malfunctioning, attempt to restart it.
4. **Update the Prometheus exporter configuration**: If configuration issues are found, update the Prometheus exporter configuration to correctly connect to the Oxidized instance.
5. **Monitor the alert**: Continue to monitor the alert to ensure that the issue is resolved and the Prometheus exporter can connect to Oxidized.