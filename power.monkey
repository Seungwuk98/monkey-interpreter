let t = fn(x) { 
    if (x < 1) { return 1 } 
    else { return x * t(x-1) } 
}

puts(t(10));