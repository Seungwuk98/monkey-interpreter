let fibonacci = fn(x) {
    if (x <= 1) {
        return x;
    } else {
        return fibonacci(x - 1) + fibonacci(x - 2);
    }
}

let ok = for (let i=0; i<13; i = i + 1) {
    if (i == 10) {
        break;
    }
    puts(fibonacci(i));
};

if ()

let a = 0;
puts(a);