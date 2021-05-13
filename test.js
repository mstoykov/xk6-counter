import counter from "k6/x/counter";

export let options = {
    vus: 5,
    iterations: 20,
}
export default function() {
    console.log(counter.up(), __VU, __ITER);
}
