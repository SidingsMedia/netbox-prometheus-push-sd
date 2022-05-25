<!-- 
SPDX-FileCopyrightText: Copyright (c) 2021-2022 Sidings Media 
SPDX-License-Identifier: MIT
-->


# Netbox Prometheus Push SD

This is a simple project built with Go to make use of Netbox's built in
webhook functionality to make Netbox a source of truth for the
Prometheus monitoring system

## Setting up

### Webhooks

In order for this project to work, a webhook should be created in Netbox
to fire on the creation, update or deletion of either a device or
virtual machine. The URL should be in the following form:
```
<domain>:8081/v1/target
```
where `<domain>` is the domain or IP address of your prometheus server.
If you have changed the port of the sidecar, you will have to update
this.

### Custom fields

Currently there are two custom fields that are supported. These are
`UUID` and `fqdn`. If they are present and not empty then they will be
included with the request. `UUID` will be added as a Prometheus tag and
if `fqdn` is set then instead of using the primary IP address, the fully
qualified domain name with DNS resolution will be used instead. In order
to select the target type, a custom `prometheus_target` field should be
added. If this is not added then it is assumed that the target is node
exporter running on port `9100`. To specify a custom target, you should
use the form `<target name>://<host>:<port>`. Note: `<host>` is optional
and is only really useful when using monitors such as SNMP. If it is not
included then either the FQDN or IP of the target host will be used.

### Labels

Below is a list of items that are and are not included in Prometheus
target labels.
#### Tags

No tags are included in the data that is sent to Prometheus. This is to
make it less cluttered and complex. 

#### Role

In order to differentiate between different types of devices, the role
will be added as a Prometheus label for the target.

#### Location

The location of the target is also added in order to assist in
administration.
### Body template

In order to correctly add new targets, the following template should be
used.

```jinja
{
    "type": "{{ event }}",
    "name": "{{ data['name'] }}",
    {% if (data["custom_fields"]["UUID"] is defined) and  (data["custom_fields"]["UUID"] is not null)%}
        "uuid": "{{ data['custom_fields']['UUID'] }}",
    {% endif %}
    {% if (data["custom_fields"]["fqdn"] is defined) and (data["custom_fields"]["fqdn"] is not none) %}
        "fqdn": "{{ data['custom_fields']['fqdn'] }}",
    {% endif %}
    "target":
        {% if (data["custom_fields"]["prometheus_target"] is defined) and (data["custom_fields"]["prometheus_target"] is not none) %}
            "{{ data["custom_fields"]["prometheus_target"] }}",
        {% else %}
            "node://:9100",
        {% endif %}
    "location": "{{ data['site']["slug"] }}",
    "ipv4": "{{ data["primary_ip4"]["address"] }}",
    "ipv6": "{{ data["primary_ip6"]["address"] }}",
    "platform": "{{ data["platform"]["slug"] }}",
    "status": "{{ data["status"]["value"] }}",
    "role": "{{ data["role"]["slug"] }}",
    "vm": 
        {% if data["vcpus"] is not defined %} 
            false
        {% else %}
            true,
        "cluster_name": "{{ data["cluster"]["name"] }}"
        {% endif %}
}
```

## Licence
This repo uses the [REUSE](https://reuse.software) standard in order to communicate the correct licence for the file. For those unfamiliar with the standard the licence for each file can be found in one of three places. The licence will either be in a comment block at the top of the file, in a `.license` file with the same name as the file, or in the dep5 file located in the `.reuse` directory. If you are unsure of the licencing terms please contact [contact@sidingsmedia.com](mailto:contact@sidingsmedia.com?subject=SMRC%20Licence). All files committed to this repo must contain valid licencing information or the pull request can not be accepted.
