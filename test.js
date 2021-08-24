import counter from "k6/x/counter";

export let options = {
    vus: 5,
    iterations: 20,
}

export default function() {
    // anonymous counter:
    console.log(counter.up(), __VU, __ITER);

    // named counter:
    const key = "myKey";
    console.log(key, counter.upNamed(key), __VU, __ITER);
}
