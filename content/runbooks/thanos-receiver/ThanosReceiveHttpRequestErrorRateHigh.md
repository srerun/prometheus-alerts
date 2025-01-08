---
title: ThanosReceiveHttpRequestErrorRateHigh
description: Troubleshooting for alert ThanosReceiveHttpRequestErrorRateHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHttpRequestErrorRateHigh

Thanos Receive {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveHttpRequestErrorRateHigh" %}}

{{% comment %}}

```yaml
alert: ThanosReceiveHttpRequestErrorRateHigh
expr: (sum by (job) (rate(http_requests_total{code=~"5..", job=~".*thanos-receive.*", handler="receive"}[5m]))/  sum by (job) (rate(http_requests_total{job=~".*thanos-receive.*", handler="receive"}[5m]))) * 100 > 5
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Receive Http Request Error Rate High (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveHttpRequestErrorRateHigh.md

```

{{% /comment %}}

</details>


## Meaning

The `ThanosReceiveHttpRequestErrorRateHigh` alert is triggered when the rate of HTTP requests with a 5xx status code (indicating an error) exceeds 5% of the total requests received by the Thanos Receive component. This alert indicates that Thanos Receive is experiencing issues handling incoming requests, which can lead to data loss or inconsistencies.

## Impact

* Data loss: Failed requests may result in missing data points, affecting the accuracy and completeness of metrics.
* Increased latency: Errors in request handling can cause delays in data ingestion, leading to slower query performance and delayed alerting.
* System instability: Prolonged periods of high error rates can lead to system instability, further exacerbating the issue.

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check Thanos Receive logs**: Review the logs of the affected Thanos Receive instance(s) to identify the root cause of the errors. Look for patterns in the error messages or request payloads.
2. **Verify request patterns**: Analyze the request patterns and payloads to determine if there are any unusual or malformed requests that may be contributing to the errors.
3. **Investigate Thanos Receive configuration**: Review the Thanos Receive configuration to ensure it is correctly set up and functioning as expected.
4. **Check underlying infrastructure**: Verify that the underlying infrastructure, including servers, networks, and storage, is functioning correctly and not experiencing issues.

## Mitigation

To mitigate the issue, follow these steps:

1. **Restart Thanos Receive**: Restart the affected Thanos Receive instance(s) to reset the request handling and clear any temporary errors.
2. **Adjust Thanos Receive configuration**: Review and adjust the Thanos Receive configuration to ensure it is optimized for the current request volume and patterns.
3. **Investigate and fix underlying issues**: Address any underlying infrastructure issues, such as server or network problems, that may be contributing to the errors.
4. **Implement request rate limiting**: Consider implementing request rate limiting or throttling to prevent excessive requests and reduce the load on Thanos Receive.
5. **Monitor and review**: Continuously monitor the alert and review the mitigation steps to ensure the issue is completely resolved and does not reoccur.