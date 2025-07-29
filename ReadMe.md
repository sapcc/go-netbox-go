<!--
SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company

SPDX-License-Identifier: Apache-2.0
-->

go-netbox-go
============

Package go-netbox-go implements a custom client for the netbox [API][1]

This is a minimal implementation due to the ever-changing api of netbox.
Please do not implement unnecessary details or unused api calls due to the added complexity if the API changes.

Development
-----------
Use `NETBOX_URL` (e.g. https://netbox.corp) and `NETBOX_TOKEN` to direct the tests against your instance of netbox.

To test against a new version:
- create the new branch with the netbox api version
- delete the govcr-fixtures
- rerun tests against the new api version and see what breaks
- fix tests



[1]: https://netbox.readthedocs.io/en/stable/
