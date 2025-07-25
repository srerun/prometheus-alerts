---
title: sysLoginFailure
description: Troubleshooting for SNMP trap sysLoginFailure
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-111-TRAPS-MIB::sysLoginFailure 

User login is failed. 



Here is a runbook for the SNMP trap `E5-111-TRAPS-MIB::sysLoginFailure`:

## Meaning
The `E5-111-TRAPS-MIB::sysLoginFailure` SNMP trap is generated when a user login attempt fails. This trap indicates that a user has attempted to log in to a device or system, but the login credentials provided were invalid.

## Impact
The impact of this trap is that a user is unable to access the device or system, which may prevent them from performing their job functions or accessing critical resources. Additionally, repeated login failures could be a sign of a security threat, such as a brute-force attack, which could compromise the security of the device or system.

## Diagnosis
To diagnose the cause of this trap, follow these steps:

* Check the device or system logs to identify the username and IP address of the user who attempted to log in.
* Verify that the user's login credentials are correct and match the expected format.
* Check the device or system configuration to ensure that the user account is not locked out or disabled.
* Review the device or system security logs to see if there are any signs of a brute-force attack or other security threats.

## Mitigation
To mitigate the effects of this trap, follow these steps:

* Inform the user of the login failure and provide guidance on how to correct their login credentials.
* If the user is unable to log in due to a forgotten password, follow the organization's password reset procedures.
* If the login failure is due to a security threat, take immediate action to lock out the offending IP address and notify the security team.
* Consider implementing additional security measures, such as two-factor authentication or rate limiting, to prevent future login failures.
---




