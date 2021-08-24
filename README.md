# xk6-counter

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/grafana/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
| ---------------------------------------------------------------------------------------------------------------------------- |

This projects implements an anonymized counter (`counter.up()`) and keyed counters (`counter.upNamed('myKey')`). The functions return the current value before
increasing it, and each VU will get a different value. This means it can be used to iterate over an array, where:
1. only one element will be used by each VU 
2. the array doesn't need to be sharded between the VUs; it will "dynamically" balance between them even if some elements take longer to process

| This totally doesn't work in distributed manner, so if there are multiple k6 instances you will need a separate service (API endpoint) which to do it. |
| ---------------------------------------------------------------------------------------------------------------------------- |

Predominantly because of the above this is very unlikely to ever get in k6 in it's current form, so please don't open issues :D. 

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [gvm](https://github.com/moovweb/gvm)
- [Git](https://git-scm.com/)

Then, install [xk6](https://github.com/grafana/xk6) and build your custom k6 binary with the Kafka extension:

1. Install `xk6`:
  ```shell
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  $ xk6 build --with github.com/mstoykov/xk6-counter@latest
  ```

# example

```javascript
import counter from "k6/x/counter";

export default function() {
    // anonymous counter:
    console.log(counter.up(), __VU, __ITER);

    // named counter:
    const key = "myKey";
    console.log(key, counter.upNamed(key), __VU, __ITER);
}
```
