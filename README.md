# consul-service

Returns random service from consul requested by name

## Usage

When one call:

```
consul-service demo.consul.io redis
```

It should return something similar to this:

```
Requesting: http://demo.consul.io...

[{"Node":"nyc3-worker-1","Address":"104.131.86.92","ServiceID":"redis","ServiceName":"redis","ServiceTags":null,"ServicePort":6379},{"Node":"nyc3-worker-3","Address":"104.131.59.59","ServiceID":"redis","ServiceName":"redis","ServiceTags":null,"ServicePort":6379},{"Node":"nyc3-worker-2","Address":"104.131.109.224","ServiceID":"redis","ServiceName":"redis","ServiceTags":null,"ServicePort":6379}]%
```

Where:
- **demo.consul.io** is Consul url: http://demo.consul.io
- **redis** is a service available in Consul
