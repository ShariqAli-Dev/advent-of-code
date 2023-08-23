use std::fs;
use std::cmp;

fn main() {
    let elf_calories_file = fs
        ::read_to_string("../../../inputs/1.txt")
        .expect("could not read the input file");

    let mut elf_calories = 0;
    let mut max_elf_calories = 0;
    for calorie in elf_calories_file.lines() {
        let is_whitespace = calorie.trim().is_empty();
        if is_whitespace {
            max_elf_calories = cmp::max(max_elf_calories, elf_calories);
            elf_calories = 0;
            continue;
        }
        elf_calories += calorie.parse::<i32>().unwrap();
    }

    println!(" max elf calories: {max_elf_calories}");
}
