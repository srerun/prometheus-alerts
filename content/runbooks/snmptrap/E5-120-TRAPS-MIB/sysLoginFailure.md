---
title: sysLoginFailure
description: Troubleshooting for SNMP trap sysLoginFailure
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-120-TRAPS-MIB::sysLoginFailure 

User login is failed. 



## Meaning

This SNMP trap, E5-120-TRAPS-MIB::sysLoginFailure, is triggered when a user login attempt to a network device fails. This could be due to various reasons such as incorrect username, password, or account lockout.

## Impact

The impact of this trap can be significant as it may indicate a security threat or an attempt to gain unauthorized access to the network device. Repeated login failures can also lead to account lockouts, which can cause disruption to legitimate users. Additionally, if the failed login attempts are successful, it can lead to unauthorized access to sensitive information and potential security breaches.

## Diagnosis

To diagnose the root cause of this issue, perform the following steps:

* Check the device logs to identify the username and IP address of the device attempting to login.
* Verify the username and password combination used for the login attempt.
* Check if the account is locked out due to previous failed login attempts.
* Review the device configuration to ensure that the login credentials are correct and up-to-date.
* If the issue persists, contact the user who attempted to login to verify their credentials and intentions.

## Mitigation

To mitigate this issue, perform the following steps:

* Implement a robust password policy to ensure strong and unique passwords are used.
* Enable account lockout policies to prevent brute-force attacks.
* Configure the device to log and alert on failed login attempts.
* Implement rate-limiting on login attempts to prevent rapid-fire login attempts.
* Perform regular security audits to identify and address potential security vulnerabilities.
* Educate users on the importance of password security and the consequences of unauthorized access attempts.
---




