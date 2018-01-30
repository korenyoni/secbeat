# Secbeat

A heartbeat shipped as a Golang binary, publishing security-related events to ElasticSearch.

## Motivation

Beats such as [filebeat](https://www.elastic.co/products/beats/filebeat) and [auditbeat](https://www.elastic.co/products/beats/auditbeat) exist, but their scope limits the flexibility
for field creation. Perhaps it is possible to do so with the addition of logstash,
but that is one more component to worry about.

Furthermore, there's no guarantee that filebeat and/or auditbeat can capture
all the events necessary for making HIPAA and/or SOC-2 security logs.

Secbeat is a heartbeat designed to be shipped as a single binary and make use of environment variables
to connect to Elasticsearch and publish security related events.

## State of project

WIP
