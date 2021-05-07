

## Guide

https://www.elastic.co/guide/index.html

## Install

source:

```sh
curl -L -O https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.12.1-linux-x86_64.tar.gz
tar -xzvf elasticsearch-7.12.1-linux-x86_64.tar.gz
cd elasticsearch-7.12.1
./bin/elasticsearch
```

rpm:

```sh
curl -L -O https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.12.1-x86_64.rpm
sudo rpm -i elasticsearch-7.12.1-x86_64.rpm
sudo service elasticsearch start
```



## Simple Use

####  Make sure Elasticsearch is up and running

To test that the Elasticsearch daemon is up and running, try sending an HTTP GET request on port 9200.

```shell
curl http://127.0.0.1:9200
```

You should see a response similar to this:

```sh
{
  "name" : "QtI5dUu",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "DMXhqzzjTGqEtDlkaMOzlA",
  "version" : {
    "number" : "7.12.1",
    "build_flavor" : "default",
    "build_type" : "tar",
    "build_hash" : "00d8bc1",
    "build_date" : "2018-06-06T16:48:02.249996Z",
    "build_snapshot" : false,
    "lucene_version" : "7.3.1",
    "minimum_wire_compatibility_version" : "5.6.0",
    "minimum_index_compatibility_version" : "5.0.0"
  },
  "tagline" : "You Know, for Search"
}
```



### Install Kibana

[Kibana](https://www.elastic.co/products/kibana) is an open source analytics and visualization platform designed to work with Elasticsearch. 

**deb, rpm, or linux:**

```sh
curl -L -O https://artifacts.elastic.co/downloads/kibana/kibana-7.12.1-linux-x86_64.tar.gz
tar xzvf kibana-7.12.1-linux-x86_64.tar.gz
cd kibana-7.12.1-linux-x86_64/
./bin/kibana
```



## Quick Start



This guide helps beginners how to:

- Install and run Elasticsearch in a test environment
- Add data to Elasticsearch
- Search and sort data
- Extract fields from unstructured content during a search



