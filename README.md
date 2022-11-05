# GETTING STARTED


You only need the `./release/run` and `./release/.env` to use kibanalert.  Here's how:

```
cd ./release
cp env.sample .env
```

Fill out the details of `.env`.

Type `./run` to start kibanalert.

-----

I assume you use the following index mappings:

```
PUT /kibanalert/_mapping
{
    "properties": {
      "alert_id": {
        "type": "keyword"
      },
      "date": {
        "type": "date"
      },
      "reason": {
        "type": "text"
      },
      "rule_id": {
        "type": "keyword"
      },
      "service_name": {
        "type": "text"
      }
    }
}
```

And that alerts populate index with this document template:

```
{
  "alert_id": "{{alert.id}}",
  "rule_id": "{{rule.id}}",
  "reason": "{{context.reason}}",
  "service_name": "{{context.serviceName}}",
  "date": "{{date}}"
}
```
