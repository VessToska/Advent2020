#![allow(unused_variables)]

use std::time::Instant;

mod read;

mod day1;
mod day2;
mod day3;

enum FnSig<D1, D2, DN> {
    U16U16(fn() -> D1),
    U32U32(fn() -> D2),
    NULNUL(fn() -> DN),
}

fn main() {
    use FnSig::*;
    let vec_days = vec![FnSig::U16U16(day1::main), FnSig::U32U32(day2::main), FnSig::NULNUL(day3::main)];
    let test_count = 1000;
    let mut test_ret = String::new();
    let time_total = Instant::now();
    for itr_day in vec_days.iter().enumerate() {
        let time_now = Instant::now();
        for _temp_test in 0..test_count {
            test_ret = match itr_day.1 {
                U16U16(func_ref) => format!("{:?}", func_ref()),
                U32U32(func_ref) => format!("{:?}", func_ref()),
                NULNUL(func_ref) => format!("{:?}", func_ref()),
            };
        }
        let time_elapsed = time_now.elapsed().as_secs_f64();
        let time_day = (itr_day.0)+1;
        println!("Day {}: {}", time_day, test_ret);
        println!("{} trials of day {}: {:.5}s\n", test_count, time_day, (time_elapsed/test_count as f64));
    }
    let time_elapsed = time_total.elapsed().as_secs_f64();
    println!("{} trials of all, averages: {:.5} seconds.", test_count, (time_elapsed/test_count as f64));
}