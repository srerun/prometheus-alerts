---
title: FreeRADIUSHighMalformedRequests
description: Troubleshooting for alert FreeRADIUSHighMalformedRequests
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSHighMalformedRequests

There are malformed requests in the FreeRADIUS server. Investigate the source of the malformed requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSHighMalformedRequests" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSHighMalformedRequests
expr: freeradius_total_acct_malformed_requests > 0 or freeradius_total_auth_malformed_requests > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Malformed Requests Detected
    description: There are malformed requests in the FreeRADIUS server. Investigate the source of the malformed requests.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiushighmalformedrequests/

```

{{% /comment %}}

</details>


Here is a sample runbook for the FreeRADIUSHighMalformedRequests alert:

## Meaning

The FreeRADIUSHighMalformedRequests alert is triggered when the FreeRADIUS server reports one or more malformed requests within a 5-minute window. This alert indicates that there is an issue with the authentication or accounting requests being sent to the FreeRADIUS server, which can lead to authentication failures, system instability, or even security vulnerabilities.

## Impact

The impact of this alert can be significant, as it may prevent legitimate users from authenticating to the network or accessing network resources. Additionally, malformed requests can cause performance issues, leading to slow authentication times or even crashes. In extreme cases, malformed requests can be exploited by attackers to gain unauthorized access to the network.

## Diagnosis

To diagnose the root cause of the malformed requests, follow these steps:

1. **Check the FreeRADIUS logs**: Review the FreeRADIUS logs to identify the source of the malformed requests. Look for any errors or warnings related to authentication or accounting requests.
2. **Verify NAS configuration**: Check the configuration of the Network Access Server (NAS) devices that are sending requests to the FreeRADIUS server. Ensure that they are properly configured and sending valid requests.
3. **Check for misconfigured clients**: Identify any misconfigured clients that may be sending malformed requests to the FreeRADIUS server.
4. **Review firewall and network configurations**: Verify that there are no firewall or network configuration issues that could be causing requests to be malformed or blocked.

## Mitigation

To mitigate the effects of the malformed requests, follow these steps:

1. **Identify and block the source**: Identify the source of the malformed requests and block them at the firewall or NAS level to prevent further requests from being sent.
2. **Adjust NAS configuration**: Adjust the NAS configuration to ensure that valid requests are being sent to the FreeRADIUS server.
3. **Implement request filtering**: Implement request filtering on the FreeRADIUS server to drop or reject malformed requests.
4. **Monitor and analyze requests**: Continuously monitor and analyze requests to identify trends and patterns that may indicate a potential security issue.

By following these steps, you should be able to diagnose and mitigate the root cause of the malformed requests, reducing the risk of authentication failures, system instability, or security vulnerabilities.