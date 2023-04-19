let fibonacci = fn(x) {
    if (x <= 1) {
        return x;
    } else {
        return fibonacci(x - 1) + fibonacci(x - 2);
    }
}

for (let i=0; i<10; i = i + 1) {
    if (i == 5) {
        break;
    }
    puts(fibonacci(i));
};

let a = 0;
puts(a);