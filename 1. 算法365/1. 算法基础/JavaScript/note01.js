function cal1(n) {
    let stat_time = new Date();
    let sum = 0;
    for (let i = 0; i <= n; i++) {
        sum = sum + Math.pow(-1, i);
    }
    let stop_time = new Date();
    let time = stop_time.getTime() - stat_time.getTime();
    console.log("cal1 use time: ", time);
    return sum;
}

let n = 100000000;
console.log(cal1(n));

function cal2(n) {
    let start_time = new Date();
    if (n % 2 == 0){
        let stop_time = new Date();
        console.log("cal2 use time: ", stop_time.getTime() - start_time.getTime());
        return 1;
    } else {
        let stop_time = new Date();
        console.log("cal2 use time: ", stop_time.getTime() - start_time.getTime());
        return 0;
    }
}

console.log(cal2(n));