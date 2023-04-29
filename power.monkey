let fibonacci = fn(x) {
    if (x <= 1) {
        return 1;
    }
    return fibonacci(x - 1) + fibonacci(x - 2);
};

puts(fibonacci(10))

let make_fn = fn(x) {
    return fn(u) {
        return x + u;
    };
}

let my_fn = make_fn(3)
puts(my_fn(2))